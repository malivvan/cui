package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// frameText holds information about a line of text shown in the frame.
type frameText struct {
	Text   string      // The text to be displayed.
	Header bool        // true = place in header, false = place in footer.
	Align  int         // One of the Align constants.
	Color  tcell.Color // The text color.
}

// Frame is a wrapper which adds space around another widget. In addition,
// the top area (header) and the bottom area (footer) may also contain text.
type Frame struct {
	box *Box

	// The contained widget.
	widget Widget

	// The lines of text to be displayed.
	text []*frameText

	// Border spacing.
	top, bottom, header, footer, left, right int

	mu sync.RWMutex
}

// NewFrame returns a new frame around the given widget. The widget's
// size will be changed to fit within this frame.
func NewFrame() *Frame {
	f := &Frame{
		box:    NewBox(),
		top:    1,
		bottom: 1,
		header: 1,
		footer: 1,
		left:   1,
		right:  1,
	}
	f.box.focus = f
	return f
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (f *Frame) set(setter func(f *Frame)) *Frame {
	f.mu.Lock()
	setter(f)
	f.mu.Unlock()
	return f
}

func (f *Frame) get(getter func(f *Frame)) {
	f.mu.RLock()
	getter(f)
	f.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Frame.
func (f *Frame) GetTitle() string {
	return f.box.GetTitle()
}

// SetTitle sets the title of this Frame.
func (f *Frame) SetTitle(title string) *Frame {
	f.box.SetTitle(title)
	return f
}

// GetTitleAlign returns the title alignment of this Frame.
func (f *Frame) GetTitleAlign() int {
	return f.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Frame.
func (f *Frame) SetTitleAlign(align int) *Frame {
	f.box.SetTitleAlign(align)
	return f
}

// GetBorder returns whether this Frame has a border.
func (f *Frame) GetBorder() bool {
	return f.box.GetBorder()
}

// SetBorder sets whether this Frame has a border.
func (f *Frame) SetBorder(show bool) *Frame {
	f.box.SetBorder(show)
	return f
}

// GetBorderColor returns the border color of this Frame.
func (f *Frame) GetBorderColor() tcell.Color {
	return f.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Frame.
func (f *Frame) SetBorderColor(color tcell.Color) *Frame {
	f.box.SetBorderColor(color)
	return f
}

// GetBorderAttributes returns the border attributes of this Frame.
func (f *Frame) GetBorderAttributes() tcell.AttrMask {
	return f.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Frame.
func (f *Frame) SetBorderAttributes(attr tcell.AttrMask) *Frame {
	f.box.SetBorderAttributes(attr)
	return f
}

// GetBorderColorFocused returns the border color of this Frame when focusef.
func (f *Frame) GetBorderColorFocused() tcell.Color {
	return f.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Frame when focusef.
func (f *Frame) SetBorderColorFocused(color tcell.Color) *Frame {
	f.box.SetBorderColorFocused(color)
	return f
}

// GetTitleColor returns the title color of this Frame.
func (f *Frame) GetTitleColor() tcell.Color {
	return f.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Frame.
func (f *Frame) SetTitleColor(color tcell.Color) *Frame {
	f.box.SetTitleColor(color)
	return f
}

// GetDrawFunc returns the custom draw function of this Frame.
func (f *Frame) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return f.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Frame.
func (f *Frame) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Frame {
	f.box.SetDrawFunc(handler)
	return f
}

// ShowFocus sets whether this Frame should show a focus indicator when focusef.
func (f *Frame) ShowFocus(showFocus bool) *Frame {
	f.box.ShowFocus(showFocus)
	return f
}

// GetMouseCapture returns the mouse capture function of this Frame.
func (f *Frame) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return f.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Frame.
func (f *Frame) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Frame {
	f.box.SetMouseCapture(capture)
	return f
}

// GetBackgroundColor returns the background color of this Frame.
func (f *Frame) GetBackgroundColor() tcell.Color {
	return f.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Frame.
func (f *Frame) SetBackgroundColor(color tcell.Color) *Frame {
	f.box.SetBackgroundColor(color)
	return f
}

// GetBackgroundTransparent returns whether the background of this Frame is transparent.
func (f *Frame) GetBackgroundTransparent() bool {
	return f.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Frame is transparent.
func (f *Frame) SetBackgroundTransparent(transparent bool) *Frame {
	f.box.SetBackgroundTransparent(transparent)
	return f
}

// GetInputCapture returns the input capture function of this Frame.
func (f *Frame) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return f.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Frame.
func (f *Frame) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Frame {
	f.box.SetInputCapture(capture)
	return f
}

// GetPadding returns the padding of this Frame.
func (f *Frame) GetPadding() (top, bottom, left, right int) {
	return f.box.GetPadding()
}

// SetPadding sets the padding of this Frame.
func (f *Frame) SetPadding(top, bottom, left, right int) *Frame {
	f.box.SetPadding(top, bottom, left, right)
	return f
}

// InRect returns whether the given screen coordinates are within this Frame.
func (f *Frame) InRect(x, y int) bool {
	return f.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Frame.
func (f *Frame) GetInnerRect() (x, y, width, height int) {
	return f.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Frame is preservef.
func (f *Frame) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return f.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Frame is preservef.
func (f *Frame) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return f.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Frame.
func (f *Frame) GetRect() (x, y, width, height int) {
	return f.box.GetRect()
}

// SetRect sets the rectangle occupied by this Frame.
func (f *Frame) SetRect(x, y, width, height int) {
	f.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Frame is visible.
func (f *Frame) GetVisible() bool {
	return f.box.GetVisible()
}

// SetVisible sets whether this Frame is visible.
func (f *Frame) SetVisible(visible bool) {
	f.box.SetVisible(visible)
}

// Focus is called when this widget receives focus.
func (f *Frame) Focus(delegate func(p Widget)) {
	f.mu.RLock()
	widget := f.widget
	f.mu.RUnlock()

	delegate(widget)
}

// HasFocus returns whether this widget has focus.
func (f *Frame) HasFocus() bool {
	f.mu.RLock()
	defer f.mu.RUnlock()

	focusable, ok := f.widget.(Focusable)
	if ok {
		return focusable.HasFocus()
	}
	return false
}

// GetFocusable returns the focusable primitive of this Frame.
func (f *Frame) GetFocusable() Focusable {
	return f.box.GetFocusable()
}

// Blur is called when this Frame loses focus.
func (f *Frame) Blur() {
	f.box.Blur()
}

/////////////////////////////////////// <API> ///////////////////////////////////////

func (f *Frame) GetWidget() (widget Widget) {
	f.get(func(f *Frame) { widget = f.widget })
	return
}

func (f *Frame) SetWidget(widget Widget) *Frame {
	return f.set(func(f *Frame) { f.widget = widget })
}

// AddText adds text to the frame. Set "header" to true if the text is to appear
// in the header, above the contained widget. Set it to false for it to
// appear in the footer, below the contained widget. "align" must be one of
// the Align constants. Rows in the header are printed top to bottom, rows in
// the footer are printed bottom to top. Note that long text can overlap as
// different alignments will be placed on the same row.
func (f *Frame) AddText(text string, header bool, align int, color tcell.Color) *Frame {
	return f.set(func(f *Frame) {
		f.text = append(f.text, &frameText{
			Text:   text,
			Header: header,
			Align:  align,
			Color:  color,
		})
	})
}

// Clear removes all text from the frame.
func (f *Frame) Clear() *Frame {
	return f.set(func(f *Frame) { f.text = nil })
}

// SetBorders sets the width of the frame borders as well as "header" and
// "footer", the vertical space between the header and footer text and the
// contained widget (does not apply if there is no text).
func (f *Frame) SetBorders(top, bottom, header, footer, left, right int) *Frame {
	return f.set(func(f *Frame) {
		f.top, f.bottom, f.header, f.footer, f.left, f.right = top, bottom, header, footer, left, right
	})
}

// Draw draws this widget onto the screen.
func (f *Frame) Draw(screen tcell.Screen) {
	if !f.GetVisible() {
		return
	}

	f.box.Draw(screen)

	f.mu.Lock()
	defer f.mu.Unlock()

	// Calculate start positions.
	x, top, width, height := f.GetInnerRect()
	bottom := top + height - 1
	x += f.left
	top += f.top
	bottom -= f.bottom
	width -= f.left + f.right
	if width <= 0 || top >= bottom {
		return // No space left.
	}

	// Draw text.
	var rows [6]int // top-left, top-center, top-right, bottom-left, bottom-center, bottom-right.
	topMax := top
	bottomMin := bottom
	for _, text := range f.text {
		// Where do we place this text?
		var y int
		if text.Header {
			y = top + rows[text.Align]
			rows[text.Align]++
			if y >= bottomMin {
				continue
			}
			if y+1 > topMax {
				topMax = y + 1
			}
		} else {
			y = bottom - rows[3+text.Align]
			rows[3+text.Align]++
			if y <= topMax {
				continue
			}
			if y-1 < bottomMin {
				bottomMin = y - 1
			}
		}

		// Draw text.
		Print(screen, []byte(text.Text), x, y, width, text.Align, text.Color)
	}

	// Set the size of the contained widget.
	if topMax > top {
		top = topMax + f.header
	}
	if bottomMin < bottom {
		bottom = bottomMin - f.footer
	}
	if top > bottom {
		return // No space for the widget.
	}
	f.widget.SetRect(x, top, width, bottom+1-top)

	// Finally, draw the contained widget.
	f.widget.Draw(screen)
}

// InputHandler returns the input handler for this widget.
func (f *Frame) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return f.box.InputHandler()
}

// MouseHandler returns the mouse handler for this widget.
func (f *Frame) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return f.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if !f.InRect(event.Position()) {
			return false, nil
		}

		// Pass mouse events on to contained widget.
		return f.widget.MouseHandler()(action, event, setFocus)
	})
}
