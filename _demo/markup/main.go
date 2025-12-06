package main

import (
	_ "embed"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/markup"
)

func main() {
	app := cui.New()
	defer app.HandlePanic()

	log := cui.NewTextView().SetBackgroundColor(tcell.ColorDimGrey)
	idx, err := markup.NewFile("index.ctml", `<style>
	button {color: black;}
	#b11 {color: orange;}
	#b12 {color: green;}
	#b21 {color: orange;}
	#b22 {color: green;}
	#b33 {color: orange;}
	#b31 {color: green;}
	#b32 {color: orange;}
	.xxx {color: blue;}
	.blob {color: purple;}
</style>
<script>
const abc = () => {
aaa("xxx")
	console.log("===================================================================")
};

</script>
<flex direction="row">
	<button id="b11" size="1" style="color: red;" onclick="abc()">xxxx</button>
	<button id="b12" class="xxx" size="3" onclick="console.log('xxxxx')">xxxxzzzzzz</button>

	<flex direction="column" grow="1">
		<button id="b21"  class="xxx"  size="10">aaaaa</button>
		<button id="b22" class="blob" grow="1" >MID</button>
		<button id="b23"  class="xxx" size="10" style="color: blue;">bbbbbb</button>
	</flex>
</flex>
<flex direction="row">
	<button id="b31" grow="1" style="color: red;" >xxxx</button>
	<button id="b32" grow="1" style="color: red;" >yyyyy</button>
</flex>
<script>
    console.log("Layout test initialized.");
</script>
`)
	if err != nil {
		panic(err)
	}
	//idx.Button("b11").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b11 clicked\n") })
	idx.Button("b21").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b21 clicked\n") })
	idx.Button("b22").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b22 clicked\n") })
	idx.Button("b23").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b23 clicked\n") })
	idx.Button("b31").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b31 clicked\n") })
	idx.Button("b32").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b32 clicked\n") })

	//
	cur := idx.RootCount()

	setRoot := func(i int) {
		root := idx.Root(i)
		if app.GetScreen() != nil {
			app.GetScreen().SetTitle(root.ID())
		}
		app.SetRoot(cui.NewFlex().SetDirection(cui.FlexRow).
			AddItem(log, 5, 0, false).
			AddItem(root.Widget(), 0, 3, true), true)
	}
	setRoot(0)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyLeft {
			if cur > 0 {
				cur--
			} else {
				cur = idx.RootCount() - 1
			}
			setRoot(cur)
		} else if event.Key() == tcell.KeyRight {
			if cur < idx.RootCount()-1 {
				cur++
			} else {
				cur = 0
			}
			setRoot(cur)
		} else if event.Key() == tcell.KeyCtrlQ {
			app.Stop()
		}
		return event
	})

	if err := app.EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
