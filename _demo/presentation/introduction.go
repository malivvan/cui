package main

import "github.com/malivvan/cui"

// Introduction returns a cui.List with the highlights of the cview package.
func Introduction(nextSlide func()) (title string, info string, content cui.Primitive) {
	list := cui.NewList()

	listText := [][]string{
		{"A Go package for terminal based UIs", "with a special focus on rich interactive widgets"},
		{"Based on github.com/gdamore/tcell", "Supports Linux, FreeBSD, MacOS and Windows"},
		{"Designed to be simple", `"Hello world" is less than 20 lines of code`},
		{"Good for data entry", `For charts, use "termui" - for low-level views, use "gocui" - ...`},
		{"Supports context menus", "Right click on one of these items or press Alt+Enter"},
		{"Extensive documentation", "Demo code is available for each widget"},
	}

	reset := func() {
		list.Clear()

		for i, itemText := range listText {
			item := cui.NewListItem(itemText[0])
			item.SetSecondaryText(itemText[1])
			item.SetShortcut(rune('1' + i))
			item.SetSelectedFunc(nextSlide)
			list.AddItem(item)
		}

		list.ContextMenuList().SetItemEnabled(3, false)
	}

	list.AddContextItem("Delete item", 'i', func(index int) {
		list.RemoveItem(index)

		if list.GetItemCount() == 0 {
			list.ContextMenuList().SetItemEnabled(0, false)
			list.ContextMenuList().SetItemEnabled(1, false)
		}
		list.ContextMenuList().SetItemEnabled(3, true)
	})

	list.AddContextItem("Delete all", 'a', func(index int) {
		list.Clear()

		list.ContextMenuList().SetItemEnabled(0, false)
		list.ContextMenuList().SetItemEnabled(1, false)
		list.ContextMenuList().SetItemEnabled(3, true)
	})

	list.AddContextItem("", 0, nil)

	list.AddContextItem("Reset", 'r', func(index int) {
		reset()

		list.ContextMenuList().SetItemEnabled(0, true)
		list.ContextMenuList().SetItemEnabled(1, true)
		list.ContextMenuList().SetItemEnabled(3, false)
	})

	reset()
	return "Introduction", listInfo, Center(80, 12, list)
}
