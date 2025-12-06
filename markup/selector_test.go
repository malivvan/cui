package markup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

var validSelectors []validSelector

func init() {
	c, err := ioutil.ReadFile("testdata/valid_selectors.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(c, &validSelectors); err != nil {
		log.Fatal(err)
	}
}

type selectorTest struct {
	HTML, selector string
	results        []string
}

func nodeString(n *Node) string {
	buf := bytes.NewBufferString("")
	if err := Render(buf, n); err != nil {
		log.Fatal(err)
	}
	return buf.String()
}

var selectorTests = []selectorTest{
	{
		`<body><address>This address...</address></body>`,
		"address",
		[]string{
			"<address>This address...</address>",
		},
	},
	{
		`<!-- comment --><html><head></head><body>text</body></html>`,
		"*",
		[]string{
			"<html><head></head><body>text</body></html>",
			"<head></head>",
			"<body>text</body>",
		},
	},
	{
		`<html><head></head><body></body></html>`,
		"*",
		[]string{
			"<html><head></head><body></body></html>",
			"<head></head>",
			"<body></body>",
		},
	},
	{
		`<p id="foo"><p id="bar">`,
		"#foo",
		[]string{
			`<p id="foo"></p>`,
		},
	},
	{
		`<ul><li id="t1"><p id="t1">`,
		"li#t1",
		[]string{
			`<li id="t1"><p id="t1"></p></li>`,
		},
	},
	{
		`<ol><li id="t4"><li id="t44">`,
		"*#t4",
		[]string{
			`<li id="t4"></li>`,
		},
	},
	{
		`<ul><li class="t1"><li class="t2">`,
		".t1",
		[]string{
			`<li class="t1"></li>`,
		},
	},
	{
		`<p class="t1 t2">`,
		"p.t1",
		[]string{
			`<p class="t1 t2"></p>`,
		},
	},
	{
		`<div class="test">`,
		"div.teST",
		[]string{},
	},
	{
		`<p class="t1 t2">`,
		".t1.fail",
		[]string{},
	},
	{
		`<p class="t1 t2">`,
		"p.t1.t2",
		[]string{
			`<p class="t1 t2"></p>`,
		},
	},
	{
		`<p class="--t1 --t2">`,
		"p.--t1",
		[]string{
			`<p class="--t1 --t2"></p>`,
		},
	},
	{
		`<p class="--t1 --t2">`,
		"p.--t1.--t2",
		[]string{
			`<p class="--t1 --t2"></p>`,
		},
	},
	{
		`<p><p title="title">`,
		"p[title]",
		[]string{
			`<p title="title"></p>`,
		},
	},
	{
		`<div><div class="Red">`,
		`div[class="red" i]`,
		[]string{
			`<div class="Red"></div>`,
		},
	},
	{
		`<address><address title="foo"><address title="bar">`,
		`address[title="foo"]`,
		[]string{
			`<address title="foo"><address title="bar"></address></address>`,
		},
	},
	{
		`<address><address title="fooIgnoreCase"><address title="bar">`,
		`address[title="FoOIgnoRECaSe" i]`,
		[]string{
			`<address title="fooIgnoreCase"><address title="bar"></address></address>`,
		},
	},
	{
		`<address><address title="foo"><address title="bar">`,
		`address[title!="foo"]`,
		[]string{
			`<address><address title="foo"><address title="bar"></address></address></address>`,
			`<address title="bar"></address>`,
		},
	},
	{
		`<address><address title="FOO"><address title="bar">`,
		`address[title!="foo" i]`,
		[]string{
			`<address><address title="FOO"><address title="bar"></address></address></address>`,
			`<address title="bar"></address>`,
		},
	},
	{
		`<p title="fooBARuFOO"><p title="varfoo">`,
		`p[title!="FooBarUFoo" i]`,
		[]string{
			`<p title="varfoo"></p>`,
		},
	},
	{
		`<p title="tot foo bar">`,
		`[    	title        ~=       foo    ]`,
		[]string{
			`<p title="tot foo bar"></p>`,
		},
	},
	{
		`<p title="tot foo bar">`,
		`p[title~="FOO" i]`,
		[]string{
			`<p title="tot foo bar"></p>`,
		},
	},
	{
		`<p title="tot foo bar">`,
		`p[title~=toofoo i]`,
		[]string{},
	},
	{
		`<p title="hello world">`,
		`[title~="hello world"]`,
		[]string{},
	},
	{
		`<p title="HELLO world">`,
		`[title~="hello" i]`,
		[]string{
			`<p title="HELLO world"></p>`,
		},
	},
	{
		`<p title="HELLO world">`,
		`[title~="hello"          I]`,
		[]string{
			`<p title="HELLO world"></p>`,
		},
	},
	{
		`<p lang="en"><p lang="en-gb"><p lang="enough"><p lang="fr-en">`,
		`[lang|="en"]`,
		[]string{
			`<p lang="en"></p>`,
			`<p lang="en-gb"></p>`,
		},
	},
	{
		`<p lang="en"><p lang="En-gb"><p lang="enough"><p lang="fr-en">`,
		`[lang|="EN" i]`,
		[]string{
			`<p lang="en"></p>`,
			`<p lang="En-gb"></p>`,
		},
	},
	{
		`<p lang="en"><p lang="En-gb"><p lang="enough"><p lang="fr-en">`,
		`[lang|="EN"     i]`,
		[]string{
			`<p lang="en"></p>`,
			`<p lang="En-gb"></p>`,
		},
	},
	{
		`<p title="foobar"><p title="barfoo">`,
		`[title^="foo"]`,
		[]string{
			`<p title="foobar"></p>`,
		},
	},
	{
		`<p title="FooBAR"><p title="barfoo">`,
		`[title^="foo" i]`,
		[]string{
			`<p title="FooBAR"></p>`,
		},
	},
	{
		`<p title="foobar"><p title="barfoo">`,
		`[title$="bar"]`,
		[]string{
			`<p title="foobar"></p>`,
		},
	},
	{
		`<p title="foobar"><p title="barfoo">`,
		`[title$="BAR" i]`,
		[]string{
			`<p title="foobar"></p>`,
		},
	},
	{
		`<p title="foobarufoo">`,
		`[title*="bar"]`,
		[]string{
			`<p title="foobarufoo"></p>`,
		},
	},
	{
		`<p title="foobarufoo">`,
		`[title*="BaRu" i]`,
		[]string{
			`<p title="foobarufoo"></p>`,
		},
	},
	{
		`<p title="foobarufoo">`,
		`[title*="BaRu" I]`,
		[]string{
			`<p title="foobarufoo"></p>`,
		},
	},
	{
		`<p class=" ">This text should be green.</p><p>This text should be green.</p>`,
		`p[class$=" "]`,
		[]string{},
	},
	{
		`<p class="">This text should be green.</p><p>This text should be green.</p>`,
		`p[class$=""]`,
		[]string{},
	},
	{
		`<p class=" ">This text should be green.</p><p>This text should be green.</p>`,
		`p[class^=" "]`,
		[]string{},
	},
	{
		`<p class="">This text should be green.</p><p>This text should be green.</p>`,
		`p[class^=""]`,
		[]string{},
	},
	{
		`<p class=" ">This text should be green.</p><p>This text should be green.</p>`,
		`p[class*=" "]`,
		[]string{},
	},
	{
		`<p class="">This text should be green.</p><p>This text should be green.</p>`,
		`p[class*=""]`,
		[]string{},
	},
	{
		`<input type="radio" name="Sex" value="F"/>`,
		`input[name=Sex][value=F]`,
		[]string{
			`<input type="radio" name="Sex" value="F"/>`,
		},
	},
	{
		`<table border="0" cellpadding="0" cellspacing="0" style="table-layout: fixed; width: 100%; border: 0 dashed; border-color: #FFFFFF"><tr style="height:64px">aaa</tr></table>`,
		`table[border="0"][cellpadding="0"][cellspacing="0"]`,
		[]string{
			`<table border="0" cellpadding="0" cellspacing="0" style="table-layout: fixed; width: 100%; border: 0 dashed; border-color: #FFFFFF"><tbody><tr style="height:64px"></tr></tbody></table>`,
		},
	},
	{
		`<p class="t1 t2">`,
		".t1:not(.t2)",
		[]string{},
	},
	{
		`<div class="t3">`,
		`div:not(.t1)`,
		[]string{
			`<div class="t3"></div>`,
		},
	},
	{
		`<div><div class="t2"><div class="t3">`,
		`div:not([class="t2"])`,
		[]string{
			`<div><div class="t2"><div class="t3"></div></div></div>`,
			`<div class="t3"></div>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3></ol>`,
		`li:nth-child(odd)`,
		[]string{
			`<li id="1"></li>`,
			`<li id="3"></li>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3></ol>`,
		`li:nth-child(even)`,
		[]string{
			`<li id="2"></li>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3></ol>`,
		`li:nth-child(-n+2)`,
		[]string{
			`<li id="1"></li>`,
			`<li id="2"></li>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3></ol>`,
		`li:nth-child(3n+1)`,
		[]string{
			`<li id="1"></li>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3><li id=4></ol>`,
		`li:nth-last-child(odd)`,
		[]string{
			`<li id="2"></li>`,
			`<li id="4"></li>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3><li id=4></ol>`,
		`li:nth-last-child(even)`,
		[]string{
			`<li id="1"></li>`,
			`<li id="3"></li>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3><li id=4></ol>`,
		`li:nth-last-child(-n+2)`,
		[]string{
			`<li id="3"></li>`,
			`<li id="4"></li>`,
		},
	},
	{
		`<ol><li id=1><li id=2><li id=3><li id=4></ol>`,
		`li:nth-last-child(3n+1)`,
		[]string{
			`<li id="1"></li>`,
			`<li id="4"></li>`,
		},
	},
	{
		`<p>some text <span id="1">and a span</span><span id="2"> and another</span></p>`,
		`span:first-child`,
		[]string{
			`<span id="1">and a span</span>`,
		},
	},
	{
		`<span>a span</span> and some text`,
		`span:last-child`,
		[]string{
			`<span>a span</span>`,
		},
	},
	{
		`<address></address><p id=1><p id=2>`,
		`p:nth-of-type(2)`,
		[]string{
			`<p id="2"></p>`,
		},
	},
	{
		`<address></address><p id=1><p id=2></p><a>`,
		`p:nth-last-of-type(2)`,
		[]string{
			`<p id="1"></p>`,
		},
	},
	{
		`<address></address><p id=1><p id=2></p><a>`,
		`p:last-of-type`,
		[]string{
			`<p id="2"></p>`,
		},
	},
	{
		`<address></address><p id=1><p id=2></p><a>`,
		`p:first-of-type`,
		[]string{
			`<p id="1"></p>`,
		},
	},
	{
		`<div><p id="1"></p><a></a></div><div><p id="2"></p></div>`,
		`p:only-child`,
		[]string{
			`<p id="2"></p>`,
		},
	},
	{
		`<div><p id="1"></p><a></a></div><div><p id="2"></p><p id="3"></p></div>`,
		`p:only-of-type`,
		[]string{
			`<p id="1"></p>`,
		},
	},
	{
		`<p id="1"><!-- --><p id="2">Hello<p id="3"><span>`,
		`:empty`,
		[]string{
			`<head></head>`,
			`<p id="1"><!-- --></p>`,
			`<span></span>`,
		},
	},
	{
		`<div><p id="1"><table><tr><td><p id="2"></table></div><p id="3">`,
		`div p`,
		[]string{
			`<p id="1"><table><tbody><tr><td><p id="2"></p></td></tr></tbody></table></p>`,
			`<p id="2"></p>`,
		},
	},
	{
		`<div><p id="1"><table><tr><td><p id="2"></table></div><p id="3">`,
		`div table p`,
		[]string{
			`<p id="2"></p>`,
		},
	},
	{
		`<div><p id="1"><div><p id="2"></div><table><tr><td><p id="3"></table></div>`,
		`div > p`,
		[]string{
			`<p id="1"></p>`,
			`<p id="2"></p>`,
		},
	},
	{
		`<p id="1"><p id="2"></p><address></address><p id="3">`,
		`p ~ p`,
		[]string{
			`<p id="2"></p>`,
			`<p id="3"></p>`,
		},
	},
	{
		`<p id="1"></p>
		 <!--comment-->
		 <p id="2"></p><address></address><p id="3">`,
		`p + p`,
		[]string{
			`<p id="2"></p>`,
		},
	},
	{
		`<ul><li></li><li></li></ul><p>`,
		`li, p`,
		[]string{
			"<li></li>",
			"<li></li>",
			"<p></p>",
		},
	},
	{
		`<p id="1"><p id="2"></p><address></address><p id="3">`,
		`p +/*This is a comment*/ p`,
		[]string{
			`<p id="2"></p>`,
		},
	},
	{
		`<p>Text block that <span>wraps inner text</span> and continues</p>`,
		`p:contains("that wraps")`,
		[]string{
			`<p>Text block that <span>wraps inner text</span> and continues</p>`,
		},
	},
	{
		`<p>Text block that <span>wraps inner text</span> and continues</p>`,
		`p:containsOwn("that wraps")`,
		[]string{},
	},
	{
		`<p>Text block that <span>wraps inner text</span> and continues</p>`,
		`:containsOwn("inner")`,
		[]string{
			`<span>wraps inner text</span>`,
		},
	},
	{
		`<p>Text block that <span>wraps inner text</span> and continues</p>`,
		`p:containsOwn("block")`,
		[]string{
			`<p>Text block that <span>wraps inner text</span> and continues</p>`,
		},
	},
	{
		`<div id="d1"><p id="p1"><span>text content</span></p></div><div id="d2"/>`,
		`div:has(#p1)`,
		[]string{
			`<div id="d1"><p id="p1"><span>text content</span></p></div>`,
		},
	},
	{
		`<div id="d1"><p id="p1"><span>contents 1</span></p></div>
		<div id="d2"><p>contents <em>2</em></p></div>`,
		`div:has(:containsOwn("2"))`,
		[]string{
			`<div id="d2"><p>contents <em>2</em></p></div>`,
		},
	},
	{
		`<body><div id="d1"><p id="p1"><span>contents 1</span></p></div>
		<div id="d2"><p id="p2">contents <em>2</em></p></div></body>`,
		`body :has(:containsOwn("2"))`,
		[]string{
			`<div id="d2"><p id="p2">contents <em>2</em></p></div>`,
			`<p id="p2">contents <em>2</em></p>`,
		},
	},
	{
		`<body><div id="d1"><p id="p1"><span>contents 1</span></p></div>
		<div id="d2"><p id="p2">contents <em>2</em></p></div></body>`,
		`body :haschild(:containsOwn("2"))`,
		[]string{
			`<p id="p2">contents <em>2</em></p>`,
		},
	},
	{
		`<p id="p1">0123456789</p><p id="p2">abcdef</p><p id="p3">0123ABCD</p>`,
		`p:matches([\d])`,
		[]string{
			`<p id="p1">0123456789</p>`,
			`<p id="p3">0123ABCD</p>`,
		},
	},
	{
		`<p id="p1">0123456789</p><p id="p2">abcdef</p><p id="p3">0123ABCD</p>`,
		`p:matches([a-z])`,
		[]string{
			`<p id="p2">abcdef</p>`,
		},
	},
	{
		`<p id="p1">0123456789</p><p id="p2">abcdef</p><p id="p3">0123ABCD</p>`,
		`p:matches([a-zA-Z])`,
		[]string{
			`<p id="p2">abcdef</p>`,
			`<p id="p3">0123ABCD</p>`,
		},
	},
	{
		`<p id="p1">0123456789</p><p id="p2">abcdef</p><p id="p3">0123ABCD</p>`,
		`p:matches([^\d])`,
		[]string{
			`<p id="p2">abcdef</p>`,
			`<p id="p3">0123ABCD</p>`,
		},
	},
	{
		`<p id="p1">0123456789</p><p id="p2">abcdef</p><p id="p3">0123ABCD</p>`,
		`p:matches(^(0|a))`,
		[]string{
			`<p id="p1">0123456789</p>`,
			`<p id="p2">abcdef</p>`,
			`<p id="p3">0123ABCD</p>`,
		},
	},
	{
		`<p id="p1">0123456789</p><p id="p2">abcdef</p><p id="p3">0123ABCD</p>`,
		`p:matches(^\d+$)`,
		[]string{
			`<p id="p1">0123456789</p>`,
		},
	},
	{
		`<p id="p1">0123456789</p><p id="p2">abcdef</p><p id="p3">0123ABCD</p>`,
		`p:not(:matches(^\d+$))`,
		[]string{
			`<p id="p2">abcdef</p>`,
			`<p id="p3">0123ABCD</p>`,
		},
	},
	{
		`<div><p id="p1">01234<em>567</em>89</p><div>`,
		`div :matchesOwn(^\d+$)`,
		[]string{
			`<p id="p1">01234<em>567</em>89</p>`,
			`<em>567</em>`,
		},
	},
	{
		`<ul>
			<li><a id="a1" href="http://www.google.com/finance"></a>
			<li><a id="a2" href="http://finance.yahoo.com/"></a>
			<li><a id="a2" href="http://finance.untrusted.com/"/>
			<li><a id="a3" href="https://www.google.com/news"/>
			<li><a id="a4" href="http://news.yahoo.com"/>
		</ul>`,
		`[href#=(fina)]:not([href#=(\/\/[^\/]+untrusted)])`,
		[]string{
			`<a id="a1" href="http://www.google.com/finance"></a>`,
			`<a id="a2" href="http://finance.yahoo.com/"></a>`,
		},
	},
	{
		`<ul>
			<li><a id="a1" href="http://www.google.com/finance"/>
			<li><a id="a2" href="http://finance.yahoo.com/"/>
			<li><a id="a3" href="https://www.google.com/news"></a>
			<li><a id="a4" href="http://news.yahoo.com"/>
		</ul>`,
		`[href#=(^https:\/\/[^\/]*\/?news)]`,
		[]string{
			`<a id="a3" href="https://www.google.com/news"></a>`,
		},
	},
	{
		`<form>
			<label>Username <input type="text" name="username" /></label>
			<label>Password <input type="password" name="password" /></label>
			<label>Country
				<select name="country">
					<option value="ca">Canada</option>
					<option value="us">United States</option>
				</select>
			</label>
			<label>Bio <textarea name="bio"></textarea></label>
			<button>Sign up</button>
		</form>`,
		`:input`,
		[]string{
			`<input type="text" name="username"/>`,
			`<input type="password" name="password"/>`,
			`<select name="country">
					<option value="ca">Canada</option>
					<option value="us">United States</option>
				</select>`,
			`<textarea name="bio"></textarea>`,
			`<button>Sign up</button>`,
		},
	},
	{
		`<html><head></head><body></body></html>`,
		":root",
		[]string{
			"<html><head></head><body></body></html>",
		},
	},
	{
		`<html><head></head><body></body></html>`,
		"*:root",
		[]string{
			"<html><head></head><body></body></html>",
		},
	},
	{
		`<html><head></head><body></body></html>`,
		"html:nth-child(1)",
		[]string{
			"<html><head></head><body></body></html>",
		},
	},
	{
		`<html><head></head><body></body></html>`,
		"*:root:first-child",
		[]string{
			`<html><head></head><body></body></html>`,
		},
	},
	{
		`<html><head></head><body></body></html>`,
		"*:root:nth-child(1)",
		[]string{
			`<html><head></head><body></body></html>`,
		},
	},
	{
		`<html><head></head><body><a href="http://www.foo.com"></a></body></html>`,
		"a:not(:root)",
		[]string{
			`<a href="http://www.foo.com"></a>`,
		},
	},
	{
		`<html><head></head><body><p></p><div></div><span></span><a></a><form></form></body></html>`,
		"body > *:nth-child(3n+2)",
		[]string{
			"<div></div>",
			"<form></form>",
		},
	},
	{
		`<html><head></head><body><fieldset disabled><legend id="1"><input id="i1"/></legend><legend id="2"><input id="i2"/></legend></fieldset></body></html>`,
		"input:disabled",
		[]string{
			`<input id="i2"/>`,
		},
	},
	{
		`<html><head></head><body><fieldset disabled></fieldset></body></html>`,
		":disabled",
		[]string{
			`<fieldset disabled=""></fieldset>`,
		},
	},
	{
		`<html><head></head><body><fieldset></fieldset></body></html>`,
		":enabled",
		[]string{
			`<fieldset></fieldset>`,
		},
	},
	{
		`<div class=class1></div><div class=class2></div><div class=class3></div>`,
		"div.class1, div.class2",
		[]string{
			`<div class="class1"></div>`,
			`<div class="class2"></div>`,
		},
	},
}

func setup(selector, testHTML string) (Selector, *Node, error) {
	s, err := CompileSelector(selector)
	if err != nil {
		return nil, nil, fmt.Errorf("error compiling %q: %s", selector, err)
	}

	doc, err := Parse(strings.NewReader(testHTML))
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing %q: %s", testHTML, err)
	}
	return s, doc, nil
}

func TestSelectors(t *testing.T) {
	for _, test := range selectorTests {
		s, doc, err := setup(test.selector, test.HTML)
		if err != nil {
			t.Error(err)
			continue
		}

		matches := s.MatchAll(doc)
		if len(matches) != len(test.results) {
			t.Errorf("selector %s wanted %d elements, got %d instead", test.selector, len(test.results), len(matches))
			continue
		}

		for i, m := range matches {
			got := nodeString(m)
			if got != test.results[i] {
				t.Errorf("selector %s wanted %s, got %s instead", test.selector, test.results[i], got)
			}
		}

		firstMatch := s.MatchFirst(doc)
		if len(test.results) == 0 {
			if firstMatch != nil {
				t.Errorf("MatchFirst: selector %s want nil, got %s", test.selector, nodeString(firstMatch))
			}
		} else {
			got := nodeString(firstMatch)
			if got != test.results[0] {
				t.Errorf("MatchFirst: selector %s want %s, got %s", test.selector, test.results[0], got)
			}
		}
	}
}

func setupMatcher(selector, testHTML string) (Matcher, *Node, error) {
	s, err := ParseSelectorGroup(selector)
	if err != nil {
		return nil, nil, fmt.Errorf("error compiling %q: %s", selector, err)
	}

	doc, err := Parse(strings.NewReader(testHTML))
	if err != nil {
		return nil, nil, fmt.Errorf("error parsing %q: %s", testHTML, err)
	}
	return s, doc, nil
}

func TestMatchers(t *testing.T) {
	for _, test := range selectorTests {
		s, doc, err := setupMatcher(test.selector, test.HTML)
		if err != nil {
			t.Error(err)
			continue
		}

		matches := QueryAll(doc, s)
		if len(matches) != len(test.results) {
			t.Errorf("selector %s wanted %d elements, got %d instead", test.selector, len(test.results), len(matches))
			continue
		}

		for i, m := range matches {
			got := nodeString(m)
			if got != test.results[i] {
				t.Errorf("selector %s wanted %s, got %s instead", test.selector, test.results[i], got)
			}
		}

		firstMatch := Query(doc, s)
		if len(test.results) == 0 {
			if firstMatch != nil {
				t.Errorf("Query: selector %s want nil, got %s", test.selector, nodeString(firstMatch))
			}
		} else {
			got := nodeString(firstMatch)
			if got != test.results[0] {
				t.Errorf("Query: selector %s want %s, got %s", test.selector, test.results[0], got)
			}
		}

		if !reflect.DeepEqual(matches, Selector(s.Match).Filter(matches)) {
			t.Fatalf("inconsistent Filter result")
		}
	}
}

type testPseudo struct {
	HTML, selector string
	spec           Specificity
	pseudo         string
}

var testsPseudo = []testPseudo{
	{
		HTML:     `<html><body><ul><ol><li id="s12" class="red level"></li></ol></ul></body></html>`,
		selector: "#s12:not(FOO)::before",
		spec:     Specificity{1, 0, 2},
		pseudo:   "before",
	},
	{
		HTML:     `<html><body><ul><ol><li id="s12" class="red level"></li></ol></ul></body></html>`,
		selector: "#s12::first-line",
		spec:     Specificity{1, 0, 1},
		pseudo:   "first-line",
	},
	{
		HTML:     `<html><body><ul><ol><li id="s12" class="red level"></li></ol></ul></body></html>`,
		selector: "ol > #s12:first-line",
		spec:     Specificity{1, 0, 2},
		pseudo:   "first-line",
	},
	{
		HTML:     `<html><body><ul><ol><li id="s12" class="red level"></li></ol></ul></body></html>`,
		selector: "#s12:not(FOO)::after",
		spec:     Specificity{1, 0, 2},
		pseudo:   "after",
	},
	{
		HTML:     `<html><body><ul><ol><li id="s12" class="red level"></li></ol></ul></body></html>`,
		selector: "LI.red.level:before",
		spec:     Specificity{0, 2, 2},
		pseudo:   "before",
	},
}

func TestPseudoElement(t *testing.T) {
	for _, test := range testsPseudo {
		s, err := ParseSelectorWithPseudoElement(test.selector)
		if err != nil {
			t.Fatalf("error compiling %q: %s", test.selector, err)
		}

		if _, err = ParseSelector(test.selector); err == nil {
			t.Fatalf("selector %s with pseudo-element should not compile", test.selector)
		}

		doc, err := Parse(strings.NewReader(test.HTML))
		if err != nil {
			t.Fatalf("error parsing %q: %s", test.HTML, err)
		}

		body := doc.FirstChild.LastChild
		testNode := body.FirstChild.FirstChild.LastChild
		if !s.Match(testNode) {
			t.Errorf("%s didn't match (html tree : \n %s) \n", test.selector, nodeString(doc))
			continue
		}
		if s.Specificity() != test.spec {
			t.Errorf("wrong specificity : expected %v got %v", test.spec, s.Specificity())
		}
		if s.PseudoElement() != test.pseudo {
			t.Errorf("wrong pseudo-element : expected %s got %s", test.pseudo, s.PseudoElement())
		}
	}
}

type invalidSelector struct {
	Name     string `json:"name,omitempty"`
	Selector string `json:"selector,omitempty"`
}

type validSelector struct {
	invalidSelector
	Expect  []string `json:"expect,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
	Level   int      `json:"level,omitempty"`
	Xfail   bool     `json:"xfail,omitempty"`
}

func TestShakespeare(t *testing.T) {
	doc := parseReference("testdata/shakespeare.html")
	body := doc.FirstChild.NextSibling.LastChild
	assertCount := func(selector string, expected int) {
		sel, err := ParseSelectorGroup(selector)
		if err != nil {
			t.Errorf("invalid selector %s", selector)
		}
		if l := len(Selector(sel.Match).MatchAll(body)); l != expected {
			t.Errorf("%s -> expected %d, got %d", selector, expected, l)
		}
	}

	// Tag borrowed from https://github.com/Kozea/cssselect2
	assertCount("*", 246)
	assertCount("div:only-child", 22) // ?
	assertCount("div:nth-child(even)", 106)
	assertCount("div:nth-child(2n)", 106)
	assertCount("div:nth-child(odd)", 137)
	assertCount("div:nth-child(2n+1)", 137)
	assertCount("div:nth-child(n)", 243)
	assertCount("div:last-child", 53)
	assertCount("div:first-child", 51)
	assertCount("div > div", 242)
	assertCount("div + div", 190)
	assertCount("div ~ div", 190)
	assertCount("body", 1)
	assertCount("body div", 243)
	assertCount("div", 243)
	assertCount("div div", 242)
	assertCount("div div div", 241)
	assertCount("div, div, div", 243)
	assertCount("div, a, span", 243)
	assertCount(".dialog", 51)
	assertCount("div.dialog", 51)
	assertCount("div .dialog", 51)
	assertCount("div.character, div.dialog", 99)
	assertCount("div.direction.dialog", 0)
	assertCount("div.dialog.direction", 0)
	assertCount("div.dialog.scene", 1)
	assertCount("div.scene.scene", 1)
	assertCount("div.scene .scene", 0)
	assertCount("div.direction .dialog ", 0)
	assertCount("div .dialog .direction", 4)
	assertCount("div.dialog .dialog .direction", 4)
	assertCount("#speech5", 1)
	assertCount("div#speech5", 1)
	assertCount("div #speech5", 1)
	assertCount("div.scene div.dialog", 49)
	assertCount("div#scene1 div.dialog div", 142)
	assertCount("#scene1 #speech1", 1)
	assertCount("div[class]", 103)
	assertCount("div[class=dialog]", 50)
	assertCount("div[class^=dia]", 51)
	assertCount("div[class$=log]", 50)
	assertCount("div[class*=sce]", 1)
	assertCount("div[class|=dialog]", 50)
	assertCount("div[class~=dialog]", 51)
}

var identifierTests = map[string]string{
	"x":             "x",
	"96":            "",
	"-x":            "-x",
	`r\e9 sumé`:     "résumé",
	`r\0000e9 sumé`: "résumé",
	`r\0000e9sumé`:  "résumé",
	`a\"b`:          `a"b`,
}

func TestParseIdentifier(t *testing.T) {
	for source, want := range identifierTests {
		p := &selectorParser{s: source}
		got, err := p.parseIdentifier()
		if err != nil {
			if want == "" {
				// It was supposed to be an error.
				continue
			}
			t.Errorf("parsing %q: got error (%s), want %q", source, err, want)
			continue
		}

		if want == "" {
			if err == nil {
				t.Errorf("parsing %q: got %q, want error", source, got)
			}
			continue
		}

		if p.i < len(source) {
			t.Errorf("parsing %q: %d bytes left over", source, len(source)-p.i)
			continue
		}

		if got != want {
			t.Errorf("parsing %q: got %q, want %q", source, got, want)
		}
	}
}

var stringTests = map[string]string{
	`"x"`:             "x",
	`'x'`:             "x",
	`'x`:              "",
	"'x\\\r\nx'":      "xx",
	`"r\e9 sumé"`:     "résumé",
	`"r\0000e9 sumé"`: "résumé",
	`"r\0000e9sumé"`:  "résumé",
	`"a\"b"`:          `a"b`,
}

func TestParseString(t *testing.T) {
	for source, want := range stringTests {
		p := &selectorParser{s: source}
		got, err := p.parseString()
		if err != nil {
			if want == "" {
				// It was supposed to be an error.
				continue
			}
			t.Errorf("parsing %q: got error (%s), want %q", source, err, want)
			continue
		}

		if want == "" {
			if err == nil {
				t.Errorf("parsing %q: got %q, want error", source, got)
			}
			continue
		}

		if p.i < len(source) {
			t.Errorf("parsing %q: %d bytes left over", source, len(source)-p.i)
			continue
		}

		if got != want {
			t.Errorf("parsing %q: got %q, want %q", source, got, want)
		}
	}
}

func MustParseHTML(doc string) *Node {
	dom, err := Parse(strings.NewReader(doc))
	if err != nil {
		panic(err)
	}
	return dom
}

var selector = MustCompileSelector(`div.matched`)
var doc = `<!DOCTYPE html>
<html>
<body>
<div class="matched">
  <div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
    <div class="matched"></div>
  </div>
</div>
</body>
</html>
`
var dom = MustParseHTML(doc)

func BenchmarkMatchAll(b *testing.B) {
	var matches []*Node
	for i := 0; i < b.N; i++ {
		matches = selector.MatchAll(dom)
	}
	_ = matches
}

// TestComments exhaustively tests every 'interesting' N-byte string is
// correctly parsed as a comment. N ranges from 4+1 to 4+maxSuffixLen
// inclusive. 4 is the length of the "<!--" prefix that starts an HTML comment.
//
// 'Interesting' means that the N-4 byte suffix consists entirely of bytes
// sampled from the interestingCommentBytes const string, below. These cover
// all of the possible state transitions from comment-related parser states, as
// listed in the HTML spec (https://spec.whatwg.org/#comment-start-state
// and subsequent sections).
//
// The spec is written as an explicit state machine that, as a side effect,
// accumulates "the comment token's data" to a separate buffer.
// Tokenizer.readComment in this package does not have an explicit state
// machine and usually returns the comment text as a sub-slice of the input,
// between the opening '<' and closing '>' or EOF. This test confirms that the
// two algorithms match.
func TestComments(t *testing.T) {
	const prefix = "<!--"
	const maxSuffixLen = 6
	buffer := make([]byte, 0, len(prefix)+maxSuffixLen)
	testAllComments(t, append(buffer, prefix...))
}

// NUL isn't in this list, even though the HTML spec sections 13.2.5.43 -
// 13.2.5.52 mentions it. It's not interesting in terms of state transitions.
// It's equivalent to any other non-interesting byte (other than being replaced
// by U+FFFD REPLACEMENT CHARACTER).
//
// EOF isn't in this list. The HTML spec treats EOF as "an input character" but
// testOneComment below breaks the loop instead.
//
// 'x' represents all other "non-interesting" comment bytes.
var interestingCommentBytes = [...]byte{
	'!', '-', '<', '>', 'x',
}

// testAllComments recursively fills in buffer[len(buffer):cap(buffer)] with
// interesting bytes and then tests that this package's tokenization matches
// the HTML spec.
//
// Precondition: len(buffer) < cap(buffer)
// Precondition: string(buffer[:4]) == "<!--"
func testAllComments(t *testing.T, buffer []byte) {
	for _, interesting := range interestingCommentBytes {
		b := append(buffer, interesting)
		testOneComment(t, b)
		if len(b) < cap(b) {
			testAllComments(t, b)
		}
	}
}

func testOneComment(t *testing.T, b []byte) {
	z := NewTokenizer(bytes.NewReader(b))
	if next := z.Next(); next != CommentToken {
		t.Fatalf("Next(%q): got %v, want %v", b, next, CommentToken)
	}
	gotRemainder := string(b[len(z.Raw()):])
	gotComment := string(z.Text())

	i := len("<!--")
	wantBuffer := []byte(nil)
loop:
	for state := 43; ; {
		// Consume the next input character, handling EOF.
		if i >= len(b) {
			break
		}
		nextInputCharacter := b[i]
		i++

		switch state {
		case 43: // 13.2.5.43 Comment start state.
			switch nextInputCharacter {
			case '-':
				state = 44
			case '>':
				break loop
			default:
				i-- // Reconsume.
				state = 45
			}

		case 44: // 13.2.5.44 Comment start dash state.
			switch nextInputCharacter {
			case '-':
				state = 51
			case '>':
				break loop
			default:
				wantBuffer = append(wantBuffer, '-')
				i-- // Reconsume.
				state = 45
			}

		case 45: // 13.2.5.45 Comment state.
			switch nextInputCharacter {
			case '-':
				state = 50
			case '<':
				wantBuffer = append(wantBuffer, '<')
				state = 46
			default:
				wantBuffer = append(wantBuffer, nextInputCharacter)
			}

		case 46: // 13.2.5.46 Comment less-than sign state.
			switch nextInputCharacter {
			case '!':
				wantBuffer = append(wantBuffer, '!')
				state = 47
			case '<':
				wantBuffer = append(wantBuffer, '<')
				state = 46
			default:
				i-- // Reconsume.
				state = 45
			}

		case 47: // 13.2.5.47 Comment less-than sign bang state.
			switch nextInputCharacter {
			case '-':
				state = 48
			default:
				i-- // Reconsume.
				state = 45
			}

		case 48: // 13.2.5.48 Comment less-than sign bang dash state.
			switch nextInputCharacter {
			case '-':
				state = 49
			default:
				i-- // Reconsume.
				state = 50
			}

		case 49: // 13.2.5.49 Comment less-than sign bang dash dash state.
			switch nextInputCharacter {
			case '>':
				break loop
			default:
				i-- // Reconsume.
				state = 51
			}

		case 50: // 13.2.5.50 Comment end dash state.
			switch nextInputCharacter {
			case '-':
				state = 51
			default:
				wantBuffer = append(wantBuffer, '-')
				i-- // Reconsume.
				state = 45
			}

		case 51: // 13.2.5.51 Comment end state.
			switch nextInputCharacter {
			case '!':
				state = 52
			case '-':
				wantBuffer = append(wantBuffer, '-')
			case '>':
				break loop
			default:
				wantBuffer = append(wantBuffer, "--"...)
				i-- // Reconsume.
				state = 45
			}

		case 52: // 13.2.5.52 Comment end bang state.
			switch nextInputCharacter {
			case '-':
				wantBuffer = append(wantBuffer, "--!"...)
				state = 50
			case '>':
				break loop
			default:
				wantBuffer = append(wantBuffer, "--!"...)
				i-- // Reconsume.
				state = 45
			}

		default:
			t.Fatalf("input=%q: unexpected state %d", b, state)
		}
	}

	wantRemainder := ""
	if i < len(b) {
		wantRemainder = string(b[i:])
	}
	wantComment := string(wantBuffer)
	if (gotComment != wantComment) || (gotRemainder != wantRemainder) {
		t.Errorf("input=%q\ngot:  %q + %q\nwant: %q + %q",
			b, gotComment, gotRemainder, wantComment, wantRemainder)
		return
	}

	// suffix is the "N-4 byte suffix" per the TestComments comment.
	suffix := string(b[4:])

	// Test that a round trip, rendering (escaped) and re-parsing, of a comment
	// token (with that suffix as the Token.Tag) preserves that string.
	tok := Token{
		Type: CommentToken,
		Data: suffix,
	}
	z2 := NewTokenizer(strings.NewReader(tok.String()))
	if next := z2.Next(); next != CommentToken {
		t.Fatalf("round-trip Next(%q): got %v, want %v", suffix, next, CommentToken)
	}
	gotComment2 := string(z2.Text())
	if gotComment2 != suffix {
		t.Errorf("round-trip\ngot:  %q\nwant: %q", gotComment2, suffix)
		return
	}
}

// This table below summarizes the HTML-comment-related state machine from
// 13.2.5.43 "Comment start state" and subsequent sections.
// https://spec.whatwg.org/#comment-start-state
//
// Get to state 13.2.5.43 after seeing "<!--". Specifically, starting from the
// initial 13.2.5.1 "Tag state":
//   - "<"  moves to 13.2.5.6  "Tag open state",
//   - "!"  moves to 13.2.5.42 "Markup declaration open state",
//   - "--" moves to 13.2.5.43 "Comment start state".
// Each of these transitions are the only way to get to the 6/42/43 states.
//
// State   !         -         <         >         NUL       EOF       default   HTML spec section
// 43      ...       s44       ...       s01.T.E0  ...       ...       r45       13.2.5.43 Comment start state
// 44      ...       s51       ...       s01.T.E0  ...       T.Z.E1    r45.A-    13.2.5.44 Comment start dash state
// 45      ...       s50       s46.A<    ...       t45.A?.E2 T.Z.E1    t45.Ax    13.2.5.45 Comment state
// 46      s47.A!    ...       t46.A<    ...       ...       ...       r45       13.2.5.46 Comment less-than sign state
// 47      ...       s48       ...       ...       ...       ...       r45       13.2.5.47 Comment less-than sign bang state
// 48      ...       s49       ...       ...       ...       ...       r50       13.2.5.48 Comment less-than sign bang dash state
// 49      ...       ...       ...       s01.T     ...       T.Z.E1    r51.E3    13.2.5.49 Comment less-than sign bang dash dash state
// 50      ...       s51       ...       ...       ...       T.Z.E1    r45.A-    13.2.5.50 Comment end dash state
// 51      s52       t51.A-    ...       s01.T     ...       T.Z.E1    r45.A--   13.2.5.51 Comment end state
// 52      ...       s50.A--!  ...       s01.T.E4  ...       T.Z.E1    r45.A--!  13.2.5.52 Comment end bang state
//
// State 43 is the "Comment start state" meaning that we've only seen "<!--"
// and nothing else. Similarly, state 44 means that we've only seen "<!---",
// with three dashes, and nothing else. For the other states, we deduce
// (working backwards) that the immediate prior input must be:
//   - 45  something that's not '-'
//   - 46  "<"
//   - 47  "<!"
//   - 48  "<!-"
//   - 49  "<!--"  not including the opening "<!--"
//   - 50  "-"     not including the opening "<!--" and also not "--"
//   - 51  "--"    not including the opening "<!--"
//   - 52  "--!"
//
// The table cell actions:
//   - ...   do the default action
//   - A!    append "!"      to the comment token's data.
//   - A-    append "-"      to the comment token's data.
//   - A--   append "--"     to the comment token's data.
//   - A--!  append "--!"    to the comment token's data.
//   - A<    append "<"      to the comment token's data.
//   - A?    append "\uFFFD" to the comment token's data.
//   - Ax    append the current input character to the comment token's data.
//   - E0    parse error (abrupt-closing-of-empty-comment).
//   - E1    parse error (eof-in-comment).
//   - E2    parse error (unexpected-null-character).
//   - E3    parse error (nested-comment).
//   - E4    parse error (incorrectly-closed-comment).
//   - T     emit the current comment token.
//   - Z     emit an end-of-file token.
//   - rNN   reconsume in the 13.2.5.NN     state (after any A* or E* operations).
//   - s01   switch to the    13.2.5.1 Tag state (after any A* or E* operations).
//   - sNN   switch to the    13.2.5.NN     state (after any A* or E* operations).
//   - tNN   stay in the      13.2.5.NN     state (after any A* or E* operations).
//
// The E* actions are called errors in the HTML spec but they are not fatal
// (https://spec.whatwg.org/#parse-errors says "may [but not must] abort
// the parser"). They are warnings that, in practice, browsers simply ignore.

func TestInvalidSelectors(t *testing.T) {
	c, err := os.ReadFile("testdata/invalid_selectors.json")
	if err != nil {
		t.Fatal(err)
	}
	var tests []invalidSelector
	if err = json.Unmarshal(c, &tests); err != nil {
		t.Fatal(err)
	}
	for _, test := range tests {
		_, err := ParseSelectorGroupWithPseudoElements(test.Selector)
		if err == nil {
			t.Fatalf("%s -> expected error on invalid selector : %s", test.Name, test.Selector)
		}
	}
}

func parseReference(filename string) *Node {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	node, err := Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	return node
}

func getId(n *Node) string {
	for _, attr := range n.Attrs {
		if attr.Key == "id" {
			return attr.Val
		}
	}
	return ""
}

func isEqual(m map[string]int, l []string) bool {
	expected := map[string]int{}
	for _, s := range l {
		expected[s]++
	}
	return reflect.DeepEqual(m, expected)
}

func loadValidSelectors(t *testing.T) []validSelector {
	c, err := os.ReadFile("testdata/valid_selectors.json")
	if err != nil {
		t.Fatal(err)
	}
	var tests []validSelector
	if err = json.Unmarshal(c, &tests); err != nil {
		t.Fatal(err)
	}
	return tests
}

func TestValidSelectors(t *testing.T) {
	tests := loadValidSelectors(t)
	doc := parseReference("testdata/content.xhtml")
	for i, test := range tests {
		if test.Xfail {
			t.Logf("skiped test %s", test.Name)
			continue
		}
		sels, err := ParseSelectorGroupWithPseudoElements(test.Selector)
		if err != nil {
			t.Fatalf("%s -> unable to parse valid selector : %s : %s", test.Name, test.Selector, err)
		}
		matchingNodes := map[*Node]bool{}
		for _, sel := range sels {
			if sel.PseudoElement() != "" {
				continue // pseudo-element doesn't count as a match in this test since they are not part of the document
			}
			for _, node := range Selector(sel.Match).MatchAll(doc) {
				matchingNodes[node] = true
			}
		}
		matchingIds := map[string]int{}
		for node := range matchingNodes {
			matchingIds[getId(node)]++
		}
		if !isEqual(matchingIds, test.Expect) {
			t.Fatalf("test %d %s : expected %v got %v", i, test.Name, test.Expect, matchingIds)
		}
	}
}
