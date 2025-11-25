package main

import (
	"fmt"
	"strings"

	"github.com/beevik/etree"
	"github.com/malivvan/cui"
)

type View struct {
	root   *Node[cui.Widget]
	flex   map[string]*Node[*cui.Flex]
	button map[string]*Node[*cui.Button]
}

type Node[WIDGET cui.Widget] struct {
	*etree.Element
	widget   WIDGET
	parent   *Node[cui.Widget]
	children []*Node[cui.Widget]
}

var (
	ErrNoRootNode = fmt.Errorf("no root node found")
)

func NewView(s string) (view *View, err error) {

	doc := etree.NewDocument()
	if err := doc.ReadFromString(s); err != nil {
		panic(err)
	}
	view = &View{
		flex:   make(map[string]*Node[*cui.Flex]),
		button: make(map[string]*Node[*cui.Button]),
	}
	view.root, err = view.traverse(doc.Root())
	if err != nil {
		return nil, err
	}
	return view, nil
}

func (view *View) traverse(elem *etree.Element) (*Node[cui.Widget], error) {
	switch elem.Tag {
	case "flex":
		flex := cui.NewFlex().SetDirection(cui.FlexRow)
		for _, child := range elem.ChildElements() {
			childNode, err := view.traverse(child)
			if err != nil {
				return nil, err
			}
			flex.AddItem(childNode.widget, 1, 0, false)
		}
		return &Node[cui.Widget]{
			Element: elem,
			widget:  flex,
		}, nil

	case "button":
		button := cui.NewButton()
		button.SetLabel(elem.Text())
		return &Node[cui.Widget]{
			Element: elem,
			widget:  button,
		}, nil
	}
	return nil, fmt.Errorf("unknown node type: %s", strings.ToLower(elem.Tag))
}

func main() {
	s := `<flex>
	<button>Click Me!</button>
	<button>Click Me2</button>
	<button>Click Me3</button>
</flex>
`

	view, err := NewView(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(view)

	app := cui.New()
	defer app.HandlePanic()

	if err := app.SetRoot(view.root.widget, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
