// Demo code for the Button primitive.
package main

import "github.com/malivvan/cui"

func main() {
	app := cui.New()
	defer app.HandlePanic()

	button := cui.NewButton().
		SetLabel("Hit Enter to close").
		OnClick(func() {
			app.Stop()
		})
	button.SetRect(0, 0, 22, 3)

	if err := app.EnableMouse(true).SetRoot(button, false).Run(); err != nil {
		panic(err)
	}
}
