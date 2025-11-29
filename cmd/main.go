package main

import (
	"fmt"
	"strings"

	"github.com/malivvan/cui"
	"github.com/malivvan/cui/internal/etree"
)

type View struct {
	root   *Node[cui.Widget]
	flex   map[string]*Node[*cui.Flex]
	button map[string]*Node[*cui.Button]
}

type Element[WIDGET cui.Widget] struct {
	*etree.Element
	widget WIDGET
	parent *Container[cui.Widget]
	styles []*Style
}

type Container[WIDGET cui.Widget] struct {
	Element[WIDGET]
	children []*Node[cui.Widget]
}

type Style struct {
}

type Node[WIDGET cui.Widget] struct {
	*etree.Element
	widget   WIDGET
	parent   *Node[cui.Widget]
	children []*Node[cui.Widget]
}

type Stylex struct {
	BackgroundColor string
	ForegroundColor string
	BorderColor     string
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

type Type int

const (
	// ELEMENT is a Widget drawing the ui.
	ELEMENT Type = iota

	// CONTAINER is a Widget that can contain other Widgets.
	// A container can draw itself and control elements.
	// the layout of its children through attributes (e.g. flex)
	CONTAINER

	// STYLE is a Widget that defines styles for other Widgets.
	// Styles are prioritized according to their CONTAINER scope.
	// Styles defined for id supersede styles defined for class.
	// Styles defined for ELEMENT supersede styles defined for id.
	STYLE
)

func main() {
	s := `<style>
#warning {
	background: red;
	foreground: white;
	border: yellow;
}
#info {
	background: blue;
	foreground: white;
	border: cyan;
}
</style>
<flex>
	<button id="first" style="background:#222222;">Click Me!</button>
	<button class="warning">Click Me2</button>
	<button class="info">Click Me3</button>
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
