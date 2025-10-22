// Demo code for the Box primitive.
package main

import (
	"bytes"
	_ "embed"
	"image/png"

	"github.com/malivvan/cui"
)

//go:embed image.png
var imagePNG []byte

func main() {
	app := cui.NewApplication()
	defer app.HandlePanic()

	dec, err := png.Decode(bytes.NewReader(imagePNG))
	if err != nil {
		panic(err)
	}
	img := cui.NewImage()
	img.SetImage(dec)
	app.SetRoot(img, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
