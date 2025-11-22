package cui

import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

// panel represents a single panel of a Panels object.
type panel struct {
	Name    string // The panel's name.
	Item    Widget // The panel's primitive.
	Resize  bool   // Whether to resize the panel when it is drawn.
	Visible bool   // Whether this panel is visible.
}

// Panels is a container for other primitives often used as the application's
// root primitive. It allows to easily switch the visibility of the contained
// primitives.
type Panels struct {
	box *Box

	// The contained panels. (Visible) panels are drawn from back to front.
	panels []*panel

	// We keep a reference to the function which allows us to set the focus to
	// a newly visible panel.
	setFocus func(p Widget)

	// An optional handler which is called whenever the visibility or the order of
	// panels changes.
	changed func()

	mu sync.RWMutex
}

// NewPanels returns a new Panels object.
func NewPanels() *Panels {
	p := &Panels{
		box: NewBox(),
	}
	p.box.focus = p
	return p
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (p *Panels) set(setter func(p *Panels)) *Panels {
	p.mu.Lock()
	setter(p)
	p.mu.Unlock()
	return p
}

func (p *Panels) get(getter func(p *Panels)) {
	p.mu.RLock()
	getter(p)
	p.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this Panels.
func (p *Panels) GetTitle() string {
	return p.box.GetTitle()
}

// SetTitle sets the title of this Panels.
func (p *Panels) SetTitle(title string) *Panels {
	p.box.SetTitle(title)
	return p
}

// GetTitleAlign returns the title alignment of this Panels.
func (p *Panels) GetTitleAlign() int {
	return p.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this Panels.
func (p *Panels) SetTitleAlign(align int) *Panels {
	p.box.SetTitleAlign(align)
	return p
}

// GetBorder returns whether this Panels has a border.
func (p *Panels) GetBorder() bool {
	return p.box.GetBorder()
}

// SetBorder sets whether this Panels has a border.
func (p *Panels) SetBorder(show bool) *Panels {
	p.box.SetBorder(show)
	return p
}

// GetBorderColor returns the border color of this Panels.
func (p *Panels) GetBorderColor() tcell.Color {
	return p.box.GetBorderColor()
}

// SetBorderColor sets the border color of this Panels.
func (p *Panels) SetBorderColor(color tcell.Color) *Panels {
	p.box.SetBorderColor(color)
	return p
}

// GetBorderAttributes returns the border attributes of this Panels.
func (p *Panels) GetBorderAttributes() tcell.AttrMask {
	return p.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this Panels.
func (p *Panels) SetBorderAttributes(attr tcell.AttrMask) *Panels {
	p.box.SetBorderAttributes(attr)
	return p
}

// GetBorderColorFocused returns the border color of this Panels when focusel.
func (p *Panels) GetBorderColorFocused() tcell.Color {
	return p.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this Panels when focusel.
func (p *Panels) SetBorderColorFocused(color tcell.Color) *Panels {
	p.box.SetBorderColorFocused(color)
	return p
}

// GetTitleColor returns the title color of this Panels.
func (p *Panels) GetTitleColor() tcell.Color {
	return p.box.GetTitleColor()
}

// SetTitleColor sets the title color of this Panels.
func (p *Panels) SetTitleColor(color tcell.Color) *Panels {
	p.box.SetTitleColor(color)
	return p
}

// GetDrawFunc returns the custom draw function of this Panels.
func (p *Panels) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return p.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this Panels.
func (p *Panels) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *Panels {
	p.box.SetDrawFunc(handler)
	return p
}

// ShowFocus sets whether this Panels should show a focus indicator when focusel.
func (p *Panels) ShowFocus(showFocus bool) *Panels {
	p.box.ShowFocus(showFocus)
	return p
}

// GetMouseCapture returns the mouse capture function of this Panels.
func (p *Panels) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return p.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this Panels.
func (p *Panels) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *Panels {
	p.box.SetMouseCapture(capture)
	return p
}

// GetBackgroundColor returns the background color of this Panels.
func (p *Panels) GetBackgroundColor() tcell.Color {
	return p.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this Panels.
func (p *Panels) SetBackgroundColor(color tcell.Color) *Panels {
	p.box.SetBackgroundColor(color)
	return p
}

// GetBackgroundTransparent returns whether the background of this Panels is transparent.
func (p *Panels) GetBackgroundTransparent() bool {
	return p.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this Panels is transparent.
func (p *Panels) SetBackgroundTransparent(transparent bool) *Panels {
	p.box.SetBackgroundTransparent(transparent)
	return p
}

// GetInputCapture returns the input capture function of this Panels.
func (p *Panels) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return p.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this Panels.
func (p *Panels) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *Panels {
	p.box.SetInputCapture(capture)
	return p
}

// GetPadding returns the padding of this Panels.
func (p *Panels) GetPadding() (top, bottom, left, right int) {
	return p.box.GetPadding()
}

// SetPadding sets the padding of this Panels.
func (p *Panels) SetPadding(top, bottom, left, right int) *Panels {
	p.box.SetPadding(top, bottom, left, right)
	return p
}

// InRect returns whether the given screen coordinates are within this Panels.
func (p *Panels) InRect(x, y int) bool {
	return p.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this Panels.
func (p *Panels) GetInnerRect() (x, y, width, height int) {
	return p.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the Panels is preservel.
func (p *Panels) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the Panels is preservel.
func (p *Panels) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this Panels.
func (p *Panels) GetRect() (x, y, width, height int) {
	return p.box.GetRect()
}

// SetRect sets the rectangle occupied by this Panels.
func (p *Panels) SetRect(x, y, width, height int) {
	p.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this Panels is visible.
func (p *Panels) GetVisible() bool {
	return p.box.GetVisible()
}

// SetVisible sets whether this Panels is visible.
func (p *Panels) SetVisible(visible bool) {
	p.box.SetVisible(visible)
}

// Focus is called by the application when the primitive receives focus.
func (p *Panels) Focus(delegate func(p Widget)) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if delegate == nil {
		return // We cannot delegate so we cannot focus.
	}
	p.setFocus = delegate
	var topItem Widget
	for _, panel := range p.panels {
		if panel.Visible {
			topItem = panel.Item
		}
	}
	if topItem != nil {
		p.mu.Unlock()
		delegate(topItem)
		p.mu.Lock()
	}
}

// HasFocus returns whether or not this primitive has focus.
func (p *Panels) HasFocus() bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, panel := range p.panels {
		if panel.Item.GetFocusable().HasFocus() {
			return true
		}
	}
	return false
}

// GetFocusable returns the focusable primitive of this Panels.
func (p *Panels) GetFocusable() Focusable {
	return p.box.GetFocusable()
}

// Blur is called when this Panels loses focus.
func (p *Panels) Blur() {
	p.box.Blur()
}

// SetChangedFunc sets a handler which is called whenever the visibility or the
// order of any visible panels changes. This can be used to redraw the panels.
func (p *Panels) SetChangedFunc(handler func()) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.changed = handler
}

// GetPanelCount returns the number of panels currently stored in this object.
func (p *Panels) GetPanelCount() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.panels)
}

// AddPanel adds a new panel with the given name and primitive. If there was
// previously a panel with the same name, it is overwritten. Leaving the name
// empty may cause conflicts in other functions so always specify a non-empty
// name.
//
// Visible panels will be drawn in the order they were added (unless that order
// was changed in one of the other functions). If "resize" is set to true, the
// primitive will be set to the size available to the Panels primitive whenever
// the panels are drawn.
func (p *Panels) AddPanel(name string, item Widget, resize, visible bool) {
	hasFocus := p.HasFocus()

	p.mu.Lock()
	defer p.mu.Unlock()

	var added bool
	for i, pg := range p.panels {
		if pg.Name == name {
			p.panels[i] = &panel{Item: item, Name: name, Resize: resize, Visible: visible}
			added = true
			break
		}
	}
	if !added {
		p.panels = append(p.panels, &panel{Item: item, Name: name, Resize: resize, Visible: visible})
	}
	if p.changed != nil {
		p.mu.Unlock()
		p.changed()
		p.mu.Lock()
	}
	if hasFocus {
		p.mu.Unlock()
		p.Focus(p.setFocus)
		p.mu.Lock()
	}
}

// RemovePanel removes the panel with the given name. If that panel was the only
// visible panel, visibility is assigned to the last panel.
func (p *Panels) RemovePanel(name string) {
	hasFocus := p.HasFocus()

	p.mu.Lock()
	defer p.mu.Unlock()

	var isVisible bool
	for index, panel := range p.panels {
		if panel.Name == name {
			isVisible = panel.Visible
			p.panels = append(p.panels[:index], p.panels[index+1:]...)
			if panel.Visible && p.changed != nil {
				p.mu.Unlock()
				p.changed()
				p.mu.Lock()
			}
			break
		}
	}
	if isVisible {
		for index, panel := range p.panels {
			if index < len(p.panels)-1 {
				if panel.Visible {
					break // There is a remaining visible panel.
				}
			} else {
				panel.Visible = true // We need at least one visible panel.
			}
		}
	}
	if hasFocus {
		p.mu.Unlock()
		p.Focus(p.setFocus)
		p.mu.Lock()
	}
}

// HasPanel returns true if a panel with the given name exists in this object.
func (p *Panels) HasPanel(name string) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for _, panel := range p.panels {
		if panel.Name == name {
			return true
		}
	}
	return false
}

// ShowPanel sets a panel's visibility to "true" (in addition to any other panels
// which are already visible).
func (p *Panels) ShowPanel(name string) {
	hasFocus := p.HasFocus()

	p.mu.Lock()
	defer p.mu.Unlock()

	for _, panel := range p.panels {
		if panel.Name == name {
			panel.Visible = true
			if p.changed != nil {
				p.mu.Unlock()
				p.changed()
				p.mu.Lock()
			}
			break
		}
	}
	if hasFocus {
		p.mu.Unlock()
		p.Focus(p.setFocus)
		p.mu.Lock()
	}
}

// HidePanel sets a panel's visibility to "false".
func (p *Panels) HidePanel(name string) {
	hasFocus := p.HasFocus()

	p.mu.Lock()
	defer p.mu.Unlock()

	for _, panel := range p.panels {
		if panel.Name == name {
			panel.Visible = false
			if p.changed != nil {
				p.mu.Unlock()
				p.changed()
				p.mu.Lock()
			}
			break
		}
	}
	if hasFocus {
		p.mu.Unlock()
		p.Focus(p.setFocus)
		p.mu.Lock()
	}
}

// SetCurrentPanel sets a panel's visibility to "true" and all other panels'
// visibility to "false".
func (p *Panels) SetCurrentPanel(name string) {
	hasFocus := p.HasFocus()

	p.mu.Lock()
	defer p.mu.Unlock()

	for _, panel := range p.panels {
		if panel.Name == name {
			panel.Visible = true
		} else {
			panel.Visible = false
		}
	}
	if p.changed != nil {
		p.mu.Unlock()
		p.changed()
		p.mu.Lock()
	}
	if hasFocus {
		p.mu.Unlock()
		p.Focus(p.setFocus)
		p.mu.Lock()
	}
}

// SendToFront changes the order of the panels such that the panel with the given
// name comes last, causing it to be drawn last with the next update (if
// visible).
func (p *Panels) SendToFront(name string) {
	hasFocus := p.HasFocus()

	p.mu.Lock()
	defer p.mu.Unlock()

	for index, panel := range p.panels {
		if panel.Name == name {
			if index < len(p.panels)-1 {
				p.panels = append(append(p.panels[:index], p.panels[index+1:]...), panel)
			}
			if panel.Visible && p.changed != nil {
				p.mu.Unlock()
				p.changed()
				p.mu.Lock()
			}
			break
		}
	}
	if hasFocus {
		p.mu.Unlock()
		p.Focus(p.setFocus)
		p.mu.Lock()
	}
}

// SendToBack changes the order of the panels such that the panel with the given
// name comes first, causing it to be drawn first with the next update (if
// visible).
func (p *Panels) SendToBack(name string) {
	hasFocus := p.HasFocus()

	p.mu.Lock()
	defer p.mu.Unlock()

	for index, pg := range p.panels {
		if pg.Name == name {
			if index > 0 {
				p.panels = append(append([]*panel{pg}, p.panels[:index]...), p.panels[index+1:]...)
			}
			if pg.Visible && p.changed != nil {
				p.mu.Unlock()
				p.changed()
				p.mu.Lock()
			}
			break
		}
	}
	if hasFocus {
		p.mu.Unlock()
		p.Focus(p.setFocus)
		p.mu.Lock()
	}
}

// GetFrontPanel returns the front-most visible panel. If there are no visible
// panels, ("", nil) is returned.
func (p *Panels) GetFrontPanel() (name string, item Widget) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	for index := len(p.panels) - 1; index >= 0; index-- {
		if p.panels[index].Visible {
			return p.panels[index].Name, p.panels[index].Item
		}
	}
	return
}

// Draw draws this primitive onto the screen.
func (p *Panels) Draw(screen tcell.Screen) {
	if !p.GetVisible() {
		return
	}

	p.box.Draw(screen)

	p.mu.Lock()
	defer p.mu.Unlock()

	x, y, width, height := p.GetInnerRect()

	for _, panel := range p.panels {
		if !panel.Visible {
			continue
		}
		if panel.Resize {
			panel.Item.SetRect(x, y, width, height)
		}
		panel.Item.Draw(screen)
	}
}

func (p *Panels) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return p.box.InputHandler()
}

// MouseHandler returns the mouse handler for this primitive.
func (p *Panels) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return p.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		if !p.InRect(event.Position()) {
			return false, nil
		}

		// Pass mouse events along to the last visible panel item that takes it.
		for index := len(p.panels) - 1; index >= 0; index-- {
			panel := p.panels[index]
			if panel.Visible {
				consumed, capture = panel.Item.MouseHandler()(action, event, setFocus)
				if consumed {
					return
				}
			}
		}

		return
	})
}

//
//// Support backwards compatibility with Pages.
//type page = panel
//
//// Pages is a wrapper around Panels.
////
//// Deprecated: This type is provided for backwards compatibility.
//// Developers should use Panels instead.
//type Pages struct {
//	*Panels
//}
//
//// NewPages returns a new Panels object.
////
//// Deprecated: This function is provided for backwards compatibility.
//// Developers should use NewPanels instead.
//func NewPages() *Pages {
//	return &Pages{NewPanels()}
//}
//
//// GetPageCount returns the number of panels currently stored in this object.
//func (p *Pages) GetPageCount() int {
//	return p.GetPanelCount()
//}
//
//// AddPage adds a new panel with the given name and primitive.
//func (p *Pages) AddPage(name string, item Widget, resize, visible bool) {
//	p.AddPanel(name, item, resize, visible)
//}
//
//// AddAndSwitchToPage calls Add(), then SwitchTo() on that newly added panel.
//func (p *Pages) AddAndSwitchToPage(name string, item Widget, resize bool) {
//	p.AddPanel(name, item, resize, true)
//	p.SetCurrentPanel(name)
//}
//
//// RemovePage removes the panel with the given name.
//func (p *Pages) RemovePage(name string) {
//	p.RemovePanel(name)
//}
//
//// HasPage returns true if a panel with the given name exists in this object.
//func (p *Pages) HasPage(name string) bool {
//	return p.HasPanel(name)
//}
//
//// ShowPage sets a panel's visibility to "true".
//func (p *Pages) ShowPage(name string) {
//	p.ShowPanel(name)
//}
//
//// HidePage sets a panel's visibility to "false".
//func (p *Pages) HidePage(name string) {
//	p.HidePanel(name)
//}
//
//// SwitchToPage sets a panel's visibility to "true" and all other panels'
//// visibility to "false".
//func (p *Pages) SwitchToPage(name string) {
//	p.SetCurrentPanel(name)
//}
//
//// GetFrontPage returns the front-most visible panel.
//func (p *Pages) GetFrontPage() (name string, item Widget) {
//	return p.GetFrontPanel()
//}
