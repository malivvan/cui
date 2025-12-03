package cui

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Gauge represents utilisation mode gauge Widget.
type Gauge struct {
	box *Box
	// pc percentage value
	pc float64
	// warn percentage value
	warnPc float64
	// critical percentage value
	critPc float64
	// okColor ok color
	okColor tcell.Color
	// warnColor warning block color
	warnColor tcell.Color
	// critColor critical block color
	critColor tcell.Color
	// emptyColor empty block color
	emptyColor tcell.Color
	// label prints label on the left of the gauge
	label string
	// labelColor label and percentage text color
	labelColor tcell.Color

	mu sync.RWMutex
}

func (g *Gauge) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return g.box.InputHandler()
}

func (g *Gauge) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return g.box.MouseHandler()
}

// NewGauge returns new utilisation mode gauge permitive.
func NewGauge() *Gauge {
	gauge := &Gauge{
		box:        NewBox(),
		pc:         gaugeMinPc,
		warnPc:     gaugeWarnPc,
		critPc:     gaugeCritPc,
		warnColor:  tcell.ColorOrange,
		critColor:  tcell.ColorRed,
		okColor:    tcell.ColorGreen,
		emptyColor: tcell.ColorWhite,
		labelColor: Styles.PrimaryTextColor,
		label:      "",
	}

	return gauge
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (g *Gauge) set(setter func(g *Gauge)) *Gauge {
	g.mu.Lock()
	setter(g)
	g.mu.Unlock()
	return g
}

func (g *Gauge) get(getter func(g *Gauge)) {
	g.mu.RLock()
	getter(g)
	g.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Gauge.
func (g *Gauge) GetTitle() string {
	return g.box.GetTitle()
}

// SetTitle sets the title of this Gauge.
func (g *Gauge) SetTitle(title string) *Gauge {
	g.box.SetTitle(title)
	return g
}

// GetTitleAlign returns the title alignment of this Gauge.
func (g *Gauge) GetTitleAlign() int {
	return g.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Gauge.
func (g *Gauge) SetTitleAlign(align int) *Gauge {
	g.box.SetTitleAlign(align)
	return g
}

// GetBorder returns whether this Gauge has a border.
func (g *Gauge) GetBorder() bool {
	return g.box.GetBorder()
}

// SetBorder sets whether this Gauge has a border.
func (g *Gauge) SetBorder(show bool) *Gauge {
	g.box.SetBorder(show)
	return g
}

// GetBorderColor returns the border color of this Gauge.
func (g *Gauge) GetBorderColor() tcell.Color {
	return g.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Gauge.
func (g *Gauge) SetBorderColor(color tcell.Color) *Gauge {
	g.box.SetBorderColor(color)
	return g
}

// GetBorderAttributes returns the border attributes of this Gauge.
func (g *Gauge) GetBorderAttributes() tcell.AttrMask {
	return g.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Gauge.
func (g *Gauge) SetBorderAttributes(attr tcell.AttrMask) *Gauge {
	g.box.SetBorderAttributes(attr)
	return g
}

// GetBorderColorFocused returns the border color of this Gauge when focused.
func (g *Gauge) GetBorderColorFocused() tcell.Color {
	return g.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Gauge when focused.
func (g *Gauge) SetBorderColorFocused(color tcell.Color) *Gauge {
	g.box.SetBorderColorFocused(color)
	return g
}

// GetTitleColor returns the title color of this Gauge.
func (g *Gauge) GetTitleColor() tcell.Color {
	return g.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Gauge.
func (g *Gauge) SetTitleColor(color tcell.Color) *Gauge {
	g.box.SetTitleColor(color)
	return g
}

// GetDrawFunc returns the custom draw function of this Gauge.
func (g *Gauge) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return g.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Gauge.
func (g *Gauge) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Gauge {
	g.box.SetDrawFunc(handler)
	return g
}

// ShowFocus sets whether this Gauge should show a focus indicator when focused.
func (g *Gauge) ShowFocus(showFocus bool) *Gauge {
	g.box.ShowFocus(showFocus)
	return g
}

// GetMouseCapture returns the mouse capture function of this Gauge.
func (g *Gauge) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return g.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Gauge.
func (g *Gauge) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Gauge {
	g.box.SetMouseCapture(capture)
	return g
}

// GetBackgroundColor returns the background color of this Gauge.
func (g *Gauge) GetBackgroundColor() tcell.Color {
	return g.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Gauge.
func (g *Gauge) SetBackgroundColor(color tcell.Color) *Gauge {
	g.box.SetBackgroundColor(color)
	return g
}

// GetBackgroundTransparent returns whether the background of this Gauge is transparent.
func (g *Gauge) GetBackgroundTransparent() bool {
	return g.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Gauge is transparent.
func (g *Gauge) SetBackgroundTransparent(transparent bool) *Gauge {
	g.box.SetBackgroundTransparent(transparent)
	return g
}

// GetInputCapture returns the input capture function of this Gauge.
func (g *Gauge) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return g.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Gauge.
func (g *Gauge) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Gauge {
	g.box.SetInputCapture(capture)
	return g
}

// GetPadding returns the padding of this Gauge.
func (g *Gauge) GetPadding() (top, bottom, left, right int) {
	return g.box.GetPadding()
}

// SetPadding sets the padding of this Gauge.
func (g *Gauge) SetPadding(top, bottom, left, right int) *Gauge {
	g.box.SetPadding(top, bottom, left, right)
	return g
}

// InRect returns whether the given screen coordinates are within this Gauge.
func (g *Gauge) InRect(x, y int) bool {
	return g.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Gauge.
func (g *Gauge) GetInnerRect() (x, y, width, height int) {
	return g.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Gauge is preserved.
func (g *Gauge) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return g.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Gauge is preserved.
func (g *Gauge) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return g.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Gauge.
func (g *Gauge) GetRect() (x, y, width, height int) {
	return g.box.GetRect()
}

// SetRect sets the rectangle occupied by this Gauge.
func (g *Gauge) SetRect(x, y, width, height int) {
	g.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Gauge is visible.
func (g *Gauge) GetVisible() bool {
	return g.box.GetVisible()
}

// SetVisible sets whether this Gauge is visible.
func (g *Gauge) SetVisible(visible bool) {
	g.box.SetVisible(visible)
}

// Focus is called when this Gauge receives focus.
func (g *Gauge) Focus(delegate func(p Widget)) {
	g.box.Focus(delegate)
}

// HasFocus returns whether this Gauge has focus.
func (g *Gauge) HasFocus() bool {
	return g.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Gauge.
func (g *Gauge) GetFocusable() Focusable {
	return g.box.GetFocusable()
}

// Blur is called when this Gauge loses focus.
func (g *Gauge) Blur() {
	g.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

// SetLabel sets label for this Widget.
func (g *Gauge) SetLabel(label string) {
	g.label = label
}

// SetLabelColor sets label text color.
func (g *Gauge) SetLabelColor(color tcell.Color) {
	g.labelColor = color
}

// SetValue update the gauge progress.
func (g *Gauge) SetValue(value float64) {
	if value <= float64(gaugeMaxPc) {
		g.pc = value
	}
}

// GetValue returns current gauge value.
func (g *Gauge) GetValue() float64 {
	return g.pc
}

// Draw draws this Widget onto the screen.
func (g *Gauge) Draw(screen tcell.Screen) {
	if !g.GetVisible() {
		return
	}
	g.box.Draw(screen)

	g.mu.Lock()
	defer g.mu.Unlock()

	x, y, width, height := g.box.GetInnerRect()
	labelPCWidth := 7
	labelWidth := len(g.label)
	barWidth := width - labelPCWidth - labelWidth

	for i := range barWidth {
		for j := range height {
			value := float64(i * 100 / barWidth)
			color := g.getBarColor(value)

			if value > g.pc {
				color = g.emptyColor
			}

			Print(screen, []byte(prgCell), x+labelWidth+i, y+j, 1, AlignCenter, color)
		}
	}
	// draw label
	tY := y + (height / emptySpaceParts)
	if labelWidth > 0 {
		Print(screen, []byte(g.label), x, tY, labelWidth, AlignLeft, g.labelColor)
	}

	// draw percentage text
	Print(screen, []byte(fmt.Sprintf("%6.2f%%", g.pc)),
		x+barWidth+labelWidth,
		tY,
		labelPCWidth,
		AlignLeft,
		Styles.PrimaryTextColor)
}

// SetWarnPercentage sets warning percentage start range.
func (g *Gauge) SetWarnPercentage(percentage float64) {
	if percentage > 0 && percentage < 100 {
		g.warnPc = percentage
	}
}

// SetCritPercentage sets critical percentage start range.
func (g *Gauge) SetCritPercentage(percentage float64) {
	if percentage > 0 && percentage < 100 && percentage > g.warnPc {
		g.critPc = percentage
	}
}

func (g *Gauge) getBarColor(percentage float64) tcell.Color {
	if percentage < g.warnPc {
		return g.okColor
	} else if percentage < g.critPc {
		return g.warnColor
	}

	return g.critColor
}

// SetEmptyColor sets empty gauge color.
func (g *Gauge) SetEmptyColor(color tcell.Color) {
	g.emptyColor = color
}
