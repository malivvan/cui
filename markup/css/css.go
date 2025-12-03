package css

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/gorilla/css/scanner"
)

type Stylesheet []*Rule

func ToString(rules []*Rule) string {
	result := ""

	for _, rule := range rules {
		if result != "" {
			result += "\n"
		}
		result += rule.String()
	}

	return result
}

// Declaration represents a parsed style property
type Declaration struct {
	Property  string
	Value     string
	Important bool
}

// NewDeclaration instanciates a new Declaration
func NewDeclaration() *Declaration {
	return &Declaration{}
}

// Returns string representation of the Declaration
func (decl *Declaration) String() string {
	return decl.StringWithImportant(true)
}

// StringWithImportant returns string representation with optional !important part
func (decl *Declaration) StringWithImportant(option bool) string {
	result := fmt.Sprintf("%s: %s", decl.Property, decl.Value)

	if option && decl.Important {
		result += " !important"
	}

	result += ";"

	return result
}

// Equal returns true if both Declarations are equals
func (decl *Declaration) Equal(other *Declaration) bool {
	return (decl.Property == other.Property) && (decl.Value == other.Value) && (decl.Important == other.Important)
}

//
// DeclarationsByProperty
//

// DeclarationsByProperty represents sortable style declarations
type DeclarationsByProperty []*Declaration

// Len Implements sort.Interface
func (declarations DeclarationsByProperty) Len() int {
	return len(declarations)
}

// Swap Implements sort.Interface
func (declarations DeclarationsByProperty) Swap(i, j int) {
	declarations[i], declarations[j] = declarations[j], declarations[i]
}

// Less Implements sort.Interface
func (declarations DeclarationsByProperty) Less(i, j int) bool {
	return declarations[i].Property < declarations[j].Property
}

const (
	importantSuffixRegexp = `(?i)\s*!important\s*$`
)

var (
	importantRegexp *regexp.Regexp
)

func init() {
	importantRegexp = regexp.MustCompile(importantSuffixRegexp)
}

// Parser represents a CSS parser
type Parser struct {
	scan *scanner.Scanner // Tokenizer

	// Tokens parsed but not consumed yet
	tokens []*scanner.Token

	// Rule embedding level
	embedLevel int
}

// Parse parses a whole stylesheet
func Parse(s string) (Stylesheet, error) {
	return (&Parser{scan: scanner.New(s)}).ParseStylesheet()
}

// ParseDeclarations parses CSS declarations
func ParseDeclarations(s string) ([]*Declaration, error) {
	return (&Parser{scan: scanner.New(s)}).ParseDeclarations()
}

// ParseStylesheet parses a stylesheet
func (parser *Parser) ParseStylesheet() (Stylesheet, error) {
	if _, err := parser.parseBOM(); err != nil {
		return nil, err
	}
	rules, err := parser.ParseRules()
	if err != nil {
		return nil, err
	}
	return rules, nil
}

// ParseRules parses a list of rules
func (parser *Parser) ParseRules() ([]*Rule, error) {
	var result []*Rule

	inBlock := false
	if parser.tokenChar("{") {
		// parsing a block of rules
		inBlock = true
		parser.embedLevel++

		parser.shiftToken()
	}

	for parser.tokenParsable() {
		if parser.tokenIgnorable() {
			parser.shiftToken()
		} else if parser.tokenChar("}") {
			if !inBlock {
				errMsg := fmt.Sprintf("Unexpected } character: %s", parser.nextToken().String())
				return result, errors.New(errMsg)
			}

			parser.shiftToken()
			parser.embedLevel--

			// finished
			break
		} else {
			rule, err := parser.ParseRule()
			if err != nil {
				return result, err
			}

			rule.EmbedLevel = parser.embedLevel
			result = append(result, rule)
		}
	}

	return result, parser.err()
}

// ParseRule parses a rule
func (parser *Parser) ParseRule() (*Rule, error) {
	if parser.tokenAtKeyword() {
		return parser.parseAtRule()
	}

	return parser.parseQualifiedRule()
}

// ParseDeclarations parses a list of declarations
func (parser *Parser) ParseDeclarations() ([]*Declaration, error) {
	result := []*Declaration{}

	if parser.tokenChar("{") {
		parser.shiftToken()
	}

	for parser.tokenParsable() {
		if parser.tokenIgnorable() {
			parser.shiftToken()
		} else if parser.tokenChar("}") {
			// end of block
			parser.shiftToken()
			break
		} else {
			declaration, err := parser.ParseDeclaration()
			if err != nil {
				return result, err
			}

			result = append(result, declaration)
		}
	}

	return result, parser.err()
}

// ParseDeclaration parses a declaration
func (parser *Parser) ParseDeclaration() (*Declaration, error) {
	result := NewDeclaration()
	curValue := ""

	for parser.tokenParsable() {
		if parser.tokenChar(":") {
			result.Property = strings.TrimSpace(curValue)
			curValue = ""

			parser.shiftToken()
		} else if parser.tokenChar(";") || parser.tokenChar("}") {
			if result.Property == "" {
				errMsg := fmt.Sprintf("Unexpected ; character: %s", parser.nextToken().String())
				return result, errors.New(errMsg)
			}

			if importantRegexp.MatchString(curValue) {
				result.Important = true
				curValue = importantRegexp.ReplaceAllString(curValue, "")
			}

			result.Value = strings.TrimSpace(curValue)

			if parser.tokenChar(";") {
				parser.shiftToken()
			}

			// finished
			break
		} else {
			token := parser.shiftToken()
			curValue += token.Value
		}
	}

	// log.Printf("[parsed] Declaration: %s", result.String())

	return result, parser.err()
}

// Parse an At Rule
func (parser *Parser) parseAtRule() (*Rule, error) {
	// parse rule name (eg: "@import")
	token := parser.shiftToken()

	result := NewRule(AtRule)
	result.Name = token.Value

	for parser.tokenParsable() {
		if parser.tokenChar(";") {
			parser.shiftToken()

			// finished
			break
		} else if parser.tokenChar("{") {
			if result.EmbedsRules() {
				// parse rules block
				rules, err := parser.ParseRules()
				if err != nil {
					return result, err
				}

				result.Rules = rules
			} else {
				// parse declarations block
				declarations, err := parser.ParseDeclarations()
				if err != nil {
					return result, err
				}

				result.Declarations = declarations
			}

			// finished
			break
		} else {
			// parse prelude
			prelude, err := parser.parsePrelude()
			if err != nil {
				return result, err
			}
			if result.Prelude != "" && prelude == "" {
				return result, fmt.Errorf("replacing a non-empty value %q with an empty one", result.Prelude)
			}
			result.Prelude = prelude
		}
	}

	// log.Printf("[parsed] Rule: %s", result.String())

	return result, parser.err()
}

// Parse a Qualified Rule
func (parser *Parser) parseQualifiedRule() (*Rule, error) {
	result := NewRule(QualifiedRule)

	for parser.tokenParsable() {
		if parser.tokenChar(";") {
			parser.shiftToken()
		} else if parser.tokenChar("{") {
			if result.Prelude == "" {
				errMsg := fmt.Sprintf("Unexpected { character: %s", parser.nextToken().String())
				return result, errors.New(errMsg)
			}

			// parse declarations block
			declarations, err := parser.ParseDeclarations()
			if err != nil {
				return result, err
			}

			result.Declarations = declarations

			// finished
			break
		} else {
			// parse prelude
			prelude, err := parser.parsePrelude()
			if err != nil {
				return result, err
			}
			if result.Prelude != "" && prelude == "" {
				return result, fmt.Errorf("replacing a non-empty value %q with an empty one", result.Prelude)
			}

			result.Prelude = prelude
		}
	}

	result.Selectors = strings.Split(result.Prelude, ",")
	for i, sel := range result.Selectors {
		result.Selectors[i] = strings.TrimSpace(sel)
	}

	// log.Printf("[parsed] Rule: %s", result.String())

	return result, parser.err()
}

// Parse Rule prelude
func (parser *Parser) parsePrelude() (string, error) {
	result := ""

	for parser.tokenParsable() && !parser.tokenEndOfPrelude() {
		token := parser.shiftToken()
		result += token.Value
	}

	result = strings.TrimSpace(result)

	if result == "" && parser.tokenChar(";") {
		parser.shiftToken()
	}

	// log.Printf("[parsed] prelude: %s", result)

	return result, parser.err()
}

// Parse BOM
func (parser *Parser) parseBOM() (bool, error) {
	if parser.nextToken().Type == scanner.TokenBOM {
		parser.shiftToken()
		return true, nil
	}

	return false, parser.err()
}

// Returns next token without removing it from tokens buffer
func (parser *Parser) nextToken() *scanner.Token {
	if len(parser.tokens) == 0 {
		// fetch next token
		nextToken := parser.scan.Next()

		// log.Printf("[token] %s => %v", nextToken.Type.String(), nextToken.Value)

		// queue it
		parser.tokens = append(parser.tokens, nextToken)
	}

	return parser.tokens[0]
}

// Returns next token and remove it from the tokens buffer
func (parser *Parser) shiftToken() *scanner.Token {
	var result *scanner.Token

	result, parser.tokens = parser.tokens[0], parser.tokens[1:]
	return result
}

// Returns tokenizer error, or nil if no error
func (parser *Parser) err() error {
	if parser.tokenError() {
		token := parser.nextToken()
		return fmt.Errorf("tokenizer error: %s", token.String())
	}

	return nil
}

// Returns true if next token is Error
func (parser *Parser) tokenError() bool {
	return parser.nextToken().Type == scanner.TokenError
}

// Returns true if next token is EOF
func (parser *Parser) tokenEOF() bool {
	return parser.nextToken().Type == scanner.TokenEOF
}

// Returns true if next token is a whitespace
func (parser *Parser) tokenWS() bool {
	return parser.nextToken().Type == scanner.TokenS
}

// Returns true if next token is a comment
func (parser *Parser) tokenComment() bool {
	return parser.nextToken().Type == scanner.TokenComment
}

// Returns true if next token is a CDO or a CDC
func (parser *Parser) tokenCDOorCDC() bool {
	switch parser.nextToken().Type {
	case scanner.TokenCDO, scanner.TokenCDC:
		return true
	default:
		return false
	}
}

// Returns true if next token is ignorable
func (parser *Parser) tokenIgnorable() bool {
	return parser.tokenWS() || parser.tokenComment() || parser.tokenCDOorCDC()
}

// Returns true if next token is parsable
func (parser *Parser) tokenParsable() bool {
	return !parser.tokenEOF() && !parser.tokenError()
}

// Returns true if next token is an At Rule keyword
func (parser *Parser) tokenAtKeyword() bool {
	return parser.nextToken().Type == scanner.TokenAtKeyword
}

// Returns true if next token is given character
func (parser *Parser) tokenChar(value string) bool {
	token := parser.nextToken()
	return (token.Type == scanner.TokenChar) && (token.Value == value)
}

// Returns true if next token marks the end of a prelude
func (parser *Parser) tokenEndOfPrelude() bool {
	return parser.tokenChar(";") || parser.tokenChar("{")
}
