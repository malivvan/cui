package cui

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/gdamore/tcell/v2"
)

// TabbedPanels is a tabbed container for other primitives. The tab switcher
// may be positioned vertically or horizontally, before or after the content.
type TabbedPanels struct {
	flex     *Flex
	switcher *Text
	panels   *Panels

	tabLabels  map[string]string
	currentTab string

	dividerStart string
	dividerMid   string
	dividerEnd   string

	switcherVertical     bool
	switcherAfterContent bool
	switcherHeight       int

	width, lastWidth int

	setFocus func(Widget)

	mu sync.RWMutex
}

// NewTabbedPanels returns a new TabbedPanels object.
func NewTabbedPanels() *TabbedPanels {
	t := &TabbedPanels{
		flex:       NewFlex(),
		switcher:   NewTextView(),
		panels:     NewPanels(),
		dividerMid: string(BoxDrawingsDoubleVertical),
		dividerEnd: string(BoxDrawingsLightVertical),
		tabLabels:  make(map[string]string),
	}

	s := t.switcher
	s.SetDynamicColors(true)
	s.SetHighlightForegroundColor(Styles.InverseTextColor)
	s.SetHighlightBackgroundColor(Styles.PrimaryTextColor)
	s.SetRegions(true)
	s.SetScrollable(true)
	s.SetWrap(true)
	s.SetWordWrap(true)
	s.SetHighlightedFunc(func(added, removed, remaining []string) {
		if len(added) == 0 {
			return
		}

		s.ScrollToHighlight()
		t.SetCurrentTab(added[0])
		if t.setFocus != nil {
			t.setFocus(t.panels)
		}
	})

	t.rebuild()

	return t
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (t *TabbedPanels) set(setter func(t *TabbedPanels)) *TabbedPanels {
	t.mu.Lock()
	setter(t)
	t.mu.Unlock()
	return t
}

func (t *TabbedPanels) get(getter func(t *TabbedPanels)) {
	t.mu.RLock()
	getter(t)
	t.mu.RUnlock()
}

///////////////////////////////////// <FLEX> ////////////////////////////////////

func (t *TabbedPanels) GetTitle() string {
	return t.flex.GetTitle()
}
func (t *TabbedPanels) SetTitle(title string) *TabbedPanels {
	t.flex.SetTitle(title)
	return t
}

func (t *TabbedPanels) GetTitleColor() tcell.Color {
	return t.flex.GetTitleColor()
}
func (t *TabbedPanels) SetTitleColor(color tcell.Color) *TabbedPanels {
	t.flex.SetTitleColor(color)
	return t
}

func (t *TabbedPanels) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return t.flex.GetDrawFunc()
}

func (t *TabbedPanels) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *TabbedPanels {
	t.flex.SetDrawFunc(handler)
	return t
}

func (t *TabbedPanels) HasFocus() bool {
	return t.flex.HasFocus()
}
func (t *TabbedPanels) ShowFocus(showFocus bool) *TabbedPanels {
	t.flex.ShowFocus(showFocus)
	return t
}

func (t *TabbedPanels) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return t.flex.GetMouseCapture()
}

func (t *TabbedPanels) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *TabbedPanels {
	t.flex.SetMouseCapture(capture)
	return t
}

func (t *TabbedPanels) GetBorder() (border bool) {
	return t.flex.GetBorder()
}
func (t *TabbedPanels) SetBorder(show bool) *TabbedPanels {
	t.flex.SetBorder(show)
	return t
}

func (t *TabbedPanels) GetBorderColorFocused() (color tcell.Color) {
	return t.flex.GetBorderColorFocused()
}
func (t *TabbedPanels) SetBorderColorFocused(color tcell.Color) *TabbedPanels {
	t.flex.SetBorderColorFocused(color)
	return t
}
func (t *TabbedPanels) GetTitleAlign() (align int) {
	return t.flex.GetTitleAlign()
}
func (t *TabbedPanels) SetTitleAlign(align int) *TabbedPanels {
	t.flex.SetTitleAlign(align)
	return t
}

func (t *TabbedPanels) GetBorderAttributes() (attr tcell.AttrMask) {
	return t.flex.GetBorderAttributes()
}
func (t *TabbedPanels) SetBorderAttributes(attr tcell.AttrMask) *TabbedPanels {
	t.flex.SetBorderAttributes(attr)
	return t
}

func (t *TabbedPanels) GetBorderColor() (color tcell.Color) {
	return t.flex.GetBorderColor()
}
func (t *TabbedPanels) SetBorderColor(color tcell.Color) *TabbedPanels {
	t.flex.SetBorderColor(color)
	return t
}

func (t *TabbedPanels) GetBackgroundTransparent() (transparent bool) {
	return t.flex.GetBackgroundTransparent()
}

func (t *TabbedPanels) SetBackgroundTransparent(transparent bool) *TabbedPanels {
	t.flex.SetBackgroundTransparent(transparent)
	return t
}

func (t *TabbedPanels) GetPadding() (top, bottom, left, right int) {
	return t.flex.GetPadding()
}
func (t *TabbedPanels) SetPadding(top, bottom, left, right int) *TabbedPanels {
	t.flex.SetPadding(top, bottom, left, right)
	return t
}

func (t *TabbedPanels) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return t.flex.GetInputCapture()
}

func (t *TabbedPanels) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *TabbedPanels {
	t.flex.SetInputCapture(capture)
	return t
}

func (t *TabbedPanels) WrapMouseHandler(mouseHandler func(MouseAction, *tcell.EventMouse, func(p Widget)) (bool, Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return t.flex.WrapMouseHandler(mouseHandler)
}
func (t *TabbedPanels) WrapInputHandler(inputHandler func(*tcell.EventKey, func(p Widget))) func(*tcell.EventKey, func(p Widget)) {
	return t.flex.WrapInputHandler(inputHandler)
}
func (t *TabbedPanels) GetInnerRect() (innerX, innerY, innerW, innerH int) {
	return t.flex.GetInnerRect()
}
func (t *TabbedPanels) InRect(x, y int) bool {
	return t.flex.InRect(x, y)
}

func (t *TabbedPanels) GetRect() (x, y, width, height int) {
	return t.flex.GetRect()
}
func (t *TabbedPanels) SetRect(x, y, width, height int) {
	t.flex.SetRect(x, y, width, height)
}

func (t *TabbedPanels) GetVisible() bool {
	return t.flex.GetVisible()
}
func (t *TabbedPanels) SetVisible(visible bool) {
	t.flex.SetVisible(visible)
}

func (t *TabbedPanels) Focus(delegate func(p Widget)) {
	t.flex.Focus(delegate)
}

func (t *TabbedPanels) GetFocusable() Focusable {
	return t.flex.GetFocusable()
}
func (t *TabbedPanels) Blur() {
	t.flex.Blur()
}

/////////////////////////////////////// <API> ///////////////////////////////////////

// SetChangedFunc sets a handler which is called whenever a tab is added,
// selected, reordered or removed.
func (t *TabbedPanels) SetChangedFunc(handler func()) *TabbedPanels {
	t.panels.SetChangedFunc(handler)
	return t
}

// AddTab adds a new tab. Tab names should consist only of letters, numbers
// and spaces.
func (t *TabbedPanels) AddTab(name, label string, item Widget) *TabbedPanels {
	t.mu.Lock()
	t.tabLabels[name] = label
	t.mu.Unlock()

	t.panels.AddPanel(name, item, true, false)

	t.updateAll()
	return t
}

// RemoveTab removes a tab.
func (t *TabbedPanels) RemoveTab(name string) *TabbedPanels {
	t.panels.RemovePanel(name)

	t.updateAll()
	return t
}

// HasTab returns true if a tab with the given name exists in this object.
func (t *TabbedPanels) HasTab(name string) bool {
	t.mu.RLock()
	defer t.mu.RUnlock()

	for _, panel := range t.panels.panels {
		if panel.Name == name {
			return true
		}
	}
	return false
}

// SetCurrentTab sets the currently visible tab.
func (t *TabbedPanels) SetCurrentTab(name string) *TabbedPanels {
	t.mu.Lock()

	if t.currentTab == name {
		t.mu.Unlock()
		return t
	}

	t.currentTab = name

	t.updateAll()

	t.mu.Unlock()

	h := t.switcher.GetHighlights()
	var found bool
	for _, hl := range h {
		if hl == name {
			found = true
			break
		}
	}
	if !found {
		t.switcher.Highlight(t.currentTab)
	}
	t.switcher.ScrollToHighlight()
	return t
}

// GetCurrentTab returns the currently visible tab.
func (t *TabbedPanels) GetCurrentTab() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.currentTab
}

// SetTabLabel sets the label of a tab.
func (t *TabbedPanels) SetTabLabel(name, label string) *TabbedPanels {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.tabLabels[name] == label {
		return t
	}

	t.tabLabels[name] = label
	t.updateTabLabels()
	return t
}

// SetBackgroundColor sets the background color of the tabbed panels.
func (t *TabbedPanels) SetBackgroundColor(color tcell.Color) *TabbedPanels {
	t.panels.box.SetBackgroundColor(color)
	return t
}

// GetBackgroundColor returns the background color of the tabbed panels.
func (t *TabbedPanels) GetBackgroundColor() tcell.Color {
	return t.panels.box.GetBackgroundColor()
}

// SetTabTextColor sets the color of the tab text.
func (t *TabbedPanels) SetTabTextColor(color tcell.Color) *TabbedPanels {
	t.switcher.SetTextColor(color)
	return t
}

// SetTabTextColorFocused sets the color of the tab text when the tab is in focus.
func (t *TabbedPanels) SetTabTextColorFocused(color tcell.Color) *TabbedPanels {
	t.switcher.SetHighlightForegroundColor(color)
	return t
}

// SetTabBackgroundColor sets the background color of the tab.
func (t *TabbedPanels) SetTabBackgroundColor(color tcell.Color) *TabbedPanels {
	t.switcher.SetBackgroundColor(color)
	return t
}

// SetTabBackgroundColorFocused sets the background color of the tab when the
// tab is in focus.
func (t *TabbedPanels) SetTabBackgroundColorFocused(color tcell.Color) *TabbedPanels {
	t.switcher.SetHighlightBackgroundColor(color)
	return t
}

// SetTabSwitcherDivider sets the tab switcher divider text. Color tags are supported.
func (t *TabbedPanels) SetTabSwitcherDivider(start, mid, end string) *TabbedPanels {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.dividerStart, t.dividerMid, t.dividerEnd = start, mid, end
	return t
}

// SetTabSwitcherHeight sets the tab switcher height. This setting only applies
// when rendering horizontally. A value of 0 (the default) indicates the height
// should automatically adjust to fit all of the tab labels.
func (t *TabbedPanels) SetTabSwitcherHeight(height int) *TabbedPanels {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.switcherHeight = height
	t.rebuild()
	return t
}

// SetTabSwitcherVertical sets the orientation of the tab switcher.
func (t *TabbedPanels) SetTabSwitcherVertical(vertical bool) *TabbedPanels {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.switcherVertical == vertical {
		return t
	}

	t.switcherVertical = vertical
	t.rebuild()
	return t
}

// SetTabSwitcherAfterContent sets whether the tab switcher is positioned after content.
func (t *TabbedPanels) SetTabSwitcherAfterContent(after bool) *TabbedPanels {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.switcherAfterContent == after {
		return t
	}

	t.switcherAfterContent = after
	t.rebuild()
	return t
}

func (t *TabbedPanels) rebuild() {
	f := t.flex
	if t.switcherVertical {
		f.SetDirection(FlexColumn)
	} else {
		f.SetDirection(FlexRow)
	}
	f.RemoveItem(t.panels)
	f.RemoveItem(t.switcher)
	if t.switcherAfterContent {
		f.AddItem(t.panels, 0, 1, true)
		f.AddItem(t.switcher, 1, 1, false)
	} else {
		f.AddItem(t.switcher, 1, 1, false)
		f.AddItem(t.panels, 0, 1, true)
	}

	t.updateTabLabels()

	t.switcher.SetMaxLines(t.switcherHeight)
}

func (t *TabbedPanels) updateTabLabels() {
	if len(t.panels.panels) == 0 {
		t.switcher.SetText("")
		t.flex.ResizeItem(t.switcher, 0, 0)
		return
	}

	maxWidth := 0
	for _, panel := range t.panels.panels {
		label := t.tabLabels[panel.Name]
		if len(label) > maxWidth {
			maxWidth = len(label)
		}
	}

	var b bytes.Buffer
	if !t.switcherVertical {
		b.WriteString(t.dividerStart)
	}
	l := len(t.panels.panels)
	spacer := []byte(" ")
	for i, panel := range t.panels.panels {
		if i > 0 && t.switcherVertical {
			b.WriteRune('\n')
		}

		if t.switcherVertical && t.switcherAfterContent {
			b.WriteString(t.dividerMid)
			b.WriteRune(' ')
		}

		label := t.tabLabels[panel.Name]
		if !t.switcherVertical {
			label = " " + label
		}

		if t.switcherVertical {
			spacer = bytes.Repeat([]byte(" "), maxWidth-len(label)+1)
		}

		b.WriteString(fmt.Sprintf(`["%s"]%s%s[""]`, panel.Name, label, spacer))

		if i == l-1 && !t.switcherVertical {
			b.WriteString(t.dividerEnd)
		} else if !t.switcherAfterContent {
			b.WriteString(t.dividerMid)
		}
	}
	t.switcher.SetText(b.String())

	var reqLines int
	if t.switcherVertical {
		reqLines = maxWidth + 2
	} else {
		if t.switcherHeight > 0 {
			reqLines = t.switcherHeight
		} else {
			reqLines = len(WordWrap(t.switcher.GetText(true), t.width))
			if reqLines < 1 {
				reqLines = 1
			}
		}
	}
	t.flex.ResizeItem(t.switcher, reqLines, 1)
}

func (t *TabbedPanels) updateVisibleTabs() {
	allPanels := t.panels.panels

	var newTab string

	var foundCurrent bool
	for _, panel := range allPanels {
		if panel.Name == t.currentTab {
			newTab = panel.Name
			foundCurrent = true
			break
		}
	}
	if !foundCurrent {
		for _, panel := range allPanels {
			if panel.Name != "" {
				newTab = panel.Name
				break
			}
		}
	}

	if t.currentTab != newTab {
		t.SetCurrentTab(newTab)
		return
	}

	for _, panel := range allPanels {
		if panel.Name == t.currentTab {
			t.panels.ShowPanel(panel.Name)
		} else {
			t.panels.HidePanel(panel.Name)
		}
	}
}

func (t *TabbedPanels) updateAll() {
	t.updateTabLabels()
	t.updateVisibleTabs()
}

// Draw draws this primitive onto the screen.
func (t *TabbedPanels) Draw(screen tcell.Screen) {
	if !t.GetVisible() {
		return
	}

	t.flex.box.Draw(screen)

	_, _, t.width, _ = t.GetInnerRect()
	if t.width != t.lastWidth {
		t.updateTabLabels()
	}
	t.lastWidth = t.width

	t.flex.Draw(screen)
}

// InputHandler returns the handler for this primitive.
func (t *TabbedPanels) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		if t.setFocus == nil {
			t.setFocus = setFocus
		}
		t.flex.InputHandler()(event, setFocus)
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (t *TabbedPanels) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return t.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if t.setFocus == nil {
			t.setFocus = setFocus
		}

		x, y := event.Position()
		if !t.InRect(x, y) {
			return false, nil
		}

		if t.switcher.InRect(x, y) {
			if t.setFocus != nil {
				defer t.setFocus(t.panels)
			}
			defer t.switcher.MouseHandler()(action, event, setFocus)
			return true, nil
		}

		return t.flex.MouseHandler()(action, event, setFocus)
	})
}
