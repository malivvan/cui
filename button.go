package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Button is labeled box that triggers an action when selected.
type Button struct {
	box *Box

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
	box.SetRect(0, 0, 0, 1)
	box.SetBackgroundColor(Styles.MoreContrastBackgroundColor)
	return &Button{
		box:                    box,
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

// GetTitle returns the title of this Button.
func (b *Button) GetTitle() string {
	return b.box.GetTitle()
}

// SetTitle sets the title of this Button.
func (b *Button) SetTitle(title string) *Button {
	b.box.SetTitle(title)
	return b
}

// GetTitleAlign returns the title alignment of this Button.
func (b *Button) GetTitleAlign() int {
	return b.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Button.
func (b *Button) SetTitleAlign(align int) *Button {
	b.box.SetTitleAlign(align)
	return b
}

// GetBorder returns whether this Button has a border.
func (b *Button) GetBorder() bool {
	return b.box.GetBorder()
}

// SetBorder sets whether this Button has a border.
func (b *Button) SetBorder(show bool) *Button {
	b.box.SetBorder(show)
	return b
}

// GetBorderColor returns the border color of this Button.
func (b *Button) GetBorderColor() tcell.Color {
	return b.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Button.
func (b *Button) SetBorderColor(color tcell.Color) *Button {
	b.box.SetBorderColor(color)
	return b
}

// GetBorderAttributes returns the border attributes of this Button.
func (b *Button) GetBorderAttributes() tcell.AttrMask {
	return b.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Button.
func (b *Button) SetBorderAttributes(attr tcell.AttrMask) *Button {
	b.box.SetBorderAttributes(attr)
	return b
}

// GetBorderColorFocused returns the border color of this Button when focused.
func (b *Button) GetBorderColorFocused() tcell.Color {
	return b.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Button when focused.
func (b *Button) SetBorderColorFocused(color tcell.Color) *Button {
	b.box.SetBorderColorFocused(color)
	return b
}

// GetTitleColor returns the title color of this Button.
func (b *Button) GetTitleColor() tcell.Color {
	return b.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Button.
func (b *Button) SetTitleColor(color tcell.Color) *Button {
	b.box.SetTitleColor(color)
	return b
}

// GetDrawFunc returns the custom draw function of this Button.
func (b *Button) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return b.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Button.
func (b *Button) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Button {
	b.box.SetDrawFunc(handler)
	return b
}

// ShowFocus sets whether this Button should show a focus indicator when focused.
func (b *Button) ShowFocus(showFocus bool) *Button {
	b.box.ShowFocus(showFocus)
	return b
}

// GetMouseCapture returns the mouse capture function of this Button.
func (b *Button) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return b.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Button.
func (b *Button) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Button {
	b.box.SetMouseCapture(capture)
	return b
}

// GetBackgroundColor returns the background color of this Button.
func (b *Button) GetBackgroundColor() tcell.Color {
	return b.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Button.
func (b *Button) SetBackgroundColor(color tcell.Color) *Button {
	b.box.SetBackgroundColor(color)
	return b
}

// GetBackgroundTransparent returns whether the background of this Button is transparent.
func (b *Button) GetBackgroundTransparent() bool {
	return b.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Button is transparent.
func (b *Button) SetBackgroundTransparent(transparent bool) *Button {
	b.box.SetBackgroundTransparent(transparent)
	return b
}

// GetInputCapture returns the input capture function of this Button.
func (b *Button) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return b.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Button.
func (b *Button) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Button {
	b.box.SetInputCapture(capture)
	return b
}

// GetPadding returns the padding of this Button.
func (b *Button) GetPadding() (top, bottom, left, right int) {
	return b.box.GetPadding()
}

// SetPadding sets the padding of this Button.
func (b *Button) SetPadding(top, bottom, left, right int) *Button {
	b.box.SetPadding(top, bottom, left, right)
	return b
}

// InRect returns whether the given screen coordinates are within this Button.
func (b *Button) InRect(x, y int) bool {
	return b.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Button.
func (b *Button) GetInnerRect() (x, y, width, height int) {
	return b.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Button is preserved.
func (b *Button) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return b.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Button is preserved.
func (b *Button) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return b.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Button.
func (b *Button) GetRect() (x, y, width, height int) {
	return b.box.GetRect()
}

// SetRect sets the rectangle occupied by this Button.
func (b *Button) SetRect(x, y, width, height int) {
	b.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Button is visible.
func (b *Button) GetVisible() bool {
	return b.box.GetVisible()
}

// SetVisible sets whether this Button is visible.
func (b *Button) SetVisible(visible bool) {
	b.box.SetVisible(visible)
}

// Focus is called when this Button receives focus.
func (b *Button) Focus(delegate func(p Widget)) {
	b.box.Focus(delegate)
}

// HasFocus returns whether this Button has focus.
func (b *Button) HasFocus() bool {
	return b.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Button.
func (b *Button) GetFocusable() Focusable {
	return b.box.GetFocusable()
}

// Blur is called when this Button loses focus.
func (b *Button) Blur() {
	b.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

// SetLabel sets the button text.
func (b *Button) SetLabel(label string) *Button {
	b.box.SetRect(0, 0, TaggedStringWidth(label)+4, 1)
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
	borderColor := b.box.borderColor
	backgroundColor := b.box.backgroundColor
	if b.box.focus.HasFocus() {
		b.box.backgroundColor = b.backgroundColorFocused
		b.box.borderColor = b.labelColorFocused
		defer func() {
			b.box.borderColor = borderColor
		}()
	}
	b.mu.Unlock()
	b.box.Draw(screen)
	b.mu.Lock()
	b.box.backgroundColor = backgroundColor

	// Draw label.
	x, y, width, height := b.GetInnerRect()
	if width > 0 && height > 0 {
		y = y + height/2
		labelColor := b.labelColor
		if b.box.focus.HasFocus() {
			labelColor = b.labelColorFocused
		}
		_, pw := Print(screen, b.label, x, y, width, AlignCenter, labelColor)

		// Draw cursor.
		if b.box.focus.HasFocus() && b.cursorRune != 0 {
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
