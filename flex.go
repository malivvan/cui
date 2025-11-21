package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// Configuration values.
const (
	FlexRow = iota
	FlexColumn
)

// flexItem holds layout options for one item.
type flexItem struct {
	Item       Widget // The item to be positioned. May be nil for an empty item.
	FixedSize  int    // The item's fixed size which may not be changed, 0 if it has no fixed size.
	Proportion int    // The item's proportion.
	Focus      bool   // Whether this item attracts the layout's focus.
}

// Flex is a basic implementation of the Flexbox layout. The contained
// primitives are arranged horizontally or vertically. The way they are
// distributed along that dimension depends on their layout settings, which is
// either a fixed length or a proportional length. See AddItem() for details.
type Flex struct {
	box *Box

	// The items to be positioned.
	items []*flexItem

	// FlexRow or FlexColumn.
	direction int

	// If set to true, Flex will use the entire screen as its available space
	// instead its box dimensions.
	fullScreen bool

	mu sync.RWMutex
}

// NewFlex returns a new flexbox layout container with no primitives and its
// direction set to FlexColumn. To add primitives to this layout, see AddItem().
// To change the direction, see SetDirection().
//
// Note that Flex will have a transparent background by default so that any nil
// flex items will show primitives behind the Flex.
// To disable this transparency:
//
//	flex.SetBackgroundTransparent(false)
func NewFlex() *Flex {
	f := &Flex{
		box:       NewBox(),
		direction: FlexColumn,
	}
	f.SetBackgroundTransparent(true)
	f.box.focus = f
	return f
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (f *Flex) set(setter func(f *Flex)) *Flex {
	f.mu.Lock()
	setter(f)
	f.mu.Unlock()
	return f
}

func (f *Flex) get(getter func(f *Flex)) {
	f.mu.RLock()
	getter(f)
	f.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Flex.
func (f *Flex) GetTitle() string {
	return f.box.GetTitle()
}

// SetTitle sets the title of this Flex.
func (f *Flex) SetTitle(title string) *Flex {
	f.box.SetTitle(title)
	return f
}

// GetTitleAlign returns the title alignment of this Flex.
func (f *Flex) GetTitleAlign() int {
	return f.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Flex.
func (f *Flex) SetTitleAlign(align int) *Flex {
	f.box.SetTitleAlign(align)
	return f
}

// GetBorder returns whether this Flex has a border.
func (f *Flex) GetBorder() bool {
	return f.box.GetBorder()
}

// SetBorder sets whether this Flex has a border.
func (f *Flex) SetBorder(show bool) *Flex {
	f.box.SetBorder(show)
	return f
}

// GetBorderColor returns the border color of this Flex.
func (f *Flex) GetBorderColor() tcell.Color {
	return f.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Flex.
func (f *Flex) SetBorderColor(color tcell.Color) *Flex {
	f.box.SetBorderColor(color)
	return f
}

// GetBorderAttributes returns the border attributes of this Flex.
func (f *Flex) GetBorderAttributes() tcell.AttrMask {
	return f.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Flex.
func (f *Flex) SetBorderAttributes(attr tcell.AttrMask) *Flex {
	f.box.SetBorderAttributes(attr)
	return f
}

// GetBorderColorFocused returns the border color of this Flex when focusef.
func (f *Flex) GetBorderColorFocused() tcell.Color {
	return f.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Flex when focusef.
func (f *Flex) SetBorderColorFocused(color tcell.Color) *Flex {
	f.box.SetBorderColorFocused(color)
	return f
}

// GetTitleColor returns the title color of this Flex.
func (f *Flex) GetTitleColor() tcell.Color {
	return f.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Flex.
func (f *Flex) SetTitleColor(color tcell.Color) *Flex {
	f.box.SetTitleColor(color)
	return f
}

// GetDrawFunc returns the custom draw function of this Flex.
func (f *Flex) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return f.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Flex.
func (f *Flex) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Flex {
	f.box.SetDrawFunc(handler)
	return f
}

// ShowFocus sets whether this Flex should show a focus indicator when focusef.
func (f *Flex) ShowFocus(showFocus bool) *Flex {
	f.box.ShowFocus(showFocus)
	return f
}

// GetMouseCapture returns the mouse capture function of this Flex.
func (f *Flex) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return f.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Flex.
func (f *Flex) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Flex {
	f.box.SetMouseCapture(capture)
	return f
}

// GetBackgroundColor returns the background color of this Flex.
func (f *Flex) GetBackgroundColor() tcell.Color {
	return f.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Flex.
func (f *Flex) SetBackgroundColor(color tcell.Color) *Flex {
	f.box.SetBackgroundColor(color)
	return f
}

// GetBackgroundTransparent returns whether the background of this Flex is transparent.
func (f *Flex) GetBackgroundTransparent() bool {
	return f.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Flex is transparent.
func (f *Flex) SetBackgroundTransparent(transparent bool) *Flex {
	f.box.SetBackgroundTransparent(transparent)
	return f
}

// GetInputCapture returns the input capture function of this Flex.
func (f *Flex) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return f.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Flex.
func (f *Flex) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Flex {
	f.box.SetInputCapture(capture)
	return f
}

// GetPadding returns the padding of this Flex.
func (f *Flex) GetPadding() (top, bottom, left, right int) {
	return f.box.GetPadding()
}

// SetPadding sets the padding of this Flex.
func (f *Flex) SetPadding(top, bottom, left, right int) *Flex {
	f.box.SetPadding(top, bottom, left, right)
	return f
}

// InRect returns whether the given screen coordinates are within this Flex.
func (f *Flex) InRect(x, y int) bool {
	return f.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Flex.
func (f *Flex) GetInnerRect() (x, y, width, height int) {
	return f.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Flex is preservef.
func (f *Flex) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return f.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Flex is preservef.
func (f *Flex) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return f.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Flex.
func (f *Flex) GetRect() (x, y, width, height int) {
	return f.box.GetRect()
}

// SetRect sets the rectangle occupied by this Flex.
func (f *Flex) SetRect(x, y, width, height int) {
	f.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Flex is visible.
func (f *Flex) GetVisible() bool {
	return f.box.GetVisible()
}

// SetVisible sets whether this Flex is visible.
func (f *Flex) SetVisible(visible bool) {
	f.box.SetVisible(visible)
}

// Focus is called when this primitive receives focus.
func (f *Flex) Focus(delegate func(p Widget)) {
	f.mu.Lock()
	for _, item := range f.items {
		if item.Item != nil && item.Focus {
			f.mu.Unlock()
			delegate(item.Item)
			return
		}
	}
	f.mu.Unlock()
}

// HasFocus returns whether this primitive has focus.
func (f *Flex) HasFocus() (hasFocus bool) {
	f.get(func(f *Flex) {
		for _, item := range f.items {
			if item.Item != nil && item.Item.GetFocusable().HasFocus() {
				hasFocus = true
			}
		}
	})
	return
}

// GetFocusable returns the focusable primitive of this Flex.
func (f *Flex) GetFocusable() Focusable {
	return f.box.GetFocusable()
}

// Blur is called when this Flex loses focus.
func (f *Flex) Blur() {
	f.box.Blur()
}

/////////////////////////////////////// <API> ///////////////////////////////////////

// GetDirection returns the direction in which the contained primitives are
// distributed. This can be either FlexColumn (default) or FlexRow.
func (f *Flex) GetDirection() (direction int) {
	f.get(func(f *Flex) { direction = f.direction })
	return
}

// SetDirection sets the direction in which the contained primitives are
// distributed. This can be either FlexColumn (default) or FlexRow.
func (f *Flex) SetDirection(direction int) *Flex {
	return f.set(func(f *Flex) { f.direction = direction })
}

// IsFullScreen returns whether the flex layout is using the entire screen space
// instead of whatever size it is currently assigned to.
func (f *Flex) IsFullScreen() (fullScreen bool) {
	f.get(func(f *Flex) { fullScreen = f.fullScreen })
	return
}

// SetFullScreen sets the flag which, when true, causes the flex layout to use
// the entire screen space instead of whatever size it is currently assigned to.
func (f *Flex) SetFullScreen(fullScreen bool) *Flex {
	return f.set(func(f *Flex) { f.fullScreen = fullScreen })
}

// Clear removes all items from the container.
func (f *Flex) Clear() *Flex {
	return f.set(func(f *Flex) { f.items = nil })
}

// AddItem adds a new item to the container. The "fixedSize" argument is a width
// or height that may not be changed by the layout algorithm. A value of 0 means
// that its size is flexible and may be changed. The "proportion" argument
// defines the relative size of the item compared to other flexible-size items.
// For example, items with a proportion of 2 will be twice as large as items
// with a proportion of 1. The proportion must be at least 1 if fixedSize == 0
// (ignored otherwise).
//
// If "focus" is set to true, the item will receive focus when the Flex
// primitive receives focus. If multiple items have the "focus" flag set to
// true, the first one will receive focus.
//
// A nil value for the primitive represents empty space.
func (f *Flex) AddItem(item Widget, fixedSize, proportion int, focus bool) *Flex {
	if item == nil {
		item = NewBox()
		item.SetVisible(false)
	}
	return f.set(func(f *Flex) {
		f.items = append(f.items, &flexItem{Item: item, FixedSize: fixedSize, Proportion: proportion, Focus: focus})
	})
}

// AddItemAtIndex adds an item to the flex at a given index.
// For more information see AddItem.
func (f *Flex) AddItemAtIndex(index int, item Widget, fixedSize, proportion int, focus bool) *Flex {
	newItem := &flexItem{Item: item, FixedSize: fixedSize, Proportion: proportion, Focus: focus}
	return f.set(func(f *Flex) {
		if index == 0 {
			f.items = append([]*flexItem{newItem}, f.items...)
		} else {
			f.items = append(f.items[:index], append([]*flexItem{newItem}, f.items[index:]...)...)
		}
	})
}

// RemoveItem removes all items for the given primitive from the container,
// keeping the order of the remaining items intact.
func (f *Flex) RemoveItem(p Widget) *Flex {
	return f.set(func(f *Flex) {
		for index := len(f.items) - 1; index >= 0; index-- {
			if f.items[index].Item == p {
				f.items = append(f.items[:index], f.items[index+1:]...)
			}
		}
	})
}

// ResizeItem sets a new size for the item(s) with the given primitive. If there
// are multiple Flex items with the same primitive, they will all receive the
// same size. For details regarding the size parameters, see AddItem().
func (f *Flex) ResizeItem(p Widget, fixedSize, proportion int) *Flex {
	return f.set(func(f *Flex) {
		for _, item := range f.items {
			if item.Item == p {
				item.FixedSize = fixedSize
				item.Proportion = proportion
			}
		}
	})
}

// Draw draws this primitive onto the screen.
func (f *Flex) Draw(screen tcell.Screen) {
	if !f.GetVisible() {
		return
	}

	f.box.Draw(screen)

	f.mu.Lock()
	defer f.mu.Unlock()

	// Calculate size and position of the items.

	// Do we use the entire screen?
	if f.fullScreen {
		width, height := screen.Size()
		f.SetRect(0, 0, width, height)
	}

	// How much space can we distribute?
	x, y, width, height := f.GetInnerRect()
	var proportionSum int
	distSize := width
	if f.direction == FlexRow {
		distSize = height
	}
	for _, item := range f.items {
		if item.FixedSize > 0 {
			distSize -= item.FixedSize
		} else {
			proportionSum += item.Proportion
		}
	}

	// Calculate positions and draw items.
	pos := x
	if f.direction == FlexRow {
		pos = y
	}
	for _, item := range f.items {
		size := item.FixedSize
		if size <= 0 {
			if proportionSum > 0 {
				size = distSize * item.Proportion / proportionSum
				distSize -= size
				proportionSum -= item.Proportion
			} else {
				size = 0
			}
		}
		if item.Item != nil {
			if f.direction == FlexColumn {
				item.Item.SetRect(pos, y, size, height)
			} else {
				item.Item.SetRect(x, pos, width, size)
			}
		}
		pos += size

		if item.Item != nil {
			if item.Item.GetFocusable().HasFocus() {
				defer item.Item.Draw(screen)
			} else {
				item.Item.Draw(screen)
			}
		}
	}
}

// InputHandler returns the input handler for this primitive.
func (f *Flex) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return f.box.InputHandler()
}

// MouseHandler returns the mouse handler for this primitive.
func (f *Flex) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return f.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if !f.InRect(event.Position()) {
			return false, nil
		}

		// Pass mouse events along to the first child item that takes it.
		for _, item := range f.items {
			if item.Item == nil {
				continue
			}

			consumed, capture = item.Item.MouseHandler()(action, event, setFocus)
			if consumed {
				return
			}
		}

		return
	})
}
