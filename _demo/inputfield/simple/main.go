// Demo code for the InputField primitive.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
)

type formItemLabelWidthSetter[T cui.Widget] interface{ SetLabelWidth(int) T }

// setFormItemLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func setFormItemLabelWidth[T cui.Widget](widget T, width int) bool {
	if setter, ok := cui.Widget(widget).(formItemLabelWidthSetter[T]); ok {
		setter.SetLabelWidth(width)
		return true
	}
	return false
}

func main() {
	app := cui.New()
	defer app.HandlePanic()

	app.EnableMouse(true)

	inputField := cui.NewInputField()
	inputField.SetLabel("Enter a number: ")
	inputField.SetPlaceholder("E.g. 1234")
	inputField.SetFieldWidth(10)
	inputField.SetAcceptanceFunc(cui.InputFieldInteger)
	inputField.SetDoneFunc(func(key tcell.Key) {
		app.Stop()
	})
	setFormItemLabelWidth(inputField, 30)

	app.SetRoot(inputField, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
