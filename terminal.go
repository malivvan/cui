package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/malivvan/cui/terminal"
	"github.com/malivvan/cui/terminal/pty"
)

type Terminal struct {
	box *Box

	term    *terminal.VT
	running bool
	opt     pty.Options
	app     *App
	w       int
	h       int

	mu sync.RWMutex
}

func NewTerminal(app *App, opt pty.Options) *Terminal {
	t := &Terminal{
		box:  NewBox(),
		term: terminal.New(),
		app:  app,
		opt:  opt,
	}
	return t
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (t *Terminal) set(setter func(t *Terminal)) *Terminal {
	t.mu.Lock()
	setter(t)
	t.mu.Unlock()
	return t
}

func (t *Terminal) get(getter func(t *Terminal)) {
	t.mu.RLock()
	getter(t)
	t.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Terminal.
func (t *Terminal) GetTitle() string {
	return t.box.GetTitle()
}

// SetTitle sets the title of this Terminal.
func (t *Terminal) SetTitle(title string) *Terminal {
	t.box.SetTitle(title)
	return t
}

// GetTitleAlign returns the title alignment of this Terminal.
func (t *Terminal) GetTitleAlign() int {
	return t.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Terminal.
func (t *Terminal) SetTitleAlign(align int) *Terminal {
	t.box.SetTitleAlign(align)
	return t
}

// GetBorder returns whether this Terminal has a border.
func (t *Terminal) GetBorder() bool {
	return t.box.GetBorder()
}

// SetBorder sets whether this Terminal has a border.
func (t *Terminal) SetBorder(show bool) *Terminal {
	t.box.SetBorder(show)
	return t
}

// GetBorderColor returns the border color of this Terminal.
func (t *Terminal) GetBorderColor() tcell.Color {
	return t.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Terminal.
func (t *Terminal) SetBorderColor(color tcell.Color) *Terminal {
	t.box.SetBorderColor(color)
	return t
}

// GetBorderAttributes returns the border attributes of this Terminal.
func (t *Terminal) GetBorderAttributes() tcell.AttrMask {
	return t.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Terminal.
func (t *Terminal) SetBorderAttributes(attr tcell.AttrMask) *Terminal {
	t.box.SetBorderAttributes(attr)
	return t
}

// GetBorderColorFocused returns the border color of this Terminal when focused.
func (t *Terminal) GetBorderColorFocused() tcell.Color {
	return t.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Terminal when focused.
func (t *Terminal) SetBorderColorFocused(color tcell.Color) *Terminal {
	t.box.SetBorderColorFocused(color)
	return t
}

// GetTitleColor returns the title color of this Terminal.
func (t *Terminal) GetTitleColor() tcell.Color {
	return t.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Terminal.
func (t *Terminal) SetTitleColor(color tcell.Color) *Terminal {
	t.box.SetTitleColor(color)
	return t
}

// GetDrawFunc returns the custom draw function of this Terminal.
func (t *Terminal) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return t.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Terminal.
func (t *Terminal) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Terminal {
	t.box.SetDrawFunc(handler)
	return t
}

// ShowFocus sets whether this Terminal should show a focus indicator when focused.
func (t *Terminal) ShowFocus(showFocus bool) *Terminal {
	t.box.ShowFocus(showFocus)
	return t
}

// GetMouseCapture returns the mouse capture function of this Terminal.
func (t *Terminal) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return t.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Terminal.
func (t *Terminal) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Terminal {
	t.box.SetMouseCapture(capture)
	return t
}

// GetBackgroundColor returns the background color of this Terminal.
func (t *Terminal) GetBackgroundColor() tcell.Color {
	return t.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Terminal.
func (t *Terminal) SetBackgroundColor(color tcell.Color) *Terminal {
	t.box.SetBackgroundColor(color)
	return t
}

// GetBackgroundTransparent returns whether the background of this Terminal is transparent.
func (t *Terminal) GetBackgroundTransparent() bool {
	return t.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Terminal is transparent.
func (t *Terminal) SetBackgroundTransparent(transparent bool) *Terminal {
	t.box.SetBackgroundTransparent(transparent)
	return t
}

// GetInputCapture returns the input capture function of this Terminal.
func (t *Terminal) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return t.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Terminal.
func (t *Terminal) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Terminal {
	t.box.SetInputCapture(capture)
	return t
}

// GetPadding returns the padding of this Terminal.
func (t *Terminal) GetPadding() (top, bottom, left, right int) {
	return t.box.GetPadding()
}

// SetPadding sets the padding of this Terminal.
func (t *Terminal) SetPadding(top, bottom, left, right int) *Terminal {
	t.box.SetPadding(top, bottom, left, right)
	return t
}

// InRect returns whether the given screen coordinates are within this Terminal.
func (t *Terminal) InRect(x, y int) bool {
	return t.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Terminal.
func (t *Terminal) GetInnerRect() (x, y, width, height int) {
	return t.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Terminal is preserved.
func (t *Terminal) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return t.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Terminal is preserved.
func (t *Terminal) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return t.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Terminal.
func (t *Terminal) GetRect() (x, y, width, height int) {
	return t.box.GetRect()
}

// SetRect sets the rectangle occupied by this Terminal.
func (t *Terminal) SetRect(x, y, width, height int) {
	t.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Terminal is visible.
func (t *Terminal) GetVisible() bool {
	return t.box.GetVisible()
}

// SetVisible sets whether this Terminal is visible.
func (t *Terminal) SetVisible(visible bool) {
	t.box.SetVisible(visible)
}

// Focus is called when this Terminal receives focus.
func (t *Terminal) Focus(delegate func(p Widget)) {
	t.box.Focus(delegate)
}

// HasFocus returns whether this Terminal has focus.
func (t *Terminal) HasFocus() bool {
	return t.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Terminal.
func (t *Terminal) GetFocusable() Focusable {
	return t.box.GetFocusable()
}

// Blur is called when this Terminal loses focus.
func (t *Terminal) Blur() {
	t.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

func (t *Terminal) Draw(s tcell.Screen) {
	if !t.GetVisible() {
		return
	}

	t.box.Draw(s)

	t.mu.Lock()
	defer t.mu.Unlock()

	x, y, w, h := t.GetInnerRect()
	view := views.NewViewPort(s, x, y, w, h)
	t.term.SetSurface(view)
	if w != t.w || h != t.h {
		t.w = w
		t.h = h
		t.term.Resize(w, h)
	}

	if !t.running {
		err := t.term.Start(t.opt)
		if err != nil {
			panic(err)
		}
		t.term.Attach(t.HandleEvent)
		t.running = true
	}
	if t.HasFocus() {
		cy, cx, style, vis := t.term.Cursor()
		if vis {
			s.ShowCursor(cx+x, cy+y)
			s.SetCursorStyle(style)
		} else {
			s.HideCursor()
		}
	}
	t.term.Draw()
}

func (t *Terminal) HandleEvent(ev tcell.Event) {
	switch ev.(type) {
	case *terminal.EventRedraw:
		go func() {
			t.app.QueueUpdateDraw(func() {})
		}()
	}
}

func (t *Terminal) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		t.term.HandleEvent(event)
	})
}

func (t *Terminal) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return t.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if action == MouseLeftClick && t.InRect(event.Position()) {
			setFocus(t)
			return t.term.HandleEvent(event), nil
		}
		return false, nil
	})
}
