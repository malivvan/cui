package main

import "github.com/malivvan/cui"

const loremIpsumText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

func main() {
	wm := cui.NewWindowManager()

	list := cui.NewList()
	list.ShowSecondaryText(false)
	list.AddItem(cui.NewListItem("Item #1"))
	list.AddItem(cui.NewListItem("Item #2"))
	list.AddItem(cui.NewListItem("Item #3"))
	list.AddItem(cui.NewListItem("Item #4"))
	list.AddItem(cui.NewListItem("Item #5"))
	list.AddItem(cui.NewListItem("Item #6"))
	list.AddItem(cui.NewListItem("Item #7"))

	loremIpsum := cui.NewTextView()
	loremIpsum.SetText(loremIpsumText)

	w1 := cui.NewWindow(list).
		SetRect(2, 2, 20, 7)

	w2 := cui.NewWindow(loremIpsum)
	w2.SetRect(27, 4, 32, 12)

	w1.SetTitle("List")
	w2.SetTitle("Lorem Ipsum")

	wm.Add(w1, w2)

	app := cui.New()
	app.SetRoot(wm, true)
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
