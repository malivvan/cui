package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// CheckBox implements a simple box for boolean values which can be checked and
// unchecked.
type CheckBox struct {
	*Box

	// Whether this box is checked.
	checked bool

	// The text to be displayed before the checkbox.
	label []byte

	// The text to be displayed after the checkbox.
	message []byte

	// The screen width of the label area. A value of 0 means use the width of
	// the label text.
	labelWidth int

	// The label color.
	labelColor tcell.Color

	// The label color when focused.
	labelColorFocused tcell.Color

	// The background color of the input area.
	fieldBackgroundColor tcell.Color

	// The background color of the input area when focused.
	fieldBackgroundColorFocused tcell.Color

	// The text color of the input area.
	fieldTextColor tcell.Color

	// The text color of the input area when focused.
	fieldTextColorFocused tcell.Color

	// An optional function which is called when the user changes the checked
	// state of this checkbox.
	changed func(checked bool)

	// An optional function which is called when the user indicated that they
	// are done entering text. The key which was pressed is provided (tab,
	// shift-tab, or escape).
	done func(tcell.Key)

	// A callback function set by the Form class and called when the user leaves
	// this form item.
	finished func(tcell.Key)

	// The rune to show when the checkbox is checked
	checkedRune rune

	// An optional rune to show within the checkbox when it is focused
	cursorRune rune

	mu sync.RWMutex
}

// NewCheckBox returns a new input field.
func NewCheckBox() *CheckBox {
	return &CheckBox{
		Box:                         NewBox(),
		labelColor:                  Styles.SecondaryTextColor,
		fieldBackgroundColor:        Styles.MoreContrastBackgroundColor,
		fieldBackgroundColorFocused: Styles.ContrastBackgroundColor,
		fieldTextColor:              Styles.PrimaryTextColor,
		checkedRune:                 Styles.CheckBoxCheckedRune,
		cursorRune:                  Styles.CheckBoxCursorRune,
		labelColorFocused:           ColorUnset,
		fieldTextColorFocused:       ColorUnset,
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (c *CheckBox) set(setter func(c *CheckBox)) *CheckBox {
	c.mu.Lock()
	setter(c)
	c.mu.Unlock()
	return c
}

func (c *CheckBox) get(getter func(c *CheckBox)) {
	c.mu.RLock()
	getter(c)
	c.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

func (c *CheckBox) SetTitle(title string) *CheckBox {
	c.Box.SetTitle(title)
	return c
}

func (c *CheckBox) SetTitleAlign(align int) *CheckBox {
	c.Box.SetTitleAlign(align)
	return c
}

func (c *CheckBox) SetBorder(show bool) *CheckBox {
	c.Box.SetBorder(show)
	return c
}

func (c *CheckBox) SetBorderColor(color tcell.Color) *CheckBox {
	c.Box.SetBorderColor(color)
	return c
}

func (c *CheckBox) SetBorderAttributes(attr tcell.AttrMask) *CheckBox {
	c.Box.SetBorderAttributes(attr)
	return c
}

func (c *CheckBox) SetBorderColorFocused(color tcell.Color) *CheckBox {
	c.Box.SetBorderColorFocused(color)
	return c
}

func (c *CheckBox) SetTitleColor(color tcell.Color) *CheckBox {
	c.Box.SetTitleColor(color)
	return c
}

func (c *CheckBox) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *CheckBox {
	c.Box.SetDrawFunc(handler)
	return c
}

func (c *CheckBox) ShowFocus(showFocus bool) *CheckBox {
	c.Box.ShowFocus(showFocus)
	return c
}

func (c *CheckBox) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *CheckBox {
	c.Box.SetMouseCapture(capture)
	return c
}

func (c *CheckBox) SetBackgroundColor(color tcell.Color) *CheckBox {
	c.Box.SetBackgroundColor(color)
	return c
}

func (c *CheckBox) SetBackgroundTransparent(transparent bool) *CheckBox {
	c.Box.SetBackgroundTransparent(transparent)
	return c
}

func (c *CheckBox) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *CheckBox {
	c.Box.SetInputCapture(capture)
	return c
}

func (c *CheckBox) SetPadding(top, bottom, left, right int) *CheckBox {
	c.Box.SetPadding(top, bottom, left, right)
	return c
}

/////////////////////////////////////// <API> ///////////////////////////////////////

// SetChecked sets the state of the checkbox.
func (c *CheckBox) SetChecked(checked bool) *CheckBox {
	return c.set(func(c *CheckBox) { c.checked = checked })
}

// SetCheckedRune sets the rune to show when the checkbox is checked.
func (c *CheckBox) SetCheckedRune(r rune) *CheckBox {
	return c.set(func(c *CheckBox) { c.checkedRune = r })
}

// SetCursorRune sets the rune to show within the checkbox when it is focused.
func (c *CheckBox) SetCursorRune(r rune) *CheckBox {
	return c.set(func(c *CheckBox) { c.cursorRune = r })
}

// IsChecked returns whether the box is checked.
func (c *CheckBox) IsChecked() (checked bool) {
	c.get(func(c *CheckBox) { checked = c.checked })
	return
}

// SetLabel sets the text to be displayed before the input area.
func (c *CheckBox) SetLabel(label string) *CheckBox {
	return c.set(func(c *CheckBox) { c.label = []byte(label) })
}

// GetLabel returns the text to be displayed before the input area.
func (c *CheckBox) GetLabel() (label string) {
	c.get(func(c *CheckBox) { label = string(c.label) })
	return
}

// SetMessage sets the text to be displayed after the checkbox
func (c *CheckBox) SetMessage(message string) *CheckBox {
	return c.set(func(c *CheckBox) { c.message = []byte(message) })
}

// GetMessage returns the text to be displayed after the checkbox
func (c *CheckBox) GetMessage() (msg string) {
	c.get(func(c *CheckBox) { msg = string(c.message) })
	return
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (c *CheckBox) SetLabelWidth(width int) *CheckBox {
	return c.set(func(c *CheckBox) { c.labelWidth = width })
}

// SetLabelColor sets the color of the label.
func (c *CheckBox) SetLabelColor(color tcell.Color) *CheckBox {
	return c.set(func(c *CheckBox) { c.labelColor = color })
}

// SetLabelColorFocused sets the color of the label when focused.
func (c *CheckBox) SetLabelColorFocused(color tcell.Color) *CheckBox {
	return c.set(func(c *CheckBox) { c.labelColorFocused = color })
}

// SetFieldBackgroundColor sets the background color of the input area.
func (c *CheckBox) SetFieldBackgroundColor(color tcell.Color) *CheckBox {
	return c.set(func(c *CheckBox) { c.fieldBackgroundColor = color })
}

// SetFieldBackgroundColorFocused sets the background color of the input area when focused.
func (c *CheckBox) SetFieldBackgroundColorFocused(color tcell.Color) *CheckBox {
	return c.set(func(c *CheckBox) { c.fieldBackgroundColorFocused = color })
}

// SetFieldTextColor sets the text color of the input area.
func (c *CheckBox) SetFieldTextColor(color tcell.Color) *CheckBox {
	return c.set(func(c *CheckBox) { c.fieldTextColor = color })
}

// SetFieldTextColorFocused sets the text color of the input area when focused.
func (c *CheckBox) SetFieldTextColorFocused(color tcell.Color) *CheckBox {
	return c.set(func(c *CheckBox) { c.fieldTextColorFocused = color })
}

// GetFieldHeight returns the height of the field.
func (c *CheckBox) GetFieldHeight() int {
	return 1
}

// GetFieldWidth returns this primitive's field width.
func (c *CheckBox) GetFieldWidth() (w int) {
	c.get(func(c *CheckBox) {
		if len(c.message) == 0 {
			w = 1
		}
		w = 2 + len(c.message)
	})
	return
}

// SetChangedFunc sets a handler which is called when the checked state of this
// checkbox was changed by the user. The handler function receives the new
// state.
func (c *CheckBox) SetChangedFunc(handler func(checked bool)) *CheckBox {
	return c.set(func(c *CheckBox) { c.changed = handler })
}

// SetDoneFunc sets a handler which is called when the user is done using the
// checkbox. The callback function is provided with the key that was pressed,
// which is one of the following:
//
//   - KeyEscape: Abort text input.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (c *CheckBox) SetDoneFunc(handler func(key tcell.Key)) *CheckBox {
	return c.set(func(c *CheckBox) { c.done = handler })
}

// SetFinishedFunc sets a callback invoked when the user leaves this form item.
func (c *CheckBox) SetFinishedFunc(handler func(key tcell.Key)) *CheckBox {
	return c.set(func(c *CheckBox) { c.finished = handler })
}

////////////////////////////////////// <WIDGET> ///////////////////////////////////////

// Draw draws this primitive onto the screen.
func (c *CheckBox) Draw(screen tcell.Screen) {
	if !c.GetVisible() {
		return
	}

	c.Box.Draw(screen)

	c.mu.Lock()
	defer c.mu.Unlock()

	hasFocus := c.GetFocusable().HasFocus()

	// Select colors
	labelColor := c.labelColor
	fieldBackgroundColor := c.fieldBackgroundColor
	fieldTextColor := c.fieldTextColor
	if hasFocus {
		if c.labelColorFocused != ColorUnset {
			labelColor = c.labelColorFocused
		}
		if c.fieldBackgroundColorFocused != ColorUnset {
			fieldBackgroundColor = c.fieldBackgroundColorFocused
		}
		if c.fieldTextColorFocused != ColorUnset {
			fieldTextColor = c.fieldTextColorFocused
		}
	}

	// Prepare
	x, y, width, height := c.GetInnerRect()
	rightLimit := x + width
	if height < 1 || rightLimit <= x {
		return
	}

	// Draw label.
	if c.labelWidth > 0 {
		labelWidth := c.labelWidth
		if labelWidth > rightLimit-x {
			labelWidth = rightLimit - x
		}
		Print(screen, c.label, x, y, labelWidth, AlignLeft, labelColor)
		x += labelWidth
	} else {
		_, drawnWidth := Print(screen, c.label, x, y, rightLimit-x, AlignLeft, labelColor)
		x += drawnWidth
	}

	// Draw checkbox.
	fieldStyle := tcell.StyleDefault.Background(fieldBackgroundColor).Foreground(fieldTextColor)

	checkedRune := c.checkedRune
	if !c.checked {
		checkedRune = ' '
	}
	rightRune := ' '
	if c.cursorRune != 0 && hasFocus {
		rightRune = c.cursorRune
	}
	screen.SetContent(x, y, ' ', nil, fieldStyle)
	screen.SetContent(x+1, y, checkedRune, nil, fieldStyle)
	screen.SetContent(x+2, y, rightRune, nil, fieldStyle)

	if len(c.message) > 0 {
		Print(screen, c.message, x+4, y, len(c.message), AlignLeft, labelColor)
	}
}

// InputHandler returns the handler for this primitive.
func (c *CheckBox) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return c.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		if HitShortcut(event, Keys.Select, Keys.Select2) {
			c.mu.Lock()
			c.checked = !c.checked
			c.mu.Unlock()
			if c.changed != nil {
				c.changed(c.checked)
			}
		} else if HitShortcut(event, Keys.Cancel, Keys.MovePreviousField, Keys.MoveNextField) {
			if c.done != nil {
				c.done(event.Key())
			}
			if c.finished != nil {
				c.finished(event.Key())
			}
		}
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (c *CheckBox) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return c.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		x, y := event.Position()
		_, rectY, _, _ := c.GetInnerRect()
		if !c.InRect(x, y) {
			return false, nil
		}

		// Process mouse event.
		if action == MouseLeftClick && y == rectY {
			setFocus(c)
			c.checked = !c.checked
			if c.changed != nil {
				c.changed(c.checked)
			}
			consumed = true
		}

		return
	})
}
