package chart

import (
	"fmt"
	"github.com/malivvan/cui"

	"github.com/gdamore/tcell/v2"
)

// UtilModeGauge represents utilisation mode gauge permitive.
type UtilModeGauge struct {
	*cui.Box
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
}

// NewUtilModeGauge returns new utilisation mode gauge permitive.
func NewUtilModeGauge() *UtilModeGauge {
	gauge := &UtilModeGauge{
		Box:        cui.NewBox(),
		pc:         gaugeMinPc,
		warnPc:     gaugeWarnPc,
		critPc:     gaugeCritPc,
		warnColor:  tcell.ColorOrange,
		critColor:  tcell.ColorRed,
		okColor:    tcell.ColorGreen,
		emptyColor: tcell.ColorWhite,
		labelColor: cui.Styles.PrimaryTextColor,
		label:      "",
	}

	return gauge
}

// SetLabel sets label for this primitive.
func (g *UtilModeGauge) SetLabel(label string) {
	g.label = label
}

// SetLabelColor sets label text color.
func (g *UtilModeGauge) SetLabelColor(color tcell.Color) {
	g.labelColor = color
}

// Focus is called when this primitive receives focus.
func (g *UtilModeGauge) Focus(delegate func(p cui.Primitive)) { //nolint:revive
	delegate(g)
}

// HasFocus returns whether or not this primitive has focus.
func (g *UtilModeGauge) HasFocus() bool {
	return g.Box.HasFocus()
}

// GetRect return primitive current rect.
func (g *UtilModeGauge) GetRect() (int, int, int, int) {
	return g.Box.GetRect()
}

// SetRect sets rect for this primitive.
func (g *UtilModeGauge) SetRect(x, y, width, height int) {
	g.Box.SetRect(x, y, width, height)
}

// SetValue update the gauge progress.
func (g *UtilModeGauge) SetValue(value float64) {
	if value <= float64(gaugeMaxPc) {
		g.pc = value
	}
}

// GetValue returns current gauge value.
func (g *UtilModeGauge) GetValue() float64 {
	return g.pc
}

// Draw draws this primitive onto the screen.
func (g *UtilModeGauge) Draw(screen tcell.Screen) {
	g.Box.Draw(screen)
	x, y, width, height := g.Box.GetInnerRect()
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

			cui.Print(screen, []byte(prgCell), x+labelWidth+i, y+j, 1, cui.AlignCenter, color)
		}
	}
	// draw label
	tY := y + (height / emptySpaceParts)
	if labelWidth > 0 {
		cui.Print(screen, []byte(g.label), x, tY, labelWidth, cui.AlignLeft, g.labelColor)
	}

	// draw percentage text
	cui.Print(screen, []byte(fmt.Sprintf("%6.2f%%", g.pc)),
		x+barWidth+labelWidth,
		tY,
		labelPCWidth,
		cui.AlignLeft,
		cui.Styles.PrimaryTextColor)
}

// SetWarnPercentage sets warning percentage start range.
func (g *UtilModeGauge) SetWarnPercentage(percentage float64) {
	if percentage > 0 && percentage < 100 {
		g.warnPc = percentage
	}
}

// SetCritPercentage sets critical percentage start range.
func (g *UtilModeGauge) SetCritPercentage(percentage float64) {
	if percentage > 0 && percentage < 100 && percentage > g.warnPc {
		g.critPc = percentage
	}
}

func (g *UtilModeGauge) getBarColor(percentage float64) tcell.Color {
	if percentage < g.warnPc {
		return g.okColor
	} else if percentage < g.critPc {
		return g.warnColor
	}

	return g.critColor
}

// SetEmptyColor sets empty gauge color.
func (g *UtilModeGauge) SetEmptyColor(color tcell.Color) {
	g.emptyColor = color
}
