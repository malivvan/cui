package markup

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/malivvan/cui"
	"github.com/malivvan/cui/markup/css"
	"github.com/malivvan/cui/markup/xml"
)

type File struct {
	path string
	doc  *xml.Document
	root []*Node

	widget map[string]cui.Widget
	button map[string]*cui.Button
	flex   map[string]*cui.Flex
	mu     sync.Mutex
}

func OpenFile(path string) (f *File, err error) {
	var data []byte
	data, err = os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewFile(path, string(data))
}

func NewFile(path string, data string) (f *File, err error) {
	f = &File{
		path:   path,
		doc:    xml.NewDocument(),
		widget: make(map[string]cui.Widget),
		button: make(map[string]*cui.Button),
		flex:   make(map[string]*cui.Flex),
	}
	err = f.doc.ReadFromString(data)
	if err != nil {
		return nil, err
	}
	var node *Node
	var rules []*css.Rule
	for elem := f.doc.Root(); elem != nil; elem = node.element.NextSibling() {
		node, rules, err = f.traverse(nil, rules, elem)
		if err != nil {
			return nil, err
		}
		f.root = append(f.root, node)
	}
	return f, nil
}

func (f *File) RootCount() int {
	return len(f.root)
}

func (f *File) Root(i int) *Node {
	if i < 0 || i >= len(f.root) {
		return nil
	}
	return f.root[i]
}

func (f *File) Widget(id string) cui.Widget {
	if w, ok := f.widget[id]; ok {
		return w
	}
	return nil
}

func (f *File) Button(id string) *cui.Button {
	if b, ok := f.button[id]; ok {
		return b
	}
	return nil
}

func (f *File) Flex(id string) *cui.Flex {
	if f, ok := f.flex[id]; ok {
		return f
	}
	return nil
}

func (f *File) traverse(parent *Node, rules []*css.Rule, element *xml.Element) (*Node, []*css.Rule, error) {

	create := func(element *xml.Element, rules []*css.Rule, defaults map[string]string) (node *Node, err error) {

		tag := element.Tag

		var id string
		if idVal := element.SelectAttr("id"); idVal != nil {
			id = idVal.Value
		}
		if id == "" {
			for i := 1; i <= 1000; i++ {
				id = fmt.Sprintf("%s-%d", element.Tag, i)
				if _, exists := f.widget[id]; !exists {
					break
				}
			}

		}

		var class []string
		if classVal := element.SelectAttr("class"); classVal != nil {
			class = strings.Split(classVal.Value, " ")
		}

		var style []*css.Declaration
		if styleVal := element.SelectAttr(string(xml.STYLE)); styleVal != nil {
			style, err = css.ParseDeclarations(styleVal.Value)
			if err != nil {
				return nil, err
			}
		}
		_ = style

		node = &Node{
			id:      id,
			tag:     tag,
			class:   class,
			parent:  parent,
			element: element,
			props:   make(map[css.Property]string),
			attrs:   make(map[xml.Attribute]string),
		}

		//	for key, val := range defaults {
		//		node.attrs[key] = element.SelectAttrValue(key, val)
		//	}
		return
	}

	var err error
	var node *Node
	switch element.Tag {
	case "style":
		var style []*css.Rule
		if src := strings.TrimSpace(element.SelectAttrValue("src", "")); src != "" {
			if !strings.HasPrefix(src, "/") {
				src = filepath.Join(filepath.Dir(f.path), src)
			}
			data, err := os.ReadFile(src)
			if err != nil {
				return nil, nil, err
			}
			style, err = css.Parse(string(data))
			if err != nil {
				return nil, nil, err
			}
		} else {
			style, err = css.Parse(element.Text())
			if err != nil {
				return nil, nil, err
			}
		}
		return f.traverse(parent, append(rules, style...), element.NextSibling())

	case "button":
		node, err = create(element, rules, map[string]string{})
		if err != nil {
			return nil, nil, err
		}
		//	col, _ := node.getStyle("color")

		button := cui.NewButton().
			SetLabel(element.Text()) //.
		//		SetLabelColor(tcell.GetColor(col))
		node.widget = button

		f.button[node.id] = button
		f.widget[node.id] = button

	case "flex":
		node, err = create(element, rules, map[string]string{
			"direction": "row",
		})

		flex := cui.NewFlex()
		if node.attrs["direction"] == "column" {
			flex.SetDirection(cui.FlexColumn)
		} else {
			flex.SetDirection(cui.FlexRow)
		}
		node.widget = flex

		for _, childElement := range element.ChildElements() {
			var grow, size int
			if growStr := strings.TrimSpace(childElement.SelectAttrValue("grow", "")); growStr != "" {
				if growInt, err := strconv.Atoi(growStr); err == nil {
					grow = growInt
				}
			}
			if sizeStr := strings.TrimSpace(childElement.SelectAttrValue("size", "")); sizeStr != "" {
				if sizeInt, err := strconv.Atoi(sizeStr); err == nil {
					size = sizeInt
				}
			}

			var err error
			var childNode *Node
			childNode, rules, err = f.traverse(node, rules, childElement)
			if err != nil {
				return nil, nil, err
			}
			flex.AddItem(childNode.widget, size, grow, false)

		}
		f.flex[node.id] = flex
		f.widget[node.id] = flex

	}

	if parent != nil {
		parent.children = append(parent.children, node)
	}
	return node, rules, nil
}
