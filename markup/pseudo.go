package markup

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/malivvan/cui/markup/atom"
)

// This file implements the pseudo classes selectors,
// which share the implementation of PseudoElement() and Specificity()

type abstractPseudoClass struct{}

func (apc abstractPseudoClass) Specificity() Specificity {
	return Specificity{0, 1, 0}
}

func (apc abstractPseudoClass) PseudoElement() string {
	return ""
}

type relativePseudoClassSelector struct {
	name  string // one of "not", "has", "haschild"
	match SelectorGroup
}

func (c relativePseudoClassSelector) Match(n *Node) bool {
	if n.Type != ElementNode {
		return false
	}
	switch c.name {
	case "not":
		// matches elements that do not match a.
		return !c.match.Match(n)
	case "has":
		//  matches elements with any descendant that matches a.
		return hasDescendantMatch(n, c.match)
	case "haschild":
		// matches elements with a child that matches a.
		return hasChildMatch(n, c.match)
	default:
		panic(fmt.Sprintf("unsupported relative pseudo class selector : %s", c.name))
	}
}

// hasChildMatch returns whether n has any child that matches a.
func hasChildMatch(n *Node, a Matcher) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if a.Match(c) {
			return true
		}
	}
	return false
}

// hasDescendantMatch performs a depth-first search of n's descendants,
// testing whether any of them match a. It returns true as soon as a match is
// found, or false if no match is found.
func hasDescendantMatch(n *Node, a Matcher) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if a.Match(c) || (c.Type == ElementNode && hasDescendantMatch(c, a)) {
			return true
		}
	}
	return false
}

// Specificity returns the specificity of the most specific selectors
// in the pseudo-class arguments.
// See https://www.w3.org/TR/selectors/#specificity-rules
func (c relativePseudoClassSelector) Specificity() Specificity {
	var max Specificity
	for _, sel := range c.match {
		newSpe := sel.Specificity()
		if max.Less(newSpe) {
			max = newSpe
		}
	}
	return max
}

func (c relativePseudoClassSelector) PseudoElement() string {
	return ""
}

type containsPseudoClassSelector struct {
	abstractPseudoClass
	value string
	own   bool
}

func (c containsPseudoClassSelector) Match(n *Node) bool {
	var text string
	if c.own {
		// matches nodes that directly contain the given text
		text = strings.ToLower(nodeOwnText(n))
	} else {
		// matches nodes that contain the given text.
		text = strings.ToLower(nodeText(n))
	}
	return strings.Contains(text, c.value)
}

type regexpPseudoClassSelector struct {
	abstractPseudoClass
	regexp *regexp.Regexp
	own    bool
}

func (c regexpPseudoClassSelector) Match(n *Node) bool {
	var text string
	if c.own {
		// matches nodes whose text directly matches the specified regular expression
		text = nodeOwnText(n)
	} else {
		// matches nodes whose text matches the specified regular expression
		text = nodeText(n)
	}
	return c.regexp.MatchString(text)
}

// writeNodeText writes the text contained in n and its descendants to b.
func writeNodeText(n *Node, b *bytes.Buffer) {
	switch n.Type {
	case TextNode:
		b.WriteString(n.Tag)
	case ElementNode:
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			writeNodeText(c, b)
		}
	}
}

// nodeText returns the text contained in n and its descendants.
func nodeText(n *Node) string {
	var b bytes.Buffer
	writeNodeText(n, &b)
	return b.String()
}

// nodeOwnText returns the contents of the text nodes that are direct
// children of n.
func nodeOwnText(n *Node) string {
	var b bytes.Buffer
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == TextNode {
			b.WriteString(c.Tag)
		}
	}
	return b.String()
}

type nthPseudoClassSelector struct {
	abstractPseudoClass
	a, b         int
	last, ofType bool
}

func (c nthPseudoClassSelector) Match(n *Node) bool {
	if c.a == 0 {
		if c.last {
			return simpleNthLastChildMatch(c.b, c.ofType, n)
		} else {
			return simpleNthChildMatch(c.b, c.ofType, n)
		}
	}
	return nthChildMatch(c.a, c.b, c.last, c.ofType, n)
}

// nthChildMatch implements :nth-child(an+b).
// If last is true, implements :nth-last-child instead.
// If ofType is true, implements :nth-of-type instead.
func nthChildMatch(a, b int, last, ofType bool, n *Node) bool {
	if n.Type != ElementNode {
		return false
	}

	parent := n.Parent
	if parent == nil {
		return false
	}

	i := -1
	count := 0
	for c := parent.FirstChild; c != nil; c = c.NextSibling {
		if (c.Type != ElementNode) || (ofType && c.Tag != n.Tag) {
			continue
		}
		count++
		if c == n {
			i = count
			if !last {
				break
			}
		}
	}

	if i == -1 {
		// This shouldn't happen, since n should always be one of its parent's children.
		return false
	}

	if last {
		i = count - i + 1
	}

	i -= b
	if a == 0 {
		return i == 0
	}

	return i%a == 0 && i/a >= 0
}

// simpleNthChildMatch implements :nth-child(b).
// If ofType is true, implements :nth-of-type instead.
func simpleNthChildMatch(b int, ofType bool, n *Node) bool {
	if n.Type != ElementNode {
		return false
	}

	parent := n.Parent
	if parent == nil {
		return false
	}

	count := 0
	for c := parent.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != ElementNode || (ofType && c.Tag != n.Tag) {
			continue
		}
		count++
		if c == n {
			return count == b
		}
		if count >= b {
			return false
		}
	}
	return false
}

// simpleNthLastChildMatch implements :nth-last-child(b).
// If ofType is true, implements :nth-last-of-type instead.
func simpleNthLastChildMatch(b int, ofType bool, n *Node) bool {
	if n.Type != ElementNode {
		return false
	}

	parent := n.Parent
	if parent == nil {
		return false
	}

	count := 0
	for c := parent.LastChild; c != nil; c = c.PrevSibling {
		if c.Type != ElementNode || (ofType && c.Tag != n.Tag) {
			continue
		}
		count++
		if c == n {
			return count == b
		}
		if count >= b {
			return false
		}
	}
	return false
}

type onlyChildPseudoClassSelector struct {
	abstractPseudoClass
	ofType bool
}

// Match implements :only-child.
// If `ofType` is true, it implements :only-of-type instead.
func (c onlyChildPseudoClassSelector) Match(n *Node) bool {
	if n.Type != ElementNode {
		return false
	}

	parent := n.Parent
	if parent == nil {
		return false
	}

	count := 0
	for node := parent.FirstChild; node != nil; node = node.NextSibling {
		if (node.Type != ElementNode) || (c.ofType && node.Tag != n.Tag) {
			continue
		}
		count++
		if count > 1 {
			return false
		}
	}

	return count == 1
}

type inputPseudoClassSelector struct {
	abstractPseudoClass
}

// Match matches input, select, textarea and button elements.
func (c inputPseudoClassSelector) Match(n *Node) bool {
	return n.Type == ElementNode && (n.Tag == "input" || n.Tag == "select" || n.Tag == "textarea" || n.Tag == "button")
}

type emptyElementPseudoClassSelector struct {
	abstractPseudoClass
}

// Match matches empty elements.
func (c emptyElementPseudoClassSelector) Match(n *Node) bool {
	if n.Type != ElementNode {
		return false
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case ElementNode:
			return false
		case TextNode:
			if strings.TrimSpace(nodeText(c)) == "" {
				continue
			} else {
				return false
			}
		}
	}

	return true
}

type rootPseudoClassSelector struct {
	abstractPseudoClass
}

// Match implements :root
func (c rootPseudoClassSelector) Match(n *Node) bool {
	if n.Type != ElementNode {
		return false
	}
	if n.Parent == nil {
		return false
	}
	return n.Parent.Type == DocumentNode
}

func hasAttr(n *Node, attr string) bool {
	return matchAttribute(n, attr, func(string) bool { return true })
}

type linkPseudoClassSelector struct {
	abstractPseudoClass
}

// Match implements :link
func (c linkPseudoClassSelector) Match(n *Node) bool {
	return (n.Atom == atom.A || n.Atom == atom.Area || n.Atom == atom.Link) && hasAttr(n, "href")
}

type langPseudoClassSelector struct {
	abstractPseudoClass
	lang string
}

func (c langPseudoClassSelector) Match(n *Node) bool {
	own := matchAttribute(n, "lang", func(val string) bool {
		return val == c.lang || strings.HasPrefix(val, c.lang+"-")
	})
	if n.Parent == nil {
		return own
	}
	return own || c.Match(n.Parent)
}

type enabledPseudoClassSelector struct {
	abstractPseudoClass
}

func (c enabledPseudoClassSelector) Match(n *Node) bool {
	if n.Type != ElementNode {
		return false
	}
	switch n.Atom {
	case atom.A, atom.Area, atom.Link:
		return hasAttr(n, "href")
	case atom.Optgroup, atom.Menuitem, atom.Fieldset:
		return !hasAttr(n, "disabled")
	case atom.Button, atom.Input, atom.Select, atom.Textarea, atom.Option:
		return !hasAttr(n, "disabled") && !inDisabledFieldset(n)
	}
	return false
}

type disabledPseudoClassSelector struct {
	abstractPseudoClass
}

func (c disabledPseudoClassSelector) Match(n *Node) bool {
	if n.Type != ElementNode {
		return false
	}
	switch n.Atom {
	case atom.Optgroup, atom.Menuitem, atom.Fieldset:
		return hasAttr(n, "disabled")
	case atom.Button, atom.Input, atom.Select, atom.Textarea, atom.Option:
		return hasAttr(n, "disabled") || inDisabledFieldset(n)
	}
	return false
}

func hasLegendInPreviousSiblings(n *Node) bool {
	for s := n.PrevSibling; s != nil; s = s.PrevSibling {
		if s.Atom == atom.Legend {
			return true
		}
	}
	return false
}

func inDisabledFieldset(n *Node) bool {
	if n.Parent == nil {
		return false
	}
	if n.Parent.Atom == atom.Fieldset && hasAttr(n.Parent, "disabled") &&
		(n.Atom != atom.Legend || hasLegendInPreviousSiblings(n)) {
		return true
	}
	return inDisabledFieldset(n.Parent)
}

type checkedPseudoClassSelector struct {
	abstractPseudoClass
}

func (c checkedPseudoClassSelector) Match(n *Node) bool {
	if n.Type != ElementNode {
		return false
	}
	switch n.Atom {
	case atom.Input, atom.Menuitem:
		return hasAttr(n, "checked") && matchAttribute(n, "type", func(val string) bool {
			t := toLowerASCII(val)
			return t == "checkbox" || t == "radio"
		})
	case atom.Option:
		return hasAttr(n, "selected")
	}
	return false
}
