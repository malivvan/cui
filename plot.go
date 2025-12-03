package cui

import (
	"fmt"
	"image"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Marker represents plot drawing marker (braille or dot).
type Marker uint

const (
	PlotMarkerBraille Marker = iota
	PlotMarkerDot
)

// PlotYAxisLabelDataType represents plot y-axis type (integer or float).
type PlotYAxisLabelDataType uint

const (
	PlotYAxisLabelDataInt PlotYAxisLabelDataType = iota
	PlotYAxisLabelDataFloat
)

// PlotType represents plot type (line chart or scatter).
type PlotType uint

const (
	PlotTypeLineChart PlotType = iota
	PlotTypeScatter
)

const (
	plotHorizontalScale   = 1
	plotXAxisLabelsHeight = 1
	plotXAxisLabelsGap    = 2
	plotYAxisLabelsGap    = 1

	gapRune = " "
)

type brailleCell struct {
	cRune rune
	color tcell.Color
}

// Plot represents a plot primitive used for different charts.
type Plot struct {
	box  *Box
	data [][]float64
	// maxVal is the maximum y-axis (vertical) value found in any of the lines in the data set.
	maxVal float64
	// minVal is the minimum y-axis (vertical) value found in any of the lines in the data set.
	minVal             float64
	marker             Marker
	plotType           PlotType
	dotMarkerRune      rune
	lineColors         []tcell.Color
	axesColor          tcell.Color
	axesLabelColor     tcell.Color
	drawAxes           bool
	drawXAxisLabel     bool
	xAxisLabelFunc     func(int) string
	drawYAxisLabel     bool
	yAxisLabelDataType PlotYAxisLabelDataType
	yAxisAutoScaleMin  bool
	yAxisAutoScaleMax  bool
	brailleCellMap     map[image.Point]brailleCell
	mu                 sync.RWMutex
}

func (p *Plot) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.InputHandler()
}

func (p *Plot) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.box.MouseHandler()
}

// NewPlot returns a plot widget.
func NewPlot() *Plot {
	return &Plot{
		box:                NewBox(),
		marker:             PlotMarkerDot,
		plotType:           PlotTypeLineChart,
		dotMarkerRune:      dotRune,
		axesColor:          tcell.ColorDimGray,
		axesLabelColor:     tcell.ColorDimGray,
		drawAxes:           true,
		drawXAxisLabel:     true,
		xAxisLabelFunc:     strconv.Itoa,
		drawYAxisLabel:     true,
		yAxisLabelDataType: PlotYAxisLabelDataFloat,
		yAxisAutoScaleMin:  false,
		yAxisAutoScaleMax:  true,
		lineColors: []tcell.Color{
			tcell.ColorSteelBlue,
		},
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (p *Plot) set(setter func(p *Plot)) *Plot {
	p.mu.Lock()
	setter(p)
	p.mu.Unlock()
	return p
}

func (p *Plot) get(getter func(p *Plot)) {
	p.mu.RLock()
	getter(p)
	p.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Plot.
func (p *Plot) GetTitle() string {
	return p.box.GetTitle()
}

// SetTitle sets the title of this Plot.
func (p *Plot) SetTitle(title string) *Plot {
	p.box.SetTitle(title)
	return p
}

// GetTitleAlign returns the title alignment of this Plot.
func (p *Plot) GetTitleAlign() int {
	return p.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Plot.
func (p *Plot) SetTitleAlign(align int) *Plot {
	p.box.SetTitleAlign(align)
	return p
}

// GetBorder returns whether this Plot has a border.
func (p *Plot) GetBorder() bool {
	return p.box.GetBorder()
}

// SetBorder sets whether this Plot has a border.
func (p *Plot) SetBorder(show bool) *Plot {
	p.box.SetBorder(show)
	return p
}

// GetBorderColor returns the border color of this Plot.
func (p *Plot) GetBorderColor() tcell.Color {
	return p.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Plot.
func (p *Plot) SetBorderColor(color tcell.Color) *Plot {
	p.box.SetBorderColor(color)
	return p
}

// GetBorderAttributes returns the border attributes of this Plot.
func (p *Plot) GetBorderAttributes() tcell.AttrMask {
	return p.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Plot.
func (p *Plot) SetBorderAttributes(attr tcell.AttrMask) *Plot {
	p.box.SetBorderAttributes(attr)
	return p
}

// GetBorderColorFocused returns the border color of this Plot when focused.
func (p *Plot) GetBorderColorFocused() tcell.Color {
	return p.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Plot when focused.
func (p *Plot) SetBorderColorFocused(color tcell.Color) *Plot {
	p.box.SetBorderColorFocused(color)
	return p
}

// GetTitleColor returns the title color of this Plot.
func (p *Plot) GetTitleColor() tcell.Color {
	return p.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Plot.
func (p *Plot) SetTitleColor(color tcell.Color) *Plot {
	p.box.SetTitleColor(color)
	return p
}

// GetDrawFunc returns the custom draw function of this Plot.
func (p *Plot) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return p.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Plot.
func (p *Plot) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Plot {
	p.box.SetDrawFunc(handler)
	return p
}

// ShowFocus sets whether this Plot should show a focus indicator when focused.
func (p *Plot) ShowFocus(showFocus bool) *Plot {
	p.box.ShowFocus(showFocus)
	return p
}

// GetMouseCapture returns the mouse capture function of this Plot.
func (p *Plot) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return p.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Plot.
func (p *Plot) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Plot {
	p.box.SetMouseCapture(capture)
	return p
}

// GetBackgroundColor returns the background color of this Plot.
func (p *Plot) GetBackgroundColor() tcell.Color {
	return p.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Plot.
func (p *Plot) SetBackgroundColor(color tcell.Color) *Plot {
	p.box.SetBackgroundColor(color)
	return p
}

// GetBackgroundTransparent returns whether the background of this Plot is transparent.
func (p *Plot) GetBackgroundTransparent() bool {
	return p.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Plot is transparent.
func (p *Plot) SetBackgroundTransparent(transparent bool) *Plot {
	p.box.SetBackgroundTransparent(transparent)
	return p
}

// GetInputCapture returns the input capture function of this Plot.
func (p *Plot) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return p.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Plot.
func (p *Plot) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Plot {
	p.box.SetInputCapture(capture)
	return p
}

// GetPadding returns the padding of this Plot.
func (p *Plot) GetPadding() (top, bottom, left, right int) {
	return p.box.GetPadding()
}

// SetPadding sets the padding of this Plot.
func (p *Plot) SetPadding(top, bottom, left, right int) *Plot {
	p.box.SetPadding(top, bottom, left, right)
	return p
}

// InRect returns whether the given screen coordinates are within this Plot.
func (p *Plot) InRect(x, y int) bool {
	return p.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Plot.
func (p *Plot) GetInnerRect() (x, y, width, height int) {
	return p.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Plot is preserved.
func (p *Plot) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Plot is preserved.
func (p *Plot) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Plot.
func (p *Plot) GetRect() (x, y, width, height int) {
	return p.box.GetRect()
}

// SetRect sets the rectangle occupied by this Plot.
func (p *Plot) SetRect(x, y, width, height int) {
	p.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Plot is visible.
func (p *Plot) GetVisible() bool {
	return p.box.GetVisible()
}

// SetVisible sets whether this Plot is visible.
func (p *Plot) SetVisible(visible bool) {
	p.box.SetVisible(visible)
}

// Focus is called when this Plot receives focus.
func (p *Plot) Focus(delegate func(p Widget)) {
	p.box.Focus(delegate)
}

// HasFocus returns whether this Plot has focus.
func (p *Plot) HasFocus() bool {
	return p.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Plot.
func (p *Plot) GetFocusable() Focusable {
	return p.box.GetFocusable()
}

// Blur is called when this Plot loses focus.
func (p *Plot) Blur() {
	p.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

// Draw draws this primitive onto the screen.
func (plot *Plot) Draw(screen tcell.Screen) {
	if !plot.GetVisible() {
		return
	}

	plot.box.Draw(screen)

	switch plot.marker {
	case PlotMarkerDot:
		plot.drawDotMarkerToScreen(screen)
	case PlotMarkerBraille:
		plot.drawBrailleMarkerToScreen(screen)
	}

	plot.drawAxesToScreen(screen)
}

// SetLineColor sets chart line color.
func (plot *Plot) SetLineColor(color []tcell.Color) {
	plot.lineColors = color
}

// SetYAxisLabelDataType sets Y axis label data type (integer or float).
func (plot *Plot) SetYAxisLabelDataType(dataType PlotYAxisLabelDataType) {
	plot.yAxisLabelDataType = dataType
}

// SetYAxisAutoScaleMin enables YAxis min value autoscale.
func (plot *Plot) SetYAxisAutoScaleMin(autoScale bool) {
	plot.yAxisAutoScaleMin = autoScale
}

// SetYAxisAutoScaleMax enables YAxix max value autoscale.
func (plot *Plot) SetYAxisAutoScaleMax(autoScale bool) {
	plot.yAxisAutoScaleMax = autoScale
}

// SetAxesColor sets axes x and y lines color.
func (plot *Plot) SetAxesColor(color tcell.Color) {
	plot.axesColor = color
}

// SetAxesLabelColor sets axes x and y label color.
func (plot *Plot) SetAxesLabelColor(color tcell.Color) {
	plot.axesLabelColor = color
}

// SetDrawAxes set true in order to draw axes to screen.
func (plot *Plot) SetDrawAxes(draw bool) {
	plot.drawAxes = draw
}

// SetDrawXAxisLabel set true in order to draw x-axis label to screen.
func (plot *Plot) SetDrawXAxisLabel(draw bool) {
	plot.drawXAxisLabel = draw
}

// SetXAxisLabelFunc sets x axis label function.
func (plot *Plot) SetXAxisLabelFunc(f func(int) string) {
	plot.xAxisLabelFunc = f
}

// SetDrawYAxisLabel set true in order to draw y-axis label to screen.
func (plot *Plot) SetDrawYAxisLabel(draw bool) {
	plot.drawYAxisLabel = draw
}

// SetMarker sets marker type braille or dot mode.
func (plot *Plot) SetMarker(marker Marker) {
	plot.marker = marker
}

// SetPlotType sets plot type (line chart or scatter).
func (plot *Plot) SetPlotType(ptype PlotType) {
	plot.plotType = ptype
}

// SetData sets plot data.
func (plot *Plot) SetData(data [][]float64) {
	plot.mu.Lock()
	defer plot.mu.Unlock()

	plot.brailleCellMap = make(map[image.Point]brailleCell)
	plot.data = data

	if plot.yAxisAutoScaleMax {
		plot.maxVal = getMaxFloat64From2dSlice(data)
	}

	if plot.yAxisAutoScaleMin {
		plot.minVal = getMinFloat64From2dSlice(data)
	}
}

func (plot *Plot) SetMaxVal(maxVal float64) {
	plot.maxVal = maxVal
}

func (plot *Plot) SetMinVal(minVal float64) {
	plot.minVal = minVal
}

func (plot *Plot) SetYRange(minVal float64, maxVal float64) {
	plot.minVal = minVal
	plot.maxVal = maxVal
}

// SetDotMarkerRune sets dot marker rune.
func (plot *Plot) SetDotMarkerRune(r rune) {
	plot.dotMarkerRune = r
}

// Figure out the text width necessary to display the largest data value.
func (plot *Plot) getYAxisLabelsWidth() int {
	return len(fmt.Sprintf("%.2f", plot.maxVal))
}

// GetPlotRect returns the rect for the inner part of the plot, ie not including axes.
func (plot *Plot) GetPlotRect() (int, int, int, int) {
	x, y, width, height := plot.box.GetInnerRect()
	plotYAxisLabelsWidth := plot.getYAxisLabelsWidth()

	if plot.drawAxes {
		x = x + plotYAxisLabelsWidth + 1
		width = width - plotYAxisLabelsWidth - 1
		height = height - plotXAxisLabelsHeight - 1
	} else {
		x++
		width--
	}

	return x, y, width, height
}

func (plot *Plot) getData() [][]float64 {
	plot.mu.Lock()
	data := plot.data
	plot.mu.Unlock()

	return data
}

func (plot *Plot) drawAxesToScreen(screen tcell.Screen) {
	if !plot.drawAxes {
		return
	}

	x, y, width, height := plot.box.GetInnerRect()
	plotYAxisLabelsWidth := plot.getYAxisLabelsWidth()

	// draw Y axis line
	drawLine(screen,
		x+plotYAxisLabelsWidth,
		y,
		height-plotXAxisLabelsHeight-1,
		verticalLine, plot.axesColor)

	// draw X axis line
	drawLine(screen,
		x+plotYAxisLabelsWidth+1,
		y+height-plotXAxisLabelsHeight-1,
		width-plotYAxisLabelsWidth-1,
		horizontalLine, plot.axesColor)

	PrintJoinedSemigraphics(screen,
		x+plotYAxisLabelsWidth,
		y+height-plotXAxisLabelsHeight-1,
		BoxDrawingsLightUpAndRight, plot.axesColor)

	if plot.drawXAxisLabel {
		plot.drawXAxisLabelsToScreen(screen, plotYAxisLabelsWidth, x, y, width, height)
	}

	if plot.drawYAxisLabel {
		plot.drawYAxisLabelsToScreen(screen, plotYAxisLabelsWidth, x, y, height)
	}
}

func (plot *Plot) drawXAxisLabelsToScreen(
	screen tcell.Screen, plotYAxisLabelsWidth int, x int, y int, width int, height int,
) {
	xAxisAreaStartX := x + plotYAxisLabelsWidth + 1
	xAxisAreaEndX := x + width
	xAxisAvailableWidth := xAxisAreaEndX - xAxisAreaStartX

	labelMap := map[int]string{}
	labelStartMap := map[int]int{}

	maxDataPoints := 0
	for _, d := range plot.data {
		maxDataPoints = max(maxDataPoints, len(d))
	}

	// determine the width needed for the largest label
	maxXAxisLabelWidth := 0

	for _, d := range plot.data {
		for i := range d {
			label := plot.xAxisLabelFunc(i)
			labelMap[i] = label
			maxXAxisLabelWidth = max(maxXAxisLabelWidth, len(label))
		}
	}

	// determine the start position for each label, if they were
	// to be centered below the data point.
	// Note: not all of these labels will be printed, as they would
	// overlap with each other
	for i, label := range labelMap {
		expectedLabelWidth := len(label)
		if i == 0 {
			expectedLabelWidth += plotXAxisLabelsGap / 2 //nolint:mnd
		} else {
			expectedLabelWidth += plotXAxisLabelsGap
		}

		currentLabelStart := i - int(math.Round(float64(expectedLabelWidth)/2)) //nolint:mnd
		labelStartMap[i] = currentLabelStart
	}

	// print the labels, skipping those that would overlap,
	// stopping when there is no more space
	lastUsedLabelEnd := math.MinInt
	initialOffset := xAxisAreaStartX

	for i := range maxDataPoints {
		labelStart := labelStartMap[i]
		if labelStart < lastUsedLabelEnd {
			// the label would overlap with the previous label
			continue
		}

		rawLabel := labelMap[i]
		labelWithGap := rawLabel

		if i == 0 {
			labelWithGap += strings.Repeat(gapRune, plotXAxisLabelsGap/2) //nolint:mnd
		} else {
			labelWithGap = strings.Repeat(gapRune, plotXAxisLabelsGap/2) + labelWithGap + strings.Repeat(gapRune, plotXAxisLabelsGap/2) //nolint:lll,mnd
		}

		expectedLabelWidth := len(labelWithGap)
		remainingWidth := xAxisAvailableWidth - labelStart

		if expectedLabelWidth > remainingWidth {
			// the label would be too long to fit in the remaining space
			if expectedLabelWidth-1 <= remainingWidth {
				// if we omit the last gap, it fits, so we draw that before stopping
				labelWithoutGap := labelWithGap[:len(labelWithGap)-1]
				plot.printXAxisLabel(screen, labelWithoutGap, initialOffset+labelStart, y+height-plotXAxisLabelsHeight)
			}

			break
		}

		lastUsedLabelEnd = labelStart + expectedLabelWidth
		plot.printXAxisLabel(screen, labelWithGap, initialOffset+labelStart, y+height-plotXAxisLabelsHeight)
	}
}

func (plot *Plot) printXAxisLabel(screen tcell.Screen, label string, x, y int) {
	Print(screen, []byte(label), x, y, len(label), AlignLeft, plot.axesLabelColor)
}

func (plot *Plot) drawYAxisLabelsToScreen(screen tcell.Screen, plotYAxisLabelsWidth int, x int, y int, height int) {
	verticalOffset := plot.minVal
	verticalScale := (plot.maxVal - plot.minVal) / float64(height-plotXAxisLabelsHeight-1)
	previousLabel := ""

	for i := 0; i*(plotYAxisLabelsGap+1) < height-1; i++ {
		var label string
		if plot.yAxisLabelDataType == PlotYAxisLabelDataFloat {
			label = fmt.Sprintf("%.2f", float64(i)*verticalScale*(plotYAxisLabelsGap+1)+verticalOffset)
		} else {
			label = strconv.Itoa(int(float64(i)*verticalScale*(plotYAxisLabelsGap+1) + verticalOffset))
		}

		// Prevent same label being shown twice.
		// Mainly relevant for integer labels with small data sets (in value)
		if label == previousLabel {
			continue
		}

		previousLabel = label

		Print(screen,
			[]byte(label),
			x,
			y+height-(i*(plotYAxisLabelsGap+1))-2, //nolint:mnd
			plotYAxisLabelsWidth,
			AlignLeft, plot.axesLabelColor)
	}
}

func (plot *Plot) drawDotMarkerToScreen(screen tcell.Screen) {
	x, y, width, height := plot.GetPlotRect()
	chartData := plot.getData()
	verticalOffset := -plot.minVal

	switch plot.plotType {
	case PlotTypeLineChart:
		for i, line := range chartData {

			for j := 0; j < len(line) && j*plotHorizontalScale < width; j++ {
				val := line[j]
				if math.IsNaN(val) {
					continue
				}

				lHeight := int(((val + verticalOffset) / plot.maxVal) * float64(height-1))
				if lHeight > height {
					continue
				}

				if (x+(j*plotHorizontalScale) < x+width) && (y+height-1-lHeight < y+height) {
					PrintJoinedSemigraphics(screen, x+(j*plotHorizontalScale), y+height-1-lHeight, plot.dotMarkerRune, plot.lineColors[i])
				}
			}
		}

	case PlotTypeScatter:
		for i, line := range chartData {

			for j, val := range line {
				if math.IsNaN(val) {
					continue
				}

				lHeight := int(((val + verticalOffset) / plot.maxVal) * float64(height-1))
				if lHeight > height {
					continue
				}

				if (x+(j*plotHorizontalScale) < x+width) && (y+height-1-lHeight < y+height) {
					PrintJoinedSemigraphics(screen, x+(j*plotHorizontalScale), y+height-1-lHeight, plot.dotMarkerRune, plot.lineColors[i])
				}
			}
		}
	}
}

func (plot *Plot) drawBrailleMarkerToScreen(screen tcell.Screen) {
	x, y, width, height := plot.GetPlotRect()

	plot.calcBrailleLines()

	// print to screen
	for point, cell := range plot.getBrailleCells() {
		if point.X < x+width && point.Y < y+height {
			PrintJoinedSemigraphics(screen, point.X, point.Y, cell.cRune, cell.color)
		}
	}
}

func calcDataPointHeight(val, maxVal, minVal float64, height int) int {
	return int(((val - minVal) / (maxVal - minVal)) * float64(height-1))
}

func calcDataPointHeightIfInBounds(val float64, maxVal float64, minVal float64, height int) (int, bool) {
	if math.IsNaN(val) {
		return 0, false
	}

	result := calcDataPointHeight(val, maxVal, minVal, height)
	if (val > maxVal) || (val < minVal) || (result > height) {
		return result, false
	}

	return result, true
}

func (plot *Plot) calcBrailleLines() {
	x, y, _, height := plot.GetPlotRect()
	chartData := plot.getData()

	for i, line := range chartData {
		if len(line) <= 1 {
			continue
		}

		previousHeight := 0
		lastValWasOk := false

		for j, val := range line {
			lHeight, currentValIsOk := calcDataPointHeightIfInBounds(val, plot.maxVal, plot.minVal, height)

			if !lastValWasOk && !currentValIsOk {
				// nothing valid to draw, skip to next data point
				continue
			}

			if !lastValWasOk { //nolint:gocritic
				// current data point is single valid data point, draw it individually
				plot.setBraillePoint(
					calcBraillePoint(x, j+1, y, height, lHeight),
					plot.lineColors[i],
				)
			} else if !currentValIsOk {
				// last data point was single valid data point, draw it individually
				plot.setBraillePoint(
					calcBraillePoint(x, j, y, height, previousHeight),
					plot.lineColors[i],
				)
			} else {
				// we have two valid data points, draw a line between them
				plot.setBrailleLine(
					calcBraillePoint(x, j, y, height, previousHeight),
					calcBraillePoint(x, j+1, y, height, lHeight),
					plot.lineColors[i],
				)
			}

			lastValWasOk = currentValIsOk
			previousHeight = lHeight
		}
	}
}

func calcBraillePoint(x, j, y, maxY, height int) image.Point {
	return image.Pt(
		(x+(j*plotHorizontalScale))*2, //nolint:mnd
		(y+maxY-height-1)*4,           //nolint:mnd
	)
}

func (plot *Plot) setBraillePoint(p image.Point, color tcell.Color) {
	if p.X < 0 || p.Y < 0 {
		return
	}

	point := image.Pt(p.X/2, p.Y/4) //nolint:mnd
	plot.brailleCellMap[point] = brailleCell{
		plot.brailleCellMap[point].cRune | brailleRune[p.Y%4][p.X%2],
		color,
	}
}

func (plot *Plot) setBrailleLine(p0, p1 image.Point, color tcell.Color) {
	for _, p := range plot.brailleLine(p0, p1) {
		plot.setBraillePoint(p, color)
	}
}

func (plot *Plot) getBrailleCells() map[image.Point]brailleCell {
	cellMap := make(map[image.Point]brailleCell)
	for point, cvCell := range plot.brailleCellMap {
		cellMap[point] = brailleCell{cvCell.cRune + brailleOffsetRune, cvCell.color}
	}

	return cellMap
}

func (plot *Plot) brailleLine(p0, p1 image.Point) []image.Point {
	points := []image.Point{}
	leftPoint, rightPoint := p0, p1

	if leftPoint.X > rightPoint.X {
		leftPoint, rightPoint = rightPoint, leftPoint
	}

	xDistance := absInt(leftPoint.X - rightPoint.X)
	yDistance := absInt(leftPoint.Y - rightPoint.Y)
	slope := float64(yDistance) / float64(xDistance)
	slopeSign := 1

	if rightPoint.Y < leftPoint.Y {
		slopeSign = -1
	}

	targetYCoordinate := float64(leftPoint.Y)
	currentYCoordinate := leftPoint.Y

	for i := leftPoint.X; i < rightPoint.X; i++ {
		points = append(points, image.Pt(i, currentYCoordinate))
		targetYCoordinate += slope * float64(slopeSign)

		for currentYCoordinate != int(targetYCoordinate) {
			points = append(points, image.Pt(i, currentYCoordinate))

			currentYCoordinate += slopeSign
		}
	}

	return points
}
