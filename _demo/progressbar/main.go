// Demo code for the ProgressBar primitive.
package main

import (
	"time"

	"github.com/malivvan/cui"
)

func main() {
	app := cui.New()
	defer app.HandlePanic()

	verticalProgressBar := cui.NewProgressBar().
		SetBorder(true).
		SetVertical(true)

	horizontalProgressBar := cui.NewProgressBar().
		SetBorder(true).
		SetMax(150)

	padding := cui.NewTextView()

	grid := cui.NewGrid().
		SetColumns(-1, 6, 4, 30, -1).
		SetRows(-1, 12, 4, 4, -1).
		SetBackgroundColor(cui.Styles.PrimitiveBackgroundColor).
		AddItem(padding, 0, 0, 1, 5, 0, 0, false).
		AddItem(padding, 1, 0, 1, 1, 0, 0, false).
		AddItem(verticalProgressBar, 1, 1, 2, 1, 0, 0, false).
		AddItem(padding, 1, 2, 1, 1, 0, 0, false).
		AddItem(padding, 2, 0, 1, 5, 0, 0, false).
		AddItem(horizontalProgressBar, 3, 3, 1, 1, 0, 0, false).
		AddItem(padding, 1, 4, 1, 1, 0, 0, false).
		AddItem(padding, 4, 0, 1, 5, 0, 0, false)

	go func() {
		t := time.NewTicker(100 * time.Millisecond)
		for range t.C {
			if verticalProgressBar.Complete() {
				verticalProgressBar.SetProgress(0)
			} else {
				verticalProgressBar.AddProgress(1)
			}

			if horizontalProgressBar.Complete() {
				horizontalProgressBar.SetProgress(0)
			} else {
				horizontalProgressBar.AddProgress(1)
			}

			// Queue draw
			app.QueueUpdateDraw(func() {})
		}
	}()

	if err := app.SetRoot(grid, true).Run(); err != nil {
		panic(err)
	}
}
