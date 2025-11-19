// Demo code for the Form primitive.
package main

import (
	"github.com/malivvan/cui"
)

func main() {
	app := cui.New()
	defer app.HandlePanic()

	form := cui.NewForm().
		AddDropDownSimple("Title", 0, nil, "Mr.", "Ms.", "Mrs.", "Dr.", "Prof.").
		AddInputField("First name", "", 20, nil, nil).
		AddInputField("Last name", "", 20, nil, nil).
		AddFormItem(cui.NewInputField().
			SetLabel("Address").
			SetFieldWidth(30).
			SetFieldNote("Your complete address")).
		AddPasswordField("Password", "", 10, '*', nil).
		AddCheckBox("", "Age 18+", false, nil).
		AddButton("Save", nil).
		AddButton("Quit", func() {
			app.Stop()
		}).
		SetBorder(true).
		SetTitle("Enter some data").
		SetTitleAlign(cui.AlignLeft)

	if err := app.EnableMouse(true).SetRoot(form, true).Run(); err != nil {
		panic(err)
	}
}
