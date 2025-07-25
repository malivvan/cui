package main

import (
	"github.com/malivvan/cui"
)

const form = `[green]package[white] main

[green]import[white] (
    [red]"github.com/malivvan/cui"[white]
)

[green]func[white] [yellow]main[white]() {
    form := cui.[yellow]NewForm[white]().
        [yellow]AddInputField[white]([red]"First name:"[white], [red]""[white], [red]20[white], nil, nil).
        [yellow]AddInputField[white]([red]"Last name:"[white], [red]""[white], [red]20[white], nil, nil).
        [yellow]AddDropDown[white]([red]"Role:"[white], [][green]string[white]{
            [red]"Engineer"[white],
            [red]"Manager"[white],
            [red]"Administration"[white],
        }, [red]0[white], nil).
        [yellow]AddCheckBox[white]([red]"On vacation:"[white], false, nil).
        [yellow]AddPasswordField[white]([red]"Password:"[white], [red]""[white], [red]10[white], [red]'*'[white], nil).
        [yellow]AddButton[white]([red]"Save"[white], [yellow]func[white]() { [blue]/* Save data */[white] }).
        [yellow]AddButton[white]([red]"Cancel"[white], [yellow]func[white]() { [blue]/* Cancel */[white] })
    cui.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](form, true).
        [yellow]Run[white]()
}`

// Form demonstrates forms.
func Form(nextSlide func()) (title string, info string, content cui.Primitive) {
	f := cui.NewForm()
	f.AddInputField("First name:", "", 20, nil, nil)
	f.AddInputField("Last name:", "", 20, nil, nil)
	f.AddDropDownSimple("Role:", 0, nil, "Engineer", "Manager", "Administration")
	f.AddPasswordField("Password:", "", 10, '*', nil)
	f.AddCheckBox("", "On vacation", false, nil)
	f.AddButton("Save", nextSlide)
	f.AddButton("Cancel", nextSlide)
	f.SetBorder(true)
	f.SetTitle("Employee Information")
	return "Form", formInfo, Code(f, 36, 15, form)
}
