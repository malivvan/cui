// Package cui provides a collection of terminal user interface (TUI) components built on top of the tcell library.
// It offers various widgets such as buttons, input fields, text views, tables, and more, allowing developers
// to create rich and interactive terminal applications with ease.
package cui

import (
	"github.com/gdamore/tcell/v2"
)

const (
	IconFolder  = '🖿'
	IconBolt    = '🗲'
	IconHDD     = '🖴'
	IconChat    = '🗨'
	IconRefresh = '🗘'
	IconRun     = '⯈'
	IconStop    = '⏹'
	IconPause   = '⏸'
	IconStep    = '⏭'
	IconPin     = '🖈'
	IconClose   = '🗙'
	IconFont    = '🗛'
	IconBook    = '🕮'
	IconWindow  = '🗗'
	IconMax     = '🗖'
	IconMin     = '🗕'
	IconFlag    = '🏲'
	IconNote    = '🎜'
	IconMenu    = '☰'
)

var (
	_ widget[*Box]           = (*Box)(nil)
	_ widget[*Button]        = (*Button)(nil)
	_ widget[*CheckBox]      = (*CheckBox)(nil)
	_ widget[*DropDown]      = (*DropDown)(nil)
	_ widget[*Editor]        = (*Editor)(nil)
	_ widget[*Flex]          = (*Flex)(nil)
	_ widget[*Form]          = (*Form)(nil)
	_ widget[*Frame]         = (*Frame)(nil)
	_ widget[*Grid]          = (*Grid)(nil)
	_ widget[*Image]         = (*Image)(nil)
	_ widget[*Input]         = (*Input)(nil)
	_ widget[*Layout]        = (*Layout)(nil)
	_ widget[*List]          = (*List)(nil)
	_ widget[*Modal]         = (*Modal)(nil)
	_ widget[*Panels]        = (*Panels)(nil)
	_ widget[*Progress]      = (*Progress)(nil)
	_ widget[*MenuBar]       = (*MenuBar)(nil)
	_ widget[*SubMenu]       = (*SubMenu)(nil)
	_ widget[*MenuItem]      = (*MenuItem)(nil)
	_ widget[*BarChart]      = (*BarChart)(nil)
	_ widget[*Spinner]       = (*Spinner)(nil)
	_ widget[*Gauge]         = (*Gauge)(nil)
	_ widget[*Plot]          = (*Plot)(nil)
	_ widget[*Sparkline]     = (*Sparkline)(nil)
	_ widget[*Slider]        = (*Slider)(nil)
	_ widget[*TabbedPanels]  = (*TabbedPanels)(nil)
	_ widget[*Table]         = (*Table)(nil)
	_ widget[*Text]          = (*Text)(nil)
	_ widget[*Tree]          = (*Tree)(nil)
	_ widget[*Terminal]      = (*Terminal)(nil)
	_ widget[*Window]        = (*Window)(nil)
	_ widget[*WindowManager] = (*WindowManager)(nil)
)

func (f *Flex) Children() []Widget {
	f.mu.RLock()
	defer f.mu.RUnlock()

	children := make([]Widget, len(f.items))
	for i, item := range f.items {
		children[i] = item.Item
	}
	return children
}

func (g *Grid) Children() []Widget {
	g.mu.RLock()
	defer g.mu.RUnlock()

	children := make([]Widget, 0, len(g.items))
	for _, item := range g.items {
		children = append(children, item.Item)
	}
	return children
}

//////////////////////////////////////////////////////////////////////

type widget[T Widget] interface {
	mutex[T]
	box[T]
}

type mutex[T Widget] interface {
	set(setter func(b T)) T
	get(getter func(b T))
}

type box[T Widget] interface {
	GetTitle() (title string)
	SetTitle(title string) T

	GetTitleColor() (color tcell.Color)
	SetTitleColor(color tcell.Color) T

	GetTitleAlign() (align int)
	SetTitleAlign(align int) T

	GetPadding() (top, bottom, left, right int)
	SetPadding(top, bottom, left, right int) T

	GetBorder() (border bool)
	SetBorder(show bool) T

	GetBorderColor() (color tcell.Color)
	SetBorderColor(color tcell.Color) T

	GetBorderColorFocused() (color tcell.Color)
	SetBorderColorFocused(color tcell.Color) T

	GetBorderAttributes() (attr tcell.AttrMask)
	SetBorderAttributes(attr tcell.AttrMask) T

	GetDrawFunc() (draw func(screen tcell.Screen, x, y, width, height int) (int, int, int, int))
	SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) T

	GetInputCapture() (fn func(event *tcell.EventKey) *tcell.EventKey)
	SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) T

	GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)
	SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) T

	GetBackgroundColor() (backgroundColor tcell.Color)
	SetBackgroundColor(color tcell.Color) T

	GetBackgroundTransparent() (transparent bool)
	SetBackgroundTransparent(transparent bool) T

	InRect(x, y int) bool
	ShowFocus(showFocus bool) T
	GetInnerRect() (innerX, innerY, innerW, innerH int)
	WrapInputHandler(inputHandler func(*tcell.EventKey, func(p Widget))) func(*tcell.EventKey, func(p Widget))
	WrapMouseHandler(mouseHandler func(MouseAction, *tcell.EventMouse, func(p Widget)) (bool, Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)
}

//////////////////////////////////////////////////////////////////////
//
//type View struct {
//	doc    *etree.Document
//	root   *Node[Widget]
//	flex   map[string]*Flex
//	button map[string]*Button
//}
//
//func NewEditor(s string) (view *View, err error) {
//	doc := etree.NewDocument()
//	if err := doc.ReadFromString(s); err != nil {
//		panic(err)
//	}
//	view = &View{
//		doc:    doc,
//		flex:   make(map[string]*Flex),
//		button: make(map[string]*Button),
//	}
//	view.root, err = view.traverse(doc.Root())
//	if err != nil {
//		return nil, err
//	}
//	return view, nil
//}
//
//func (view *View) Doc() *etree.Document {
//	return view.doc
//}
//
//func (view *View) Root() Widget {
//	return view.root.widget
//}
//
//func (view *View) Flex(id string) *Flex {
//	if flex, ok := view.flex[id]; ok {
//		return flex
//	}
//	return nil
//}
//
//func (view *View) Button(id string) *Button {
//	if button, ok := view.button[id]; ok {
//		return button
//	}
//	return nil
//}
//
//type Node[WIDGET Widget] struct {
//	element  *etree.Element
//	widget   WIDGET
//	parent   *Node[Widget]
//	children []*Node[Widget]
//}
//
//func (view *View) traverse(element *etree.Element) (*Node[Widget], error) {
//	id := element.SelectAttrValue("id", "")
//	switch element.Tag {
//	case "flex":
//		flex := NewFlex().
//			SetDirection(FlexRow)
//		if id != "" {
//			view.flex[id] = flex
//		}
//		flexNode := &Node[Widget]{
//			element: element,
//			widget:  flex,
//		}
//		for _, child := range element.ChildElements() {
//			childNode, err := view.traverse(child)
//			if err != nil {
//				return nil, err
//			}
//			childNode.parent = flexNode
//			flexNode.children = append(flexNode.children, childNode)
//			flex.AddItem(childNode.widget, 1, 0, false)
//		}
//		return flexNode, nil
//
//	case "button":
//		button := NewButton().SetLabel(element.Text())
//		if id != "" {
//			view.button[id] = button
//		}
//		return &Node[Widget]{
//			element: element,
//			widget:  button,
//		}, nil
//	}
//	return nil, fmt.Errorf("unknown node type: %s", strings.ToLower(element.Tag))
//}
