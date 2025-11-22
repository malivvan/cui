package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Modal is a centered message window used to inform the user or prompt them
// for an immediate decision. It needs to have at least one button (added via
// AddButtons) or it will never disappear. You may change the title and
// appearance of the window by modifying the Frame returned by GetFrame. You
// may include additional elements within the window by modifying the Form
// returned by GetForm.
type Modal struct {
	box *Box

	// The Frame embedded in the Modal.
	frame *Frame

	// The Form embedded in the Modal's Frame.
	form *Form

	// The message text (original, not word-wrapped).
	text string

	// The text color.
	textColor tcell.Color

	// The text alignment.
	textAlign int

	// The optional callback for when the user clicked one of the buttons. It
	// receives the index of the clicked button and the button's label.
	done func(buttonIndex int, buttonLabel string)

	mu sync.RWMutex
}

// NewModal returns a new centered message window.
func NewModal() *Modal {
	m := &Modal{
		box:       NewBox(),
		textColor: Styles.PrimaryTextColor,
		textAlign: AlignCenter,
	}

	m.form = NewForm()
	m.form.SetButtonsAlign(AlignCenter)
	m.form.SetPadding(0, 0, 0, 0)
	m.form.SetCancelFunc(func() {
		if m.done != nil {
			m.done(-1, "")
		}
	})

	m.frame = NewFrame().SetWidget(m.form)
	m.frame.SetBorder(true)
	m.frame.SetBorders(0, 0, 1, 0, 0, 0)
	m.frame.SetPadding(1, 1, 1, 1)

	m.box.focus = m
	return m
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (m *Modal) set(setter func(m *Modal)) *Modal {
	m.mu.Lock()
	setter(m)
	m.mu.Unlock()
	return m
}

func (m *Modal) get(getter func(m *Modal)) {
	m.mu.RLock()
	getter(m)
	m.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Modal.
func (m *Modal) GetTitle() string {
	return m.box.GetTitle()
}

// SetTitle sets the title of this Modal.
func (m *Modal) SetTitle(title string) *Modal {
	m.box.SetTitle(title)
	return m
}

// GetTitleAlign returns the title alignment of this Modal.
func (m *Modal) GetTitleAlign() int {
	return m.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Modal.
func (m *Modal) SetTitleAlign(align int) *Modal {
	m.box.SetTitleAlign(align)
	return m
}

// GetBorder returns whether this Modal has a border.
func (m *Modal) GetBorder() bool {
	return m.box.GetBorder()
}

// SetBorder sets whether this Modal has a border.
func (m *Modal) SetBorder(show bool) *Modal {
	m.box.SetBorder(show)
	return m
}

// GetBorderColor returns the border color of this Modal.
func (m *Modal) GetBorderColor() tcell.Color {
	return m.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Modal.
func (m *Modal) SetBorderColor(color tcell.Color) *Modal {
	m.box.SetBorderColor(color)
	return m
}

// GetBorderAttributes returns the border attributes of this Modal.
func (m *Modal) GetBorderAttributes() tcell.AttrMask {
	return m.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Modal.
func (m *Modal) SetBorderAttributes(attr tcell.AttrMask) *Modal {
	m.box.SetBorderAttributes(attr)
	return m
}

// GetBorderColorFocused returns the border color of this Modal when focusel.
func (m *Modal) GetBorderColorFocused() tcell.Color {
	return m.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Modal when focusel.
func (m *Modal) SetBorderColorFocused(color tcell.Color) *Modal {
	m.box.SetBorderColorFocused(color)
	return m
}

// GetTitleColor returns the title color of this Modal.
func (m *Modal) GetTitleColor() tcell.Color {
	return m.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Modal.
func (m *Modal) SetTitleColor(color tcell.Color) *Modal {
	m.box.SetTitleColor(color)
	return m
}

// GetDrawFunc returns the custom draw function of this Modal.
func (m *Modal) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return m.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Modal.
func (m *Modal) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Modal {
	m.box.SetDrawFunc(handler)
	return m
}

// ShowFocus sets whether this Modal should show a focus indicator when focusel.
func (m *Modal) ShowFocus(showFocus bool) *Modal {
	m.box.ShowFocus(showFocus)
	return m
}

// GetMouseCapture returns the mouse capture function of this Modal.
func (m *Modal) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return m.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Modal.
func (m *Modal) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Modal {
	m.box.SetMouseCapture(capture)
	return m
}

// GetBackgroundColor returns the background color of this Modal.
func (m *Modal) GetBackgroundColor() tcell.Color {
	return m.box.GetBackgroundColor()
}

// SetBackgroundColor sets the color of the Modal Frame background.
func (m *Modal) SetBackgroundColor(color tcell.Color) *Modal {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.form.SetBackgroundColor(color)
	m.frame.SetBackgroundColor(color)
	return m
}

// GetBackgroundTransparent returns whether the background of this Modal is transparent.
func (m *Modal) GetBackgroundTransparent() bool {
	return m.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Modal is transparent.
func (m *Modal) SetBackgroundTransparent(transparent bool) *Modal {
	m.box.SetBackgroundTransparent(transparent)
	return m
}

// GetInputCapture returns the input capture function of this Modal.
func (m *Modal) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return m.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Modal.
func (m *Modal) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Modal {
	m.box.SetInputCapture(capture)
	return m
}

// GetPadding returns the padding of this Modal.
func (m *Modal) GetPadding() (top, bottom, left, right int) {
	return m.box.GetPadding()
}

// SetPadding sets the padding of this Modal.
func (m *Modal) SetPadding(top, bottom, left, right int) *Modal {
	m.box.SetPadding(top, bottom, left, right)
	return m
}

// InRect returns whether the given screen coordinates are within this Modal.
func (m *Modal) InRect(x, y int) bool {
	return m.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Modal.
func (m *Modal) GetInnerRect() (x, y, width, height int) {
	return m.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Modal is preservel.
func (m *Modal) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return m.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Modal is preservel.
func (m *Modal) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return m.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Modal.
func (m *Modal) GetRect() (x, y, width, height int) {
	return m.box.GetRect()
}

// SetRect sets the rectangle occupied by this Modal.
func (m *Modal) SetRect(x, y, width, height int) {
	m.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Modal is visible.
func (m *Modal) GetVisible() bool {
	return m.box.GetVisible()
}

// SetVisible sets whether this Modal is visible.
func (m *Modal) SetVisible(visible bool) {
	m.box.SetVisible(visible)
}

// Focus is called when this primitive receives focus.
func (m *Modal) Focus(delegate func(p Widget)) {
	delegate(m.form)
}

// HasFocus returns whether or not this primitive has focus.
func (m *Modal) HasFocus() bool {
	return m.GetForm().HasFocus()
}

// GetFocusable returns the focusable primitive of this Modal.
func (m *Modal) GetFocusable() Focusable {
	return m.box.GetFocusable()
}

// Blur is called when this Modal loses focus.
func (m *Modal) Blur() {
	m.box.Blur()
}

// SetTextColor sets the color of the message text.
func (m *Modal) SetTextColor(color tcell.Color) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.textColor = color
}

// SetButtonBackgroundColor sets the background color of the buttons.
func (m *Modal) SetButtonBackgroundColor(color tcell.Color) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.form.SetButtonBackgroundColor(color)
}

// SetButtonTextColor sets the color of the button texts.
func (m *Modal) SetButtonTextColor(color tcell.Color) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.form.SetButtonTextColor(color)
}

// SetButtonsAlign sets the horizontal alignment of the buttons. This must be
// either AlignLeft, AlignCenter (the default), or AlignRight.
func (m *Modal) SetButtonsAlign(align int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.form.SetButtonsAlign(align)
}

// SetDoneFunc sets a handler which is called when one of the buttons was
// pressed. It receives the index of the button as well as its label text. The
// handler is also called when the user presses the Escape key. The index will
// then be negative and the label text an empty string.
func (m *Modal) SetDoneFunc(handler func(buttonIndex int, buttonLabel string)) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.done = handler
}

// SetText sets the message text of the window. The text may contain line
// breaks. Note that words are wrapped, too, based on the final size of the
// window.
func (m *Modal) SetText(text string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.text = text
}

// SetTextAlign sets the horizontal alignment of the text. This must be either
// AlignLeft, AlignCenter (the default), or AlignRight.
func (m *Modal) SetTextAlign(align int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.textAlign = align
}

// GetForm returns the Form embedded in the window. The returned Form may be
// modified to include additional elements (e.g. AddInputField, AddFormItem).
func (m *Modal) GetForm() *Form {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.form
}

// GetFrame returns the Frame embedded in the window.
func (m *Modal) GetFrame() *Frame {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.frame
}

// AddButtons adds buttons to the window. There must be at least one button and
// a "done" handler so the window can be closed again.
func (m *Modal) AddButtons(labels []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for index, label := range labels {
		func(i int, l string) {
			m.form.AddButton(label, func() {
				if m.done != nil {
					m.done(i, l)
				}
			})
			button := m.form.GetButton(m.form.GetButtonCount() - 1)
			button.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
				switch event.Key() {
				case tcell.KeyDown, tcell.KeyRight:
					return tcell.NewEventKey(tcell.KeyTab, 0, tcell.ModNone)
				case tcell.KeyUp, tcell.KeyLeft:
					return tcell.NewEventKey(tcell.KeyBacktab, 0, tcell.ModNone)
				}
				return event
			})
		}(index, label)
	}
}

// ClearButtons removes all buttons from the window.
func (m *Modal) ClearButtons() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.form.ClearButtons()
}

// SetFocus shifts the focus to the button with the given index.
func (m *Modal) SetFocus(index int) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.form.SetFocus(index)
}

// Draw draws this primitive onto the screen.
func (m *Modal) Draw(screen tcell.Screen) {
	if !m.GetVisible() {
		return
	}

	formItemCount := m.form.GetFormItemCount()

	m.mu.Lock()
	defer m.mu.Unlock()

	// Calculate the width of this Modal.
	buttonsWidth := 0
	for _, button := range m.form.buttons {
		buttonsWidth += TaggedTextWidth(button.label) + 4 + 2
	}
	buttonsWidth -= 2
	screenWidth, screenHeight := screen.Size()
	width := screenWidth / 3
	if width < buttonsWidth {
		width = buttonsWidth
	}
	// width is now without the box border.

	// Reset the text and find out how wide it is.
	m.frame.Clear()
	lines := WordWrap(m.text, width)
	for _, line := range lines {
		m.frame.AddText(line, true, m.textAlign, m.textColor)
	}

	// Set the Modal's position and size.
	height := len(lines) + (formItemCount * 2) + 6
	width += 4
	x := (screenWidth - width) / 2
	y := (screenHeight - height) / 2
	m.SetRect(x, y, width, height)

	// Draw the frame.
	m.frame.SetRect(x, y, width, height)
	m.frame.Draw(screen)
}

func (m *Modal) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return m.box.InputHandler()
}

// MouseHandler returns the mouse handler for this primitive.
func (m *Modal) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return m.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		// Pass mouse events on to the form.
		consumed, capture = m.form.MouseHandler()(action, event, setFocus)
		if !consumed && action == MouseLeftClick && m.InRect(event.Position()) {
			setFocus(m.box)
			consumed = true
		}
		return
	})
}
