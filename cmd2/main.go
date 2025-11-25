package main

import (
	"fmt"
	"strings"

	"github.com/agreffard/douceur/parser"
)

func main() {
	input := `body {
		x {
		aaa: bbb; 
		}
    background-color: black;
}

`

	stylesheet, err := parser.Parse(input)
	if err != nil {
		panic("Please fill a bug :)")
	}

	for _, rule := range stylesheet.Rules {
		fmt.Printf("%s\n", strings.Join(rule.Selectors, ", "))
		for _, rule2 := range rule.Rules {
			fmt.Printf("%s\n", strings.Join(rule2.Selectors, ", "))
		}
	}
}
