package markup

import (
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/markup/css"
	"github.com/malivvan/cui/markup/xml"
)

type Node struct {
	id       string
	tag      string
	class    []string
	attrs    map[xml.Attribute]string
	props    map[css.Property]string
	parent   *Node
	children []*Node
	widget   cui.Widget
	element  *xml.Element
}

func (node *Node) ID() string            { return node.id }
func (node *Node) Tag() string           { return node.tag }
func (node *Node) Class() []string       { return node.class }
func (node *Node) Widget() cui.Widget    { return node.widget }
func (node *Node) Element() *xml.Element { return node.element }
func (node *Node) Parent() *Node         { return node.parent }
func (node *Node) Children() []*Node     { return node.children }

//
//func (node *Node) getStyle(property string) (value string, found bool) {
//
//	// 1. Use inline styles
//	for _, decl := range node.style {
//		if decl.Property == property {
//			return decl.Value, true
//		}
//	}
//
//	candidates := []string{"#" + node.id}
//	for _, class := range node.class {
//		candidates = append(candidates, "."+class)
//	}
//	candidates = append(candidates, node.tag)
//	fmt.Println("Candidates:", candidates)
//	for _, cand := range candidates {
//
//		for _, rule := range node.rules {
//
//			for _, sel := range rule.Selectors {
//
//				if sel == cand {
//
//					for _, decl := range rule.Declarations {
//						if decl.Property == property {
//							return decl.Value, true
//						}
//					}
//				}
//			}
//		}
//	}
//	return "", false
//}
