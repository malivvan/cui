// Package cui provides a collection of terminal user interface (TUI) components built on top of the tcell library.
// It offers various widgets such as buttons, input fields, text views, tables, and more, allowing developers
// to create rich and interactive terminal applications with ease.
package cui

import "github.com/gdamore/tcell/v2"

var (
	_ widget[*Box]           = (*Box)(nil)
	_ widget[*Button]        = (*Button)(nil)
	_ widget[*CheckBox]      = (*CheckBox)(nil)
	_ widget[*DropDown]      = (*DropDown)(nil)
	_ widget[*Flex]          = (*Flex)(nil)
	_ widget[*Form]          = (*Form)(nil)
	_ widget[*Frame]         = (*Frame)(nil)
	_ widget[*Grid]          = (*Grid)(nil)
	_ widget[*Image]         = (*Image)(nil)
	_ widget[*InputField]    = (*InputField)(nil)
	_ widget[*Layout]        = (*Layout)(nil)
	_ widget[*List]          = (*List)(nil)
	_ widget[*Modal]         = (*Modal)(nil)
	_ widget[*Panels]        = (*Panels)(nil)
	_ widget[*ProgressBar]   = (*ProgressBar)(nil)
	_ widget[*Slider]        = (*Slider)(nil)
	_ widget[*Spinner]       = (*Spinner)(nil)
	_ widget[*TabbedPanels]  = (*TabbedPanels)(nil)
	_ widget[*Table]         = (*Table)(nil)
	_ widget[*TextView]      = (*TextView)(nil)
	_ widget[*TreeView]      = (*TreeView)(nil)
	_ widget[*Window]        = (*Window)(nil)
	_ widget[*WindowManager] = (*WindowManager)(nil)
)

///////////////////////////

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
