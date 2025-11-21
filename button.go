package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Button is labeled box that triggers an action when selected.
type Button struct {
	*Box

	// The text to be displayed before the input area.
	label []byte

	// The label color.
	labelColor tcell.Color

	// The label color when the button is in focus.
	labelColorFocused tcell.Color

	// The background color when the button is in focus.
	backgroundColorFocused tcell.Color

	// An optional function which is called when the button was selected.
	selected func()

	// An optional function which is called when the user leaves the button. A
	// key is provided indicating which key was pressed to leave (tab or backtab).
	blur func(tcell.Key)

	// An optional rune which is drawn after the label when the button is focused.
	cursorRune rune

	mu sync.RWMutex
}

// NewButton returns a new input field.
func NewButton() *Button {
	box := NewBox()
	box.SetRect(0, 0, TaggedStringWidth(label)+4, 1)
	box.SetBackgroundColor(Styles.MoreContrastBackgroundColor)
	return &Button{
		Box:                    box,
		labelColor:             Styles.PrimaryTextColor,
		labelColorFocused:      Styles.PrimaryTextColor,
		cursorRune:             Styles.ButtonCursorRune,
		backgroundColorFocused: Styles.ContrastBackgroundColor,
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (b *Button) set(setter func(b *Button)) *Button {
	b.mu.Lock()
	setter(b)
	b.mu.Unlock()
	return b
}

func (b *Button) get(getter func(b *Button)) {
	b.mu.RLock()
	getter(b)
	b.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// SetTitle sets the title of the Button.
func (b *Button) SetTitle(title string) *Button {
	b.Box.SetTitle(title)
	return b
}

// SetTitleColor sets the title color of the Button.
func (b *Button) SetTitleColor(color tcell.Color) *Button {
	b.Box.SetTitleColor(color)
	return b
}

// SetTitleAlign sets the title alignment of the Button.
func (b *Button) SetTitleAlign(align int) *Button {
	b.Box.SetTitleAlign(align)
	return b
}

// SetPadding sets the padding of the Button.
func (b *Button) SetPadding(top, bottom, left, right int) *Button {
	b.Box.SetPadding(top, bottom, left, right)
	return b
}

// SetBorder sets whether the Button has a border.
func (b *Button) SetBorder(show bool) *Button {
	b.Box.SetBorder(show)
	return b
}

// SetBorderColor sets the border color of the Button.
func (b *Button) SetBorderColor(color tcell.Color) *Button {
	b.Box.SetBorderColor(color)
	return b
}

// SetBorderColorFocused sets the border color of the Button when focused.
func (b *Button) SetBorderColorFocused(color tcell.Color) *Button {
	b.Box.SetBorderColorFocused(color)
	return b
}

// SetBorderAttributes sets the border attributes of the Button.
func (b *Button) SetBorderAttributes(attr tcell.AttrMask) *Button {
	b.Box.SetBorderAttributes(attr)
	return b
}

func (b *Button) SetBackgroundColor(color tcell.Color) *Button {
	b.Box.SetBackgroundColor(color)
	return b
}

func (b *Button) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Button {
	b.Box.SetDrawFunc(handler)
	return b
}

func (b *Button) ShowFocus(showFocus bool) *Button {
	b.Box.ShowFocus(showFocus)
	return b
}

func (b *Button) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Button {
	b.Box.SetMouseCapture(capture)
	return b
}

func (b *Button) SetBackgroundTransparent(transparent bool) *Button {
	b.Box.SetBackgroundTransparent(transparent)
	return b
}

func (b *Button) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Button {
	b.Box.SetInputCapture(capture)
	return b
}

////////////////////////////////// <API> ////////////////////////////////////

// SetLabel sets the button text.
func (b *Button) SetLabel(label string) *Button {
	return b.set(func(b *Button) { b.label = []byte(label) })
}

// GetLabel returns the button text.
func (b *Button) GetLabel() (label string) {
	b.get(func(b *Button) { label = string(b.label) })
	return
}

// SetLabelColor sets the color of the button text.
func (b *Button) SetLabelColor(color tcell.Color) *Button {
	return b.set(func(b *Button) { b.labelColor = color })
}

// SetLabelColorFocused sets the color of the button text when the button is
// in focus.
func (b *Button) SetLabelColorFocused(color tcell.Color) *Button {
	return b.set(func(b *Button) { b.labelColorFocused = color })
}

// SetCursorRune sets the rune to show within the button when it is focused.
func (b *Button) SetCursorRune(rune rune) *Button {
	return b.set(func(b *Button) { b.cursorRune = rune })
}

// SetBackgroundColorFocused sets the background color of the button text when
// the button is in focus.
func (b *Button) SetBackgroundColorFocused(color tcell.Color) *Button {
	return b.set(func(b *Button) { b.backgroundColorFocused = color })
}

// SetSelectedFunc sets a handler which is called when the button was selected.
func (b *Button) SetSelectedFunc(handler func()) *Button {
	return b.set(func(b *Button) { b.selected = handler })
}

// SetBlurFunc sets a handler which is called when the user leaves the button.
// The callback function is provided with the key that was pressed, which is one
// of the following:
//
//   - KeyEscape: Leaving the button with no specific direction.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (b *Button) SetBlurFunc(handler func(key tcell.Key)) *Button {
	return b.set(func(b *Button) { b.blur = handler })
}

// ////////////////////////////////// <WIDGET> ////////////////////////////////////
var _ Widget = (*Button)(nil)

// Draw draws this primitive onto the screen.
func (b *Button) Draw(screen tcell.Screen) {
	if !b.GetVisible() {
		return
	}

	b.mu.Lock()
	defer b.mu.Unlock()

	// Draw the box.
	borderColor := b.borderColor
	backgroundColor := b.backgroundColor
	if b.focus.HasFocus() {
		b.backgroundColor = b.backgroundColorFocused
		b.borderColor = b.labelColorFocused
		defer func() {
			b.borderColor = borderColor
		}()
	}
	b.mu.Unlock()
	b.Box.Draw(screen)
	b.mu.Lock()
	b.backgroundColor = backgroundColor

	// Draw label.
	x, y, width, height := b.GetInnerRect()
	if width > 0 && height > 0 {
		y = y + height/2
		labelColor := b.labelColor
		if b.focus.HasFocus() {
			labelColor = b.labelColorFocused
		}
		_, pw := Print(screen, b.label, x, y, width, AlignCenter, labelColor)

		// Draw cursor.
		if b.focus.HasFocus() && b.cursorRune != 0 {
			cursorX := x + int(float64(width)/2+float64(pw)/2)
			if cursorX > x+width-1 {
				cursorX = x + width - 1
			} else if cursorX < x+width {
				cursorX++
			}
			Print(screen, []byte(string(b.cursorRune)), cursorX, y, width, AlignLeft, labelColor)
		}
	}
}

// InputHandler returns the handler for this primitive.
func (b *Button) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return b.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		// Process key event.
		if HitShortcut(event, Keys.Select, Keys.Select2) {
			if b.selected != nil {
				b.selected()
			}
		} else if HitShortcut(event, Keys.Cancel, Keys.MovePreviousField, Keys.MoveNextField) {
			if b.blur != nil {
				b.blur(event.Key())
			}
		}
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (b *Button) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return b.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if !b.InRect(event.Position()) {
			return false, nil
		}

		// Process mouse event.
		if action == MouseLeftClick {
			setFocus(b)
			if b.selected != nil {
				b.selected()
			}
			consumed = true
		}

		return
	})
}
