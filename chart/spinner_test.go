package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/malivvan/cui/chart"
)

var _ = Describe("Spinner", Ordered, func() {
	var (
		app       *cui.Application
		headerBox *cui.Box
		spinner   *chart.Spinner
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = cui.NewApplication()
		headerBox = cui.NewBox()
		headerBox.SetBorder(true)
		spinner = chart.NewSpinner()
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := cui.NewFlex()
			appLayout.SetDirection(cui.FlexRow)
			appLayout.AddItem(headerBox, 1, 0, true)
			appLayout.AddItem(spinner, 50, 0, true)
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
			Expect(spinner.HasFocus()).To(Equal(false))

			app.SetFocus(spinner)
			app.Draw()
			Expect(spinner.HasFocus()).To(Equal(true))
		})
	})

	Describe("Style", func() {
		It("checks style", func() {
			spinner.SetStyle(chart.SpinnerGrowHorizontal)
			spinner.Reset()
			app.Draw()

			prune, _, _, _ := screen.GetContent(0, 1)
			Expect(prune).To(Equal('▉'))

			spinner.Pulse()
			app.Draw()
			prune, _, _, _ = screen.GetContent(0, 1)
			Expect(prune).To(Equal('▊'))
		})
	})

	Describe("CustomStyle", func() {
		It("checks custom style", func() {
			customStyle := []rune{'\u2705', '\u274C'}
			spinner.SetCustomStyle(customStyle)
			spinner.Reset()

			app.Draw()
			prune, _, _, _ := screen.GetContent(0, 1)
			Expect(prune).To(Equal(customStyle[0]))

			spinner.Pulse()
			app.Draw()
			prune, _, _, _ = screen.GetContent(0, 1)
			Expect(prune).To(Equal(customStyle[1]))

			spinner.Pulse()
			app.Draw()
			prune, _, _, _ = screen.GetContent(0, 1)
			Expect(prune).To(Equal(customStyle[0]))
		})
	})
})
