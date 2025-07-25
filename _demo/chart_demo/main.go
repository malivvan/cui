// Demo code for the bar chart primitive.
package main

import (
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/chart"
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	app := cui.NewApplication()

	// spinners
	spinners := []*chart.Spinner{

		chart.NewSpinner().SetStyle(chart.SpinnerDotsCircling),
		chart.NewSpinner().SetStyle(chart.SpinnerDotsUpDown),
		chart.NewSpinner().SetStyle(chart.SpinnerBounce),
		chart.NewSpinner().SetStyle(chart.SpinnerLine),

		chart.NewSpinner().SetStyle(chart.SpinnerCircleQuarters),
		chart.NewSpinner().SetStyle(chart.SpinnerSquareCorners),
		chart.NewSpinner().SetStyle(chart.SpinnerCircleHalves),
		chart.NewSpinner().SetStyle(chart.SpinnerCorners),

		chart.NewSpinner().SetStyle(chart.SpinnerArrows),
		chart.NewSpinner().SetStyle(chart.SpinnerHamburger),
		chart.NewSpinner().SetStyle(chart.SpinnerStack),
		chart.NewSpinner().SetStyle(chart.SpinnerStar),

		chart.NewSpinner().SetStyle(chart.SpinnerGrowHorizontal),
		chart.NewSpinner().SetStyle(chart.SpinnerGrowVertical),
		chart.NewSpinner().SetStyle(chart.SpinnerBoxBounce),
		chart.NewSpinner().SetCustomStyle([]rune{'🕛', '🕐', '🕑', '🕒', '🕓', '🕔', '🕕', '🕖', '🕗', '🕘', '🕙', '🕚'}),
	}

	spinnerRow := cui.NewFlex()
	spinnerRow.SetDirection(cui.FlexColumn)
	spinnerRow.SetBorder(true)
	spinnerRow.SetTitle("spinners")

	for _, spinner := range spinners {
		spinnerRow.AddItem(spinner, 0, 1, false)
	}

	// bar graph
	barGraph := newBarChart()
	barGraph.SetMaxValue(100)
	barGraph.SetAxesColor(tcell.ColorAntiqueWhite)
	barGraph.SetAxesLabelColor(tcell.ColorAntiqueWhite)

	// activity mode gauge
	amGauge := chart.NewActivityModeGauge()
	amGauge.SetTitle("activity mode gauge")
	amGauge.SetPgBgColor(tcell.ColorOrange)
	amGauge.SetBorder(true)

	// percentage mode gauge
	pmGauge := chart.NewPercentageModeGauge()
	pmGauge.SetTitle("percentage mode gauge")
	pmGauge.SetBorder(true)
	pmGauge.SetMaxValue(50)

	// cpu usage gauge
	cpuGauge := chart.NewUtilModeGauge()
	cpuGauge.SetLabel("cpu usage:   ")
	cpuGauge.SetLabelColor(tcell.ColorLightSkyBlue)
	cpuGauge.SetBorder(false)
	// memory usage gauge
	memGauge := chart.NewUtilModeGauge()
	memGauge.SetLabel("memory usage:")
	memGauge.SetLabelColor(tcell.ColorLightSkyBlue)
	memGauge.SetBorder(false)
	// swap usage gauge
	swapGauge := chart.NewUtilModeGauge()
	swapGauge.SetLabel("swap usage:  ")
	swapGauge.SetLabelColor(tcell.ColorLightSkyBlue)
	swapGauge.SetBorder(false)

	// utilisation flex
	utilFlex := cui.NewFlex()
	utilFlex.SetDirection(cui.FlexRow)
	utilFlex.AddItem(cpuGauge, 1, 0, false)
	utilFlex.AddItem(memGauge, 1, 0, false)
	utilFlex.AddItem(swapGauge, 1, 0, false)
	utilFlex.SetTitle("utilisation mode gauge")
	utilFlex.SetBorder(true)

	// plot (line charts)
	sinData := func() [][]float64 {
		n := 220
		data := make([][]float64, 2)
		data[0] = make([]float64, n)
		data[1] = make([]float64, n)
		for i := 0; i < n; i++ {
			data[0][i] = 1 + math.Sin(float64(i)/5)
			data[1][i] = 1 + math.Cos(float64(i)/5)
		}
		return data
	}()

	bmLineChart := newBrailleModeLineChart()
	bmLineChart.SetData(sinData)

	dmLineChart := newDotModeLineChart()

	sampleData1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sampleData2 := []float64{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	dotChartData := [][]float64{sampleData1}
	dotChartData[0] = append(dotChartData[0], sampleData2...)
	dotChartData[0] = append(dotChartData[0], sampleData1[:5]...)
	dotChartData[0] = append(dotChartData[0], sampleData2[5:]...)
	dotChartData[0] = append(dotChartData[0], sampleData1[:7]...)
	dotChartData[0] = append(dotChartData[0], sampleData2[3:]...)

	dmLineChart.SetData(dotChartData)

	// sparkline
	iowaitSparkline := chart.NewSparkline()
	iowaitSparkline.SetBorder(false)
	iowaitSparkline.SetDataTitle("Disk IO (iowait)")
	iowaitSparkline.SetDataTitleColor(tcell.ColorDarkOrange)
	iowaitSparkline.SetLineColor(tcell.ColorMediumPurple)

	systemSparkline := chart.NewSparkline()
	systemSparkline.SetBorder(false)
	systemSparkline.SetDataTitle("Disk IO (system)")
	systemSparkline.SetDataTitleColor(tcell.ColorDarkOrange)
	systemSparkline.SetLineColor(tcell.ColorSteelBlue)

	iowaitData := []float64{4, 2, 1, 6, 3, 9, 1, 4, 2, 15, 14, 9, 8, 6, 10, 13, 15, 12, 10, 5, 3, 6, 1, 7, 10, 10, 14, 13, 6}
	systemData := []float64{0, 0, 1, 2, 9, 5, 3, 1, 2, 0, 6, 3, 2, 2, 6, 8, 5, 2, 1, 5, 8, 6, 1, 4, 1, 1, 4, 3, 6}

	ioSparkLineData := func() []float64 {
		for i := 0; i < 5; i++ {
			iowaitData = append(iowaitData, iowaitData...)
		}

		return iowaitData
	}()

	systemSparklineData := func() []float64 {
		for i := 0; i < 5; i++ {
			systemData = append(systemData, systemData...)
		}

		return systemData
	}()

	iowaitSparkline.SetData(ioSparkLineData)
	systemSparkline.SetData(systemSparklineData)

	sparklineGroupLayout := cui.NewFlex()
	sparklineGroupLayout.SetDirection(cui.FlexColumn)
	sparklineGroupLayout.SetBorder(true)
	sparklineGroupLayout.SetTitle("sparkline")
	sparklineGroupLayout.AddItem(iowaitSparkline, 0, 1, false)
	sparklineGroupLayout.AddItem(cui.NewBox(), 1, 0, false)
	sparklineGroupLayout.AddItem(systemSparkline, 0, 1, false)

	// first row layout
	firstRowfirstCol := cui.NewFlex()
	firstRowfirstCol.SetDirection(cui.FlexRow)
	firstRowfirstCol.AddItem(barGraph, 0, 1, false)

	firstRowSecondCol := cui.NewFlex()
	firstRowSecondCol.SetDirection(cui.FlexRow)
	firstRowSecondCol.AddItem(amGauge, 0, 3, false)
	firstRowSecondCol.AddItem(pmGauge, 0, 3, false)
	firstRowSecondCol.AddItem(utilFlex, 0, 5, false)

	firstRow := cui.NewFlex()
	firstRow.SetDirection(cui.FlexColumn)
	firstRow.AddItem(firstRowfirstCol, 0, 1, false)
	firstRow.AddItem(firstRowSecondCol, 0, 1, false)

	// second row
	plotRowLayout := cui.NewFlex()
	plotRowLayout.SetDirection(cui.FlexColumn)
	plotRowLayout.AddItem(bmLineChart, 0, 1, false)
	plotRowLayout.AddItem(dmLineChart, 0, 1, false)

	screenLayout := cui.NewFlex()
	screenLayout.SetDirection(cui.FlexRow)
	screenLayout.AddItem(firstRow, 11, 0, false)
	screenLayout.AddItem(plotRowLayout, 15, 0, false)
	screenLayout.AddItem(sparklineGroupLayout, 6, 0, false)
	screenLayout.AddItem(spinnerRow, 3, 0, false)

	screenLayout.SetRect(0, 0, 100, 40)

	// upgrade datat functions
	moveDotChartData := func() {
		newData := append(dotChartData[0], dotChartData[0][0])
		dotChartData[0] = newData[1:]
	}

	moveDiskIOData := func() ([]float64, []float64) {

		newIOWaitData := ioSparkLineData[1:]
		newIOWaitData = append(newIOWaitData, ioSparkLineData[0])
		ioSparkLineData = newIOWaitData

		newSystemData := systemSparklineData[1:]
		newSystemData = append(newSystemData, systemSparklineData[0])
		systemSparklineData = newSystemData

		return newIOWaitData, newSystemData
	}

	moveSinData := func(data [][]float64) [][]float64 {
		newData := make([][]float64, 2)
		newData[0] = rotate(data[0], -1)
		newData[1] = rotate(data[1], -1)
		return newData
	}

	updateSpinner := func() {
		spinnerTick := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-spinnerTick.C:
				// update spinners
				for _, spinner := range spinners {
					spinner.Pulse()
				}
				// update gauge
				amGauge.Pulse()
				app.Draw()
			}
		}
	}

	// update screen ticker
	update := func() {
		value := 0
		maxValue := pmGauge.GetMaxValue()
		rand.Seed(time.Now().UnixNano())
		tick := time.NewTicker(500 * time.Millisecond)
		for {
			select {
			case <-tick.C:

				if value > maxValue {
					value = 0
				} else {
					value = value + 1
				}
				pmGauge.SetValue(value)

				randNum1 := float64(rand.Float64() * 100)
				randNum2 := float64(rand.Float64() * 100)
				randNum3 := float64(rand.Float64() * 100)
				randNum4 := float64(rand.Float64() * 100)

				barGraph.SetBarValue("eth0", int(randNum1))
				cpuGauge.SetValue(randNum1)
				barGraph.SetBarValue("eth1", int(randNum2))
				memGauge.SetValue(randNum2)
				barGraph.SetBarValue("eth2", int(randNum3))
				swapGauge.SetValue(randNum3)
				barGraph.SetBarValue("eth3", int(randNum4))

				// move line charts
				sinData = moveSinData(sinData)
				bmLineChart.SetData(sinData)

				moveDotChartData()
				dmLineChart.SetData(dotChartData)

				d1, d2 := moveDiskIOData()
				iowaitSparkline.SetData(d1)
				systemSparkline.SetData(d2)

				app.Draw()
			}
		}
	}

	go updateSpinner()
	go update()

	app.SetRoot(screenLayout, false)
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newDotModeLineChart() *chart.Plot {
	dmLineChart := chart.NewPlot()
	dmLineChart.SetBorder(true)
	dmLineChart.SetTitle("line chart (dot mode)")
	dmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorDarkOrange,
	})
	dmLineChart.SetAxesLabelColor(tcell.ColorGold)
	dmLineChart.SetAxesColor(tcell.ColorGold)
	dmLineChart.SetMarker(chart.PlotMarkerDot)
	dmLineChart.SetDotMarkerRune('\u25c9')

	return dmLineChart
}

func newBrailleModeLineChart() *chart.Plot {
	bmLineChart := chart.NewPlot()
	bmLineChart.SetBorder(true)
	bmLineChart.SetTitle("line chart (braille mode)")
	bmLineChart.SetLineColor([]tcell.Color{
		tcell.ColorSteelBlue,
		tcell.ColorGreen,
	})
	bmLineChart.SetMarker(chart.PlotMarkerBraille)

	return bmLineChart
}

func newBarChart() *chart.BarChart {
	barGraph := chart.NewBarChart()
	barGraph.SetBorder(true)
	barGraph.SetTitle("bar chart")
	barGraph.AddBar("eth0", 20, tcell.ColorBlue)
	barGraph.AddBar("eth1", 60, tcell.ColorRed)
	barGraph.AddBar("eth2", 80, tcell.ColorGreen)
	barGraph.AddBar("eth3", 100, tcell.ColorOrange)

	return barGraph
}

// Source: https://stackoverflow.com/questions/50833673/rotate-array-in-go/79079760#79079760
// rotate rotates the given slice by k positions to the left or right.
func rotate[T any](slice []T, k int) []T {
	if len(slice) == 0 {
		return slice
	}

	var r int
	if k > 0 {
		r = len(slice) - k%len(slice)
	} else {
		kAbs := int(math.Abs(float64(k)))
		r = kAbs % len(slice)
	}

	slice = append(slice[r:], slice[:r]...)

	return slice
}
