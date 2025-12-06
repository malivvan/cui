package fuzz

import "github.com/malivvan/cui/markup"

// Fuzz is the entrypoint used by the go-fuzz framework
func Fuzz(data []byte) int {
	sel, err := cml.CompileSelector(string(data))
	if err != nil {
		if sel != nil {
			panic("sel != nil on error")
		}
		return 0
	}
	return 1
}
