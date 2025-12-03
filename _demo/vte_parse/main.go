package main

import (
	"fmt"
	"os"

	"github.com/malivvan/cui/terminal"
)

func main() {
	fmt.Println("----- tcell-term parser example -----")
	fmt.Println("reading from stdin")
	parser := terminal.NewParser(os.Stdin)
	for {
		seq := parser.Next()
		fmt.Printf("%s\n", seq)
		switch seq.(type) {
		case terminal.EOF:
			return
		}

	}
}
