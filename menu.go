package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// MenuBar represents a menu bar widget.
type MenuBar struct {
	box           *Box
	menuItems     []*MenuItem
	subMenu       *SubMenu // sub menu if not nil will be drawn
	currentOption int
	mu            sync.RWMutex
}

func NewMenuBar() *MenuBar {
	return &MenuBar{
		box:       NewBox(),
		menuItems: make([]*MenuItem, 0),
	}
}

func (mb *MenuBar) Focus(delegate func(p Widget)) {
	//if menuBar.subMenu != nil {
	//	delegate(menuBar.subMenu)
	//} else {
	mb.box.Focus(delegate)
	mb.subMenu = nil
	//}
}

func (mb *MenuBar) AfterDraw() func(tcell.Screen) {
	return func(screen tcell.Screen) {
		if mb.subMenu != nil {
			mb.subMenu.Draw(screen)
		}
	}
}

func (mb *MenuBar) AddItem(item *MenuItem) *MenuBar {
	return mb.set(func(m *MenuBar) { m.menuItems = append(m.menuItems, item) })
}

func (mb *MenuBar) Draw(screen tcell.Screen) {
	if !mb.GetVisible() {
		return
	}

	mb.box.Draw(screen)

	mb.mu.Lock()
	defer mb.mu.Unlock()

	x, y, width, _ := mb.GetInnerRect()

	for i := 0; i < width; i += 1 {
		screen.SetContent(x+i, y, ' ', nil, tcell.StyleDefault.Background(mb.box.GetBackgroundColor()))
	}

	menuItemOffset := 1
	for _, mi := range mb.menuItems {
		itemLen := len([]rune(mi.title))
		mi.box.SetRect(menuItemOffset, y, itemLen, 1)
		mi.Draw(screen)
		menuItemOffset += itemLen + 1
	}
}

func (mb *MenuBar) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return mb.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		switch event.Key() {
		case tcell.KeyLeft:
			mb.currentOption--
			if mb.currentOption < 0 {
				mb.currentOption = -1
			}
		case tcell.KeyRight:
			mb.currentOption++
			if mb.currentOption >= len(mb.menuItems) {
				mb.currentOption = len(mb.menuItems) - 1
			}
		}
	})
}

func (mb *MenuBar) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return mb.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if mb.subMenu != nil {
			consumed, capture = mb.subMenu.MouseHandler()(action, event, setFocus)
			if consumed {
				//p.subMenu = nil
				return
			}
		}
		if !mb.InRect(event.Position()) {
			return false, nil
		}
		// Pass mouse events down.
		for _, item := range mb.menuItems {
			consumed, capture = item.box.MouseHandler()(action, event, setFocus)
			if consumed {
				mb.subMenu = NewSubMenu(mb, item.subItems)
				x, y, _, _ := item.box.GetRect()
				if mb.GetBorder() {
					x++
				}
				mb.subMenu.box.SetRect(x, y+1, 15, 10)
				return
			}
		}

		// ...handle mouse events not directed to the child primitive...
		return true, nil
	})
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (mb *MenuBar) set(setter func(m *MenuBar)) *MenuBar {
	mb.mu.Lock()
	setter(mb)
	mb.mu.Unlock()
	return mb
}

func (mb *MenuBar) get(getter func(m *MenuBar)) {
	mb.mu.RLock()
	getter(mb)
	mb.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this MenuBar.
func (mb *MenuBar) GetTitle() string {
	return mb.box.GetTitle()
}

// SetTitle sets the title of this MenuBar.
func (mb *MenuBar) SetTitle(title string) *MenuBar {
	mb.box.SetTitle(title)
	return mb
}

// GetTitleAlign returns the title alignment of this MenuBar.
func (mb *MenuBar) GetTitleAlign() int {
	return mb.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this MenuBar.
func (mb *MenuBar) SetTitleAlign(align int) *MenuBar {
	mb.box.SetTitleAlign(align)
	return mb
}

// GetBorder returns whether this MenuBar has a border.
func (mb *MenuBar) GetBorder() bool {
	return mb.box.GetBorder()
}

// SetBorder sets whether this MenuBar has a border.
func (mb *MenuBar) SetBorder(show bool) *MenuBar {
	mb.box.SetBorder(show)
	return mb
}

// GetBorderColor returns the border color of this MenuBar.
func (mb *MenuBar) GetBorderColor() tcell.Color {
	return mb.box.GetBorderColor()
}

// SetBorderColor sets the border color of this MenuBar.
func (mb *MenuBar) SetBorderColor(color tcell.Color) *MenuBar {
	mb.box.SetBorderColor(color)
	return mb
}

// GetBorderAttributes returns the border attributes of this MenuBar.
func (mb *MenuBar) GetBorderAttributes() tcell.AttrMask {
	return mb.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this MenuBar.
func (mb *MenuBar) SetBorderAttributes(attr tcell.AttrMask) *MenuBar {
	mb.box.SetBorderAttributes(attr)
	return mb
}

// GetBorderColorFocused returns the border color of this MenuBar when focused.
func (mb *MenuBar) GetBorderColorFocused() tcell.Color {
	return mb.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this MenuBar when focused.
func (mb *MenuBar) SetBorderColorFocused(color tcell.Color) *MenuBar {
	mb.box.SetBorderColorFocused(color)
	return mb
}

// GetTitleColor returns the title color of this MenuBar.
func (mb *MenuBar) GetTitleColor() tcell.Color {
	return mb.box.GetTitleColor()
}

// SetTitleColor sets the title color of this MenuBar.
func (mb *MenuBar) SetTitleColor(color tcell.Color) *MenuBar {
	mb.box.SetTitleColor(color)
	return mb
}

// GetDrawFunc returns the custom draw function of this MenuBar.
func (mb *MenuBar) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return mb.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this MenuBar.
func (mb *MenuBar) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *MenuBar {
	mb.box.SetDrawFunc(handler)
	return mb
}

// ShowFocus sets whether this MenuBar should show a focus indicator when focused.
func (mb *MenuBar) ShowFocus(showFocus bool) *MenuBar {
	mb.box.ShowFocus(showFocus)
	return mb
}

// GetMouseCapture returns the mouse capture function of this MenuBar.
func (mb *MenuBar) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return mb.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this MenuBar.
func (mb *MenuBar) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *MenuBar {
	mb.box.SetMouseCapture(capture)
	return mb
}

// GetBackgroundColor returns the background color of this MenuBar.
func (mb *MenuBar) GetBackgroundColor() tcell.Color {
	return mb.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this MenuBar.
func (mb *MenuBar) SetBackgroundColor(color tcell.Color) *MenuBar {
	mb.box.SetBackgroundColor(color)
	return mb
}

// GetBackgroundTransparent returns whether the background of this MenuBar is transparent.
func (mb *MenuBar) GetBackgroundTransparent() bool {
	return mb.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this MenuBar is transparent.
func (mb *MenuBar) SetBackgroundTransparent(transparent bool) *MenuBar {
	mb.box.SetBackgroundTransparent(transparent)
	return mb
}

// GetInputCapture returns the input capture function of this MenuBar.
func (mb *MenuBar) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return mb.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this MenuBar.
func (mb *MenuBar) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *MenuBar {
	mb.box.SetInputCapture(capture)
	return mb
}

// GetPadding returns the padding of this MenuBar.
func (mb *MenuBar) GetPadding() (top, bottom, left, right int) {
	return mb.box.GetPadding()
}

// SetPadding sets the padding of this MenuBar.
func (mb *MenuBar) SetPadding(top, bottom, left, right int) *MenuBar {
	mb.box.SetPadding(top, bottom, left, right)
	return mb
}

// InRect returns whether the given screen coordinates are within this MenuBar.
func (mb *MenuBar) InRect(x, y int) bool {
	return mb.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this MenuBar.
func (mb *MenuBar) GetInnerRect() (x, y, width, height int) {
	return mb.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the MenuBar is preserved.
func (mb *MenuBar) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return mb.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the MenuBar is preserved.
func (mb *MenuBar) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return mb.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this MenuBar.
func (mb *MenuBar) GetRect() (x, y, width, height int) {
	return mb.box.GetRect()
}

// SetRect sets the rectangle occupied by this MenuBar.
func (mb *MenuBar) SetRect(x, y, width, height int) {
	mb.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this MenuBar is visible.
func (mb *MenuBar) GetVisible() bool {
	return mb.box.GetVisible()
}

// SetVisible sets whether this MenuBar is visible.
func (mb *MenuBar) SetVisible(visible bool) {
	mb.box.SetVisible(visible)
}

// HasFocus returns whether this MenuBar has focus.
func (mb *MenuBar) HasFocus() bool {
	return mb.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this MenuBar.
func (mb *MenuBar) GetFocusable() Focusable {
	return mb.box.GetFocusable()
}

// Blur is called when this MenuBar loses focus.
func (mb *MenuBar) Blur() {
	mb.box.Blur()
}

type SubMenu struct {
	box           *Box
	items         []*MenuItem
	parent        *MenuBar
	childMenu     *SubMenu
	currentSelect int
	mu            sync.RWMutex
}

func (sm *SubMenu) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return sm.box.InputHandler()
}

func (mi *MenuItem) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return mi.box.InputHandler()
}

func NewSubMenu(parent *MenuBar, items []*MenuItem) *SubMenu {
	sm := &SubMenu{
		box:           NewBox(),
		items:         items,
		parent:        parent,
		currentSelect: -1,
	}
	sm.box.SetBorder(true)
	return sm
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (sm *SubMenu) set(setter func(sm *SubMenu)) *SubMenu {
	sm.mu.Lock()
	setter(sm)
	sm.mu.Unlock()
	return sm
}

func (sm *SubMenu) get(getter func(sm *SubMenu)) {
	sm.mu.RLock()
	getter(sm)
	sm.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this SubMenu.
func (sm *SubMenu) GetTitle() string {
	return sm.box.GetTitle()
}

// SetTitle sets the title of this SubMenu.
func (sm *SubMenu) SetTitle(title string) *SubMenu {
	sm.box.SetTitle(title)
	return sm
}

// GetTitleAlign returns the title alignment of this SubMenu.
func (sm *SubMenu) GetTitleAlign() int {
	return sm.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this SubMenu.
func (sm *SubMenu) SetTitleAlign(align int) *SubMenu {
	sm.box.SetTitleAlign(align)
	return sm
}

// GetBorder returns whether this SubMenu has a border.
func (sm *SubMenu) GetBorder() bool {
	return sm.box.GetBorder()
}

// SetBorder sets whether this SubMenu has a border.
func (sm *SubMenu) SetBorder(show bool) *SubMenu {
	sm.box.SetBorder(show)
	return sm
}

// GetBorderColor returns the border color of this SubMenu.
func (sm *SubMenu) GetBorderColor() tcell.Color {
	return sm.box.GetBorderColor()
}

// SetBorderColor sets the border color of this SubMenu.
func (sm *SubMenu) SetBorderColor(color tcell.Color) *SubMenu {
	sm.box.SetBorderColor(color)
	return sm
}

// GetBorderAttributes returns the border attributes of this SubMenu.
func (sm *SubMenu) GetBorderAttributes() tcell.AttrMask {
	return sm.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this SubMenu.
func (sm *SubMenu) SetBorderAttributes(attr tcell.AttrMask) *SubMenu {
	sm.box.SetBorderAttributes(attr)
	return sm
}

// GetBorderColorFocused returns the border color of this SubMenu when focused.
func (sm *SubMenu) GetBorderColorFocused() tcell.Color {
	return sm.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this SubMenu when focused.
func (sm *SubMenu) SetBorderColorFocused(color tcell.Color) *SubMenu {
	sm.box.SetBorderColorFocused(color)
	return sm
}

// GetTitleColor returns the title color of this SubMenu.
func (sm *SubMenu) GetTitleColor() tcell.Color {
	return sm.box.GetTitleColor()
}

// SetTitleColor sets the title color of this SubMenu.
func (sm *SubMenu) SetTitleColor(color tcell.Color) *SubMenu {
	sm.box.SetTitleColor(color)
	return sm
}

// GetDrawFunc returns the custom draw function of this SubMenu.
func (sm *SubMenu) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return sm.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this SubMenu.
func (sm *SubMenu) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *SubMenu {
	sm.box.SetDrawFunc(handler)
	return sm
}

// ShowFocus sets whether this SubMenu should show a focus indicator when focused.
func (sm *SubMenu) ShowFocus(showFocus bool) *SubMenu {
	sm.box.ShowFocus(showFocus)
	return sm
}

// GetMouseCapture returns the mouse capture function of this SubMenu.
func (sm *SubMenu) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return sm.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this SubMenu.
func (sm *SubMenu) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *SubMenu {
	sm.box.SetMouseCapture(capture)
	return sm
}

// GetBackgroundColor returns the background color of this SubMenu.
func (sm *SubMenu) GetBackgroundColor() tcell.Color {
	return sm.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this SubMenu.
func (sm *SubMenu) SetBackgroundColor(color tcell.Color) *SubMenu {
	sm.box.SetBackgroundColor(color)
	return sm
}

// GetBackgroundTransparent returns whether the background of this SubMenu is transparent.
func (sm *SubMenu) GetBackgroundTransparent() bool {
	return sm.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this SubMenu is transparent.
func (sm *SubMenu) SetBackgroundTransparent(transparent bool) *SubMenu {
	sm.box.SetBackgroundTransparent(transparent)
	return sm
}

// GetInputCapture returns the input capture function of this SubMenu.
func (sm *SubMenu) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return sm.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this SubMenu.
func (sm *SubMenu) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *SubMenu {
	sm.box.SetInputCapture(capture)
	return sm
}

// GetPadding returns the padding of this SubMenu.
func (sm *SubMenu) GetPadding() (top, bottom, left, right int) {
	return sm.box.GetPadding()
}

// SetPadding sets the padding of this SubMenu.
func (sm *SubMenu) SetPadding(top, bottom, left, right int) *SubMenu {
	sm.box.SetPadding(top, bottom, left, right)
	return sm
}

// InRect returns whether the given screen coordinates are within this SubMenu.
func (sm *SubMenu) InRect(x, y int) bool {
	return sm.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this SubMenu.
func (sm *SubMenu) GetInnerRect() (x, y, width, height int) {
	return sm.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the SubMenu is preserved.
func (sm *SubMenu) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return sm.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the SubMenu is preserved.
func (sm *SubMenu) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return sm.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this SubMenu.
func (sm *SubMenu) GetRect() (x, y, width, height int) {
	return sm.box.GetRect()
}

// SetRect sets the rectangle occupied by this SubMenu.
func (sm *SubMenu) SetRect(x, y, width, height int) {
	sm.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this SubMenu is visible.
func (sm *SubMenu) GetVisible() bool {
	return sm.box.GetVisible()
}

// SetVisible sets whether this SubMenu is visible.
func (sm *SubMenu) SetVisible(visible bool) {
	sm.box.SetVisible(visible)
}

// Focus is called when this SubMenu receives focus.
func (sm *SubMenu) Focus(delegate func(p Widget)) {
	sm.box.Focus(delegate)
}

// HasFocus returns whether this SubMenu has focus.
func (sm *SubMenu) HasFocus() bool {
	return sm.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this SubMenu.
func (sm *SubMenu) GetFocusable() Focusable {
	return sm.box.GetFocusable()
}

// Blur is called when this SubMenu loses focus.
func (sm *SubMenu) Blur() {
	sm.box.Blur()
}

////////////////////////////////// <API> ////////////////////////////////////

func (sm *SubMenu) Draw(screen tcell.Screen) {
	anySubItems := false
	maxWidth := 0
	for _, item := range sm.items {
		if itemTitleLen := len(item.title); itemTitleLen > maxWidth {
			maxWidth = itemTitleLen
		}
		if len(item.subItems) > 0 {
			anySubItems = true
		}
	}

	rectX, rectY, _, _ := sm.box.GetRect()
	rectWid := maxWidth
	if anySubItems {
		rectWid += 1
	}
	rectHig := len(sm.items)
	// +2 - add space one space for each side of rect - to fit text inside
	sm.box.SetRect(rectX, rectY, rectWid+2, rectHig+2)

	if !sm.box.GetVisible() {
		return
	}

	sm.box.Draw(screen)

	sm.mu.Lock()
	defer sm.mu.Unlock()

	x, y, _, _ := sm.box.GetInnerRect()
	for i, item := range sm.items {
		if i == sm.currentSelect {
			Print(screen, []byte(item.title), x, y+i, 20, 0, tcell.ColorBlue)
			if len(item.subItems) > 0 {
				Print(screen, []byte(">"), x+maxWidth, y+i, 20, 0, tcell.ColorBlue)
			}
			continue
		}
		PrintSimple(screen, []byte(item.title), x, y+i)
		if len(item.subItems) > 0 {
			PrintSimple(screen, []byte(">"), x+maxWidth, y+i)
		}
	}
	if sm.childMenu != nil {
		sm.childMenu.Draw(screen)
	}
}

func (sm *SubMenu) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return sm.box.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if sm.childMenu != nil {
			consumed, capture = sm.childMenu.MouseHandler()(action, event, setFocus)

			if consumed {
				return
			}
		}
		rectX, rectY, rectW, _ := sm.box.GetInnerRect()
		if !sm.box.InRect(event.Position()) {
			// Close the menu if the user clicks outside the menu box
			if action == MouseLeftClick {
				sm.parent.subMenu = nil
			}
			return false, nil
		}
		_, y := event.Position()
		index := y - rectY

		sm.currentSelect = index
		consumed = true

		if action == MouseLeftClick {
			setFocus(sm.box)
			if index >= 0 && index < len(sm.items) {
				handler := sm.items[index].onClick
				if handler != nil {
					handler(sm.items[index])
				}
				if len(sm.items[index].subItems) > 0 {
					sm.childMenu = NewSubMenu(sm.parent, sm.items[index].subItems)
					sm.childMenu.box.SetRect(rectX+rectW, y, 15, 10)
					return
				}
			}
			sm.parent.subMenu = nil
		}
		return
	})
}

////////////////////////////////////////////////////

type MenuItem struct {
	box      *Box
	title    string
	subItems []*MenuItem
	onClick  func(*MenuItem)
	mu       sync.RWMutex
}

func (mi *MenuItem) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return mi.box.MouseHandler()
}

func NewMenuItem(title string) *MenuItem {
	return &MenuItem{
		box:      NewBox(),
		title:    title,
		subItems: make([]*MenuItem, 0),
	}
}

func (mi *MenuItem) AddItem(item *MenuItem) *MenuItem {
	mi.mu.Lock()
	mi.subItems = append(mi.subItems, item)
	mi.mu.Unlock()
	return mi
}

func (mi *MenuItem) SetOnClick(fn func(*MenuItem)) *MenuItem {
	mi.mu.Lock()
	mi.onClick = fn
	mi.mu.Unlock()
	return mi
}

func (mi *MenuItem) Draw(screen tcell.Screen) {
	if !mi.box.GetVisible() {
		return
	}

	mi.box.Draw(screen)

	mi.mu.Lock()
	defer mi.mu.Unlock()

	x, y, _, _ := mi.box.GetInnerRect()

	PrintSimple(screen, []byte(mi.title), x, y)
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (mi *MenuItem) set(setter func(mi *MenuItem)) *MenuItem {
	mi.mu.Lock()
	setter(mi)
	mi.mu.Unlock()
	return mi
}

func (mi *MenuItem) get(getter func(mi *MenuItem)) {
	mi.mu.RLock()
	getter(mi)
	mi.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this MenuItem.
func (mi *MenuItem) GetTitle() string {
	return mi.box.GetTitle()
}

// SetTitle sets the title of this MenuItem.
func (mi *MenuItem) SetTitle(title string) *MenuItem {
	mi.box.SetTitle(title)
	return mi
}

// GetTitleAlign returns the title alignment of this MenuItem.
func (mi *MenuItem) GetTitleAlign() int {
	return mi.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this MenuItem.
func (mi *MenuItem) SetTitleAlign(align int) *MenuItem {
	mi.box.SetTitleAlign(align)
	return mi
}

// GetBorder returns whether this MenuItem has a border.
func (mi *MenuItem) GetBorder() bool {
	return mi.box.GetBorder()
}

// SetBorder sets whether this MenuItem has a border.
func (mi *MenuItem) SetBorder(show bool) *MenuItem {
	mi.box.SetBorder(show)
	return mi
}

// GetBorderColor returns the border color of this MenuItem.
func (mi *MenuItem) GetBorderColor() tcell.Color {
	return mi.box.GetBorderColor()
}

// SetBorderColor sets the border color of this MenuItem.
func (mi *MenuItem) SetBorderColor(color tcell.Color) *MenuItem {
	mi.box.SetBorderColor(color)
	return mi
}

// GetBorderAttributes returns the border attributes of this MenuItem.
func (mi *MenuItem) GetBorderAttributes() tcell.AttrMask {
	return mi.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this MenuItem.
func (mi *MenuItem) SetBorderAttributes(attr tcell.AttrMask) *MenuItem {
	mi.box.SetBorderAttributes(attr)
	return mi
}

// GetBorderColorFocused returns the border color of this MenuItem when focused.
func (mi *MenuItem) GetBorderColorFocused() tcell.Color {
	return mi.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this MenuItem when focused.
func (mi *MenuItem) SetBorderColorFocused(color tcell.Color) *MenuItem {
	mi.box.SetBorderColorFocused(color)
	return mi
}

// GetTitleColor returns the title color of this MenuItem.
func (mi *MenuItem) GetTitleColor() tcell.Color {
	return mi.box.GetTitleColor()
}

// SetTitleColor sets the title color of this MenuItem.
func (mi *MenuItem) SetTitleColor(color tcell.Color) *MenuItem {
	mi.box.SetTitleColor(color)
	return mi
}

// GetDrawFunc returns the custom draw function of this MenuItem.
func (mi *MenuItem) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return mi.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this MenuItem.
func (mi *MenuItem) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *MenuItem {
	mi.box.SetDrawFunc(handler)
	return mi
}

// ShowFocus sets whether this MenuItem should show a focus indicator when focused.
func (mi *MenuItem) ShowFocus(showFocus bool) *MenuItem {
	mi.box.ShowFocus(showFocus)
	return mi
}

// GetMouseCapture returns the mouse capture function of this MenuItem.
func (mi *MenuItem) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return mi.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this MenuItem.
func (mi *MenuItem) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *MenuItem {
	mi.box.SetMouseCapture(capture)
	return mi
}

// GetBackgroundColor returns the background color of this MenuItem.
func (mi *MenuItem) GetBackgroundColor() tcell.Color {
	return mi.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this MenuItem.
func (mi *MenuItem) SetBackgroundColor(color tcell.Color) *MenuItem {
	mi.box.SetBackgroundColor(color)
	return mi
}

// GetBackgroundTransparent returns whether the background of this MenuItem is transparent.
func (mi *MenuItem) GetBackgroundTransparent() bool {
	return mi.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this MenuItem is transparent.
func (mi *MenuItem) SetBackgroundTransparent(transparent bool) *MenuItem {
	mi.box.SetBackgroundTransparent(transparent)
	return mi
}

// GetInputCapture returns the input capture function of this MenuItem.
func (mi *MenuItem) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return mi.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this MenuItem.
func (mi *MenuItem) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *MenuItem {
	mi.box.SetInputCapture(capture)
	return mi
}

// GetPadding returns the padding of this MenuItem.
func (mi *MenuItem) GetPadding() (top, bottom, left, right int) {
	return mi.box.GetPadding()
}

// SetPadding sets the padding of this MenuItem.
func (mi *MenuItem) SetPadding(top, bottom, left, right int) *MenuItem {
	mi.box.SetPadding(top, bottom, left, right)
	return mi
}

// InRect returns whether the given screen coordinates are within this MenuItem.
func (mi *MenuItem) InRect(x, y int) bool {
	return mi.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this MenuItem.
func (mi *MenuItem) GetInnerRect() (x, y, width, height int) {
	return mi.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the MenuItem is preserved.
func (mi *MenuItem) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return mi.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the MenuItem is preserved.
func (mi *MenuItem) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return mi.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this MenuItem.
func (mi *MenuItem) GetRect() (x, y, width, height int) {
	return mi.box.GetRect()
}

// SetRect sets the rectangle occupied by this MenuItem.
func (mi *MenuItem) SetRect(x, y, width, height int) {
	mi.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this MenuItem is visible.
func (mi *MenuItem) GetVisible() bool {
	return mi.box.GetVisible()
}

// SetVisible sets whether this MenuItem is visible.
func (mi *MenuItem) SetVisible(visible bool) {
	mi.box.SetVisible(visible)
}

// Focus is called when this MenuItem receives focus.
func (mi *MenuItem) Focus(delegate func(p Widget)) {
	mi.box.Focus(delegate)
}

// HasFocus returns whether this MenuItem has focus.
func (mi *MenuItem) HasFocus() bool {
	return mi.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this MenuItem.
func (mi *MenuItem) GetFocusable() Focusable {
	return mi.box.GetFocusable()
}

// Blur is called when this MenuItem loses focus.
func (mi *MenuItem) Blur() {
	mi.box.Blur()
}
