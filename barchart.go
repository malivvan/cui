package cui

import (
	"strconv"
	"sync"

	"github.com/gdamore/tcell/v2"
)

const (
	barChartYAxisLabelWidth = 2
	barGap                  = 2
	barWidth                = 3
)

// BarChartItem represents a single bar in bar
type BarChartItem struct {
	label string
	value int
	color tcell.Color
}

// BarChart represents bar chart primitive.
type BarChart struct {
	box *Box
	// bar items
	bars []BarChartItem
	// maximum value of bars
	maxVal int
	// barGap gap between two bars
	barGap int
	// barWidth width of bars
	barWidth int
	// hasBorder true if primitive has border
	hasBorder      bool
	axesColor      tcell.Color
	axesLabelColor tcell.Color

	mu sync.RWMutex
}

// NewBarChart returns a new bar chart primitive.
func NewBarChart() *BarChart {
	chart := &BarChart{
		box:            NewBox(),
		barGap:         barGap,
		barWidth:       barWidth,
		axesColor:      tcell.ColorDimGray,
		axesLabelColor: tcell.ColorDimGray,
	}

	return chart
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (bc *BarChart) set(setter func(b *BarChart)) *BarChart {
	bc.mu.Lock()
	setter(bc)
	bc.mu.Unlock()
	return bc
}

func (bc *BarChart) get(getter func(b *BarChart)) {
	bc.mu.RLock()
	getter(bc)
	bc.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this BarChart.
func (bc *BarChart) GetTitle() string {
	return bc.box.GetTitle()
}

// SetTitle sets the title of this BarChart.
func (bc *BarChart) SetTitle(title string) *BarChart {
	bc.box.SetTitle(title)
	return bc
}

// GetTitleAlign returns the title alignment of this BarChart.
func (bc *BarChart) GetTitleAlign() int {
	return bc.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this BarChart.
func (bc *BarChart) SetTitleAlign(align int) *BarChart {
	bc.box.SetTitleAlign(align)
	return bc
}

// GetBorder returns whether this BarChart has a border.
func (bc *BarChart) GetBorder() bool {
	return bc.box.GetBorder()
}

// GetBorderColor returns the border color of this BarChart.
func (bc *BarChart) GetBorderColor() tcell.Color {
	return bc.box.GetBorderColor()
}

// SetBorderColor sets the border color of this BarChart.
func (bc *BarChart) SetBorderColor(color tcell.Color) *BarChart {
	bc.box.SetBorderColor(color)
	return bc
}

// GetBorderAttributes returns the border attributes of this BarChart.
func (bc *BarChart) GetBorderAttributes() tcell.AttrMask {
	return bc.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this BarChart.
func (bc *BarChart) SetBorderAttributes(attr tcell.AttrMask) *BarChart {
	bc.box.SetBorderAttributes(attr)
	return bc
}

// GetBorderColorFocused returns the border color of this BarChart when focused.
func (bc *BarChart) GetBorderColorFocused() tcell.Color {
	return bc.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this BarChart when focused.
func (bc *BarChart) SetBorderColorFocused(color tcell.Color) *BarChart {
	bc.box.SetBorderColorFocused(color)
	return bc
}

// GetTitleColor returns the title color of this BarChart.
func (bc *BarChart) GetTitleColor() tcell.Color {
	return bc.box.GetTitleColor()
}

// SetTitleColor sets the title color of this BarChart.
func (bc *BarChart) SetTitleColor(color tcell.Color) *BarChart {
	bc.box.SetTitleColor(color)
	return bc
}

// GetDrawFunc returns the custom draw function of this BarChart.
func (bc *BarChart) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return bc.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this BarChart.
func (bc *BarChart) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *BarChart {
	bc.box.SetDrawFunc(handler)
	return bc
}

// ShowFocus sets whether this BarChart should show a focus indicator when focused.
func (bc *BarChart) ShowFocus(showFocus bool) *BarChart {
	bc.box.ShowFocus(showFocus)
	return bc
}

// GetMouseCapture returns the mouse capture function of this BarChart.
func (bc *BarChart) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return bc.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this BarChart.
func (bc *BarChart) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *BarChart {
	bc.box.SetMouseCapture(capture)
	return bc
}

// GetBackgroundColor returns the background color of this BarChart.
func (bc *BarChart) GetBackgroundColor() tcell.Color {
	return bc.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this BarChart.
func (bc *BarChart) SetBackgroundColor(color tcell.Color) *BarChart {
	bc.box.SetBackgroundColor(color)
	return bc
}

// GetBackgroundTransparent returns whether the background of this BarChart is transparent.
func (bc *BarChart) GetBackgroundTransparent() bool {
	return bc.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this BarChart is transparent.
func (bc *BarChart) SetBackgroundTransparent(transparent bool) *BarChart {
	bc.box.SetBackgroundTransparent(transparent)
	return bc
}

// GetInputCapture returns the input capture function of this BarChart.
func (bc *BarChart) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return bc.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this BarChart.
func (bc *BarChart) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *BarChart {
	bc.box.SetInputCapture(capture)
	return bc
}

// GetPadding returns the padding of this BarChart.
func (bc *BarChart) GetPadding() (top, bottom, left, right int) {
	return bc.box.GetPadding()
}

// SetPadding sets the padding of this BarChart.
func (bc *BarChart) SetPadding(top, bottom, left, right int) *BarChart {
	bc.box.SetPadding(top, bottom, left, right)
	return bc
}

// InRect returns whether the given screen coordinates are within this BarChart.
func (bc *BarChart) InRect(x, y int) bool {
	return bc.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this BarChart.
func (bc *BarChart) GetInnerRect() (x, y, width, height int) {
	return bc.box.GetInnerRect()
}

func (bc *BarChart) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return bc.box.InputHandler()

}
func (bc *BarChart) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return bc.box.MouseHandler()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the BarChart is preserved.
func (bc *BarChart) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return bc.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the BarChart is preserved.
func (bc *BarChart) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return bc.box.WrapMouseHandler(mouseHandler)
}

// GetVisible returns whether this BarChart is visible.
func (bc *BarChart) GetVisible() bool {
	return bc.box.GetVisible()
}

// SetVisible sets whether this BarChart is visible.
func (bc *BarChart) SetVisible(visible bool) {
	bc.box.SetVisible(visible)
}

// Focus is called when this BarChart receives focus.
func (bc *BarChart) Focus(delegate func(p Widget)) {
	bc.box.Focus(delegate)
}

// HasFocus returns whether this BarChart has focus.
func (bc *BarChart) HasFocus() bool {
	return bc.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this BarChart.
func (bc *BarChart) GetFocusable() Focusable {
	return bc.box.GetFocusable()
}

// Blur is called when this BarChart loses focus.
func (bc *BarChart) Blur() {
	bc.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

// Draw draws this primitive onto the screen.
func (bc *BarChart) Draw(screen tcell.Screen) { //nolint:funlen,cyclop
	if !bc.GetVisible() {
		return
	}

	bc.box.Draw(screen)

	bc.mu.Lock()
	defer bc.mu.Unlock()

	x, y, width, height := bc.box.GetInnerRect()

	maxValY := y + 1
	xAxisStartY := y + height - 2 //nolint:mnd
	barStartY := y + height - 3   //nolint:mnd
	borderPadding := 0

	if bc.hasBorder {
		borderPadding = 1
	}
	// set max value if not set
	bc.initMaxValue()
	maxValueSr := strconv.Itoa(bc.maxVal)
	maxValLength := len(maxValueSr) + 1

	if maxValLength < barChartYAxisLabelWidth {
		maxValLength = barChartYAxisLabelWidth
	}

	// draw Y axis line
	drawLine(screen,
		x+maxValLength,
		y+borderPadding,
		height-borderPadding-1,
		verticalLine, bc.axesColor)

	// draw X axis line
	drawLine(screen,
		x+maxValLength+1,
		xAxisStartY,
		width-borderPadding-maxValLength-1,
		horizontalLine, bc.axesColor)

	PrintJoinedSemigraphics(screen,
		x+maxValLength,
		xAxisStartY,
		BoxDrawingsLightUpAndRight, bc.axesColor)

	PrintJoinedSemigraphics(screen, x+maxValLength-1, xAxisStartY, '0', bc.axesLabelColor)

	mxValRune := []rune(maxValueSr)
	for i := range mxValRune {
		PrintJoinedSemigraphics(screen, x+borderPadding+i, maxValY, mxValRune[i], bc.axesLabelColor)
	}

	// draw bars
	startX := x + maxValLength + bc.barGap
	labelY := y + height - 1
	valueMaxHeight := barStartY - maxValY

	for _, item := range bc.bars {
		if startX > x+width {
			return
		}
		// set labels
		r := []rune(item.label)
		for j := range r {
			PrintJoinedSemigraphics(screen, startX+j, labelY, r[j], item.color)
		}
		// bar style
		barHeight := bc.getHeight(valueMaxHeight, item.value)

		for k := range barHeight {
			for l := range bc.barWidth {
				PrintJoinedSemigraphics(screen, startX+l, barStartY-k, fullBlockRune, item.color)
			}
		}
		// bar value
		vSt := strconv.Itoa(item.value)
		vRune := []rune(vSt)

		for i := range vRune {
			PrintJoinedSemigraphics(screen, startX+i, barStartY-barHeight, vRune[i], item.color)
		}

		// calculate next startX for next bar
		rWidth := len(r)
		if rWidth < bc.barWidth {
			rWidth = bc.barWidth
		}

		startX = startX + bc.barGap + rWidth
	}
}

// SetBorder sets border for this primitive.
func (bc *BarChart) SetBorder(status bool) *BarChart {
	bc.hasBorder = status
	bc.box.SetBorder(status)
	return bc
}

// GetRect return primitive current rect.
func (bc *BarChart) GetRect() (int, int, int, int) {
	return bc.box.GetRect()
}

// SetRect sets rect for this primitive.
func (bc *BarChart) SetRect(x, y, width, height int) {
	bc.box.SetRect(x, y, width, height)
}

// SetMaxValue sets maximum value of bars.
func (bc *BarChart) SetMaxValue(maxValue int) {
	bc.maxVal = maxValue
}

// SetAxesColor sets axes x and y lines color.
func (bc *BarChart) SetAxesColor(color tcell.Color) {
	bc.axesColor = color
}

// SetAxesLabelColor sets axes x and y label color.
func (bc *BarChart) SetAxesLabelColor(color tcell.Color) {
	bc.axesLabelColor = color
}

// AddBar adds new bar item to the bar chart primitive.
func (bc *BarChart) AddBar(label string, value int, color tcell.Color) {
	bc.bars = append(bc.bars, BarChartItem{
		label: label,
		value: value,
		color: color,
	})
}

// RemoveBar removes a bar item from the bar
func (bc *BarChart) RemoveBar(label string) {
	bars := bc.bars[:0]

	for _, barItem := range bc.bars {
		if barItem.label != label {
			bars = append(bars, barItem)
		}
	}

	bc.bars = bars
}

// SetBarValue sets bar values.
func (bc *BarChart) SetBarValue(name string, value int) {
	for i := range bc.bars {
		if bc.bars[i].label == name {
			bc.bars[i].value = value
		}
	}
}

func (bc *BarChart) getHeight(maxHeight int, value int) int {
	if value >= bc.maxVal {
		return maxHeight
	}

	height := (value * maxHeight) / bc.maxVal

	return height
}

func (bc *BarChart) initMaxValue() {
	// set max value if not set
	if bc.maxVal == 0 {
		for _, b := range bc.bars {
			if b.value > bc.maxVal {
				bc.maxVal = b.value
			}
		}
	}
}
