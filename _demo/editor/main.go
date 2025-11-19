// Demo code for the Box primitive.
package main

import (
	_ "embed"
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/editor"
)

//go:embed main.go
var mainGo string

func main() {
	app := cui.New()
	defer app.HandlePanic()

	buf := editor.NewBufferFromString(mainGo, "main.go")
	view := editor.NewView(buf)
	view.SetTheme("darcula")
	view.SetBorder(true)
	view.SetBorderAttributes(tcell.AttrBold)

	app.SetRoot(view, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
