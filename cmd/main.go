package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
)

func main() {
	s := `<?xml version="1.0" encoding="UTF-8" ?>
<view id="myview">
<channel>
  <title>W3Schools Home Page</title>
  <link>https://www.w3schools.com</link>
  <description>Free web building tutorials</description>
  <item>
    <title>RSS Tutorial</title>
    <link>https://www.w3schools.com/xml/xml_rss.asp</link>
    <description>New RSS tutorial on W3Schools</description>
  </item>
  <item>
    <title>XML Tutorial</title>
    <link>https://www.w3schools.com/xml</link>
    <description>New XML tutorial on W3Schools</description>
  </item>
</channel>
</view>`

	doc, err := xmlquery.Parse(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	for _, n := range doc.ChildNodes() {
		if n.Data == "view" {
			fmt.Println("view node:", n.SelectAttr("id"))
			for _, nn := range n.ChildNodes() {
				fmt.Printf("node: %s\n", nn.Data)
			}
		}
	}

	channel := xmlquery.FindOne(doc, "//channel")
	if n := channel.SelectElement("title"); n != nil {
		fmt.Printf("title: %s\n", n.InnerText())
	}
	if n := channel.SelectElement("link"); n != nil {
		fmt.Printf("link: %s\n", n.InnerText())
	}
	for i, n := range xmlquery.Find(doc, "//item/title") {
		fmt.Printf("#%d %s\n", i, n.InnerText())
	}
}
