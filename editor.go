package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui/editor"
)

type Editor struct {
	box  *Box
	view *editor.View
	mu   sync.RWMutex
}

func (e *Editor) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return func(event *tcell.EventKey, setFocus func(p Widget)) {
		e.view.HandleEvent(event)
	}
}

func (e *Editor) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if e.view.InRect(event.Position()) {
			setFocus(e)
			e.view.HandleEvent(event)
			return true, nil
		}
		return false, nil
	}
}

func NewEditor() *Editor {
	e := &Editor{
		box:  NewBox(),
		view: editor.NewView(),
	}
	return e
}

func (e *Editor) SetTheme(theme string) *Editor {
	return e.set(func(e *Editor) { e.view.SetTheme(theme) })
}

func (e *Editor) SetBuffer(buf *editor.Buffer) *Editor {
	return e.set(func(e *Editor) { e.view.SetBuffer(buf) })
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (e *Editor) set(setter func(e *Editor)) *Editor {
	e.mu.Lock()
	setter(e)
	e.mu.Unlock()
	return e
}

func (e *Editor) get(getter func(e *Editor)) {
	e.mu.RLock()
	getter(e)
	e.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Editor.
func (e *Editor) GetTitle() string {
	return e.box.GetTitle()
}

// SetTitle sets the title of this Editor.
func (e *Editor) SetTitle(title string) *Editor {
	e.box.SetTitle(title)
	return e
}

// GetTitleAlign returns the title alignment of this Editor.
func (e *Editor) GetTitleAlign() int {
	return e.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Editor.
func (e *Editor) SetTitleAlign(align int) *Editor {
	e.box.SetTitleAlign(align)
	return e
}

// GetBorder returns whether this Editor has a border.
func (e *Editor) GetBorder() bool {
	return e.box.GetBorder()
}

// SetBorder sets whether this Editor has a border.
func (e *Editor) SetBorder(show bool) *Editor {
	e.box.SetBorder(show)
	return e
}

// GetBorderColor returns the border color of this Editor.
func (e *Editor) GetBorderColor() tcell.Color {
	return e.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Editor.
func (e *Editor) SetBorderColor(color tcell.Color) *Editor {
	e.box.SetBorderColor(color)
	return e
}

// GetBorderAttributes returns the border attributes of this Editor.
func (e *Editor) GetBorderAttributes() tcell.AttrMask {
	return e.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Editor.
func (e *Editor) SetBorderAttributes(attr tcell.AttrMask) *Editor {
	e.box.SetBorderAttributes(attr)
	return e
}

// GetBorderColorFocused returns the border color of this Editor when focused.
func (e *Editor) GetBorderColorFocused() tcell.Color {
	return e.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Editor when focused.
func (e *Editor) SetBorderColorFocused(color tcell.Color) *Editor {
	e.box.SetBorderColorFocused(color)
	return e
}

// GetTitleColor returns the title color of this Editor.
func (e *Editor) GetTitleColor() tcell.Color {
	return e.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Editor.
func (e *Editor) SetTitleColor(color tcell.Color) *Editor {
	e.box.SetTitleColor(color)
	return e
}

// GetDrawFunc returns the custom draw function of this Editor.
func (e *Editor) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return e.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Editor.
func (e *Editor) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Editor {
	e.box.SetDrawFunc(handler)
	return e
}

// ShowFocus sets whether this Editor should show a focus indicator when focused.
func (e *Editor) ShowFocus(showFocus bool) *Editor {
	e.box.ShowFocus(showFocus)
	return e
}

// GetMouseCapture returns the mouse capture function of this Editor.
func (e *Editor) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return e.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Editor.
func (e *Editor) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Editor {
	e.box.SetMouseCapture(capture)
	return e
}

// GetBackgroundColor returns the background color of this Editor.
func (e *Editor) GetBackgroundColor() tcell.Color {
	return e.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Editor.
func (e *Editor) SetBackgroundColor(color tcell.Color) *Editor {
	e.box.SetBackgroundColor(color)
	return e
}

// GetBackgroundTransparent returns whether the background of this Editor is transparent.
func (e *Editor) GetBackgroundTransparent() bool {
	return e.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Editor is transparent.
func (e *Editor) SetBackgroundTransparent(transparent bool) *Editor {
	e.box.SetBackgroundTransparent(transparent)
	return e
}

// GetInputCapture returns the input capture function of this Editor.
func (e *Editor) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return e.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Editor.
func (e *Editor) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Editor {
	e.box.SetInputCapture(capture)
	return e
}

// GetPadding returns the padding of this Editor.
func (e *Editor) GetPadding() (top, bottom, left, right int) {
	return e.box.GetPadding()
}

// SetPadding sets the padding of this Editor.
func (e *Editor) SetPadding(top, bottom, left, right int) *Editor {
	e.box.SetPadding(top, bottom, left, right)
	return e
}

// InRect returns whether the given screen coordinates are within this Editor.
func (e *Editor) InRect(x, y int) bool {
	return e.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Editor.
func (e *Editor) GetInnerRect() (x, y, width, height int) {
	return e.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Editor is preserved.
func (e *Editor) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return e.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Editor is preserved.
func (e *Editor) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return e.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Editor.
func (e *Editor) GetRect() (x, y, width, height int) {
	return e.box.GetRect()
}

// SetRect sets the rectangle occupied by this Editor.
func (e *Editor) SetRect(x, y, width, height int) {
	e.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Editor is visible.
func (e *Editor) GetVisible() bool {
	return e.box.GetVisible()
}

// SetVisible sets whether this Editor is visible.
func (e *Editor) SetVisible(visible bool) {
	e.box.SetVisible(visible)
}

// Focus is called when this Editor receives focus.
func (e *Editor) Focus(delegate func(p Widget)) {
	e.box.Focus(delegate)
}

// HasFocus returns whether this Editor has focus.
func (e *Editor) HasFocus() bool {
	return e.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Editor.
func (e *Editor) GetFocusable() Focusable {
	return e.box.GetFocusable()
}

// Blur is called when this Editor loses focus.
func (e *Editor) Blur() {
	e.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

func (e *Editor) Draw(screen tcell.Screen) {
	if !e.GetVisible() {
		return
	}
	e.box.Draw(screen)
	x, y, width, height := e.box.GetInnerRect()

	e.mu.Lock()
	e.view.SetRect(x, y, width, height)
	e.view.Draw(screen)
	e.mu.Unlock()
}
