package main

import "github.com/malivvan/cui"

const corporate = `Leverage agile frameworks to provide a robust synopsis for high level overviews. Iterative approaches to corporate strategy foster collaborative thinking to further the overall value proposition. Organically grow the holistic world view of disruptive innovation via workplace diversity and empowerment.

Bring to the table win-win survival strategies to ensure proactive domination. At the end of the day, going forward, a new normal that has evolved from generation X is on the runway heading towards a streamlined cloud solution. User generated content in real-time will have multiple touchpoints for offshoring.

Capitalize on low hanging fruit to identify a ballpark value added activity to beta test. Override the digital divide with additional clickthroughs from DevOps. Nanotechnology immersion along the information highway will close the loop on focusing solely on the bottom line.

[yellow]Press Enter, then Tab/Backtab for word selections`

func main() {
	app := cui.NewApplication()

	textView := cui.NewTextView()
	textView.SetDynamicColors(true)
	textView.SetRegions(true)
	textView.SetWordWrap(true)
	textView.SetText(corporate)

	textView2 := cui.NewTextView()
	textView2.SetDynamicColors(true)
	textView2.SetRegions(true)
	textView2.SetWordWrap(true)
	textView2.SetText(corporate)

	textView3 := cui.NewTextView()
	textView3.SetDynamicColors(true)
	textView3.SetRegions(true)
	textView3.SetWordWrap(true)
	textView3.SetText(corporate)

	layout := cui.NewLayout().
		SetDirection(cui.HorizontalLayout).
		SetSplitter(true).
		AddItem(textView, cui.AutoSize).
		AddItem(textView2, cui.AutoSize).
		AddItem(textView3, cui.AutoSize)

	app.SetRoot(layout, true)
	app.EnableMouse(true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
