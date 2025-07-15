package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/malivvan/cui/chart"
)

var _ = Describe("Barchart", Ordered, func() {
	var (
		app       *cui.Application
		headerBox *cui.Box
		barchart  *chart.BarChart
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = cui.NewApplication()
		headerBox = cui.NewBox()
		headerBox.SetBorder(true)
		barchart = chart.NewBarChart()
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := cui.NewFlex()
			appLayout.SetDirection(cui.FlexRow)
			appLayout.AddItem(headerBox, 1, 0, true)
			appLayout.AddItem(barchart, 50, 0, true)
			app.SetScreen(screen)
			app.SetRoot(appLayout, true)
			err := app.Run()
			if err != nil {
				panic(err)
			}
		}()
	})

	AfterAll(func() {
		app.Stop()
	})

	Describe("Focus", func() {
		It("checks primitivie focus", func() {
			app.SetFocus(headerBox)
			app.Draw()
			Expect(barchart.HasFocus()).To(Equal(false))

			app.SetFocus(barchart)
			app.Draw()
			Expect(barchart.HasFocus()).To(Equal(true))
		})
	})
})
