// Demo code for the Box primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
)

func main() {
	app := cui.New()
	defer app.HandlePanic()

	box := cui.NewBox().
		SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetTitle("A [red]c[yellow]o[green]l[darkcyan]o[blue]r[darkmagenta]f[red]u[yellow]l[white] [black:red]c[:yellow]o[:green]l[:darkcyan]o[:blue]r[:darkmagenta]f[:red]u[:yellow]l[white:] [::bu]title")

	if err := app.SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
