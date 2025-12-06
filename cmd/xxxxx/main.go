package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/malivvan/cui/markup"
)

var pricingHtml string = `
<html>
<head>
	<title>Layout Test</title>
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
	const abc = () => {
	aaa("xxx")
		console.log("===================================================================")
	};
	
	</script>
</head>
<body>
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
	<button id="b31" grow="1" style="color: red;" >yyyyy</button>
</flex>
<script>
    console.log("Layout test initialized.");
</script>
</body>
</html>`

func main() {
	doc, err := markup.Parse(strings.NewReader(pricingHtml))
	if err != nil {
		log.Fatal(err)
	}

	for rootNode := range doc.ChildNodes() {
		fmt.Printf("Root Node: %s\n", rootNode.Tag)
		for child := range rootNode.ChildNodes() {
			fmt.Printf(" Child Node: %s\n", child.Tag)
		}
	}

}
