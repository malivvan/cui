// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package markup

import (
	"github.com/malivvan/cui/markup/atom"
)

// A NodeType is the type of a Node.
type NodeType uint32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
	// RawNode nodes are not returned by the parser, but can be part of the
	// Node tree passed to func Render to insert raw HTML (without escaping).
	// If so, this package makes no guarantee that the rendered HTML is secure
	// (from e.g. Cross Site Scripting attacks) or well-formed.
	RawNode
	scopeMarkerNode
)

// Section 12.2.4.3 says "The markers are inserted when entering applet,
// object, marquee, template, td, th, and caption elements, and are used
// to prevent formatting from "leaking" into applet, object, marquee,
// template, td, th, and caption elements".
var scopeMarker = Node{Type: scopeMarkerNode}

// A Node consists of a NodeType and some Tag (tag name for element nodes,
// content for text) and are part of a tree of Nodes. Element nodes may also
// have a NS and contain a slice of Attrs. Tag is unescaped, so
// that it looks like "a<b" rather than "a&lt;b". For element nodes, Atom
// is the atom for Tag, or zero if Tag is not a known tag name.
//
// Node trees may be navigated using the link fields (Parent,
// FirstChild, and so on) or a range loop over iterators such as
// [Node.Descendants].
//
// An empty NS implies a "http://www.w3.org/1999/xhtml" namespace.
// Similarly, "math" is short for "http://www.w3.org/1998/Math/MathML", and
// "svg" is short for "http://www.w3.org/2000/svg".
type Node struct {
	NS          string
	Tag         string
	Type        NodeType
	Atom        atom.Atom
	Attrs       []Attribute
	Style       []Property
	Parent      *Node
	FirstChild  *Node
	LastChild   *Node
	PrevSibling *Node
	NextSibling *Node
}

type Property struct {
	Key string
	Val string
}

type Nodes []*Node

func (ns Nodes) Get(i int) *Node {
	if i < 0 || i >= len(ns) {
		return nil
	}
	return ns[i]
}

func (ns Nodes) First() *Node {
	if len(ns) == 0 {
		return nil
	}
	return ns[0]
}

func (ns Nodes) Last() *Node {
	if len(ns) == 0 {
		return nil
	}
	return ns[len(ns)-1]
}

func (ns Nodes) Len() int {
	return len(ns)
}

type Predicate func(n *Node) bool

func (ns Nodes) Filter(p Predicate) (fns Nodes) {
	for _, n := range ns {
		if p(n) {
			fns = append(fns, n)
		}
	}
	return fns
}

func (ns Nodes) Each(f func(n *Node) bool) {
	for _, n := range ns {
		if !f(n) {
			break
		}
	}
}

func (ns Nodes) Query(query string) Nodes {
	sel, err := ParseSelector(query)
	if err != nil {
		return []*Node{}
	}
	var results Nodes
	for _, n := range ns {
		matches := QueryAll(n, sel)
		results = append(results, matches...)
	}
	return results
}

//func (n *Node) Query(query string) *Node {
//	sel, err := ParseSelector(query)
//	if err != nil {
//		return &Node{}
//	}
//	return Query(n, sel)
//}

func (n *Node) Query(query string) Nodes {
	sel, err := ParseSelector(query)
	if err != nil {
		return []*Node{}
	}
	return QueryAll(n, sel)
}

func (n *Node) GetAttr(key, val string) string {
	for _, a := range n.Attrs {
		if a.Key == key {
			return a.Val
		}
	}
	return val
}

// InsertBefore inserts newChild as a child of n, immediately before oldChild
// in the sequence of n's children. oldChild may be nil, in which case newChild
// is appended to the end of n's children.
//
// It will panic if newChild already has a parent or siblings.
func (n *Node) InsertBefore(newChild, oldChild *Node) {
	if newChild.Parent != nil || newChild.PrevSibling != nil || newChild.NextSibling != nil {
		panic("html: InsertBefore called for an attached child Node")
	}
	var prev, next *Node
	if oldChild != nil {
		prev, next = oldChild.PrevSibling, oldChild
	} else {
		prev = n.LastChild
	}
	if prev != nil {
		prev.NextSibling = newChild
	} else {
		n.FirstChild = newChild
	}
	if next != nil {
		next.PrevSibling = newChild
	} else {
		n.LastChild = newChild
	}
	newChild.Parent = n
	newChild.PrevSibling = prev
	newChild.NextSibling = next
}

// AppendChild adds a node c as a child of n.
//
// It will panic if c already has a parent or siblings.
func (n *Node) AppendChild(c *Node) {
	if c.Parent != nil || c.PrevSibling != nil || c.NextSibling != nil {
		panic("html: AppendChild called for an attached child Node")
	}
	last := n.LastChild
	if last != nil {
		last.NextSibling = c
	} else {
		n.FirstChild = c
	}
	n.LastChild = c
	c.Parent = n
	c.PrevSibling = last
}

// RemoveChild removes a node c that is a child of n. Afterwards, c will have
// no parent and no siblings.
//
// It will panic if c's parent is not n.
func (n *Node) RemoveChild(c *Node) {
	if c.Parent != n {
		panic("html: RemoveChild called for a non-child Node")
	}
	if n.FirstChild == c {
		n.FirstChild = c.NextSibling
	}
	if c.NextSibling != nil {
		c.NextSibling.PrevSibling = c.PrevSibling
	}
	if n.LastChild == c {
		n.LastChild = c.PrevSibling
	}
	if c.PrevSibling != nil {
		c.PrevSibling.NextSibling = c.NextSibling
	}
	c.Parent = nil
	c.PrevSibling = nil
	c.NextSibling = nil
}

// reparentChildren reparents all of src's child nodes to dst.
func reparentChildren(dst, src *Node) {
	for {
		child := src.FirstChild
		if child == nil {
			break
		}
		src.RemoveChild(child)
		dst.AppendChild(child)
	}
}

// clone returns a new node with the same type, data and attributes.
// The clone has no parent, no siblings and no children.
func (n *Node) clone() *Node {
	m := &Node{
		Type:  n.Type,
		Atom:  n.Atom,
		Tag:   n.Tag,
		Attrs: make([]Attribute, len(n.Attrs)),
	}
	copy(m.Attrs, n.Attrs)
	return m
}

// nodeStack is a stack of nodes.
type nodeStack []*Node

// pop pops the stack. It will panic if s is empty.
func (s *nodeStack) pop() *Node {
	i := len(*s)
	n := (*s)[i-1]
	*s = (*s)[:i-1]
	return n
}

// top returns the most recently pushed node, or nil if s is empty.
func (s *nodeStack) top() *Node {
	if i := len(*s); i > 0 {
		return (*s)[i-1]
	}
	return nil
}

// index returns the index of the top-most occurrence of n in the stack, or -1
// if n is not present.
func (s *nodeStack) index(n *Node) int {
	for i := len(*s) - 1; i >= 0; i-- {
		if (*s)[i] == n {
			return i
		}
	}
	return -1
}

// contains returns whether a is within s.
func (s *nodeStack) contains(a atom.Atom) bool {
	for _, n := range *s {
		if n.Atom == a && n.NS == "" {
			return true
		}
	}
	return false
}

// insert inserts a node at the given index.
func (s *nodeStack) insert(i int, n *Node) {
	*s = append(*s, nil)
	copy((*s)[i+1:], (*s)[i:])
	(*s)[i] = n
}

// remove removes a node from the stack. It is a no-op if n is not present.
func (s *nodeStack) remove(n *Node) {
	i := s.index(n)
	if i == -1 {
		return
	}
	copy((*s)[i:], (*s)[i+1:])
	j := len(*s) - 1
	(*s)[j] = nil
	*s = (*s)[:j]
}

type insertionModeStack []insertionMode

func (s *insertionModeStack) pop() (im insertionMode) {
	i := len(*s)
	im = (*s)[i-1]
	*s = (*s)[:i-1]
	return im
}

func (s *insertionModeStack) top() insertionMode {
	if i := len(*s); i > 0 {
		return (*s)[i-1]
	}
	return nil
}
