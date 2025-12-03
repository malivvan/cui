// Demo code for the Box primitive.
package main

import (
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/terminal/pty"
)

func main() {
	app := cui.New()
	defer app.HandlePanic()

	term := cui.NewTerminal(app, pty.Options{
		Path: "/bin/bash",
	})
	term.SetBorder(true)

	if err := app.SetRoot(term, true).Run(); err != nil {
		panic(err)
	}
}
