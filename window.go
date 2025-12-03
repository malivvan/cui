package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// WindowManager provides an area which windows may be added to.
type WindowManager struct {
	box *Box

	windows []*Window

	mu sync.RWMutex
}

// NewWindowManager returns a new window manager.
func NewWindowManager() *WindowManager {
	return &WindowManager{
		box: NewBox(),
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (wm *WindowManager) set(setter func(wm *WindowManager)) *WindowManager {
	wm.mu.Lock()
	setter(wm)
	wm.mu.Unlock()
	return wm
}

func (wm *WindowManager) get(getter func(wm *WindowManager)) {
	wm.mu.RLock()
	getter(wm)
	wm.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this WindowManager.
func (wm *WindowManager) GetTitle() string {
	return wm.box.GetTitle()
}

// SetTitle sets the title of this WindowManager.
func (wm *WindowManager) SetTitle(title string) *WindowManager {
	wm.box.SetTitle(title)
	return wm
}

// GetTitleAlign returns the title alignment of this WindowManager.
func (wm *WindowManager) GetTitleAlign() int {
	return wm.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this WindowManager.
func (wm *WindowManager) SetTitleAlign(align int) *WindowManager {
	wm.box.SetTitleAlign(align)
	return wm
}

// GetBorder returns whether this WindowManager has a border.
func (wm *WindowManager) GetBorder() bool {
	return wm.box.GetBorder()
}

// SetBorder sets whether this WindowManager has a border.
func (wm *WindowManager) SetBorder(show bool) *WindowManager {
	wm.box.SetBorder(show)
	return wm
}

// GetBorderColor returns the border color of this WindowManager.
func (wm *WindowManager) GetBorderColor() tcell.Color {
	return wm.box.GetBorderColor()
}

// SetBorderColor sets the border color of this WindowManager.
func (wm *WindowManager) SetBorderColor(color tcell.Color) *WindowManager {
	wm.box.SetBorderColor(color)
	return wm
}

// GetBorderAttributes returns the border attributes of this WindowManager.
func (wm *WindowManager) GetBorderAttributes() tcell.AttrMask {
	return wm.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this WindowManager.
func (wm *WindowManager) SetBorderAttributes(attr tcell.AttrMask) *WindowManager {
	wm.box.SetBorderAttributes(attr)
	return wm
}

// GetBorderColorFocused returns the border color of this WindowManager when focused.
func (wm *WindowManager) GetBorderColorFocused() tcell.Color {
	return wm.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this WindowManager when focused.
func (wm *WindowManager) SetBorderColorFocused(color tcell.Color) *WindowManager {
	wm.box.SetBorderColorFocused(color)
	return wm
}

// GetTitleColor returns the title color of this WindowManager.
func (wm *WindowManager) GetTitleColor() tcell.Color {
	return wm.box.GetTitleColor()
}

// SetTitleColor sets the title color of this WindowManager.
func (wm *WindowManager) SetTitleColor(color tcell.Color) *WindowManager {
	wm.box.SetTitleColor(color)
	return wm
}

// GetDrawFunc returns the custom draw function of this WindowManager.
func (wm *WindowManager) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return wm.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this WindowManager.
func (wm *WindowManager) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *WindowManager {
	wm.box.SetDrawFunc(handler)
	return wm
}

// ShowFocus sets whether this WindowManager should show a focus indicator when focused.
func (wm *WindowManager) ShowFocus(showFocus bool) *WindowManager {
	wm.box.ShowFocus(showFocus)
	return wm
}

// GetMouseCapture returns the mouse capture function of this WindowManager.
func (wm *WindowManager) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return wm.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this WindowManager.
func (wm *WindowManager) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *WindowManager {
	wm.box.SetMouseCapture(capture)
	return wm
}

// GetBackgroundColor returns the background color of this WindowManager.
func (wm *WindowManager) GetBackgroundColor() tcell.Color {
	return wm.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this WindowManager.
func (wm *WindowManager) SetBackgroundColor(color tcell.Color) *WindowManager {
	wm.box.SetBackgroundColor(color)
	return wm
}

// GetBackgroundTransparent returns whether the background of this WindowManager is transparenwm.
func (wm *WindowManager) GetBackgroundTransparent() bool {
	return wm.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this WindowManager is transparenwm.
func (wm *WindowManager) SetBackgroundTransparent(transparent bool) *WindowManager {
	wm.box.SetBackgroundTransparent(transparent)
	return wm
}

// GetInputCapture returns the input capture function of this WindowManager.
func (wm *WindowManager) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return wm.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this WindowManager.
func (wm *WindowManager) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *WindowManager {
	wm.box.SetInputCapture(capture)
	return wm
}

// GetPadding returns the padding of this WindowManager.
func (wm *WindowManager) GetPadding() (top, bottom, left, right int) {
	return wm.box.GetPadding()
}

// SetPadding sets the padding of this WindowManager.
func (wm *WindowManager) SetPadding(top, bottom, left, right int) *WindowManager {
	wm.box.SetPadding(top, bottom, left, right)
	return wm
}

// InRect returns whether the given screen coordinates are within this WindowManager.
func (wm *WindowManager) InRect(x, y int) bool {
	return wm.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this WindowManager.
func (wm *WindowManager) GetInnerRect() (x, y, width, height int) {
	return wm.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the WindowManager is preserved.
func (wm *WindowManager) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return wm.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the WindowManager is preserved.
func (wm *WindowManager) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return wm.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this WindowManager.
func (wm *WindowManager) GetRect() (x, y, width, height int) {
	return wm.box.GetRect()
}

// SetRect sets the rectangle occupied by this WindowManager.
func (wm *WindowManager) SetRect(x, y, width, height int) {
	wm.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this WindowManager is visible.
func (wm *WindowManager) GetVisible() bool {
	return wm.box.GetVisible()
}

// SetVisible sets whether this WindowManager is visible.
func (wm *WindowManager) SetVisible(visible bool) {
	wm.box.SetVisible(visible)
}

// Focus is called when this primitive receives focus.
func (wm *WindowManager) Focus(delegate func(p Widget)) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	if len(wm.windows) == 0 {
		return
	}

	wm.windows[len(wm.windows)-1].Focus(delegate)
}

// HasFocus returns whether or not this primitive has focus.
func (wm *WindowManager) HasFocus() bool {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	for _, w := range wm.windows {
		if w.HasFocus() {
			return true
		}
	}

	return false
}

// GetFocusable returns the focusable primitive of this WindowManager.
func (wm *WindowManager) GetFocusable() Focusable {
	return wm.box.GetFocusable()
}

// Blur is called when this WindowManager loses focus.
func (wm *WindowManager) Blur() {
	wm.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

// Add adds a window to the manager.
func (wm *WindowManager) Add(w ...*Window) *WindowManager {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	for _, window := range w {
		window.SetBorder(true)
	}

	wm.windows = append(wm.windows, w...)
	return wm
}

// Clear removes all windows from the manager.
func (wm *WindowManager) Clear() {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	wm.windows = nil
}

// Draw draws this primitive onto the screen.
func (wm *WindowManager) Draw(screen tcell.Screen) {
	if !wm.GetVisible() {
		return
	}

	wm.mu.RLock()
	defer wm.mu.RUnlock()

	wm.box.Draw(screen)

	x, y, width, height := wm.GetInnerRect()

	var hasFullScreen bool
	for _, w := range wm.windows {
		if !w.fullscreen || !w.GetVisible() {
			continue
		}

		hasFullScreen = true
		w.SetRect(x-1, y, width+2, height+1)

		w.Draw(screen)
	}
	if hasFullScreen {
		return
	}

	for _, w := range wm.windows {
		if !w.GetVisible() {
			continue
		}

		// Reposition out of bounds windows
		margin := 3
		wx, wy, ww, wh := w.GetRect()
		ox, oy := wx, wy
		if wx > x+width-margin {
			wx = x + width - margin
		}
		if wx+ww < x+margin {
			wx = x - ww + margin
		}
		if wy > y+height-margin {
			wy = y + height - margin
		}
		if wy < y {
			wy = y // No top margin
		}
		if wx != ox || wy != oy {
			w.SetRect(wx, wy, ww, wh)
		}

		w.Draw(screen)
	}
}

func (wm *WindowManager) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return wm.box.InputHandler()
}

// MouseHandler returns the mouse handler for this primitive.
func (wm *WindowManager) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return wm.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if !wm.InRect(event.Position()) {
			return false, nil
		}

		if action == MouseMove {
			mouseX, mouseY := event.Position()

			for _, w := range wm.windows {
				if w.dragWX != -1 || w.dragWY != -1 {
					offsetX := w.box.x - mouseX
					offsetY := w.box.y - mouseY

					w.box.x -= offsetX + w.dragWX
					w.box.y -= offsetY + w.dragWY

					w.box.updateInnerRect()
					consumed = true
				}

				if w.dragX != 0 {
					if w.dragX == -1 {
						offsetX := w.box.x - mouseX

						if w.box.width+offsetX >= Styles.WindowMinWidth {
							w.box.x -= offsetX
							w.box.width += offsetX
						}
					} else {
						offsetX := mouseX - (w.box.x + w.box.width) + 1

						if w.box.width+offsetX >= Styles.WindowMinWidth {
							w.box.width += offsetX
						}
					}

					w.box.updateInnerRect()
					consumed = true
				}

				if w.dragY != 0 {
					if w.dragY == -1 {
						offsetY := mouseY - (w.box.y + w.box.height) + 1

						if w.box.height+offsetY >= Styles.WindowMinHeight {
							w.box.height += offsetY
						}
					} else {
						offsetY := w.box.y - mouseY

						if w.box.height+offsetY >= Styles.WindowMinHeight {
							w.box.y -= offsetY
							w.box.height += offsetY
						}
					}

					w.box.updateInnerRect()
					consumed = true
				}
			}
		} else if action == MouseLeftUp {
			for _, w := range wm.windows {
				w.dragX, w.dragY = 0, 0
				w.dragWX, w.dragWY = -1, -1
			}
		}

		// Focus window on mousedown
		var (
			focusWindow      *Window
			focusWindowIndex int
		)
		for i := len(wm.windows) - 1; i >= 0; i-- {
			if wm.windows[i].InRect(event.Position()) {
				focusWindow = wm.windows[i]
				focusWindowIndex = i
				break
			}
		}
		if focusWindow != nil {
			if action == MouseLeftDown || action == MouseMiddleDown || action == MouseRightDown {
				for _, w := range wm.windows {
					if w != focusWindow {
						w.Blur()
					}
				}

				wm.windows = append(append(wm.windows[:focusWindowIndex], wm.windows[focusWindowIndex+1:]...), focusWindow)
			}

			return focusWindow.MouseHandler()(action, event, setFocus)
		}

		return consumed, nil
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

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

// GetBackgroundTransparent returns whether the background of this Window is transparent.
func (w *Window) GetBackgroundTransparent() bool {
	return w.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Window is transparent.
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

// SetFullscreen sets the flag indicating whether the the window should
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
