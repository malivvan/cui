// Demo code for the Box primitive.
package main

import (
	_ "embed"

	"github.com/malivvan/cui"
	"github.com/malivvan/cui/editor"
)

//go:embed main.go
var mainGo string

func main() {
	app := cui.New()
	defer app.HandlePanic()

	view := cui.NewEditor()
	view.SetBuffer(editor.NewBufferFromString(mainGo, "main.go"))
	view.SetTheme("darcula")

	app.SetRoot(view, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
