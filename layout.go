package cui

import (
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
)

const (
	VerticalLayout = iota
	HorizontalLayout
)

const (
	AutoSize = 0
)

type LayoutItem struct {
	Primitive

	Size int
}

type Layout struct {
	// The items contained in the layout
	items []*LayoutItem

	dragX, dragY int

	focusedSplitterNumber int
	draggedSplitter       *layoutSplitter
	splitters             []*layoutSplitter

	// Whether or not a splitterFlag is drawn, reducing the box's space for content by
	// two in width and height.
	splitterFlag bool

	// The border style.
	splitterStyle tcell.Style

	direction int

	x, y, width, height int

	visible bool

	// The layout's background color.
	backgroundColor tcell.Color

	// An optional capture function which receives a key event and returns the
	// event to be forwarded to the primitive's default input handler (nil if
	// nothing should be forwarded).
	inputCapture func(event *tcell.EventKey) *tcell.EventKey

	// An optional capture function which receives a mouse event and returns the
	// event to be forwarded to the primitive's default mouse event handler (at
	// least one nil if nothing should be forwarded).
	mouseCapture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)

	l sync.RWMutex
}

func NewLayout() *Layout {
	layout := &Layout{
		backgroundColor:       Styles.PrimitiveBackgroundColor,
		focusedSplitterNumber: -1,
		splitterStyle:         tcell.StyleDefault.Foreground(Styles.BorderColor),
		visible:               true,
	}

	return layout
}

func (l *Layout) SetDirection(d int) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.direction = d
	return l
}

func (l *Layout) GetDirection() int {
	l.l.RLock()
	defer l.l.RUnlock()

	return l.direction
}

// SetBackgroundColor sets the layout's background color.
func (l *Layout) SetBackgroundColor(color tcell.Color) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.backgroundColor = color
	return l
}

// SetSplitter sets the flag indicating whether or not the layout should render a
// splitters between primitives
func (l *Layout) SetSplitter(show bool) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.splitterFlag = show
	return l
}

// SetSplitterColor sets the layout's splitter color.
func (l *Layout) SetSplitterColor(color tcell.Color) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.splitterStyle = l.splitterStyle.Foreground(color)
	return l
}

// SetSplitterAttributes sets the splitter's style attributes. You can combine
// different attributes using bitmask operations:
//
//	layout.SetSplitterAttributes(tcell.AttrUnderline | tcell.AttrBold)
func (l *Layout) SetSplitterAttributes(attr tcell.AttrMask) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.splitterStyle = l.splitterStyle.Attributes(attr)
	return l
}

// GetSplitterAttributes returns the splitter's style attributes.
func (l *Layout) GetSplitterAttributes() tcell.AttrMask {
	l.l.RLock()
	defer l.l.RUnlock()

	_, _, attr := l.splitterStyle.Decompose()
	return attr
}

// GetSplitterColor returns the layout's splitter color.
func (l *Layout) GetSplitterColor() tcell.Color {
	l.l.RLock()
	defer l.l.RUnlock()

	color, _, _ := l.splitterStyle.Decompose()
	return color
}

// GetBackgroundColor returns the layout's background color.
func (l *Layout) GetBackgroundColor() tcell.Color {
	l.l.RLock()
	defer l.l.RUnlock()

	return l.backgroundColor
}

func (l *Layout) itemsAmount() (int, int) {
	auto := 0
	for _, item := range l.items {
		if item.Size == AutoSize {
			auto += 1
		}
	}

	return len(l.items) - auto, auto
}

func (l *Layout) itemsSize() int {
	size := 0
	for _, item := range l.items {
		size += item.Size
	}

	return size
}

func (l *Layout) availableSpace() int {
	switch l.direction {
	case HorizontalLayout:
		return l.width
	case VerticalLayout:
		return l.height
	default:
		return 0
	}
}

func (l *Layout) splittersAmount() int {
	if len(l.items) <= 1 {
		return 0
	} else {
		return len(l.items) - 1
	}
}

func (l *Layout) Draw(screen tcell.Screen) {
	l.l.Lock()
	defer l.l.Unlock()

	// Don't draw anything if the box is hidden
	if !l.visible {
		return
	}

	// Don't draw anything if there is no space.
	if l.width <= 0 || l.height <= 0 {
		return
	}

	x, y, width, height := l.x, l.y, l.width, l.height
	def := tcell.StyleDefault

	// Fill background.
	background := def.Background(l.backgroundColor)
	if l.backgroundColor != tcell.ColorDefault {
		for y_ := y; y_ < y+height; y_++ {
			for x_ := x; x_ < x+width; x_++ {
				screen.SetContent(x_, y_, ' ', nil, background)
			}
		}
	}

	_, auto := l.itemsAmount()
	size := l.itemsSize()
	space := l.availableSpace()
	seps := l.splittersAmount()

	autoSize := 0
	if auto != 0 {
		autoSize = (space - size - seps) / auto
	}

	switch l.direction {
	case HorizontalLayout:
		vertical := Borders.Vertical

		for number, item := range l.items {
			if item.Size == AutoSize {
				item.Primitive.SetRect(x, y, autoSize, height)
				item.Primitive.Draw(newClipRegion(screen, x, y, autoSize, height))
				x += autoSize
			} else {
				item.Primitive.SetRect(x, y, item.Size, height)
				item.Primitive.Draw(newClipRegion(screen, x, y, item.Size, height))
				x += item.Size
			}

			if seps > 0 {
				if l.splitterFlag {
					if number == l.focusedSplitterNumber {
						vertical = Borders.VerticalFocus
					}

					for y_ := y; y_ < y+height; y_++ {
						screen.SetContent(x, y_, vertical, nil, l.splitterStyle)
					}
				}

				seps -= 1
				x += 1
			}
		}

	case VerticalLayout:
		horizontal := Borders.Horizontal

		for number, item := range l.items {
			if item.Size == AutoSize {
				item.Primitive.SetRect(x, y, width, autoSize)
				item.Primitive.Draw(newClipRegion(screen, x, y, width, autoSize))
				y += autoSize
			} else {
				item.Primitive.SetRect(x, y, width, item.Size)
				item.Primitive.Draw(newClipRegion(screen, x, y, width, item.Size))
				y += item.Size
			}

			if seps > 0 {
				if l.splitterFlag {
					if number == l.focusedSplitterNumber {
						horizontal = Borders.HorizontalFocus
					}

					for x_ := x; x_ < x+width; x_++ {
						screen.SetContent(x_, y, horizontal, nil, l.splitterStyle)
					}
				}

				seps -= 1
				y += 1
			}
		}
	}
}

func (l *Layout) GetRect() (int, int, int, int) {
	l.l.RLock()
	defer l.l.RUnlock()

	return l.x, l.y, l.width, l.height
}

func (l *Layout) SetRect(x, y, width, height int) {
	l.l.Lock()
	defer l.l.Unlock()

	l.x = x
	l.y = y
	l.width = width
	l.height = height

	l.rebuildSplitters()
}

func (l *Layout) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return l.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Primitive)) {
		for _, item := range l.items {
			if item.Primitive.HasFocus() {
				if handler := item.Primitive.InputHandler(); handler != nil {
					handler(event, setFocus)
					return
				}
			}
		}
	})
}

// WrapInputHandler wraps an input handler (see InputHandler()) with the
// functionality to capture input (see SetInputCapture()) before passing it
// on to the provided (default) input handler.
//
// This is only meant to be used by subclassing primitives.
func (l *Layout) WrapInputHandler(inputHandler func(*tcell.EventKey, func(p Primitive))) func(*tcell.EventKey, func(p Primitive)) {
	return func(event *tcell.EventKey, setFocus func(p Primitive)) {
		if l.inputCapture != nil {
			event = l.inputCapture(event)
		}
		if event != nil && inputHandler != nil {
			inputHandler(event, setFocus)
		}
	}
}

func (l *Layout) Focus(delegate func(p Primitive)) {
	l.l.Lock()

	if l.focusedSplitterNumber == -1 {
		for _, item := range l.items {
			if item.Primitive != nil {
				l.l.Unlock()
				delegate(item.Primitive)
				return
			}
		}
	}
	l.l.Unlock()
}

func (l *Layout) HasFocus() bool {
	l.l.RLock()
	defer l.l.RUnlock()

	for _, item := range l.items {
		if item.Primitive != nil && item.Primitive.HasFocus() {
			return true
		}
	}
	return false
}

func (l *Layout) GetFocusable() Focusable {
	l.l.RLock()
	defer l.l.RUnlock()

	for _, item := range l.items {
		if item.Primitive != nil && item.Primitive.HasFocus() {
			return item.Primitive
		}
	}
	return nil
}

func (l *Layout) GetVisible() bool {
	l.l.RLock()
	defer l.l.RUnlock()

	return l.visible
}

func (l *Layout) SetVisible(v bool) {
	l.l.Lock()
	defer l.l.Unlock()

	l.visible = v
}

func (l *Layout) Blur() {
	l.l.Lock()
	defer l.l.Unlock()

	l.focusedSplitterNumber = -1
	for _, item := range l.items {
		if item.Primitive != nil && item.Primitive.HasFocus() {
			item.Primitive.Blur()
		}
	}
}

// InRect returns true if the given coordinate is within the bounds of the box's
// rectangle.
func (l *Layout) InRect(x, y int) bool {
	l.l.RLock()
	defer l.l.RUnlock()

	rectX, rectY, width, height := l.x, l.y, l.width, l.height
	return x >= rectX && x < rectX+width && y >= rectY && y < rectY+height
}

// MouseHandler returns the mouse handler for this primitive.
func (l *Layout) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return l.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if !l.InRect(event.Position()) {
			return false, nil
		}

		// Pass mouse events along to the first child item that takes it.
		for _, item := range l.items {
			if item.Primitive != nil {
				consumed, capture = item.Primitive.MouseHandler()(action, event, setFocus)
				if consumed {
					l.focusedSplitterNumber = -1
					return
				}
			}
		}

		switch action {
		case MouseMove:
			if l.draggedSplitter != nil {
				x, y := event.Position()
				dx, dy := x-l.dragX, y-l.dragY
				wxa, wya, wwa, wha := l.draggedSplitter.a.GetRect()
				wxb, wyb, wwb, whb := l.draggedSplitter.b.GetRect()

				l.dragX = x
				l.dragY = y

				switch l.direction {
				case HorizontalLayout:
					l.draggedSplitter.a.SetRect(wxa, wya, wwa+dx, wha)
					l.draggedSplitter.b.SetRect(wxb+dx, wyb, wwb-dx, whb)
					l.draggedSplitter.a.Size = wwa + dx
					l.draggedSplitter.b.Size = wwb - dx
				case VerticalLayout:
					l.draggedSplitter.a.SetRect(wxa, wya, wwa, wha+dy)
					l.draggedSplitter.b.SetRect(wxb, wyb+dy, wwb, whb-dy)
					l.draggedSplitter.a.Size = wha + dy
					l.draggedSplitter.b.Size = whb - dy
				default:
					panic(fmt.Sprintf("invalid layout direction: %v", l.direction))
				}
				return true, nil
			}

		case MouseLeftDown:
			for number, splitter := range l.splitters {
				x, y := event.Position()
				if splitter.contain(x, y) {
					l.focusedSplitterNumber = number
					l.draggedSplitter = splitter
					l.dragX = x
					l.dragY = y

					for _, item := range l.items {
						if item.Primitive != nil {
							item.Primitive.Blur()
						}
					}

					return true, nil
				}
			}

		case MouseLeftUp:
			l.draggedSplitter = nil
			l.rebuildSplitters()
			return true, nil
		}

		return
	})
}

// WrapMouseHandler wraps a mouse event handler (see MouseHandler()) with the
// functionality to capture mouse events (see SetMouseCapture()) before passing
// them on to the provided (default) event handler.
//
// This is only meant to be used by subclassing primitives.
func (l *Layout) WrapMouseHandler(mouseHandler func(MouseAction, *tcell.EventMouse, func(p Primitive)) (bool, Primitive)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if l.mouseCapture != nil {
			action, event = l.mouseCapture(action, event)
		}
		if event != nil && mouseHandler != nil {
			consumed, capture = mouseHandler(action, event, setFocus)
		}
		return
	}
}

func (l *Layout) AddItem(p Primitive, size int) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.items = append(l.items, &LayoutItem{
		Primitive: p,
		Size:      size,
	})

	l.draggedSplitter = nil
	l.rebuildSplitters()

	return l
}

func (l *Layout) RemoveItem(i int) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	if i < 0 || i >= len(l.items) {
		return l
	}

	l.draggedSplitter = nil
	l.items = append(l.items[:i], l.items[i+1:]...)
	l.rebuildSplitters()

	return l
}

func (l *Layout) GetItem(i int) *LayoutItem {
	l.l.RLock()
	defer l.l.RUnlock()

	if i < 0 || i >= len(l.items) {
		return nil
	}

	return l.items[i]
}

func (l *Layout) CountItems() int {
	l.l.RLock()
	defer l.l.RUnlock()

	return len(l.items)
}

func (l *Layout) ClearItems() *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.items = nil
	l.splitters = nil
	return l
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
func (l *Layout) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.inputCapture = capture
	return l
}

// GetInputCapture returns the function installed with SetInputCapture() or nil
// if no such function has been installed.
func (l *Layout) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return l.inputCapture
}

// SetMouseCapture sets a function which captures mouse events (consisting of
// the original tcell mouse event and the semantic mouse action) before they are
// forwarded to the primitive's default mouse event handler. This function can
// then choose to forward that event (or a different one) by returning it or
// returning a nil mouse event, in which case the default handler will not be
// called.
//
// Providing a nil handler will remove a previously existing handler.
func (l *Layout) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Layout {
	l.l.Lock()
	defer l.l.Unlock()

	l.mouseCapture = capture
	return l
}

// GetMouseCapture returns the function installed with SetMouseCapture() or nil
// if no such function has been installed.
func (l *Layout) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return l.mouseCapture
}

func (l *Layout) rebuildSplitters() {
	l.splitters = nil

	x, y, width, height := l.x, l.y, l.width, l.height

	_, auto := l.itemsAmount()
	size := l.itemsSize()
	space := l.availableSpace()
	seps := l.splittersAmount()

	autoSize := 0
	if auto != 0 {
		autoSize = (space - size - seps) / auto
	}

	switch l.direction {
	case HorizontalLayout:
		for i := 0; i < len(l.items)-1; i++ {
			if l.items[i].Size == AutoSize {
				x += autoSize
			} else {
				x += l.items[i].Size
			}

			if seps > 0 {
				l.splitters = append(l.splitters, &layoutSplitter{
					x: [2]int{x, x},
					y: [2]int{y, y + height - 1},
					a: l.items[i],
					b: l.items[i+1],
				})

				seps -= 1
				x += 1
			}
		}

	case VerticalLayout:
		for i := 0; i < len(l.items)-1; i++ {
			if l.items[i].Size == AutoSize {
				y += autoSize
			} else {
				y += l.items[i].Size
			}

			if seps > 0 {
				l.splitters = append(l.splitters, &layoutSplitter{
					x: [2]int{x, x + width - 1},
					y: [2]int{y, y},
					a: l.items[i],
					b: l.items[i+1],
				})

				seps -= 1
				y += 1
			}
		}
	}
}

type layoutSplitter struct {
	x, y [2]int // begin and end points

	a, b *LayoutItem
}

func (s *layoutSplitter) contain(x, y int) bool {
	if s.x[0] == s.x[1] {
		// vertical splitter on horizontal direction
		return s.x[0] == x && (s.y[0] <= y && y <= s.y[1])
	} else {
		// horizontal splitter on vertical direction
		return s.y[0] == y && (s.x[0] <= x && x <= s.x[1])
	}
}

// clipRegion implements tcell.Screen and only allows setting content within
// a defined region
type clipRegion struct {
	tcell.Screen
	x      int
	y      int
	width  int
	height int
	style  tcell.Style
}

// newClipRegion creates a new clipped screen with the given rectangular coordinates
func newClipRegion(screen tcell.Screen, x, y, width, height int) *clipRegion {
	return &clipRegion{
		Screen: screen,
		x:      x,
		y:      y,
		width:  width,
		height: height,
		style:  tcell.StyleDefault,
	}
}

// InRect returns true if the given coordinates are within this clipped region
func (cr *clipRegion) InRect(x, y int) bool {
	return !(x < cr.x || y < cr.y || x >= cr.x+cr.width || y >= cr.y+cr.height)
}

// Fill implements tcell.Screen.Fill
func (cr *clipRegion) Fill(ch rune, style tcell.Style) {
	for x := cr.x; x < cr.width; x++ {
		for y := cr.y; y < cr.height; y++ {
			cr.SetContent(x, y, ch, nil, style)
		}
	}
}

// SetCell is an older API, and will be removed.  Please use
// SetContent instead; SetCell is implemented in terms of SetContent.
func (cr *clipRegion) SetCell(x int, y int, style tcell.Style, ch ...rune) {
	if len(ch) > 0 {
		cr.SetContent(x, y, ch[0], ch[1:], style)
	} else {
		cr.SetContent(x, y, ' ', nil, style)
	}
}

// SetContent sets the contents of the given cell location.  If
// the coordinates are out of range, then the operation is ignored.
//
// The first rune is the primary non-zero width rune.  The array
// that follows is a possible list of combining characters to append,
// and will usually be nil (no combining characters.)
//
// The results are not displayed until Show() or Sync() is called.
//
// Note that wide (East Asian full width) runes occupy two cells,
// and attempts to place character at next cell to the right will have
// undefined effects.  Wide runes that are printed in the
// last column will be replaced with a single width space on output.
func (cr *clipRegion) SetContent(x int, y int, mainc rune, combc []rune, style tcell.Style) {
	if cr.InRect(x, y) {
		cr.Screen.SetContent(x, y, mainc, combc, style)
	}
}

// ShowCursor is used to display the cursor at a given location.
// If the coordinates -1, -1 are given or are otherwise outside the
// dimensions of the screen, the cursor will be hidden.
func (cr *clipRegion) ShowCursor(x int, y int) {
	if cr.InRect(x, y) {
		cr.Screen.ShowCursor(x, y)
	}
}

// Clear clears the clipped region
func (cr *clipRegion) Clear() {
	cr.Fill(' ', cr.style)
}
