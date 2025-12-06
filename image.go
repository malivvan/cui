package cui

import (
	"image"
	"math"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Image implements a widget that displays one image. The original image
// (specified with [Image.SetImage]) is resized according to the specified size
// (see [Image.SetSize]), using the specified number of colors (see
// [Image.SetColors]), while applying dithering if necessary (see
// [Image.SetDithering]).
//
// Images are approximated by graphical characters in the terminal. The
// resolution is therefore limited by the number and type of characters that can
// be drawn in the terminal and the colors available in the terminal. The
// quality of the final image also depends on the terminal's font and spacing
// settings, none of which are under the control of this package. Results may
// vary.
type Image struct {
	box *Box

	// The image to be displayed. If nil, the widget will be empty.
	image image.Image

	// The size of the image. If a value is 0, the corresponding size is chosen
	// automatically based on the other size while preserving the image's aspect
	// ratio. If both are 0, the image uses as much space as possible. A
	// negative value represents a percentage, e.g. -50 means 50% of the
	// available space.
	width, height int

	// The number of colors to use. If 0, the number of colors is chosen based
	// on the terminal's capabilities.
	colors int

	// The width of a terminal's cell divided by its height.
	aspectRatio float64

	// Horizontal and vertical alignment, one of the "Align" constants.
	alignHorizontal, alignVertical int

	// The actual image size (in cells) when it was drawn the last time.
	lastWidth, lastHeight int

	// The actual image (in cells) when it was drawn the last time. The size of
	// this slice is lastWidth * lastHeight, indexed by y*lastWidth + x.
	pixels []pixel

	mu sync.RWMutex
}

// pixel represents a character on screen used to draw part of an image.
type pixel struct {
	style   tcell.Style
	element rune // The block element.
}

// NewImage returns a new [Image] widget with an empty image (use
// [Image.SetImage] to specify the image to be displayed). The image will use
// the widget's entire available space. The default dithering algorithm is set
// to Floyd-Steinberg dithering. The terminal's cell aspect ratio defaults to
// 0.5.
func NewImage() *Image {
	return &Image{
		box:             NewBox(),
		aspectRatio:     0.5,
		alignHorizontal: AlignCenter,
		alignVertical:   AlignCenter,
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (img *Image) set(setter func(img *Image)) *Image {
	img.mu.Lock()
	setter(img)
	img.mu.Unlock()
	return img
}

func (img *Image) get(getter func(img *Image)) {
	img.mu.RLock()
	getter(img)
	img.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Image.
func (img *Image) GetTitle() string {
	return img.box.GetTitle()
}

// SetTitle sets the title of this Image.
func (img *Image) SetTitle(title string) *Image {
	img.box.SetTitle(title)
	return img
}

// GetTitleAlign returns the title alignment of this Image.
func (img *Image) GetTitleAlign() int {
	return img.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Image.
func (img *Image) SetTitleAlign(align int) *Image {
	img.box.SetTitleAlign(align)
	return img
}

// GetBorder returns whether this Image has a border.
func (img *Image) GetBorder() bool {
	return img.box.GetBorder()
}

// SetBorder sets whether this Image has a border.
func (img *Image) SetBorder(show bool) *Image {
	img.box.SetBorder(show)
	return img
}

// GetBorderColor returns the border color of this Image.
func (img *Image) GetBorderColor() tcell.Color {
	return img.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Image.
func (img *Image) SetBorderColor(color tcell.Color) *Image {
	img.box.SetBorderColor(color)
	return img
}

// GetBorderAttributes returns the border attributes of this Image.
func (img *Image) GetBorderAttributes() tcell.AttrMask {
	return img.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Image.
func (img *Image) SetBorderAttributes(attr tcell.AttrMask) *Image {
	img.box.SetBorderAttributes(attr)
	return img
}

// GetBorderColorFocused returns the border color of this Image when focuseimg.
func (img *Image) GetBorderColorFocused() tcell.Color {
	return img.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Image when focuseimg.
func (img *Image) SetBorderColorFocused(color tcell.Color) *Image {
	img.box.SetBorderColorFocused(color)
	return img
}

// GetTitleColor returns the title color of this Image.
func (img *Image) GetTitleColor() tcell.Color {
	return img.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Image.
func (img *Image) SetTitleColor(color tcell.Color) *Image {
	img.box.SetTitleColor(color)
	return img
}

// GetDrawFunc returns the custom draw function of this Image.
func (img *Image) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return img.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Image.
func (img *Image) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Image {
	img.box.SetDrawFunc(handler)
	return img
}

// ShowFocus sets whether this Image should show a focus indicator when focuseimg.
func (img *Image) ShowFocus(showFocus bool) *Image {
	img.box.ShowFocus(showFocus)
	return img
}

// GetMouseCapture returns the mouse capture function of this Image.
func (img *Image) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return img.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Image.
func (img *Image) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Image {
	img.box.SetMouseCapture(capture)
	return img
}

// GetBackgroundColor returns the background color of this Image.
func (img *Image) GetBackgroundColor() tcell.Color {
	return img.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Image.
func (img *Image) SetBackgroundColor(color tcell.Color) *Image {
	img.box.SetBackgroundColor(color)
	return img
}

// GetBackgroundTransparent returns whether the background of this Image is transparent.
func (img *Image) GetBackgroundTransparent() bool {
	return img.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Image is transparent.
func (img *Image) SetBackgroundTransparent(transparent bool) *Image {
	img.box.SetBackgroundTransparent(transparent)
	return img
}

// GetInputCapture returns the input capture function of this Image.
func (img *Image) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return img.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Image.
func (img *Image) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Image {
	img.box.SetInputCapture(capture)
	return img
}

// GetPadding returns the padding of this Image.
func (img *Image) GetPadding() (top, bottom, left, right int) {
	return img.box.GetPadding()
}

// SetPadding sets the padding of this Image.
func (img *Image) SetPadding(top, bottom, left, right int) *Image {
	img.box.SetPadding(top, bottom, left, right)
	return img
}

// InRect returns whether the given screen coordinates are within this Image.
func (img *Image) InRect(x, y int) bool {
	return img.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Image.
func (img *Image) GetInnerRect() (x, y, width, height int) {
	return img.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Image is preserveimg.
func (img *Image) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return img.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Image is preserveimg.
func (img *Image) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return img.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Image.
func (img *Image) GetRect() (x, y, width, height int) {
	return img.box.GetRect()
}

// SetRect sets the rectangle occupied by this Image.
func (img *Image) SetRect(x, y, width, height int) {
	img.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Image is visible.
func (img *Image) GetVisible() bool {
	return img.box.GetVisible()
}

// SetVisible sets whether this Image is visible.
func (img *Image) SetVisible(visible bool) {
	img.box.SetVisible(visible)
}

// Focus is called when this Image receives focus.
func (img *Image) Focus(delegate func(p Widget)) {
	img.box.Focus(delegate)
}

// HasFocus returns whether this Image has focus.
func (img *Image) HasFocus() bool {
	return img.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Image.
func (img *Image) GetFocusable() Focusable {
	return img.box.GetFocusable()
}

// Blur is called when this Image loses focus.
func (img *Image) Blur() {
	img.box.Blur()
}

// SetImage sets the image to be displayed. If nil, the widget will be empty.
func (img *Image) SetImage(image image.Image) *Image {
	img.mu.Lock()
	defer img.mu.Unlock()

	img.image = image
	img.lastWidth, img.lastHeight = 0, 0
	return img
}

// SetSize sets the size of the image. Positive values refer to cells in the
// terminal. Negative values refer to a percentage of the available space (e.g.
// -50 means 50%). A value of 0 means that the corresponding size is chosen
// automatically based on the other size while preserving the image's aspect
// ratio. If both are 0, the image uses as much space as possible while still
// preserving the aspect ratio.
func (img *Image) SetSize(rows, columns int) *Image {
	return img.set(func(img *Image) {
		img.width = columns
		img.height = rows
	})

}

// SetColors sets the number of colors to use. This should be the number of
// colors supported by the terminal. If 0, the number of colors is chosen based
// on the TERM environment variable (which may or may not be reliable).
//
// Only the values 0, 2, 8, 256, and 16777216 ([TrueColor]) are supported. Other
// values will be rounded up to the next supported value, to a maximum of
// 16777216.
//
// The effect of using more colors than supported by the terminal is undefined.
func (img *Image) SetColors(colors int) *Image {
	return img.set(func(img *Image) {
		img.colors = colors
		img.lastWidth, img.lastHeight = 0, 0
	})
}

// GetColors returns the number of colors that will be used while drawing the
// image. This is one of the values listed in [Image.SetColors], except 0 which
// will be replaced by the actual number of colors used.
func (img *Image) GetColors() (colors int) {
	img.get(func(img *Image) {
		switch {
		case img.colors == 0:
			colors = 256
		case img.colors <= 2:
			colors = 2
		case img.colors <= 8:
			colors = 8
		case img.colors <= 256:
			colors = 256
		default:
			colors = 16777216
		}
	})
	return
}

// SetAspectRatio sets the width of a terminal's cell divided by its height.
// You may change the default of 0.5 if your terminal / font has a different
// aspect ratio. This is used to calculate the size of the image if the
// specified width or height is 0. The function will panic if the aspect ratio
// is 0 or less.
func (img *Image) SetAspectRatio(aspectRatio float64) *Image {
	if aspectRatio <= 0 {
		panic("aspect ratio must be greater than 0")
	}
	return img.set(func(img *Image) {
		img.aspectRatio = aspectRatio
		img.lastWidth, img.lastHeight = 0, 0
	})
}

// SetAlign sets the vertical and horizontal alignment of the image within the
// widget's space. The possible values are [AlignTop], [AlignCenter], and
// [AlignBottom] for vertical alignment and [AlignLeft], [AlignCenter], and
// [AlignRight] for horizontal alignment. The default is [AlignCenter] for both
// (or [AlignTop] and [AlignLeft] if the image is part of a [Form]).
func (img *Image) SetAlign(vertical, horizontal int) *Image {
	return img.set(func(img *Image) {
		img.alignHorizontal = horizontal
		img.alignVertical = vertical
	})
}

// InputHandler returns the input handler for this Image.
func (img *Image) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return img.box.InputHandler()
}

// MouseHandler returns the mouse handler for this Image.
func (img *Image) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return img.box.MouseHandler()
}

// Draw draws this primitive onto the screen.
func (img *Image) Draw(screen tcell.Screen) {
	if !img.GetVisible() {
		return
	}
	img.mu.RLock()
	defer img.mu.RUnlock()

	img.box.Draw(screen)

	// Regenerate image if necessary.
	img.render()

	// Draw label.
	viewX, viewY, viewWidth, viewHeight := img.GetInnerRect()

	// Determine image placement.
	x, y, width, height := viewX, viewY, img.lastWidth, img.lastHeight
	if img.alignHorizontal == AlignCenter {
		x += (viewWidth - width) / 2
	} else if img.alignHorizontal == AlignRight {
		x += viewWidth - width
	}
	if img.alignVertical == AlignCenter {
		y += (viewHeight - height) / 2
	} else if img.alignVertical == AlignBottom {
		y += viewHeight - height
	}

	// Draw the image.
	for row := 0; row < height; row++ {
		if y+row < viewY || y+row >= viewY+viewHeight {
			continue
		}
		for col := 0; col < width; col++ {
			if x+col < viewX || x+col >= viewX+viewWidth {
				continue
			}

			index := row*width + col
			screen.SetContent(x+col, y+row, img.pixels[index].element, nil, img.pixels[index].style)
		}
	}
}

// render re-populates the [Image.pixels] slice based on the current settings,
// if [Image.lastWidth] and [Image.lastHeight] don't match the current image's
// size. It also sets the new image size in these two variables.
func (img *Image) render() {
	// If there is no image, there are no pixels.
	if img.image == nil {
		img.pixels = nil
		return
	}

	// Calculate the new (terminal-space) image size.
	bounds := img.image.Bounds()
	imageWidth, imageHeight := bounds.Dx(), bounds.Dy()
	if img.aspectRatio != 1.0 {
		imageWidth = int(float64(imageWidth) / img.aspectRatio)
	}
	width, height := img.width, img.height
	_, _, innerWidth, innerHeight := img.GetInnerRect()
	if innerWidth <= 0 {
		img.pixels = nil
		return
	}

	if width == 0 && height == 0 {
		width, height = innerWidth, innerHeight
		if adjustedWidth := imageWidth * height / imageHeight; adjustedWidth < width {
			width = adjustedWidth
		} else {
			height = imageHeight * width / imageWidth
		}
	} else {
		if width < 0 {
			width = innerWidth * -width / 100
		}
		if height < 0 {
			height = innerHeight * -height / 100
		}
		if width == 0 {
			width = imageWidth * height / imageHeight
		} else if height == 0 {
			height = imageHeight * width / imageWidth
		}
	}

	if width <= 0 || height <= 0 {
		img.pixels = nil
		return
	}

	// If nothing has changed, we're done.
	if img.lastWidth == width && img.lastHeight == height {
		return
	}
	img.lastWidth, img.lastHeight = width, height // This could still be larger than the available space but that's ok for now.

	// Generate the initial pixels by resizing the image (8x8 per cell).
	pixels := img.resize()

	// Turn them into block elements with background/foreground colors.
	img.stamp(pixels)
}

// resize resizes the image to the current size and returns the result as a
// slice of pixels. It is assumed that [Image.lastWidth] (w) and
// [Image.lastHeight] (h) are positive, non-zero values, and the slice has a
// size of 64*w*h, with each pixel being represented by 3 float64 values in the
// range of 0-1. The factor of 64 is due to the fact that we calculate 8x8
// pixels per cell.
func (img *Image) resize() [][3]float64 {
	// Because most of the time, we will be downsizing the image, we don't even
	// attempt to do any fancy interpolation. For each target pixel, we
	// calculate a weighted average of the source pixels using their coverage
	// area.

	bounds := img.image.Bounds()
	srcWidth, srcHeight := bounds.Dx(), bounds.Dy()
	tgtWidth, tgtHeight := img.lastWidth*8, img.lastHeight*8
	coverageWidth, coverageHeight := float64(tgtWidth)/float64(srcWidth), float64(tgtHeight)/float64(srcHeight)
	pixels := make([][3]float64, tgtWidth*tgtHeight)
	weights := make([]float64, tgtWidth*tgtHeight)
	for srcY := bounds.Min.Y; srcY < bounds.Max.Y; srcY++ {
		for srcX := bounds.Min.X; srcX < bounds.Max.X; srcX++ {
			r32, g32, b32, _ := img.image.At(srcX, srcY).RGBA()
			r, g, b := float64(r32)/0xffff, float64(g32)/0xffff, float64(b32)/0xffff

			// Iterate over all target pixels. Outer loop is Y.
			startY := float64(srcY-bounds.Min.Y) * coverageHeight
			endY := startY + coverageHeight
			fromY, toY := int(startY), int(endY)
			for tgtY := fromY; tgtY <= toY && tgtY < tgtHeight; tgtY++ {
				coverageY := 1.0
				if tgtY == fromY {
					coverageY -= math.Mod(startY, 1.0)
				}
				if tgtY == toY {
					coverageY -= 1.0 - math.Mod(endY, 1.0)
				}

				// Inner loop is X.
				startX := float64(srcX-bounds.Min.X) * coverageWidth
				endX := startX + coverageWidth
				fromX, toX := int(startX), int(endX)
				for tgtX := fromX; tgtX <= toX && tgtX < tgtWidth; tgtX++ {
					coverageX := 1.0
					if tgtX == fromX {
						coverageX -= math.Mod(startX, 1.0)
					}
					if tgtX == toX {
						coverageX -= 1.0 - math.Mod(endX, 1.0)
					}

					// Add a weighted contribution to the target pixel.
					index := tgtY*tgtWidth + tgtX
					coverage := coverageX * coverageY
					pixels[index][0] += r * coverage
					pixels[index][1] += g * coverage
					pixels[index][2] += b * coverage
					weights[index] += coverage
				}
			}
		}
	}

	// Normalize the pixels.
	for index, weight := range weights {
		if weight > 0 {
			pixels[index][0] /= weight
			pixels[index][1] /= weight
			pixels[index][2] /= weight
		}
	}

	return pixels
}

// stamp takes the pixels generated by [Image.resize] and populates the
// [Image.pixels] slice accordingly.
func (img *Image) stamp(resized [][3]float64) {
	// For each 8x8 pixel block, we find the best block element to represent it,
	// given the available colors.
	img.pixels = make([]pixel, img.lastWidth*img.lastHeight)
	colors := img.GetColors()
	for row := 0; row < img.lastHeight; row++ {
		for col := 0; col < img.lastWidth; col++ {
			// Calculate an error for each potential block element + color. Keep
			// the one with the lowest error.

			// Note that the values in "resize" may lie outside [0, 1] due to
			// the error distribution during dithering.
			minMSE := math.MaxFloat64 // Mean squared error.

			// This map describes what each block element looks like. A 1 bit represents a
			// pixel that is drawn, a 0 bit represents a pixel that is not drawn. The least
			// significant bit is the top left pixel, the most significant bit is the bottom
			// right pixel, moving row by row from left to right, top to bottom.
			for element, bits := range map[rune]uint64{
				BlockLowerOneEighthBlock:            0b1111111100000000000000000000000000000000000000000000000000000000,
				BlockLowerOneQuarterBlock:           0b1111111111111111000000000000000000000000000000000000000000000000,
				BlockLowerThreeEighthsBlock:         0b1111111111111111111111110000000000000000000000000000000000000000,
				BlockLowerHalfBlock:                 0b1111111111111111111111111111111100000000000000000000000000000000,
				BlockLowerFiveEighthsBlock:          0b1111111111111111111111111111111111111111000000000000000000000000,
				BlockLowerThreeQuartersBlock:        0b1111111111111111111111111111111111111111111111110000000000000000,
				BlockLowerSevenEighthsBlock:         0b1111111111111111111111111111111111111111111111111111111100000000,
				BlockLeftSevenEighthsBlock:          0b0111111101111111011111110111111101111111011111110111111101111111,
				BlockLeftThreeQuartersBlock:         0b0011111100111111001111110011111100111111001111110011111100111111,
				BlockLeftFiveEighthsBlock:           0b0001111100011111000111110001111100011111000111110001111100011111,
				BlockLeftHalfBlock:                  0b0000111100001111000011110000111100001111000011110000111100001111,
				BlockLeftThreeEighthsBlock:          0b0000011100000111000001110000011100000111000001110000011100000111,
				BlockLeftOneQuarterBlock:            0b0000001100000011000000110000001100000011000000110000001100000011,
				BlockLeftOneEighthBlock:             0b0000000100000001000000010000000100000001000000010000000100000001,
				BlockQuadrantLowerLeft:              0b0000111100001111000011110000111100000000000000000000000000000000,
				BlockQuadrantLowerRight:             0b1111000011110000111100001111000000000000000000000000000000000000,
				BlockQuadrantUpperLeft:              0b0000000000000000000000000000000000001111000011110000111100001111,
				BlockQuadrantUpperRight:             0b0000000000000000000000000000000011110000111100001111000011110000,
				BlockQuadrantUpperLeftAndLowerRight: 0b1111000011110000111100001111000000001111000011110000111100001111,
			} {
				// Calculate the average color for the pixels covered by the set
				// bits and unset bits.
				var (
					bg, fg  [3]float64
					setBits float64
					bit     uint64 = 1
				)
				for y := 0; y < 8; y++ {
					for x := 0; x < 8; x++ {
						index := (row*8+y)*img.lastWidth*8 + (col*8 + x)
						if bits&bit != 0 {
							fg[0] += resized[index][0]
							fg[1] += resized[index][1]
							fg[2] += resized[index][2]
							setBits++
						} else {
							bg[0] += resized[index][0]
							bg[1] += resized[index][1]
							bg[2] += resized[index][2]
						}
						bit <<= 1
					}
				}
				for ch := 0; ch < 3; ch++ {
					fg[ch] /= setBits
					if fg[ch] < 0 {
						fg[ch] = 0
					} else if fg[ch] > 1 {
						fg[ch] = 1
					}
					bg[ch] /= 64 - setBits
					if bg[ch] < 0 {
						bg[ch] = 0
					}
					if bg[ch] > 1 {
						bg[ch] = 1
					}
				}

				// Quantize to the nearest acceptable color.
				for _, color := range []*[3]float64{&fg, &bg} {
					if colors == 2 {
						// Monochrome. The following weights correspond better
						// to human perception than the arithmetic mean.
						gray := 0.299*color[0] + 0.587*color[1] + 0.114*color[2]
						if gray < 0.5 {
							*color = [3]float64{0, 0, 0}
						} else {
							*color = [3]float64{1, 1, 1}
						}
					} else {
						for index, ch := range color {
							switch {
							case colors == 8:
								// colors vary wildly for each terminal. Expect
								// suboptimal results.
								if ch < 0.5 {
									color[index] = 0
								} else {
									color[index] = 1
								}
							case colors == 256:
								color[index] = math.Round(ch*6) / 6
							}
						}
					}
				}

				// Calculate the error (and the final pixel values).
				var (
					mse         float64
					values      [64][3]float64
					valuesIndex int
				)
				bit = 1
				for y := 0; y < 8; y++ {
					for x := 0; x < 8; x++ {
						if bits&bit != 0 {
							values[valuesIndex] = fg
						} else {
							values[valuesIndex] = bg
						}
						index := (row*8+y)*img.lastWidth*8 + (col*8 + x)
						for ch := 0; ch < 3; ch++ {
							err := resized[index][ch] - values[valuesIndex][ch]
							mse += err * err
						}
						bit <<= 1
						valuesIndex++
					}
				}

				// Do we have a better match?
				if mse < minMSE {
					// Yes. Save it.
					minMSE = mse
					index := row*img.lastWidth + col
					img.pixels[index].element = element
					img.pixels[index].style = tcell.StyleDefault.
						Foreground(tcell.NewRGBColor(int32(math.Min(255, fg[0]*255)), int32(math.Min(255, fg[1]*255)), int32(math.Min(255, fg[2]*255)))).
						Background(tcell.NewRGBColor(int32(math.Min(255, bg[0]*255)), int32(math.Min(255, bg[1]*255)), int32(math.Min(255, bg[2]*255))))
				}
			}

			// Check if there is a shade block which results in a smaller error.
			// What's the overall average color?
			var avg [3]float64
			for y := 0; y < 8; y++ {
				for x := 0; x < 8; x++ {
					index := (row*8+y)*img.lastWidth*8 + (col*8 + x)
					for ch := 0; ch < 3; ch++ {
						avg[ch] += resized[index][ch] / 64
					}
				}
			}
			for ch := 0; ch < 3; ch++ {
				if avg[ch] < 0 {
					avg[ch] = 0
				} else if avg[ch] > 1 {
					avg[ch] = 1
				}
			}

			// Quantize and choose shade element.
			element := BlockFullBlock
			var fg, bg tcell.Color
			shades := []rune{' ', BlockLightShade, BlockMediumShade, BlockDarkShade, BlockFullBlock}
			if colors == 2 {
				// Monochrome.
				gray := 0.299*avg[0] + 0.587*avg[1] + 0.114*avg[2] // See above for details.
				shade := int(math.Round(gray * 4))
				element = shades[shade]
				for ch := 0; ch < 3; ch++ {
					avg[ch] = float64(shade) / 4
				}
				bg = tcell.ColorBlack
				fg = tcell.ColorWhite
			} else if colors == 16777216 {
				// True color.
				fg = tcell.NewRGBColor(int32(math.Min(255, avg[0]*255)), int32(math.Min(255, avg[1]*255)), int32(math.Min(255, avg[2]*255)))
				bg = fg
			} else {
				// 8 or 256 colors.
				steps := 1.0
				if colors == 256 {
					steps = 6.0
				}
				var (
					lo, hi, pos [3]float64
					shade       float64
				)
				for ch := 0; ch < 3; ch++ {
					lo[ch] = math.Floor(avg[ch]*steps) / steps
					hi[ch] = math.Ceil(avg[ch]*steps) / steps
					if r := hi[ch] - lo[ch]; r > 0 {
						pos[ch] = (avg[ch] - lo[ch]) / r
						if math.Abs(pos[ch]-0.5) < math.Abs(shade-0.5) {
							shade = pos[ch]
						}
					}
				}
				shade = math.Round(shade * 4)
				element = shades[int(shade)]
				shade /= 4
				for ch := 0; ch < 3; ch++ { // Find the closest channel value.
					best := math.Abs(avg[ch] - (lo[ch] + (hi[ch]-lo[ch])*shade)) // Start shade from lo to hi.
					if value := math.Abs(avg[ch] - (hi[ch] - (hi[ch]-lo[ch])*shade)); value < best {
						best = value // Swap lo and hi.
						lo[ch], hi[ch] = hi[ch], lo[ch]
					}
					if value := math.Abs(avg[ch] - lo[ch]); value < best {
						best = value // Use lo.
						hi[ch] = lo[ch]
					}
					if value := math.Abs(avg[ch] - hi[ch]); value < best {
						lo[ch] = hi[ch] // Use hi.
					}
					avg[ch] = lo[ch] + (hi[ch]-lo[ch])*shade // Quantize.
				}
				bg = tcell.NewRGBColor(int32(math.Min(255, lo[0]*255)), int32(math.Min(255, lo[1]*255)), int32(math.Min(255, lo[2]*255)))
				fg = tcell.NewRGBColor(int32(math.Min(255, hi[0]*255)), int32(math.Min(255, hi[1]*255)), int32(math.Min(255, hi[2]*255)))
			}

			// Calculate the error (and the final pixel values).
			var (
				mse         float64
				values      [64][3]float64
				valuesIndex int
			)
			for y := 0; y < 8; y++ {
				for x := 0; x < 8; x++ {
					index := (row*8+y)*img.lastWidth*8 + (col*8 + x)
					for ch := 0; ch < 3; ch++ {
						err := resized[index][ch] - avg[ch]
						mse += err * err
					}
					values[valuesIndex] = avg
					valuesIndex++
				}
			}

			// Is this shade element better than the block element?
			if mse < minMSE {
				// Yes. Save it.
				index := row*img.lastWidth + col
				img.pixels[index].element = element
				img.pixels[index].style = tcell.StyleDefault.Foreground(fg).Background(bg)
			}
		}
	}
}
