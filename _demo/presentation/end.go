package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content cui.Widget) {
	textView := cui.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/malivvan/cui"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
