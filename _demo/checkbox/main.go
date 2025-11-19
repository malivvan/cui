// Demo code for the CheckBox primitive.
package main

import (
	"github.com/malivvan/cui"
)

func main() {
	app := cui.New()
	defer app.HandlePanic()

	app.EnableMouse(true)

	checkbox := cui.NewCheckBox()
	checkbox.SetLabel("Hit Enter to check box: ")

	app.SetRoot(checkbox, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
