package cui

import "sync"

// ContextMenu is a menu that appears upon user interaction, such as
// right-clicking or pressing Alt+Enter.
type ContextMenu struct {
	parent   Widget
	item     int
	open     bool
	drag     bool
	list     *List
	x, y     int
	selected func(int, string, rune)

	mu sync.RWMutex
}

// NewContextMenu returns a new context menu.
func NewContextMenu(parent Widget) *ContextMenu {
	return &ContextMenu{
		parent: parent,
	}
}

func (c *ContextMenu) initializeList() {
	if c.list != nil {
		return
	}

	c.list = NewList()
	c.list.ShowSecondaryText(false)
	c.list.SetHover(true)
	c.list.SetWrapAround(true)
	c.list.ShowFocus(false)
	c.list.SetBorder(true)
	c.list.SetPadding(
		Styles.ContextMenuPaddingTop,
		Styles.ContextMenuPaddingBottom,
		Styles.ContextMenuPaddingLeft,
		Styles.ContextMenuPaddingRight)
}

// ContextMenuList returns the underlying List of the context menu.
func (c *ContextMenu) ContextMenuList() *List {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.initializeList()

	return c.list
}

// AddContextItem adds an item to the context menu. Adding an item with no text
// or shortcut will add a divider.
func (c *ContextMenu) AddContextItem(text string, shortcut rune, selected func(index int)) *ContextMenu {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.initializeList()

	item := NewListItem(text)
	item.SetShortcut(shortcut)
	item.SetSelectedFunc(c.wrap(selected))

	c.list.AddItem(item)
	if text == "" && shortcut == 0 {
		c.list.mu.Lock()
		index := len(c.list.items) - 1
		c.list.items[index].disabled = true
		c.list.mu.Unlock()
	}

	return c
}

func (c *ContextMenu) wrap(f func(index int)) func() {
	return func() {
		f(c.item)
	}
}

// ClearContextMenu removes all items from the context menu.
func (c *ContextMenu) ClearContextMenu() *ContextMenu {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.initializeList()

	c.list.Clear()
	return c
}

// SetContextSelectedFunc sets the function which is called when the user
// selects a context menu item. The function receives the item's index in the
// menu (starting with 0), its text and its shortcut rune. OnClick must
// be called before the context menu is shown.
func (c *ContextMenu) SetContextSelectedFunc(handler func(index int, text string, shortcut rune)) *ContextMenu {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.selected = handler
	return c
}

// ShowContextMenu shows the context menu. Provide -1 for both to position on
// the selected item, or specify a 	position.
func (c *ContextMenu) ShowContextMenu(item int, x int, y int, setFocus func(Widget)) *ContextMenu {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.show(item, x, y, setFocus)
	return c
}

// HideContextMenu hides the context menu.
func (c *ContextMenu) HideContextMenu(setFocus func(Widget)) *ContextMenu {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.hide(setFocus)
	return c
}

// ContextMenuVisible returns whether the context menu is visible.
func (c *ContextMenu) ContextMenuVisible() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.open
}

func (c *ContextMenu) show(item int, x int, y int, setFocus func(Widget)) {
	c.initializeList()

	if len(c.list.items) == 0 {
		return
	}

	c.open = true
	c.item = item
	c.x, c.y = x, y

	c.list.mu.Lock()
	for i, item := range c.list.items {
		if !item.disabled {
			c.list.currentItem = i
			break
		}
	}
	c.list.mu.Unlock()

	c.list.SetSelectedFunc(func(index int, item *ListItem) {
		c.mu.Lock()

		// A context item was selected. Close the menu.
		c.hide(setFocus)

		if c.selected != nil {
			c.mu.Unlock()
			c.selected(index, string(item.mainText), item.shortcut)
		} else {
			c.mu.Unlock()
		}
	})
	c.list.SetDoneFunc(func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		c.hide(setFocus)
	})

	setFocus(c.list)
}

func (c *ContextMenu) hide(setFocus func(Widget)) {
	c.initializeList()

	c.open = false

	if c.list.HasFocus() {
		setFocus(c.parent)
	}
}
