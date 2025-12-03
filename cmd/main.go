package main

import (
	_ "embed"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/markup"
)

func main() {
	app := cui.New()
	defer app.HandlePanic()

	index, err := markup.OpenFile("_demo/markup/index.cml")
	if err != nil {
		panic(err)
	}

	idx := 0

	app.SetRoot(index.Root(0).Widget(), true)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyLeft {
			if idx > 0 {
				idx--
			} else {
				idx = index.RootCount() - 1
			}
			root := index.Root(idx)
			app.GetScreen().SetTitle(root.ID())
			app.SetRoot(root.Widget(), true)
		} else if event.Key() == tcell.KeyRight {
			if idx < index.RootCount()-1 {
				idx++
			} else {
				idx = 0
			}
			root := index.Root(idx)
			app.GetScreen().SetTitle(root.ID())
			app.SetRoot(root.Widget(), true)
		} else if event.Key() == tcell.KeyCtrlQ {
			app.Stop()
		}
		return event
	})
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
