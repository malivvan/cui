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
	Widget

	Size int
}

type Layout struct {
	box *Box

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

	mu sync.RWMutex
}

func NewLayout() *Layout {
	return &Layout{
		box:                   NewBox(),
		focusedSplitterNumber: -1,
		splitterStyle:         tcell.StyleDefault.Foreground(Styles.BorderColor),
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (l *Layout) set(setter func(l *Layout)) *Layout {
	l.mu.Lock()
	setter(l)
	l.mu.Unlock()
	return l
}

func (l *Layout) get(getter func(l *Layout)) {
	l.mu.RLock()
	getter(l)
	l.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Layout.
func (l *Layout) GetTitle() string {
	return l.box.GetTitle()
}

// SetTitle sets the title of this Layout.
func (l *Layout) SetTitle(title string) *Layout {
	l.box.SetTitle(title)
	return l
}

// GetTitleAlign returns the title alignment of this Layout.
func (l *Layout) GetTitleAlign() int {
	return l.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Layout.
func (l *Layout) SetTitleAlign(align int) *Layout {
	l.box.SetTitleAlign(align)
	return l
}

// GetBorder returns whether this Layout has a border.
func (l *Layout) GetBorder() bool {
	return l.box.GetBorder()
}

// SetBorder sets whether this Layout has a border.
func (l *Layout) SetBorder(show bool) *Layout {
	l.box.SetBorder(show)
	return l
}

// GetBorderColor returns the border color of this Layout.
func (l *Layout) GetBorderColor() tcell.Color {
	return l.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Layout.
func (l *Layout) SetBorderColor(color tcell.Color) *Layout {
	l.box.SetBorderColor(color)
	return l
}

// GetBorderAttributes returns the border attributes of this Layout.
func (l *Layout) GetBorderAttributes() tcell.AttrMask {
	return l.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Layout.
func (l *Layout) SetBorderAttributes(attr tcell.AttrMask) *Layout {
	l.box.SetBorderAttributes(attr)
	return l
}

// GetBorderColorFocused returns the border color of this Layout when focusel.
func (l *Layout) GetBorderColorFocused() tcell.Color {
	return l.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Layout when focusel.
func (l *Layout) SetBorderColorFocused(color tcell.Color) *Layout {
	l.box.SetBorderColorFocused(color)
	return l
}

// GetTitleColor returns the title color of this Layout.
func (l *Layout) GetTitleColor() tcell.Color {
	return l.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Layout.
func (l *Layout) SetTitleColor(color tcell.Color) *Layout {
	l.box.SetTitleColor(color)
	return l
}

// GetDrawFunc returns the custom draw function of this Layout.
func (l *Layout) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return l.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Layout.
func (l *Layout) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Layout {
	l.box.SetDrawFunc(handler)
	return l
}

// ShowFocus sets whether this Layout should show a focus indicator when focusel.
func (l *Layout) ShowFocus(showFocus bool) *Layout {
	l.box.ShowFocus(showFocus)
	return l
}

// GetMouseCapture returns the mouse capture function of this Layout.
func (l *Layout) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return l.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Layout.
func (l *Layout) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Layout {
	l.box.SetMouseCapture(capture)
	return l
}

// GetBackgroundColor returns the background color of this Layout.
func (l *Layout) GetBackgroundColor() tcell.Color {
	return l.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Layout.
func (l *Layout) SetBackgroundColor(color tcell.Color) *Layout {
	l.box.SetBackgroundColor(color)
	return l
}

// GetBackgroundTransparent returns whether the background of this Layout is transparent.
func (l *Layout) GetBackgroundTransparent() bool {
	return l.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Layout is transparent.
func (l *Layout) SetBackgroundTransparent(transparent bool) *Layout {
	l.box.SetBackgroundTransparent(transparent)
	return l
}

// GetInputCapture returns the input capture function of this Layout.
func (l *Layout) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return l.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Layout.
func (l *Layout) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Layout {
	l.box.SetInputCapture(capture)
	return l
}

// GetPadding returns the padding of this Layout.
func (l *Layout) GetPadding() (top, bottom, left, right int) {
	return l.box.GetPadding()
}

// SetPadding sets the padding of this Layout.
func (l *Layout) SetPadding(top, bottom, left, right int) *Layout {
	l.box.SetPadding(top, bottom, left, right)
	return l
}

// InRect returns whether the given screen coordinates are within this Layout.
func (l *Layout) InRect(x, y int) bool {
	return l.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Layout.
func (l *Layout) GetInnerRect() (x, y, width, height int) {
	return l.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Layout is preservel.
func (l *Layout) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return l.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Layout is preservel.
func (l *Layout) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return l.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Layout.
func (l *Layout) GetRect() (x, y, width, height int) {
	return l.box.GetRect()
}

// SetRect sets the rectangle occupied by this Layout.
func (l *Layout) SetRect(x, y, width, height int) {
	l.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Layout is visible.
func (l *Layout) GetVisible() bool {
	return l.box.GetVisible()
}

// SetVisible sets whether this Layout is visible.
func (l *Layout) SetVisible(visible bool) {
	l.box.SetVisible(visible)
}

func (l *Layout) Focus(delegate func(p Widget)) {
	l.mu.Lock()

	if l.focusedSplitterNumber == -1 {
		for _, item := range l.items {
			if item.Widget != nil {
				l.mu.Unlock()
				delegate(item.Widget)
				return
			}
		}
	}
	l.mu.Unlock()
}

func (l *Layout) HasFocus() (hasFocus bool) {
	l.get(func(l *Layout) {
		for _, item := range l.items {
			if item.Widget != nil && item.Widget.HasFocus() {
				hasFocus = true
				return
			}
		}
	})
	return
}

func (l *Layout) GetFocusable() (focusable Focusable) {
	l.get(func(l *Layout) {
		for _, item := range l.items {
			if item.Widget != nil && item.Widget.HasFocus() {
				focusable = item.Widget
				return
			}
		}
	})
	return
}

func (l *Layout) Blur() {
	l.set(func(l *Layout) {
		l.focusedSplitterNumber = -1
		for _, item := range l.items {
			if item.Widget != nil && item.Widget.HasFocus() {
				item.Widget.Blur()
			}
		}
	})
}

func (l *Layout) SetDirection(d int) *Layout {
	return l.set(func(l *Layout) { l.direction = d })
}

func (l *Layout) GetDirection() (direction int) {
	l.get(func(l *Layout) { direction = l.direction })
	return
}

// SetSplitter sets the flag indicating whether the layout should render a
// splitters between primitives
func (l *Layout) SetSplitter(show bool) *Layout {
	return l.set(func(l *Layout) { l.splitterFlag = show })
}

// SetSplitterColor sets the layout's splitter color.
func (l *Layout) SetSplitterColor(color tcell.Color) *Layout {
	return l.set(func(l *Layout) { l.splitterStyle = l.splitterStyle.Foreground(color) })
}

// SetSplitterAttributes sets the splitter's style attributes. You can combine
// different attributes using bitmask operations:
//
//	layout.SetSplitterAttributes(tcell.AttrUnderline | tcell.AttrBold)
func (l *Layout) SetSplitterAttributes(attr tcell.AttrMask) *Layout {
	return l.set(func(l *Layout) { l.splitterStyle = l.splitterStyle.Attributes(attr) })
}

// GetSplitterAttributes returns the splitter's style attributes.
func (l *Layout) GetSplitterAttributes() (attr tcell.AttrMask) {
	l.get(func(l *Layout) { _, _, attr = l.splitterStyle.Decompose() })
	return
}

// GetSplitterColor returns the layout's splitter color.
func (l *Layout) GetSplitterColor() (color tcell.Color) {
	l.get(func(l *Layout) { color, _, _ = l.splitterStyle.Decompose() })
	return
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
	_, _, width, height := l.box.GetInnerRect()
	switch l.direction {
	case HorizontalLayout:
		return width
	case VerticalLayout:
		return height
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
	if !l.GetVisible() {
		return
	}

	l.box.Draw(screen)

	l.mu.Lock()
	defer l.mu.Unlock()

	x, y, width, height := l.GetInnerRect()

	// Don't draw anything if there is no space.
	if width <= 0 || height <= 0 {
		return
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
				item.Widget.SetRect(x, y, autoSize, height)
				item.Widget.Draw(newClipRegion(screen, x, y, autoSize, height))
				x += autoSize
			} else {
				item.Widget.SetRect(x, y, item.Size, height)
				item.Widget.Draw(newClipRegion(screen, x, y, item.Size, height))
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
				item.Widget.SetRect(x, y, width, autoSize)
				item.Widget.Draw(newClipRegion(screen, x, y, width, autoSize))
				y += autoSize
			} else {
				item.Widget.SetRect(x, y, width, item.Size)
				item.Widget.Draw(newClipRegion(screen, x, y, width, item.Size))
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

func (l *Layout) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return l.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		for _, item := range l.items {
			if item.Widget.HasFocus() {
				if handler := item.Widget.InputHandler(); handler != nil {
					handler(event, setFocus)
					return
				}
			}
		}
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (l *Layout) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return l.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if !l.InRect(event.Position()) {
			return false, nil
		}

		// Pass mouse events along to the first child item that takes it.
		for _, item := range l.items {
			if item.Widget != nil {
				consumed, capture = item.Widget.MouseHandler()(action, event, setFocus)
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
						if item.Widget != nil {
							item.Widget.Blur()
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

func (l *Layout) AddItem(p Widget, size int) *Layout {
	return l.set(func(l *Layout) {
		l.items = append(l.items, &LayoutItem{
			Widget: p,
			Size:   size,
		})
		l.draggedSplitter = nil
		l.rebuildSplitters()
	})
}

func (l *Layout) RemoveItem(i int) *Layout {
	return l.set(func(l *Layout) {
		if i < 0 || i >= len(l.items) {
			return
		}
		l.draggedSplitter = nil
		l.items = append(l.items[:i], l.items[i+1:]...)
		l.rebuildSplitters()
	})
}

func (l *Layout) GetItem(i int) (item *LayoutItem) {
	l.get(func(l *Layout) {
		if !(i < 0 || i >= len(l.items)) {
			item = l.items[i]
		}
	})
	return
}

func (l *Layout) CountItems() (count int) {
	l.get(func(l *Layout) { count = len(l.items) })
	return
}

func (l *Layout) ClearItems() *Layout {
	return l.set(func(l *Layout) {
		l.items = nil
		l.splitters = nil
	})
}

func (l *Layout) rebuildSplitters() {
	l.splitters = nil

	x, y, width, height := l.box.GetInnerRect()

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
func (cr *clipRegion) SetContent(x int, y int, primary rune, combining []rune, style tcell.Style) {
	if cr.InRect(x, y) {
		cr.Screen.SetContent(x, y, primary, combining, style)
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
