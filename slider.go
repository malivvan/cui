package cui

import (
	"math"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Slider is a progress bar which may be modified via keyboard and mouse.
type Slider struct {
	progressBar *ProgressBar

	// The text to be displayed before the slider.
	label []byte

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

	// The amount to increment by when modified via keyboard.
	increment int

	// Set to true when mouse dragging is in progress.
	dragging bool

	// An optional function which is called when the user changes the value of
	// this slider.
	changed func(value int)

	// An optional function which is called when the user indicated that they
	// are done entering text. The key which was pressed is provided (tab,
	// shift-tab, or escape).
	done func(tcell.Key)

	// A callback function set by the Form class and called when the user leaves
	// this form item.
	finished func(tcell.Key)

	mu sync.RWMutex
}

// NewSlider returns a new slider.
func NewSlider() *Slider {
	s := &Slider{
		progressBar:                 NewProgressBar(),
		increment:                   10,
		labelColor:                  Styles.SecondaryTextColor,
		fieldBackgroundColor:        Styles.MoreContrastBackgroundColor,
		fieldBackgroundColorFocused: Styles.ContrastBackgroundColor,
		fieldTextColor:              Styles.PrimaryTextColor,
		labelColorFocused:           ColorUnset,
		fieldTextColorFocused:       ColorUnset,
	}
	return s
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (s *Slider) set(setter func(s *Slider)) *Slider {
	s.mu.Lock()
	setter(s)
	s.mu.Unlock()
	return s
}

func (s *Slider) get(getter func(s *Slider)) {
	s.mu.RLock()
	getter(s)
	s.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Slider.
func (s *Slider) GetTitle() string {
	return s.progressBar.GetTitle()
}

// SetTitle sets the title of this Slider.
func (s *Slider) SetTitle(title string) *Slider {
	s.progressBar.SetTitle(title)
	return s
}

// GetTitleAlign returns the title alignment of this Slider.
func (s *Slider) GetTitleAlign() int {
	return s.progressBar.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Slider.
func (s *Slider) SetTitleAlign(align int) *Slider {
	s.progressBar.SetTitleAlign(align)
	return s
}

// GetBorder returns whether this Slider has a border.
func (s *Slider) GetBorder() bool {
	return s.progressBar.GetBorder()
}

// SetBorder sets whether this Slider has a border.
func (s *Slider) SetBorder(show bool) *Slider {
	s.progressBar.SetBorder(show)
	return s
}

// GetBorderColor returns the border color of this Slider.
func (s *Slider) GetBorderColor() tcell.Color {
	return s.progressBar.GetBorderColor()
}

// SetBorderColor sets the border color of this Slider.
func (s *Slider) SetBorderColor(color tcell.Color) *Slider {
	s.progressBar.SetBorderColor(color)
	return s
}

// GetBorderAttributes returns the border attributes of this Slider.
func (s *Slider) GetBorderAttributes() tcell.AttrMask {
	return s.progressBar.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Slider.
func (s *Slider) SetBorderAttributes(attr tcell.AttrMask) *Slider {
	s.progressBar.SetBorderAttributes(attr)
	return s
}

// GetBorderColorFocused returns the border color of this Slider when focused.
func (s *Slider) GetBorderColorFocused() tcell.Color {
	return s.progressBar.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Slider when focused.
func (s *Slider) SetBorderColorFocused(color tcell.Color) *Slider {
	s.progressBar.SetBorderColorFocused(color)
	return s
}

// GetTitleColor returns the title color of this Slider.
func (s *Slider) GetTitleColor() tcell.Color {
	return s.progressBar.GetTitleColor()
}

// SetTitleColor sets the title color of this Slider.
func (s *Slider) SetTitleColor(color tcell.Color) *Slider {
	s.progressBar.SetTitleColor(color)
	return s
}

// GetDrawFunc returns the custom draw function of this Slider.
func (s *Slider) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return s.progressBar.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Slider.
func (s *Slider) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Slider {
	s.progressBar.SetDrawFunc(handler)
	return s
}

// ShowFocus sets whether this Slider should show a focus indicator when focused.
func (s *Slider) ShowFocus(showFocus bool) *Slider {
	s.progressBar.ShowFocus(showFocus)
	return s
}

// GetMouseCapture returns the mouse capture function of this Slider.
func (s *Slider) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return s.progressBar.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Slider.
func (s *Slider) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Slider {
	s.progressBar.SetMouseCapture(capture)
	return s
}

// GetBackgroundColor returns the background color of this Slider.
func (s *Slider) GetBackgroundColor() tcell.Color {
	return s.progressBar.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Slider.
func (s *Slider) SetBackgroundColor(color tcell.Color) *Slider {
	s.progressBar.SetBackgroundColor(color)
	return s
}

// GetBackgroundTransparent returns whether the background of this Slider is transparent.
func (s *Slider) GetBackgroundTransparent() bool {
	return s.progressBar.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Slider is transparent.
func (s *Slider) SetBackgroundTransparent(transparent bool) *Slider {
	s.progressBar.SetBackgroundTransparent(transparent)
	return s
}

// GetInputCapture returns the input capture function of this Slider.
func (s *Slider) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return s.progressBar.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Slider.
func (s *Slider) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Slider {
	s.progressBar.SetInputCapture(capture)
	return s
}

// GetPadding returns the padding of this Slider.
func (s *Slider) GetPadding() (top, bottom, left, right int) {
	return s.progressBar.GetPadding()
}

// SetPadding sets the padding of this Slider.
func (s *Slider) SetPadding(top, bottom, left, right int) *Slider {
	s.progressBar.SetPadding(top, bottom, left, right)
	return s
}

// InRect returns whether the given screen coordinates are within this Slider.
func (s *Slider) InRect(x, y int) bool {
	return s.progressBar.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Slider.
func (s *Slider) GetInnerRect() (x, y, width, height int) {
	return s.progressBar.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Slider is preserved.
func (s *Slider) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return s.progressBar.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Slider is preserved.
func (s *Slider) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return s.progressBar.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Slider.
func (s *Slider) GetRect() (x, y, width, height int) {
	return s.progressBar.GetRect()
}

// SetRect sets the rectangle occupied by this Slider.
func (s *Slider) SetRect(x, y, width, height int) {
	s.progressBar.SetRect(x, y, width, height)
}

// GetVisible returns whether this Slider is visible.
func (s *Slider) GetVisible() bool {
	return s.progressBar.GetVisible()
}

// SetVisible sets whether this Slider is visible.
func (s *Slider) SetVisible(visible bool) {
	s.progressBar.SetVisible(visible)
}

// Focus is called when this Slider receives focus.
func (s *Slider) Focus(delegate func(p Widget)) {
	s.progressBar.Focus(delegate)
}

// HasFocus returns whether this Slider has focus.
func (s *Slider) HasFocus() bool {
	return s.progressBar.HasFocus()
}

// GetFocusable returns the focusable primitive of this Slider.
func (s *Slider) GetFocusable() Focusable {
	return s.progressBar.GetFocusable()
}

// Blur is called when this Slider loses focus.
func (s *Slider) Blur() {
	s.progressBar.Blur()
}

/////////////////////////////// <PROGRESS BAR> ///////////////////////////////

func (s *Slider) SetEmptyRune(r rune) *Slider {
	s.progressBar.SetEmptyRune(r)
	return s
}
func (s *Slider) SetEmptyColor(color tcell.Color) *Slider {
	s.progressBar.SetEmptyColor(color)
	return s
}
func (s *Slider) SetFilledRune(r rune) *Slider {
	s.progressBar.SetFilledRune(r)
	return s
}
func (s *Slider) SetFilledColor(color tcell.Color) *Slider {
	s.progressBar.SetFilledColor(color)
	return s
}
func (s *Slider) SetVertical(vertical bool) *Slider {
	s.progressBar.SetVertical(vertical)
	return s
}
func (s *Slider) SetMax(max int) *Slider {
	s.progressBar.SetMax(max)
	return s
}
func (s *Slider) GetMax() int {
	return s.progressBar.GetMax()
}
func (s *Slider) AddProgress(progress int) *Slider {
	s.progressBar.AddProgress(progress)
	return s
}
func (s *Slider) Complete() bool {
	return s.progressBar.Complete()
}
func (s *Slider) GetProgress() (progress int) {
	return s.progressBar.GetProgress()
}
func (s *Slider) SetProgress(progress int) *Slider {
	s.progressBar.SetProgress(progress)
	return s
}

/////////////////////////////// <API> ///////////////////////////////

// SetLabel sets the text to be displayed before the input area.
func (s *Slider) SetLabel(label string) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.label = []byte(label)
	return s
}

// GetLabel returns the text to be displayed before the input area.
func (s *Slider) GetLabel() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return string(s.label)
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (s *Slider) SetLabelWidth(width int) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.labelWidth = width
	return s
}

// SetLabelColor sets the color of the label.
func (s *Slider) SetLabelColor(color tcell.Color) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.labelColor = color
	return s
}

// SetLabelColorFocused sets the color of the label when focused.
func (s *Slider) SetLabelColorFocused(color tcell.Color) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.labelColorFocused = color
	return s
}

// SetFieldBackgroundColor sets the background color of the input area.
func (s *Slider) SetFieldBackgroundColor(color tcell.Color) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.fieldBackgroundColor = color
	return s
}

// SetFieldBackgroundColorFocused sets the background color of the input area when focused.
func (s *Slider) SetFieldBackgroundColorFocused(color tcell.Color) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.fieldBackgroundColorFocused = color
	return s
}

// SetFieldTextColor sets the text color of the input area.
func (s *Slider) SetFieldTextColor(color tcell.Color) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.fieldTextColor = color
	return s
}

// SetFieldTextColorFocused sets the text color of the input area when focused.
func (s *Slider) SetFieldTextColorFocused(color tcell.Color) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.fieldTextColorFocused = color
	return s
}

// GetFieldHeight returns the height of the field.
func (s *Slider) GetFieldHeight() int {
	return 1
}

// GetFieldWidth returns this primitive's field width.
func (s *Slider) GetFieldWidth() int {
	return 0
}

// SetIncrement sets the amount the slider is incremented by when modified via
// keyboard.
func (s *Slider) SetIncrement(increment int) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.increment = increment
	return s
}

// SetChangedFunc sets a handler which is called when the value of this slider
// was changed by the user. The handler function receives the new value.
func (s *Slider) SetChangedFunc(handler func(value int)) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.changed = handler
	return s
}

// SetDoneFunc sets a handler which is called when the user is done using the
// slider. The callback function is provided with the key that was pressed,
// which is one of the following:
//
//   - KeyEscape: Abort text input.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (s *Slider) SetDoneFunc(handler func(key tcell.Key)) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.done = handler
	return s
}

// SetFinishedFunc sets a callback invoked when the user leaves this form item.
func (s *Slider) SetFinishedFunc(handler func(key tcell.Key)) *Slider {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.finished = handler
	return s
}

// Draw draws this primitive onto the screen.
func (s *Slider) Draw(screen tcell.Screen) {
	if !s.GetVisible() {
		return
	}

	s.progressBar.box.Draw(screen)
	hasFocus := s.GetFocusable().HasFocus()

	s.mu.Lock()

	// Select colors
	labelColor := s.labelColor
	fieldBackgroundColor := s.fieldBackgroundColor
	fieldTextColor := s.fieldTextColor
	if hasFocus {
		if s.labelColorFocused != ColorUnset {
			labelColor = s.labelColorFocused
		}
		if s.fieldBackgroundColorFocused != ColorUnset {
			fieldBackgroundColor = s.fieldBackgroundColorFocused
		}
		if s.fieldTextColorFocused != ColorUnset {
			fieldTextColor = s.fieldTextColorFocused
		}
	}

	// Prepare.
	x, y, width, height := s.GetInnerRect()
	rightLimit := x + width
	if height < 1 || rightLimit <= x {
		s.mu.Unlock()
		return
	}

	// Draw label.
	if len(s.label) > 0 {
		if s.progressBar.vertical {
			height--

			// TODO draw label on bottom
		} else {
			if s.labelWidth > 0 {
				labelWidth := s.labelWidth
				if labelWidth > rightLimit-x {
					labelWidth = rightLimit - x
				}
				Print(screen, []byte(s.label), x, y, labelWidth, AlignLeft, labelColor)
				x += labelWidth + 1
				width -= labelWidth + 1
			} else {
				_, drawnWidth := Print(screen, []byte(s.label), x, y, rightLimit-x, AlignLeft, labelColor)
				x += drawnWidth + 1
				width -= drawnWidth + 1
			}
		}
	}

	// Draw slider.
	s.mu.Unlock()
	s.progressBar.SetRect(x, y, width, height)
	s.progressBar.SetEmptyColor(fieldBackgroundColor)
	s.progressBar.SetFilledColor(fieldTextColor)
	s.progressBar.Draw(screen)
}

// InputHandler returns the handler for this primitive.
func (s *Slider) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return s.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		if HitShortcut(event, Keys.Cancel, Keys.MovePreviousField, Keys.MoveNextField) {
			if s.done != nil {
				s.done(event.Key())
			}
			if s.finished != nil {
				s.finished(event.Key())
			}
			return
		}

		previous := s.progressBar.progress

		if HitShortcut(event, Keys.MoveFirst, Keys.MoveFirst2) {
			s.progressBar.SetProgress(0)
		} else if HitShortcut(event, Keys.MoveLast, Keys.MoveLast2) {
			s.progressBar.SetProgress(s.progressBar.max)
		} else if HitShortcut(event, Keys.MoveUp, Keys.MoveUp2, Keys.MoveRight, Keys.MoveRight2, Keys.MovePreviousField) {
			s.progressBar.AddProgress(s.increment)
		} else if HitShortcut(event, Keys.MoveDown, Keys.MoveDown2, Keys.MoveLeft, Keys.MoveLeft2, Keys.MoveNextField) {
			s.progressBar.AddProgress(s.increment * -1)
		}

		if s.progressBar.progress != previous && s.changed != nil {
			s.changed(s.progressBar.progress)
		}
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (s *Slider) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return s.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		x, y := event.Position()
		if !s.InRect(x, y) {
			s.dragging = false
			return false, nil
		}

		// Process mouse event.
		if action == MouseLeftClick {
			setFocus(s)
			consumed = true
		}

		handleMouse := func() {
			if !s.progressBar.InRect(x, y) {
				s.dragging = false
				return
			}

			bx, by, bw, bh := s.GetInnerRect()
			var clickPos, clickRange int
			if s.progressBar.vertical {
				clickPos = (bh - 1) - (y - by)
				clickRange = bh - 1
			} else {
				clickPos = x - bx
				clickRange = bw - 1
			}
			setValue := int(math.Floor(float64(s.progressBar.max) * (float64(clickPos) / float64(clickRange))))
			if setValue != s.progressBar.progress {
				s.progressBar.SetProgress(setValue)
				if s.changed != nil {
					s.changed(s.progressBar.progress)
				}
			}
		}

		// HandleMessage dragging. Clicks are implicitly handled by this logic.
		switch action {
		case MouseLeftDown:
			setFocus(s)
			consumed = true
			capture = s
			s.dragging = true

			handleMouse()
		case MouseMove:
			if s.dragging {
				consumed = true
				capture = s

				handleMouse()
			}
		case MouseLeftUp:
			if s.dragging {
				consumed = true
				s.dragging = false

				handleMouse()
			}
		}

		return
	})
}
