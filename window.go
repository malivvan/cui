package cui

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Window is a draggable, resizable frame around a primitive. Windows must be
// added to a WindowManager.
type Window struct {
	*Box

	primitive Primitive

	buttons []*WindowButton

	fullscreen bool

	normalX, normalY int
	normalW, normalH int

	dragX, dragY   int
	dragWX, dragWY int

	sync.RWMutex
}

// NewWindow returns a new window around the given primitive.
func NewWindow(primitive Primitive) *Window {
	w := &Window{
		Box:       NewBox(),
		primitive: primitive,
		dragWX:    -1,
		dragWY:    -1,
	}
	w.Box.focus = w
	return w
}

// SetFullscreen sets the flag indicating whether or not the the window should
// be drawn fullscreen.
func (w *Window) SetFullscreen(fullscreen bool) {
	w.Lock()
	defer w.Unlock()

	if w.fullscreen == fullscreen {
		return
	}

	w.fullscreen = fullscreen
	if w.fullscreen {
		w.normalX, w.normalY, w.normalW, w.normalH = w.GetRect()
	} else {
		w.SetRect(w.normalX, w.normalY, w.normalW, w.normalH)
	}
}

// Focus is called when this primitive receives focus.
func (w *Window) Focus(delegate func(p Primitive)) {
	w.Lock()
	defer w.Unlock()

	w.Box.Focus(delegate)

	w.primitive.Focus(delegate)
}

// Blur is called when this primitive loses focus.
func (w *Window) Blur() {
	w.Lock()
	defer w.Unlock()

	w.Box.Blur()

	w.primitive.Blur()
}

// HasFocus returns whether or not this primitive has focus.
func (w *Window) HasFocus() bool {
	w.RLock()
	defer w.RUnlock()

	focusable := w.primitive.GetFocusable()
	if focusable != nil {
		return focusable.HasFocus()
	}

	return w.Box.HasFocus()
}

// Draw draws this primitive onto the screen.
func (w *Window) Draw(screen tcell.Screen) {
	if !w.GetVisible() {
		return
	}

	w.RLock()
	defer w.RUnlock()

	w.Box.Draw(screen)

	x, y, width, height := w.GetInnerRect()
	w.primitive.SetRect(x, y, width, height)
	w.primitive.Draw(screen)

	if w.Box.GetBorder() {
		for _, button := range w.buttons {
			buttonX, buttonY := button.offsetX+x, button.offsetY+y
			if button.offsetX < 0 {
				buttonX += width
			}
			if button.offsetY < 0 {
				buttonY += height
			}

			// render the window title buttons
			Print(screen, []byte(Escape(fmt.Sprintf("[%c]", button.Symbol))), buttonX-1, buttonY, 9, 0, tcell.ColorYellow)
		}
	}

}

// InputHandler returns the handler for this primitive.
func (w *Window) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return w.primitive.InputHandler()
}

// MouseHandler returns the mouse handler for this primitive.
func (w *Window) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return w.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if !w.InRect(event.Position()) {
			return false, nil
		}

		if action == MouseLeftDown || action == MouseMiddleDown || action == MouseRightDown {
			setFocus(w)
		}

		if action == MouseLeftDown {
			x, y := event.Position()
			wx, wy, width, height := w.GetRect()
			mouseX, mouseY := event.Position()

			// check if any window button was pressed
			// if the window does not have border, it cannot receive button events
			if y == wy && w.border {
				for _, button := range w.buttons {
					if button.offsetX >= 0 && x == wx+button.offsetX || button.offsetX < 0 && x == wx+width+button.offsetX {
						if button.OnClick != nil {
							button.OnClick(w, button)
						}
						return true, nil
					}
				}
			}

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

////////////////////////////////////

// WindowButton represents a button on the window title bar
type WindowButton struct {
	Alignment int                              // AlignLeft or AlignRight
	Symbol    rune                             // icon for the button
	OnClick   func(w *Window, b *WindowButton) // callback to be invoked when the button is clicked

	offsetX, offsetY int
}

// AddButton adds a new window button to the title bar
func (w *Window) AddButton(symbol rune, alignment int, onclick func(w *Window, b *WindowButton)) *Window {
	w.Lock()
	defer w.Unlock()

	w.buttons = append(w.buttons, &WindowButton{
		Symbol:    symbol,
		Alignment: alignment,
		OnClick:   onclick,
	})

	offsetLeft, offsetRight := 2, -3
	for _, button := range w.buttons {
		if button.Alignment == AlignRight {
			button.offsetX = offsetRight
			offsetRight -= 3
		} else {
			button.offsetX = offsetLeft
			offsetLeft += 3
		}
	}

	return w
}

// RemoveButton removes the given button from the title bar
func (w *Window) RemoveButton(i int) *Window {
	w.Lock()
	defer w.Unlock()

	if i < 0 || i >= len(w.buttons) {
		return w
	}

	w.buttons = append(w.buttons[:i], w.buttons[i+1:]...)

	return w
}

// GetButton returns the given button
func (w *Window) GetButton(i int) *WindowButton {
	w.RLock()
	defer w.RUnlock()

	if i < 0 || i >= len(w.buttons) {
		return nil
	}

	return w.buttons[i]
}

// CountButtons returns the number of buttons in the window title bar
func (w *Window) CountButtons() int {
	w.RLock()
	defer w.RUnlock()

	return len(w.buttons)
}

// ClearButtons removes all buttons from the window title bar
func (w *Window) ClearButtons() *Window {
	w.Lock()
	defer w.Unlock()

	w.buttons = nil
	return w
}
