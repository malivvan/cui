package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Box is the base Widget for all widgets. It has a background color and
// optional surrounding elements such as a border and a title. It does not have
// inner text. Widgets embed Box and draw their text over it.
type Box struct {

	// The position of the rect.
	x, y, width, height int

	// Padding.
	paddingTop, paddingBottom, paddingLeft, paddingRight int

	// The inner rect reserved for the box's content.
	innerX, innerY, innerWidth, innerHeight int

	// Whether the box is visible.
	visible bool

	// The border color when the box has focus.
	borderColorFocused tcell.Color

	// The box's background color.
	backgroundColor tcell.Color

	// Whether the box's background is transparent.
	backgroundTransparent bool

	// Whether or not a border is drawn, reducing the box's space for content by
	// two in width and height.
	border bool

	// The color of the border.
	borderColor tcell.Color

	// The style attributes of the border.
	borderAttributes tcell.AttrMask

	// The title. Only visible if there is a border, too.
	title []byte

	// The color of the title.
	titleColor tcell.Color

	// The alignment of the title.
	titleAlign int

	// Provides a way to find out if this box has focus. We always go through
	// this interface because it may be overridden by implementing classes.
	focus Focusable

	// Whether this box has focus.
	hasFocus bool

	// Whether this box shows its focus.
	showFocus bool

	// An optional capture function which receives a key event and returns the
	// event to be forwarded to the primitive's default input handler (nil if
	// nothing should be forwarded).
	inputCapture func(event *tcell.EventKey) *tcell.EventKey

	// An optional function which is called before the box is drawn.
	draw func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)

	// An optional capture function which receives a mouse event and returns the
	// event to be forwarded to the primitive's default mouse event handler (at
	// least one nil if nothing should be forwarded).
	mouseCapture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)

	mu sync.RWMutex
}

// NewBox returns a Box without a border.
func NewBox() *Box {
	b := &Box{
		width:              15,
		height:             10,
		visible:            true,
		backgroundColor:    Styles.PrimitiveBackgroundColor,
		borderColor:        Styles.BorderColor,
		titleColor:         Styles.TitleColor,
		borderColorFocused: ColorUnset,
		titleAlign:         AlignCenter,
		showFocus:          true,
	}
	b.focus = b
	b.updateInnerRect()
	return b
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (b *Box) set(setter func(b *Box)) *Box {
	b.mu.Lock()
	setter(b)
	b.mu.Unlock()
	return b
}

func (b *Box) get(getter func(b *Box)) {
	b.mu.RLock()
	getter(b)
	b.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the box's current title.
func (b *Box) GetTitle() (title string) {
	b.get(func(b *Box) { title = string(b.title) })
	return
}

// SetTitle sets the box's title.
func (b *Box) SetTitle(title string) *Box {
	return b.set(func(b *Box) { b.title = []byte(title) })
}

// GetTitleColor returns the box's title color.
func (b *Box) GetTitleColor() (color tcell.Color) {
	b.get(func(b *Box) { color = b.titleColor })
	return
}

// SetTitleColor sets the box's title color.
func (b *Box) SetTitleColor(color tcell.Color) *Box {
	return b.set(func(b *Box) { b.titleColor = color })
}

func (b *Box) GetTitleAlign() (align int) {
	b.get(func(b *Box) { align = b.titleAlign })
	return
}

// SetTitleAlign sets the alignment of the title, one of AlignLeft, AlignCenter or AlignRight.
func (b *Box) SetTitleAlign(align int) *Box {
	return b.set(func(b *Box) { b.titleAlign = align })
}

// GetPadding returns the size of the padding around the box content.
func (b *Box) GetPadding() (top, bottom, left, right int) {
	b.get(func(b *Box) { top, bottom, left, right = b.paddingTop, b.paddingBottom, b.paddingLeft, b.paddingRight })
	return
}

// SetPadding sets the size of the padding around the box content.
func (b *Box) SetPadding(top, bottom, left, right int) *Box {
	return b.set(func(b *Box) {
		b.paddingTop, b.paddingBottom, b.paddingLeft, b.paddingRight = top, bottom, left, right
		b.updateInnerRect()
	})
}

// SetBorder sets the flag indicating whether the box should have a border.
func (b *Box) SetBorder(show bool) *Box {
	return b.set(func(b *Box) { b.border = show; b.updateInnerRect() })
}

// GetBorderColor returns the box's border color.
func (b *Box) GetBorderColor() (color tcell.Color) {
	b.get(func(b *Box) { color = b.borderColor })
	return
}

// SetBorderColor sets the box's border color.
func (b *Box) SetBorderColor(color tcell.Color) *Box {
	return b.set(func(b *Box) { b.borderColor = color })
}

// GetBorderColorFocused returns the box's border color when the box is focused.
func (b *Box) GetBorderColorFocused() (color tcell.Color) {
	b.get(func(b *Box) { color = b.borderColorFocused })
	return
}

// SetBorderColorFocused sets the box's border color when the box is focused.
func (b *Box) SetBorderColorFocused(color tcell.Color) *Box {
	return b.set(func(b *Box) { b.borderColorFocused = color })
}

// GetBorderAttributes returns the border's style attributes.
func (b *Box) GetBorderAttributes() (attr tcell.AttrMask) {
	b.get(func(b *Box) { attr = b.borderAttributes })
	return
}

// SetBorderAttributes sets the border's style attributes. You can combine
// different attributes using bitmask operations:
//
// box.SetBorderAttributes(tcell.AttrUnderline | tcell.AttrBold)
func (b *Box) SetBorderAttributes(attr tcell.AttrMask) *Box {
	return b.set(func(b *Box) { b.borderAttributes = attr })
}

// GetBackgroundTransparent returns the flag indicating whether the box's
// background is transparent.
func (b *Box) GetBackgroundTransparent() (transparent bool) {
	b.get(func(b *Box) { transparent = b.backgroundTransparent })
	return
}

// SetBackgroundTransparent sets the flag indicating whether the box's
// background is transparent. The screen is not cleared before drawing the
// application. Overlaying transparent widgets directly onto the screen may
// result in artifacts. To resolve this, add a blank, non-transparent Box to
// the bottom layer of the interface via Panels, or set a handler via
// SetBeforeDrawFunc which clears the screen.
func (b *Box) SetBackgroundTransparent(transparent bool) *Box {
	return b.set(func(b *Box) { b.backgroundTransparent = transparent })
}

// GetDrawFunc returns the callback function which was installed with
// SetDrawFunc() or nil if no such function has been installed.
func (b *Box) GetDrawFunc() (draw func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) {
	b.get(func(b *Box) { b.draw = draw })
	return b.draw
}

// SetDrawFunc sets a callback function which is invoked after the box primitive
// has been drawn. This allows you to add a more individual style to the box
// (and all primitives which extend it).
//
// The function is provided with the box's dimensions (set via SetRect()). It
// must return the box's inner dimensions (x, y, width, height) which will be
// returned by GetInnerRect(), used by descendent primitives to draw their own
// content.
func (b *Box) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Box {
	return b.set(func(b *Box) { b.draw = handler })
}

// GetInputCapture returns the function installed with SetInputCapture() or nil
// if no such function has been installed.
func (b *Box) GetInputCapture() (fn func(event *tcell.EventKey) *tcell.EventKey) {
	b.get(func(b *Box) { fn = b.inputCapture })
	return
}

// SetInputCapture installs a function which captures key events before they are
// forwarded to the primitive's default key event handler. This function can
// then choose to forward that key event (or a different one) to the default
// handler by returning it. If nil is returned, the default handler will not
// be called.
//
// Providing a nil handler will remove a previously existing handler.
//
// Note that this function will not have an effect on primitives composed of
// other primitives, such as Form, Flex, or Grid. Key events are only captured
// by the primitives that have focus (e.g. InputField) and only one primitive
// can have focus at a time. Composing primitives such as Form pass the focus on
// to their contained primitives and thus never receive any key events
// themselves. Therefore, they cannot intercept key events.
func (b *Box) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Box {
	return b.set(func(b *Box) { b.inputCapture = capture })
}

// WrapInputHandler wraps an input handler (see InputHandler()) with the
// functionality to capture input (see SetInputCapture()) before passing it
// on to the provided (default) input handler.
//
// This is only meant to be used by subclassing primitives.
func (b *Box) WrapInputHandler(inputHandler func(*tcell.EventKey, func(p Widget))) func(*tcell.EventKey, func(p Widget)) {
	return func(event *tcell.EventKey, setFocus func(p Widget)) {
		if b.inputCapture != nil {
			event = b.inputCapture(event)
		}
		if event != nil && inputHandler != nil {
			inputHandler(event, setFocus)
		}
	}
}

////////////////////////////////// <API> ////////////////////////////////////

// GetInnerRect returns the position of the inner rectangle (x, y, width,
// height), without the border and without any padding. Width and height values
// will clamp to 0 and thus never be negative.
func (b *Box) GetInnerRect() (innerX, innerY, innerW, innerH int) {
	b.get(func(b *Box) { innerX, innerY, innerW, innerH = b.innerX, b.innerY, b.innerWidth, b.innerHeight })
	return
}

// WrapMouseHandler wraps a mouse event handler (see MouseHandler()) with the
// functionality to capture mouse events (see SetMouseCapture()) before passing
// them on to the provided (default) event handler.
//
// This is only meant to be used by subclassing primitives.
func (b *Box) WrapMouseHandler(mouseHandler func(MouseAction, *tcell.EventMouse, func(p Widget)) (bool, Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if b.mouseCapture != nil {
			action, event = b.mouseCapture(action, event)
		}
		if event != nil && mouseHandler != nil {
			consumed, capture = mouseHandler(action, event, setFocus)
		}
		return
	}
}

// SetMouseCapture sets a function which captures mouse events (consisting of
// the original tcell mouse event and the semantic mouse action) before they are
// forwarded to the primitive's default mouse event handler. This function can
// then choose to forward that event (or a different one) by returning it or
// returning a nil mouse event, in which case the default handler will not be
// called.
//
// Providing a nil handler will remove a previously existing handler.
func (b *Box) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Box {
	b.mouseCapture = capture
	return b
}

// InRect returns true if the given coordinate is within the bounds of the box's
// rectangle.
func (b *Box) InRect(x, y int) bool {
	rectX, rectY, width, height := b.GetRect()
	return x >= rectX && x < rectX+width && y >= rectY && y < rectY+height
}

// GetMouseCapture returns the function installed with SetMouseCapture() or nil
// if no such function has been installed.
func (b *Box) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return b.mouseCapture
}

// GetBackgroundColor returns the box's background color.
func (b *Box) GetBackgroundColor() (backgroundColor tcell.Color) {
	b.get(func(b *Box) { backgroundColor = b.backgroundColor })
	return
}

// SetBackgroundColor sets the box's background color.
func (b *Box) SetBackgroundColor(color tcell.Color) *Box {
	return b.set(func(b *Box) { b.backgroundColor = color })
}

// GetBorder returns a value indicating whether the box have a border or not.
func (b *Box) GetBorder() (border bool) {
	b.get(func(b *Box) { border = b.border })
	return
}

// ShowFocus sets the flag indicating whether the borders of this
// primitive should change thickness when focused.
func (b *Box) ShowFocus(showFocus bool) *Box {
	return b.set(func(b *Box) { b.showFocus = showFocus })
}

func (b *Box) updateInnerRect() {
	x, y, width, height := b.x, b.y, b.width, b.height

	// Subtract border space
	if b.border {
		x++
		y++
		width -= 2
		height -= 2
	}

	// Subtract padding
	x, y, width, height =
		x+b.paddingLeft,
		y+b.paddingTop,
		width-b.paddingLeft-b.paddingRight,
		height-b.paddingTop-b.paddingBottom

	if width < 0 {
		width = 0
	}
	if height < 0 {
		height = 0
	}

	b.innerX, b.innerY, b.innerWidth, b.innerHeight = x, y, width, height
}

//////////////////////////////////// <WIDGET> ////////////////////////////////////

// GetRect returns the current position of the rectangle, x, y, width, and height.
func (b *Box) GetRect() (x, y, w, h int) {
	b.get(func(b *Box) { x, y, w, h = b.x, b.y, b.width, b.height })
	return
}

// Focus is called when this primitive receives focus.
func (b *Box) Focus(delegate func(p Widget)) {
	b.set(func(b *Box) { b.hasFocus = true })
}

// Blur is called when this primitive loses focus.
func (b *Box) Blur() {
	b.set(func(b *Box) { b.hasFocus = false })
}

// HasFocus returns whether this primitive has focus.
func (b *Box) HasFocus() (hasFocus bool) {
	b.get(func(b *Box) { hasFocus = b.hasFocus })
	return
}

// GetFocusable returns the item's Focusable.
func (b *Box) GetFocusable() (focus Focusable) {
	b.get(func(b *Box) { focus = b.focus })
	return
}

// SetRect sets a new position of the primitive. Note that this has no effect
// if this primitive is part of a layout (e.g. Flex, Grid) or if it was added
// like this:
//
//	application.SetRoot(b, true)
func (b *Box) SetRect(x, y, w, h int) {
	b.set(func(b *Box) {
		b.x, b.y, b.width, b.height = x, y, w, h
		b.updateInnerRect()
	})
}

// SetVisible sets the flag indicating whether the box is visible.
func (b *Box) SetVisible(v bool) {
	_ = b.set(func(b *Box) { b.visible = v })
}

// GetVisible returns a value indicating whether the box is visible.
func (b *Box) GetVisible() (v bool) {
	b.get(func(b *Box) { v = b.visible })
	return
}

// InputHandler returns nil.
func (b *Box) InputHandler() (fn func(event *tcell.EventKey, setFocus func(p Widget))) {
	b.get(func(b *Box) { fn = b.WrapInputHandler(nil) })
	return
}

// MouseHandler returns nil.
func (b *Box) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return b.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if action == MouseLeftClick && b.InRect(event.Position()) {
			setFocus(b)
			consumed = true
		}
		return
	})
}

// Draw draws this primitive onto the screen.
func (b *Box) Draw(screen tcell.Screen) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Don't draw anything if the box is hidden
	if !b.visible {
		return
	}

	// Don't draw anything if there is no space.
	if b.width <= 0 || b.height <= 0 {
		return
	}

	def := tcell.StyleDefault

	// Fill background.
	background := def.Background(b.backgroundColor)
	if !b.backgroundTransparent {
		for y := b.y; y < b.y+b.height; y++ {
			for x := b.x; x < b.x+b.width; x++ {
				screen.SetContent(x, y, ' ', nil, background)
			}
		}
	}

	// Draw border.
	if b.border && b.width >= 2 && b.height >= 2 {
		border := SetAttributes(background.Foreground(b.borderColor), b.borderAttributes)
		var vertical, horizontal, topLeft, topRight, bottomLeft, bottomRight rune

		var hasFocus bool
		if b.focus == b {
			hasFocus = b.hasFocus
		} else {
			hasFocus = b.focus.HasFocus()
		}

		if hasFocus && b.borderColorFocused != ColorUnset {
			border = SetAttributes(background.Foreground(b.borderColorFocused), b.borderAttributes)
		}

		if hasFocus && b.showFocus {
			horizontal = Borders.HorizontalFocus
			vertical = Borders.VerticalFocus
			topLeft = Borders.TopLeftFocus
			topRight = Borders.TopRightFocus
			bottomLeft = Borders.BottomLeftFocus
			bottomRight = Borders.BottomRightFocus
		} else {
			horizontal = Borders.Horizontal
			vertical = Borders.Vertical
			topLeft = Borders.TopLeft
			topRight = Borders.TopRight
			bottomLeft = Borders.BottomLeft
			bottomRight = Borders.BottomRight
		}
		for x := b.x + 1; x < b.x+b.width-1; x++ {
			screen.SetContent(x, b.y, horizontal, nil, border)
			screen.SetContent(x, b.y+b.height-1, horizontal, nil, border)
		}
		for y := b.y + 1; y < b.y+b.height-1; y++ {
			screen.SetContent(b.x, y, vertical, nil, border)
			screen.SetContent(b.x+b.width-1, y, vertical, nil, border)
		}
		screen.SetContent(b.x, b.y, topLeft, nil, border)
		screen.SetContent(b.x+b.width-1, b.y, topRight, nil, border)
		screen.SetContent(b.x, b.y+b.height-1, bottomLeft, nil, border)
		screen.SetContent(b.x+b.width-1, b.y+b.height-1, bottomRight, nil, border)

		// Draw title.
		if len(b.title) > 0 && b.width >= 4 {
			printed, _ := Print(screen, b.title, b.x+1, b.y, b.width-2, b.titleAlign, b.titleColor)
			if len(b.title)-printed > 0 && printed > 0 {
				_, _, style, _ := screen.GetContent(b.x+b.width-2, b.y)
				fg, _, _ := style.Decompose()
				Print(screen, []byte(string(SemigraphicsHorizontalEllipsis)), b.x+b.width-2, b.y, 1, AlignLeft, fg)
			}
		}
	}

	// Call custom draw function.
	if b.draw != nil {
		b.innerX, b.innerY, b.innerWidth, b.innerHeight = b.draw(screen, b.x, b.y, b.width, b.height)
	}
}
