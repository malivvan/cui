package cui

import (
	"math"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// ProgressBar indicates the progress of an operation.
type ProgressBar struct {
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
func NewProgressBar() *ProgressBar {
	p := &ProgressBar{
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

func (p *ProgressBar) set(setter func(p *ProgressBar)) *ProgressBar {
	p.mu.Lock()
	setter(p)
	p.mu.Unlock()
	return p
}

func (p *ProgressBar) get(getter func(p *ProgressBar)) {
	p.mu.RLock()
	getter(p)
	p.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this ProgressBar.
func (p *ProgressBar) GetTitle() string {
	return p.box.GetTitle()
}

// SetTitle sets the title of this ProgressBar.
func (p *ProgressBar) SetTitle(title string) *ProgressBar {
	p.box.SetTitle(title)
	return p
}

// GetTitleAlign returns the title alignment of this ProgressBar.
func (p *ProgressBar) GetTitleAlign() int {
	return p.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this ProgressBar.
func (p *ProgressBar) SetTitleAlign(align int) *ProgressBar {
	p.box.SetTitleAlign(align)
	return p
}

// GetBorder returns whether this ProgressBar has a border.
func (p *ProgressBar) GetBorder() bool {
	return p.box.GetBorder()
}

// SetBorder sets whether this ProgressBar has a border.
func (p *ProgressBar) SetBorder(show bool) *ProgressBar {
	p.box.SetBorder(show)
	return p
}

// GetBorderColor returns the border color of this ProgressBar.
func (p *ProgressBar) GetBorderColor() tcell.Color {
	return p.box.GetBorderColor()
}

// SetBorderColor sets the border color of this ProgressBar.
func (p *ProgressBar) SetBorderColor(color tcell.Color) *ProgressBar {
	p.box.SetBorderColor(color)
	return p
}

// GetBorderAttributes returns the border attributes of this ProgressBar.
func (p *ProgressBar) GetBorderAttributes() tcell.AttrMask {
	return p.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this ProgressBar.
func (p *ProgressBar) SetBorderAttributes(attr tcell.AttrMask) *ProgressBar {
	p.box.SetBorderAttributes(attr)
	return p
}

// GetBorderColorFocused returns the border color of this ProgressBar when focusel.
func (p *ProgressBar) GetBorderColorFocused() tcell.Color {
	return p.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this ProgressBar when focusel.
func (p *ProgressBar) SetBorderColorFocused(color tcell.Color) *ProgressBar {
	p.box.SetBorderColorFocused(color)
	return p
}

// GetTitleColor returns the title color of this ProgressBar.
func (p *ProgressBar) GetTitleColor() tcell.Color {
	return p.box.GetTitleColor()
}

// SetTitleColor sets the title color of this ProgressBar.
func (p *ProgressBar) SetTitleColor(color tcell.Color) *ProgressBar {
	p.box.SetTitleColor(color)
	return p
}

// GetDrawFunc returns the custom draw function of this ProgressBar.
func (p *ProgressBar) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return p.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this ProgressBar.
func (p *ProgressBar) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *ProgressBar {
	p.box.SetDrawFunc(handler)
	return p
}

// ShowFocus sets whether this ProgressBar should show a focus indicator when focusel.
func (p *ProgressBar) ShowFocus(showFocus bool) *ProgressBar {
	p.box.ShowFocus(showFocus)
	return p
}

// GetMouseCapture returns the mouse capture function of this ProgressBar.
func (p *ProgressBar) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return p.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this ProgressBar.
func (p *ProgressBar) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *ProgressBar {
	p.box.SetMouseCapture(capture)
	return p
}

// GetBackgroundColor returns the background color of this ProgressBar.
func (p *ProgressBar) GetBackgroundColor() tcell.Color {
	return p.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this ProgressBar.
func (p *ProgressBar) SetBackgroundColor(color tcell.Color) *ProgressBar {
	p.box.SetBackgroundColor(color)
	return p
}

// GetBackgroundTransparent returns whether the background of this ProgressBar is transparent.
func (p *ProgressBar) GetBackgroundTransparent() bool {
	return p.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this ProgressBar is transparent.
func (p *ProgressBar) SetBackgroundTransparent(transparent bool) *ProgressBar {
	p.box.SetBackgroundTransparent(transparent)
	return p
}

// GetInputCapture returns the input capture function of this ProgressBar.
func (p *ProgressBar) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return p.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this ProgressBar.
func (p *ProgressBar) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *ProgressBar {
	p.box.SetInputCapture(capture)
	return p
}

// GetPadding returns the padding of this ProgressBar.
func (p *ProgressBar) GetPadding() (top, bottom, left, right int) {
	return p.box.GetPadding()
}

// SetPadding sets the padding of this ProgressBar.
func (p *ProgressBar) SetPadding(top, bottom, left, right int) *ProgressBar {
	p.box.SetPadding(top, bottom, left, right)
	return p
}

// InRect returns whether the given screen coordinates are within this ProgressBar.
func (p *ProgressBar) InRect(x, y int) bool {
	return p.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this ProgressBar.
func (p *ProgressBar) GetInnerRect() (x, y, width, height int) {
	return p.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the ProgressBar is preservel.
func (p *ProgressBar) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the ProgressBar is preservel.
func (p *ProgressBar) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this ProgressBar.
func (p *ProgressBar) GetRect() (x, y, width, height int) {
	return p.box.GetRect()
}

// SetRect sets the rectangle occupied by this ProgressBar.
func (p *ProgressBar) SetRect(x, y, width, height int) {
	p.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this ProgressBar is visible.
func (p *ProgressBar) GetVisible() bool {
	return p.box.GetVisible()
}

// SetVisible sets whether this ProgressBar is visible.
func (p *ProgressBar) SetVisible(visible bool) {
	p.box.SetVisible(visible)
}

// Focus is called when this ProgressBar receives focus.
func (p *ProgressBar) Focus(delegate func(p Widget)) {
	p.box.Focus(delegate)
}

// HasFocus returns whether this ProgressBar has focus.
func (p *ProgressBar) HasFocus() bool {
	return p.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this ProgressBar.
func (p *ProgressBar) GetFocusable() Focusable {
	return p.box.GetFocusable()
}

// Blur is called when this ProgressBar loses focus.
func (p *ProgressBar) Blur() {
	p.box.Blur()
}

// SetEmptyRune sets the rune used for the empty area of the progress bar.
func (p *ProgressBar) SetEmptyRune(empty rune) *ProgressBar {
	return p.set(func(p *ProgressBar) { p.emptyRune = empty })
}

// SetEmptyColor sets the color of the empty area of the progress bar.
func (p *ProgressBar) SetEmptyColor(empty tcell.Color) *ProgressBar {
	return p.set(func(p *ProgressBar) { p.emptyColor = empty })
}

// SetFilledRune sets the rune used for the filled area of the progress bar.
func (p *ProgressBar) SetFilledRune(filled rune) *ProgressBar {
	return p.set(func(p *ProgressBar) { p.filledRune = filled })
}

// SetFilledColor sets the color of the filled area of the progress bar.
func (p *ProgressBar) SetFilledColor(filled tcell.Color) *ProgressBar {
	return p.set(func(p *ProgressBar) { p.filledColor = filled })
}

// SetVertical sets the direction of the progress bar.
func (p *ProgressBar) SetVertical(vertical bool) *ProgressBar {
	return p.set(func(p *ProgressBar) { p.vertical = vertical })
}

// SetMax sets the progress required to fill the bar.
func (p *ProgressBar) SetMax(max int) *ProgressBar {
	return p.set(func(p *ProgressBar) { p.max = max })
}

// GetMax returns the progress required to fill the bar.
func (p *ProgressBar) GetMax() (max int) {
	p.get(func(p *ProgressBar) { max = p.max })
	return
}

// AddProgress adds to the current progress.
func (p *ProgressBar) AddProgress(progress int) *ProgressBar {
	return p.set(func(p *ProgressBar) {
		p.progress += progress
		if p.progress < 0 {
			p.progress = 0
		} else if p.progress > p.max {
			p.progress = p.max
		}
	})
}

// SetProgress sets the current progress.
func (p *ProgressBar) SetProgress(progress int) *ProgressBar {
	return p.set(func(p *ProgressBar) {
		p.progress = progress
		if p.progress < 0 {
			p.progress = 0
		} else if p.progress > p.max {
			p.progress = p.max
		}
	})
}

// GetProgress gets the current progress.
func (p *ProgressBar) GetProgress() (progress int) {
	p.get(func(p *ProgressBar) { progress = p.progress })
	return
}

// Complete returns whether the progress bar has been filled.
func (p *ProgressBar) Complete() (complete bool) {
	p.get(func(p *ProgressBar) { complete = p.progress >= p.max })
	return
}

func (p *ProgressBar) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.InputHandler()
}

func (p *ProgressBar) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.box.MouseHandler()
}

// Draw draws this primitive onto the screen.
func (p *ProgressBar) Draw(screen tcell.Screen) {
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
