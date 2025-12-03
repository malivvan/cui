package cui

import (
	"math"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Sparkline represents a sparkline widgets.
type Sparkline struct {
	box *Box

	data           []float64
	dataTitle      string
	dataTitlecolor tcell.Color
	lineColor      tcell.Color
	mu             sync.RWMutex
}

func (s *Sparkline) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return s.box.InputHandler()
}

func (s *Sparkline) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return s.box.MouseHandler()
}

// NewSparkline returns a new sparkline widget.
func NewSparkline() *Sparkline {
	return &Sparkline{
		box: NewBox(),
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (s *Sparkline) set(setter func(s *Sparkline)) *Sparkline {
	s.mu.Lock()
	setter(s)
	s.mu.Unlock()
	return s
}

func (s *Sparkline) get(getter func(s *Sparkline)) {
	s.mu.RLock()
	getter(s)
	s.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Sparkline.
func (s *Sparkline) GetTitle() string {
	return s.box.GetTitle()
}

// SetTitle sets the title of this Sparkline.
func (s *Sparkline) SetTitle(title string) *Sparkline {
	s.box.SetTitle(title)
	return s
}

// GetTitleAlign returns the title alignment of this Sparkline.
func (s *Sparkline) GetTitleAlign() int {
	return s.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Sparkline.
func (s *Sparkline) SetTitleAlign(align int) *Sparkline {
	s.box.SetTitleAlign(align)
	return s
}

// GetBorder returns whether this Sparkline has a border.
func (s *Sparkline) GetBorder() bool {
	return s.box.GetBorder()
}

// SetBorder sets whether this Sparkline has a border.
func (s *Sparkline) SetBorder(show bool) *Sparkline {
	s.box.SetBorder(show)
	return s
}

// GetBorderColor returns the border color of this Sparkline.
func (s *Sparkline) GetBorderColor() tcell.Color {
	return s.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Sparkline.
func (s *Sparkline) SetBorderColor(color tcell.Color) *Sparkline {
	s.box.SetBorderColor(color)
	return s
}

// GetBorderAttributes returns the border attributes of this Sparkline.
func (s *Sparkline) GetBorderAttributes() tcell.AttrMask {
	return s.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Sparkline.
func (s *Sparkline) SetBorderAttributes(attr tcell.AttrMask) *Sparkline {
	s.box.SetBorderAttributes(attr)
	return s
}

// GetBorderColorFocused returns the border color of this Sparkline when focused.
func (s *Sparkline) GetBorderColorFocused() tcell.Color {
	return s.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Sparkline when focused.
func (s *Sparkline) SetBorderColorFocused(color tcell.Color) *Sparkline {
	s.box.SetBorderColorFocused(color)
	return s
}

// GetTitleColor returns the title color of this Sparkline.
func (s *Sparkline) GetTitleColor() tcell.Color {
	return s.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Sparkline.
func (s *Sparkline) SetTitleColor(color tcell.Color) *Sparkline {
	s.box.SetTitleColor(color)
	return s
}

// GetDrawFunc returns the custom draw function of this Sparkline.
func (s *Sparkline) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return s.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Sparkline.
func (s *Sparkline) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Sparkline {
	s.box.SetDrawFunc(handler)
	return s
}

// ShowFocus sets whether this Sparkline should show a focus indicator when focused.
func (s *Sparkline) ShowFocus(showFocus bool) *Sparkline {
	s.box.ShowFocus(showFocus)
	return s
}

// GetMouseCapture returns the mouse capture function of this Sparkline.
func (s *Sparkline) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return s.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Sparkline.
func (s *Sparkline) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Sparkline {
	s.box.SetMouseCapture(capture)
	return s
}

// GetBackgroundColor returns the background color of this Sparkline.
func (s *Sparkline) GetBackgroundColor() tcell.Color {
	return s.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Sparkline.
func (s *Sparkline) SetBackgroundColor(color tcell.Color) *Sparkline {
	s.box.SetBackgroundColor(color)
	return s
}

// GetBackgroundTransparent returns whether the background of this Sparkline is transparent.
func (s *Sparkline) GetBackgroundTransparent() bool {
	return s.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Sparkline is transparent.
func (s *Sparkline) SetBackgroundTransparent(transparent bool) *Sparkline {
	s.box.SetBackgroundTransparent(transparent)
	return s
}

// GetInputCapture returns the input capture function of this Sparkline.
func (s *Sparkline) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return s.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Sparkline.
func (s *Sparkline) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Sparkline {
	s.box.SetInputCapture(capture)
	return s
}

// GetPadding returns the padding of this Sparkline.
func (s *Sparkline) GetPadding() (top, bottom, left, right int) {
	return s.box.GetPadding()
}

// SetPadding sets the padding of this Sparkline.
func (s *Sparkline) SetPadding(top, bottom, left, right int) *Sparkline {
	s.box.SetPadding(top, bottom, left, right)
	return s
}

// InRect returns whether the given screen coordinates are within this Sparkline.
func (s *Sparkline) InRect(x, y int) bool {
	return s.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Sparkline.
func (s *Sparkline) GetInnerRect() (x, y, width, height int) {
	return s.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Sparkline is preserved.
func (s *Sparkline) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return s.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Sparkline is preserved.
func (s *Sparkline) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return s.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Sparkline.
func (s *Sparkline) GetRect() (x, y, width, height int) {
	return s.box.GetRect()
}

// SetRect sets the rectangle occupied by this Sparkline.
func (s *Sparkline) SetRect(x, y, width, height int) {
	s.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Sparkline is visible.
func (s *Sparkline) GetVisible() bool {
	return s.box.GetVisible()
}

// SetVisible sets whether this Sparkline is visible.
func (s *Sparkline) SetVisible(visible bool) {
	s.box.SetVisible(visible)
}

// Focus is called when this Sparkline receives focus.
func (s *Sparkline) Focus(delegate func(p Widget)) {
	s.box.Focus(delegate)
}

// GetFocusable returns the focusable primitive of this Sparkline.
func (s *Sparkline) GetFocusable() Focusable {
	return s.box.GetFocusable()
}

// Blur is called when this Sparkline loses focus.
func (s *Sparkline) Blur() {
	s.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

// Draw draws this primitive onto the screen.
func (sl *Sparkline) Draw(screen tcell.Screen) {
	if !sl.box.GetVisible() {
		return
	}
	sl.box.Draw(screen)

	x, y, width, height := sl.box.GetInnerRect()
	barHeight := height

	// print label
	if sl.dataTitle != "" {
		Print(screen, []byte(sl.dataTitle), x, y, width, AlignLeft, sl.dataTitlecolor)

		barHeight--
	}

	maxVal := getMaxFloat64FromSlice(sl.data)
	if maxVal < 0 {
		return
	}

	// print lines
	for i := 0; i < len(sl.data) && i+x < x+width; i++ {
		data := sl.data[i]

		if math.IsNaN(data) {
			continue
		}

		dHeight := int((data / maxVal) * float64(barHeight))

		sparkChar := barsRune[len(barsRune)-1]

		for j := range dHeight {
			PrintJoinedSemigraphics(screen, i+x, y-1+height-j, sparkChar, sl.lineColor)
		}

		if dHeight == 0 {
			sparkChar = barsRune[1]
			PrintJoinedSemigraphics(screen, i+x, y-1+height, sparkChar, sl.lineColor)
		}
	}
}

// HasFocus returns whether this primitive has focus.
func (sl *Sparkline) HasFocus() bool {
	return sl.box.HasFocus()
}

// SetData sets sparkline data.
func (sl *Sparkline) SetData(data []float64) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	sl.data = data
}

// SetDataTitle sets sparkline data title.
func (sl *Sparkline) SetDataTitle(title string) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	sl.dataTitle = title
}

// SetDataTitleColor sets sparkline data title color.
func (sl *Sparkline) SetDataTitleColor(color tcell.Color) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	sl.dataTitlecolor = color
}

// SetLineColor sets sparkline line color.
func (sl *Sparkline) SetLineColor(color tcell.Color) {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	sl.lineColor = color
}
