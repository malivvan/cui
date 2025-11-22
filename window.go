package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Window is a draggable, resizable frame around a primitive. Windows must be
// added to a WindowManager.
type Window struct {
	box *Box

	primitive Widget

	fullscreen bool

	normalX, normalY int
	normalW, normalH int

	dragX, dragY   int
	dragWX, dragWY int

	mu sync.RWMutex
}

// NewWindow returns a new window around the given primitive.
func NewWindow() *Window {
	w := &Window{
		box:       NewBox(),
		primitive: NewBox(),
		dragWX:    -1,
		dragWY:    -1,
	}
	w.box.focus = w
	return w
}

func (w *Window) SetWidget(widget Widget) *Window {
	return w.set(func(w *Window) { w.primitive = widget })
}

func (w *Window) GetWidget() (widget Widget) {
	w.get(func(w *Window) { widget = w.primitive })
	return
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (w *Window) set(setter func(w *Window)) *Window {
	w.mu.Lock()
	setter(w)
	w.mu.Unlock()
	return w
}

func (w *Window) get(getter func(w *Window)) {
	w.mu.RLock()
	getter(w)
	w.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Window.
func (w *Window) GetTitle() string {
	return w.box.GetTitle()
}

// SetTitle sets the title of this Window.
func (w *Window) SetTitle(title string) *Window {
	w.box.SetTitle(title)
	return w
}

// GetTitleAlign returns the title alignment of this Window.
func (w *Window) GetTitleAlign() int {
	return w.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Window.
func (w *Window) SetTitleAlign(align int) *Window {
	w.box.SetTitleAlign(align)
	return w
}

// GetBorder returns whether this Window has a border.
func (w *Window) GetBorder() bool {
	return w.box.GetBorder()
}

// SetBorder sets whether this Window has a border.
func (w *Window) SetBorder(show bool) *Window {
	w.box.SetBorder(show)
	return w
}

// GetBorderColor returns the border color of this Window.
func (w *Window) GetBorderColor() tcell.Color {
	return w.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Window.
func (w *Window) SetBorderColor(color tcell.Color) *Window {
	w.box.SetBorderColor(color)
	return w
}

// GetBorderAttributes returns the border attributes of this Window.
func (w *Window) GetBorderAttributes() tcell.AttrMask {
	return w.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Window.
func (w *Window) SetBorderAttributes(attr tcell.AttrMask) *Window {
	w.box.SetBorderAttributes(attr)
	return w
}

// GetBorderColorFocused returns the border color of this Window when focused.
func (w *Window) GetBorderColorFocused() tcell.Color {
	return w.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Window when focused.
func (w *Window) SetBorderColorFocused(color tcell.Color) *Window {
	w.box.SetBorderColorFocused(color)
	return w
}

// GetTitleColor returns the title color of this Window.
func (w *Window) GetTitleColor() tcell.Color {
	return w.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Window.
func (w *Window) SetTitleColor(color tcell.Color) *Window {
	w.box.SetTitleColor(color)
	return w
}

// GetDrawFunc returns the custom draw function of this Window.
func (w *Window) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return w.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Window.
func (w *Window) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Window {
	w.box.SetDrawFunc(handler)
	return w
}

// ShowFocus sets whether this Window should show a focus indicator when focused.
func (w *Window) ShowFocus(showFocus bool) *Window {
	w.box.ShowFocus(showFocus)
	return w
}

// GetMouseCapture returns the mouse capture function of this Window.
func (w *Window) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return w.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Window.
func (w *Window) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Window {
	w.box.SetMouseCapture(capture)
	return w
}

// GetBackgroundColor returns the background color of this Window.
func (w *Window) GetBackgroundColor() tcell.Color {
	return w.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Window.
func (w *Window) SetBackgroundColor(color tcell.Color) *Window {
	w.box.SetBackgroundColor(color)
	return w
}

// GetBackgroundTransparent returns whether the background of this Window is transparenw.
func (w *Window) GetBackgroundTransparent() bool {
	return w.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Window is transparenw.
func (w *Window) SetBackgroundTransparent(transparent bool) *Window {
	w.box.SetBackgroundTransparent(transparent)
	return w
}

// GetInputCapture returns the input capture function of this Window.
func (w *Window) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return w.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Window.
func (w *Window) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Window {
	w.box.SetInputCapture(capture)
	return w
}

// GetPadding returns the padding of this Window.
func (w *Window) GetPadding() (top, bottom, left, right int) {
	return w.box.GetPadding()
}

// SetPadding sets the padding of this Window.
func (w *Window) SetPadding(top, bottom, left, right int) *Window {
	w.box.SetPadding(top, bottom, left, right)
	return w
}

// InRect returns whether the given screen coordinates are within this Window.
func (w *Window) InRect(x, y int) bool {
	return w.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Window.
func (w *Window) GetInnerRect() (x, y, width, height int) {
	return w.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Window is preserved.
func (w *Window) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return w.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Window is preserved.
func (w *Window) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return w.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Window.
func (w *Window) GetRect() (x, y, width, height int) {
	return w.box.GetRect()
}

// SetRect sets the rectangle occupied by this Window.
func (w *Window) SetRect(x, y, width, height int) {
	w.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Window is visible.
func (w *Window) GetVisible() bool {
	return w.box.GetVisible()
}

// SetVisible sets whether this Window is visible.
func (w *Window) SetVisible(visible bool) {
	w.box.SetVisible(visible)
}

// Focus is called when this primitive receives focus.
func (w *Window) Focus(delegate func(p Widget)) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.box.Focus(delegate)

	w.primitive.Focus(delegate)
}

// HasFocus returns whether this primitive has focus.
func (w *Window) HasFocus() bool {
	w.mu.RLock()
	defer w.mu.RUnlock()

	focusable := w.primitive.GetFocusable()
	if focusable != nil {
		return focusable.HasFocus()
	}

	return w.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Window.
func (w *Window) GetFocusable() Focusable {
	return w.box.GetFocusable()
}

// Blur is called when this primitive loses focus.
func (w *Window) Blur() {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.box.Blur()

	w.primitive.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

// SetFullscreen sets the flag indicating whether or not the the window should
// be drawn fullscreen.
func (w *Window) SetFullscreen(fullscreen bool) *Window {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.fullscreen == fullscreen {
		return w
	}

	w.fullscreen = fullscreen
	if w.fullscreen {
		w.normalX, w.normalY, w.normalW, w.normalH = w.GetRect()
	} else {
		w.SetRect(w.normalX, w.normalY, w.normalW, w.normalH)
	}

	return w
}

// Draw draws this primitive onto the screen.
func (w *Window) Draw(screen tcell.Screen) {
	if !w.GetVisible() {
		return
	}

	w.mu.RLock()
	defer w.mu.RUnlock()

	w.box.Draw(screen)

	x, y, width, height := w.GetInnerRect()
	w.primitive.SetRect(x, y, width, height)
	w.primitive.Draw(screen)
}

// InputHandler returns the handler for this primitive.
func (w *Window) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return w.primitive.InputHandler()
}

// MouseHandler returns the mouse handler for this primitive.
func (w *Window) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return w.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if !w.InRect(event.Position()) {
			return false, nil
		}

		if action == MouseLeftDown || action == MouseMiddleDown || action == MouseRightDown {
			setFocus(w)
		}

		if action == MouseLeftDown {
			x, y, width, height := w.GetRect()
			mouseX, mouseY := event.Position()

			leftEdge := mouseX == x
			rightEdge := mouseX == x+width-1
			bottomEdge := mouseY == y+height-1
			topEdge := mouseY == y

			if mouseY >= y && mouseY <= y+height-1 {
				if leftEdge {
					w.dragX = -1
				} else if rightEdge {
					w.dragX = 1
				}
			}

			if mouseX >= x && mouseX <= x+width-1 {
				if bottomEdge {
					w.dragY = -1
				} else if topEdge {
					if leftEdge || rightEdge {
						w.dragY = 1
					} else {
						w.dragWX = mouseX - x
						w.dragWY = mouseY - y
					}
				}
			}
		}

		_, capture = w.primitive.MouseHandler()(action, event, setFocus)
		return true, capture
	})
}
