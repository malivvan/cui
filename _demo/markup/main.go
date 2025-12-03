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
	idx, err := markup.NewFile("index.ctml", `<!DOCTYPE ctml>
<ctml lang="en">
<head>
	<title>Markup</title>

	<meta charset="UTF-8">
    <meta name="color" content="#ffffff">
	<meta name="background-color" content="#000000">

	<link rel="manifest" href="manifest.yml">
    <link rel="stylesheet" href="style.css">


	<theme id="geany">
		comment: red
		constant: default
		constant.string: [bold, yellow]
		identifier: default
		preproc: cyan
		line-number: yellow
		current-line-number: red
		diff-added: green
		diff-modified: yellow
	</theme>
    <theme src="themes/bash.yml" id="bash" />

	<syntax id="objective-c">
		filetype: objective-c
		detect:
			filename: "\\.(m|mm|h)$"
			signature: "(obj|objective)-c|#import|@(encode|end|interface|implementation|selector|protocol|synchronized|try|catch|finally|property|optional|required|import|autoreleasepool)"
	</syntax>
    <syntax src="syntaxes/bash.yml" id="bash" />

	<style>
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
		hello_world() {
		   echo 'hello, world'
		}
	</script>
    <script src="main.js" type='text/javascript' />


</head>
<body>
	<flex direction="row">
		<button id="b11" size="1" style="color: red;" >xxxx</button>
		<button id="b12" class="xxx" size="3">xxxxzzzzzz</button>
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
</body>
</ctml>
`)
	if err != nil {
		panic(err)
	}
	idx.Button("b11").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b11 clicked\n") })
	idx.Button("b12").OnClick(func() { _, _ = fmt.Fprintf(log, "Button b12 clicked\n") })
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
