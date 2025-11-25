package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Spinner represents a spinner widget.
type Spinner struct {
	box *Box

	counter      int
	currentStyle SpinnerStyle

	styles map[SpinnerStyle][]rune

	mu sync.RWMutex
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (s *Spinner) set(setter func(s *Spinner)) *Spinner {
	s.mu.Lock()
	setter(s)
	s.mu.Unlock()
	return s
}

func (s *Spinner) get(getter func(s *Spinner)) {
	s.mu.RLock()
	getter(s)
	s.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Spinner.
func (s *Spinner) GetTitle() string {
	return s.box.GetTitle()
}

// SetTitle sets the title of this Spinner.
func (s *Spinner) SetTitle(title string) *Spinner {
	s.box.SetTitle(title)
	return s
}

// GetTitleAlign returns the title alignment of this Spinner.
func (s *Spinner) GetTitleAlign() int {
	return s.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Spinner.
func (s *Spinner) SetTitleAlign(align int) *Spinner {
	s.box.SetTitleAlign(align)
	return s
}

// GetBorder returns whether this Spinner has a border.
func (s *Spinner) GetBorder() bool {
	return s.box.GetBorder()
}

// SetBorder sets whether this Spinner has a border.
func (s *Spinner) SetBorder(show bool) *Spinner {
	s.box.SetBorder(show)
	return s
}

// GetBorderColor returns the border color of this Spinner.
func (s *Spinner) GetBorderColor() tcell.Color {
	return s.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Spinner.
func (s *Spinner) SetBorderColor(color tcell.Color) *Spinner {
	s.box.SetBorderColor(color)
	return s
}

// GetBorderAttributes returns the border attributes of this Spinner.
func (s *Spinner) GetBorderAttributes() tcell.AttrMask {
	return s.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Spinner.
func (s *Spinner) SetBorderAttributes(attr tcell.AttrMask) *Spinner {
	s.box.SetBorderAttributes(attr)
	return s
}

// GetBorderColorFocused returns the border color of this Spinner when focused.
func (s *Spinner) GetBorderColorFocused() tcell.Color {
	return s.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Spinner when focused.
func (s *Spinner) SetBorderColorFocused(color tcell.Color) *Spinner {
	s.box.SetBorderColorFocused(color)
	return s
}

// GetTitleColor returns the title color of this Spinner.
func (s *Spinner) GetTitleColor() tcell.Color {
	return s.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Spinner.
func (s *Spinner) SetTitleColor(color tcell.Color) *Spinner {
	s.box.SetTitleColor(color)
	return s
}

// GetDrawFunc returns the custom draw function of this Spinner.
func (s *Spinner) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return s.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Spinner.
func (s *Spinner) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Spinner {
	s.box.SetDrawFunc(handler)
	return s
}

// ShowFocus sets whether this Spinner should show a focus indicator when focused.
func (s *Spinner) ShowFocus(showFocus bool) *Spinner {
	s.box.ShowFocus(showFocus)
	return s
}

// GetMouseCapture returns the mouse capture function of this Spinner.
func (s *Spinner) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return s.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Spinner.
func (s *Spinner) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Spinner {
	s.box.SetMouseCapture(capture)
	return s
}

// GetBackgroundColor returns the background color of this Spinner.
func (s *Spinner) GetBackgroundColor() tcell.Color {
	return s.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Spinner.
func (s *Spinner) SetBackgroundColor(color tcell.Color) *Spinner {
	s.box.SetBackgroundColor(color)
	return s
}

// GetBackgroundTransparent returns whether the background of this Spinner is transparent.
func (s *Spinner) GetBackgroundTransparent() bool {
	return s.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Spinner is transparent.
func (s *Spinner) SetBackgroundTransparent(transparent bool) *Spinner {
	s.box.SetBackgroundTransparent(transparent)
	return s
}

// GetInputCapture returns the input capture function of this Spinner.
func (s *Spinner) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return s.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Spinner.
func (s *Spinner) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Spinner {
	s.box.SetInputCapture(capture)
	return s
}

// GetPadding returns the padding of this Spinner.
func (s *Spinner) GetPadding() (top, bottom, left, right int) {
	return s.box.GetPadding()
}

// SetPadding sets the padding of this Spinner.
func (s *Spinner) SetPadding(top, bottom, left, right int) *Spinner {
	s.box.SetPadding(top, bottom, left, right)
	return s
}

// InRect returns whether the given screen coordinates are within this Spinner.
func (s *Spinner) InRect(x, y int) bool {
	return s.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Spinner.
func (s *Spinner) GetInnerRect() (x, y, width, height int) {
	return s.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Spinner is preserved.
func (s *Spinner) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return s.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Spinner is preserved.
func (s *Spinner) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return s.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Spinner.
func (s *Spinner) GetRect() (x, y, width, height int) {
	return s.box.GetRect()
}

// SetRect sets the rectangle occupied by this Spinner.
func (s *Spinner) SetRect(x, y, width, height int) {
	s.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Spinner is visible.
func (s *Spinner) GetVisible() bool {
	return s.box.GetVisible()
}

// SetVisible sets whether this Spinner is visible.
func (s *Spinner) SetVisible(visible bool) {
	s.box.SetVisible(visible)
}

// Focus is called when this Spinner receives focus.
func (s *Spinner) Focus(delegate func(p Widget)) {
	s.box.Focus(delegate)
}

// HasFocus returns whether this Spinner has focus.
func (s *Spinner) HasFocus() bool {
	return s.box.HasFocus()
}

// GetFocusable returns this Spinner as a Focusable.
func (s *Spinner) GetFocusable() Focusable {
	return s.box.GetFocusable()
}

// Blur is called when this Spinner loses focus.
func (s *Spinner) Blur() {
	s.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

type SpinnerStyle int

const (
	SpinnerDotsCircling SpinnerStyle = iota
	SpinnerDotsUpDown
	SpinnerBounce
	SpinnerLine
	SpinnerCircleQuarters
	SpinnerSquareCorners
	SpinnerCircleHalves
	SpinnerCorners
	SpinnerArrows
	SpinnerHamburger
	SpinnerStack
	SpinnerGrowHorizontal
	SpinnerGrowVertical
	SpinnerStar
	SpinnerBoxBounce
	spinnerCustom // non-public constant to indicate that a custom style has been set by the user.
)

// NewSpinner returns a new spinner widget.
func NewSpinner() *Spinner {
	return &Spinner{
		box:          NewBox(),
		currentStyle: SpinnerDotsCircling,
		styles: map[SpinnerStyle][]rune{
			SpinnerDotsCircling:   []rune(`⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`),
			SpinnerDotsUpDown:     []rune(`⠋⠙⠚⠞⠖⠦⠴⠲⠳⠓`),
			SpinnerBounce:         []rune(`⠄⠆⠇⠋⠙⠸⠰⠠⠰⠸⠙⠋⠇⠆`),
			SpinnerLine:           []rune(`|/-\`),
			SpinnerCircleQuarters: []rune(`◴◷◶◵`),
			SpinnerSquareCorners:  []rune(`◰◳◲◱`),
			SpinnerCircleHalves:   []rune(`◐◓◑◒`),
			SpinnerCorners:        []rune(`⌜⌝⌟⌞`),
			SpinnerArrows:         []rune(`⇑⇗⇒⇘⇓⇙⇐⇖`),
			SpinnerHamburger:      []rune(`☰☱☳☷☶☴`),
			SpinnerStack:          []rune(`䷀䷪䷡䷊䷒䷗䷁䷖䷓䷋䷠䷫`),
			SpinnerGrowHorizontal: []rune(`▉▊▋▌▍▎▏▎▍▌▋▊▉`),
			SpinnerGrowVertical:   []rune(`▁▃▄▅▆▇▆▅▄▃`),
			SpinnerStar:           []rune(`✶✸✹✺✹✷`),
			SpinnerBoxBounce:      []rune(`▌▀▐▄`),
		},
	}
}

// Pulse updates the spinner to the next frame.
func (s *Spinner) Pulse() *Spinner {
	return s.set(func(s *Spinner) { s.counter++ })
}

// Reset sets the frame counter to 0.
func (s *Spinner) Reset() *Spinner {
	return s.set(func(s *Spinner) { s.counter = 0 })
}

// SetStyle sets the spinner style.
func (s *Spinner) SetStyle(style SpinnerStyle) *Spinner {
	return s.set(func(s *Spinner) { s.currentStyle = style })
}

// SetCustomStyle sets a list of runes as custom frames to show as the spinner.
func (s *Spinner) SetCustomStyle(frames []rune) *Spinner {
	return s.set(func(s *Spinner) {
		s.styles[spinnerCustom] = frames
		s.currentStyle = spinnerCustom
	})
}

// InputHandler returns the input handler function for this Spinner.
func (s *Spinner) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return s.box.InputHandler()
}

// MouseHandler returns the mouse handler function for this Spinner.
func (s *Spinner) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return s.box.MouseHandler()
}

// Draw draws this Spinner onto the screen.
func (s *Spinner) Draw(screen tcell.Screen) {
	if !s.GetVisible() {
		return
	}

	s.box.Draw(screen)

	s.mu.Lock()
	defer s.mu.Unlock()

	x, y, width, _ := s.GetInnerRect()

	Print(screen, []byte(func(s *Spinner) string {
		frames := s.styles[s.currentStyle]
		if len(frames) == 0 {
			return ""
		}
		return string(frames[s.counter%len(frames)])
	}(s)), x, y, width, AlignLeft, tcell.ColorDefault)
}
