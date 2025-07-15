package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/malivvan/cui/chart"
)

var _ = Describe("GaugeUm", Ordered, func() {
	var (
		app       *cui.Application
		headerBox *cui.Box
		gaugeUm   *chart.UtilModeGauge
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = cui.NewApplication()
		headerBox = cui.NewBox()
		headerBox.SetBorder(true)
		gaugeUm = chart.NewUtilModeGauge()
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := cui.NewFlex()
			appLayout.SetDirection(cui.FlexRow)
			appLayout.AddItem(headerBox, 1, 0, true)
			appLayout.AddItem(gaugeUm, 50, 0, true)
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
			Expect(gaugeUm.HasFocus()).To(Equal(false))

			app.SetFocus(gaugeUm)
			app.Draw()
			// gauge will not get focus
			Expect(gaugeUm.HasFocus()).To(Equal(false))
		})
	})
})
