package cui

import (
	"math"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Progress indicates the progress of an operation.
type Progress struct {
	box *Box

	// Rune to use when rendering the empty area of the progress bar.
	emptyRune rune

	// Color of the empty area of the progress bar.
	emptyColor tcell.Color

	// Rune to use when rendering the filled area of the progress bar.
	filledRune rune

	// Color of the filled area of the progress bar.
	filledColor tcell.Color

	// If set to true, instead of filling from left to right, the bar is filled
	// from bottom to top.
	vertical bool

	// Current progress.
	progress int

	// Progress required to fill the bar.
	max int

	mu sync.RWMutex
}

// NewProgressBar returns a new progress bar.
func NewProgressBar() *Progress {
	p := &Progress{
		box:         NewBox(),
		emptyRune:   tcell.RuneBlock,
		emptyColor:  Styles.PrimitiveBackgroundColor,
		filledRune:  tcell.RuneBlock,
		filledColor: Styles.PrimaryTextColor,
		max:         100,
	}
	p.SetBackgroundColor(Styles.PrimitiveBackgroundColor)
	return p
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (p *Progress) set(setter func(p *Progress)) *Progress {
	p.mu.Lock()
	setter(p)
	p.mu.Unlock()
	return p
}

func (p *Progress) get(getter func(p *Progress)) {
	p.mu.RLock()
	getter(p)
	p.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Progress.
func (p *Progress) GetTitle() string {
	return p.box.GetTitle()
}

// SetTitle sets the title of this Progress.
func (p *Progress) SetTitle(title string) *Progress {
	p.box.SetTitle(title)
	return p
}

// GetTitleAlign returns the title alignment of this Progress.
func (p *Progress) GetTitleAlign() int {
	return p.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Progress.
func (p *Progress) SetTitleAlign(align int) *Progress {
	p.box.SetTitleAlign(align)
	return p
}

// GetBorder returns whether this Progress has a border.
func (p *Progress) GetBorder() bool {
	return p.box.GetBorder()
}

// SetBorder sets whether this Progress has a border.
func (p *Progress) SetBorder(show bool) *Progress {
	p.box.SetBorder(show)
	return p
}

// GetBorderColor returns the border color of this Progress.
func (p *Progress) GetBorderColor() tcell.Color {
	return p.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Progress.
func (p *Progress) SetBorderColor(color tcell.Color) *Progress {
	p.box.SetBorderColor(color)
	return p
}

// GetBorderAttributes returns the border attributes of this Progress.
func (p *Progress) GetBorderAttributes() tcell.AttrMask {
	return p.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Progress.
func (p *Progress) SetBorderAttributes(attr tcell.AttrMask) *Progress {
	p.box.SetBorderAttributes(attr)
	return p
}

// GetBorderColorFocused returns the border color of this Progress when focusel.
func (p *Progress) GetBorderColorFocused() tcell.Color {
	return p.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Progress when focusel.
func (p *Progress) SetBorderColorFocused(color tcell.Color) *Progress {
	p.box.SetBorderColorFocused(color)
	return p
}

// GetTitleColor returns the title color of this Progress.
func (p *Progress) GetTitleColor() tcell.Color {
	return p.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Progress.
func (p *Progress) SetTitleColor(color tcell.Color) *Progress {
	p.box.SetTitleColor(color)
	return p
}

// GetDrawFunc returns the custom draw function of this Progress.
func (p *Progress) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return p.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Progress.
func (p *Progress) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Progress {
	p.box.SetDrawFunc(handler)
	return p
}

// ShowFocus sets whether this Progress should show a focus indicator when focusel.
func (p *Progress) ShowFocus(showFocus bool) *Progress {
	p.box.ShowFocus(showFocus)
	return p
}

// GetMouseCapture returns the mouse capture function of this Progress.
func (p *Progress) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return p.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Progress.
func (p *Progress) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Progress {
	p.box.SetMouseCapture(capture)
	return p
}

// GetBackgroundColor returns the background color of this Progress.
func (p *Progress) GetBackgroundColor() tcell.Color {
	return p.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Progress.
func (p *Progress) SetBackgroundColor(color tcell.Color) *Progress {
	p.box.SetBackgroundColor(color)
	return p
}

// GetBackgroundTransparent returns whether the background of this Progress is transparent.
func (p *Progress) GetBackgroundTransparent() bool {
	return p.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Progress is transparent.
func (p *Progress) SetBackgroundTransparent(transparent bool) *Progress {
	p.box.SetBackgroundTransparent(transparent)
	return p
}

// GetInputCapture returns the input capture function of this Progress.
func (p *Progress) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return p.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Progress.
func (p *Progress) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Progress {
	p.box.SetInputCapture(capture)
	return p
}

// GetPadding returns the padding of this Progress.
func (p *Progress) GetPadding() (top, bottom, left, right int) {
	return p.box.GetPadding()
}

// SetPadding sets the padding of this Progress.
func (p *Progress) SetPadding(top, bottom, left, right int) *Progress {
	p.box.SetPadding(top, bottom, left, right)
	return p
}

// InRect returns whether the given screen coordinates are within this Progress.
func (p *Progress) InRect(x, y int) bool {
	return p.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Progress.
func (p *Progress) GetInnerRect() (x, y, width, height int) {
	return p.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Progress is preservel.
func (p *Progress) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Progress is preservel.
func (p *Progress) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Progress.
func (p *Progress) GetRect() (x, y, width, height int) {
	return p.box.GetRect()
}

// SetRect sets the rectangle occupied by this Progress.
func (p *Progress) SetRect(x, y, width, height int) {
	p.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Progress is visible.
func (p *Progress) GetVisible() bool {
	return p.box.GetVisible()
}

// SetVisible sets whether this Progress is visible.
func (p *Progress) SetVisible(visible bool) {
	p.box.SetVisible(visible)
}

// Focus is called when this Progress receives focus.
func (p *Progress) Focus(delegate func(p Widget)) {
	p.box.Focus(delegate)
}

// HasFocus returns whether this Progress has focus.
func (p *Progress) HasFocus() bool {
	return p.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this Progress.
func (p *Progress) GetFocusable() Focusable {
	return p.box.GetFocusable()
}

// Blur is called when this Progress loses focus.
func (p *Progress) Blur() {
	p.box.Blur()
}

// SetEmptyRune sets the rune used for the empty area of the progress bar.
func (p *Progress) SetEmptyRune(empty rune) *Progress {
	return p.set(func(p *Progress) { p.emptyRune = empty })
}

// SetEmptyColor sets the color of the empty area of the progress bar.
func (p *Progress) SetEmptyColor(empty tcell.Color) *Progress {
	return p.set(func(p *Progress) { p.emptyColor = empty })
}

// SetFilledRune sets the rune used for the filled area of the progress bar.
func (p *Progress) SetFilledRune(filled rune) *Progress {
	return p.set(func(p *Progress) { p.filledRune = filled })
}

// SetFilledColor sets the color of the filled area of the progress bar.
func (p *Progress) SetFilledColor(filled tcell.Color) *Progress {
	return p.set(func(p *Progress) { p.filledColor = filled })
}

// SetVertical sets the direction of the progress bar.
func (p *Progress) SetVertical(vertical bool) *Progress {
	return p.set(func(p *Progress) { p.vertical = vertical })
}

// SetMax sets the progress required to fill the bar.
func (p *Progress) SetMax(max int) *Progress {
	return p.set(func(p *Progress) { p.max = max })
}

// GetMax returns the progress required to fill the bar.
func (p *Progress) GetMax() (max int) {
	p.get(func(p *Progress) { max = p.max })
	return
}

// AddProgress adds to the current progress.
func (p *Progress) AddProgress(progress int) *Progress {
	return p.set(func(p *Progress) {
		p.progress += progress
		if p.progress < 0 {
			p.progress = 0
		} else if p.progress > p.max {
			p.progress = p.max
		}
	})
}

// SetProgress sets the current progress.
func (p *Progress) SetProgress(progress int) *Progress {
	return p.set(func(p *Progress) {
		p.progress = progress
		if p.progress < 0 {
			p.progress = 0
		} else if p.progress > p.max {
			p.progress = p.max
		}
	})
}

// GetProgress gets the current progress.
func (p *Progress) GetProgress() (progress int) {
	p.get(func(p *Progress) { progress = p.progress })
	return
}

// Complete returns whether the progress bar has been filled.
func (p *Progress) Complete() (complete bool) {
	p.get(func(p *Progress) { complete = p.progress >= p.max })
	return
}

func (p *Progress) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.InputHandler()
}

func (p *Progress) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.box.MouseHandler()
}

// Draw draws this primitive onto the screen.
func (p *Progress) Draw(screen tcell.Screen) {
	if !p.GetVisible() {
		return
	}

	p.box.Draw(screen)

	p.mu.Lock()
	defer p.mu.Unlock()

	x, y, width, height := p.GetInnerRect()

	barSize := height
	maxLength := width
	if p.vertical {
		barSize = width
		maxLength = height
	}

	barLength := int(math.RoundToEven(float64(maxLength) * (float64(p.progress) / float64(p.max))))
	if barLength > maxLength {
		barLength = maxLength
	}

	for i := 0; i < barSize; i++ {
		for j := 0; j < barLength; j++ {
			if p.vertical {
				screen.SetContent(x+i, y+(height-1-j), p.filledRune, nil, tcell.StyleDefault.Foreground(p.filledColor).Background(p.box.backgroundColor))
			} else {
				screen.SetContent(x+j, y+i, p.filledRune, nil, tcell.StyleDefault.Foreground(p.filledColor).Background(p.box.backgroundColor))
			}
		}
		for j := barLength; j < maxLength; j++ {
			if p.vertical {
				screen.SetContent(x+i, y+(height-1-j), p.emptyRune, nil, tcell.StyleDefault.Foreground(p.emptyColor).Background(p.box.backgroundColor))
			} else {
				screen.SetContent(x+j, y+i, p.emptyRune, nil, tcell.StyleDefault.Foreground(p.emptyColor).Background(p.box.backgroundColor))
			}
		}
	}
}
