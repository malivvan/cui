package chart_test

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/chart"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Dialog", Ordered, func() {
	var (
		app       *cui.Application
		headerBox *cui.Box
		msgDialog *chart.MessageDialog
		screen    tcell.SimulationScreen
	)

	BeforeAll(func() {
		app = cui.NewApplication()
		headerBox = cui.NewBox()
		headerBox.SetBorder(true)
		msgDialog = chart.NewMessageDialog(chart.InfoDialog)
		screen = tcell.NewSimulationScreen("UTF-8")

		if err := screen.Init(); err != nil {
			panic(err)
		}

		go func() {
			appLayout := cui.NewFlex()
			appLayout.SetDirection(cui.FlexRow)
			appLayout.AddItem(headerBox, 0, 1, true)
			appLayout.AddItem(msgDialog, 0, 1, true)
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

	Describe("NewMessageDialog", func() {
		It("returns a new message dialog primitive", func() {
			tests := []struct {
				msgType int
				bgColor tcell.Color
			}{
				{msgType: chart.InfoDialog, bgColor: tcell.ColorSteelBlue},
				{msgType: chart.ErrorDailog, bgColor: tcell.ColorOrangeRed},
			}

			for _, test := range tests {
				msgDialog.SetType(test.msgType)
				app.Draw()
				Expect(msgDialog.GetBackgroundColor()).To(Equal(test.bgColor))
			}
		})
	})

	Describe("Focus", func() {
		It("checks primitivie focus", func() {
			app.SetFocus(headerBox)
			app.Draw()
			Expect(msgDialog.HasFocus()).To(Equal(false))

			app.SetFocus(msgDialog)
			app.Draw()
			Expect(msgDialog.HasFocus()).To(Equal(true))
		})
	})
})
