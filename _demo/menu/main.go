package main

import (
	"net/textproto"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
)

func clickedMessageFn(msg string) func(*cui.MenuItem) {
	return func(*cui.MenuItem) { textView.Write([]byte("Clicked: " + textproto.TrimString(msg) + "\n")) }
}

var textView = cui.NewTextView()

func newMenu() (*cui.Flex, func(tcell.Screen)) {
	fileMenu := cui.NewMenuItem("File")

	fileMenu.AddItem(cui.NewMenuItem("Open File").SetOnClick(clickedMessageFn("Open File")))

	fileMenu.AddItem(cui.NewMenuItem("New File").SetOnClick(clickedMessageFn("New File")))

	saveSubForReal := cui.NewMenuItem("Save For Real").
		AddItem(cui.NewMenuItem("For really real").SetOnClick(clickedMessageFn("For really real"))).
		AddItem(cui.NewMenuItem("For really fake").SetOnClick(clickedMessageFn("For really fake")))
	saveSubForFake := cui.NewMenuItem("Save For Fake").SetOnClick(clickedMessageFn("Safe for fake"))

	fileMenu.AddItem(cui.NewMenuItem("Save File").
		// Add submenu items to save
		AddItem(saveSubForReal).
		AddItem(saveSubForFake).SetOnClick(clickedMessageFn("Save File")))

	fileMenu.AddItem(cui.NewMenuItem("Close File").SetOnClick(clickedMessageFn("Close File")))
	fileMenu.AddItem(cui.NewMenuItem("Exit").SetOnClick(func(*cui.MenuItem) {}))
	editMenu := cui.NewMenuItem("Edit")
	editMenu.AddItem(cui.NewMenuItem("Copy").SetOnClick(clickedMessageFn("Copy")))
	editMenu.AddItem(cui.NewMenuItem("Cut").SetOnClick(clickedMessageFn("Cut")))
	editMenu.AddItem(cui.NewMenuItem("Paste").SetOnClick(clickedMessageFn("Paste")))

	menuBar := cui.NewMenuBar().
		AddItem(fileMenu).
		AddItem(editMenu)

	menuBar.SetRect(0, 0, 100, 15)

	return cui.NewFlex().
		SetDirection(cui.FlexRow).
		AddItem(menuBar, 1, 1, false).
		AddItem(textView, 0, 4, true), menuBar.AfterDraw()
}

func main() {
	app := cui.New()

	left, leftAfterDraw := newMenu()
	right, rightAfterDraw := newMenu()
	flex := cui.NewFlex().
		SetDirection(cui.FlexRow).
		AddItem(left, 0, 1, true).
		AddItem(right, 0, 1, true)

	app.EnableMouse(true).SetRoot(flex, true).SetFocus(flex).SetAfterDrawFunc(func(screen tcell.Screen) {
		leftAfterDraw(screen)
		rightAfterDraw(screen)
	})

	if err := app.Run(); err != nil {
		panic(err)
	}
}
