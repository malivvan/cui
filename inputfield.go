package cui

import (
	"bytes"
	"math"
	"regexp"
	"sync"
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

// InputField is a one-line box (three lines if there is a title) where the
// user can enter text. Use SetAcceptanceFunc() to accept or reject input,
// SetChangedFunc() to listen for changes, and SetMaskCharacter() to hide input
// from onlookers (e.g. for password input).
//
// The following keys can be used for navigation and editing:
//
//   - Left arrow: Move left by one character.
//   - Right arrow: Move right by one character.
//   - Home, Ctrl-A, Alt-a: Move to the beginning of the line.
//   - End, Ctrl-E, Alt-e: Move to the end of the line.
//   - Alt-left, Alt-b: Move left by one word.
//   - Alt-right, Alt-f: Move right by one word.
//   - Backspace: Delete the character before the cursor.
//   - Delete: Delete the character after the cursor.
//   - Ctrl-K: Delete from the cursor to the end of the line.
//   - Ctrl-W: Delete the last word before the cursor.
//   - Ctrl-U: Delete the entire line.
type InputField struct {
	box *Box

	// The text that was entered.
	text []byte

	// The text to be displayed before the input area.
	label []byte

	// The text to be displayed in the input area when "text" is empty.
	placeholder []byte

	// The label color.
	labelColor tcell.Color

	// The label color when focused.
	labelColorFocused tcell.Color

	// The background color of the input area.
	fieldBackgroundColor tcell.Color

	// The background color of the input area when focused.
	fieldBackgroundColorFocused tcell.Color

	// The text color of the input area.
	fieldTextColor tcell.Color

	// The text color of the input area when focused.
	fieldTextColorFocused tcell.Color

	// The text color of the placeholder.
	placeholderTextColor tcell.Color

	// The text color of the placeholder when focused.
	placeholderTextColorFocused tcell.Color

	// The text color of the list items.
	autocompleteListTextColor tcell.Color

	// The background color of the autocomplete list.
	autocompleteListBackgroundColor tcell.Color

	// The text color of the selected ListItem.
	autocompleteListSelectedTextColor tcell.Color

	// The background color of the selected ListItem.
	autocompleteListSelectedBackgroundColor tcell.Color

	// The text color of the suggestion.
	autocompleteSuggestionTextColor tcell.Color

	// The text color of the note below the input field.
	fieldNoteTextColor tcell.Color

	// The note to show below the input field.
	fieldNote []byte

	// The screen width of the label area. A value of 0 means use the width of
	// the label text.
	labelWidth int

	// The screen width of the input area. A value of 0 means extend as much as
	// possible.
	fieldWidth int

	// A character to mask entered text (useful for password fields). A value of 0
	// disables masking.
	maskCharacter rune

	// The cursor position as a byte index into the text string.
	cursorPos int

	// An optional autocomplete function which receives the current text of the
	// input field and returns a slice of ListItems to be displayed in a drop-down
	// selection. Items' main text is displayed in the autocomplete list. When
	// set, items' secondary text is used as the selection value. Otherwise,
	// the main text is used.
	autocomplete func(text string) []*ListItem

	// The List object which shows the selectable autocomplete entries. If not
	// nil, the list's main texts represent the current autocomplete entries.
	autocompleteList *List

	// The suggested completion of the current autocomplete ListItem.
	autocompleteListSuggestion []byte

	// An optional function which may reject the last character that was entered.
	accept func(text string, ch rune) bool

	// An optional function which is called when the input has changed.
	changed func(text string)

	// An optional function which is called when the user indicated that they
	// are done entering text. The key which was pressed is provided (tab,
	// shift-tab, enter, or escape).
	done func(tcell.Key)

	// A callback function set by the Form class and called when the user leaves
	// this form item.
	finished func(tcell.Key)

	// The x-coordinate of the input field as determined during the last call to Draw().
	fieldX int

	// The number of bytes of the text string skipped ahead while drawing.
	offset int

	mu sync.RWMutex
}

// NewInputField returns a new input field.
func NewInputField() *InputField {
	return &InputField{
		box:                                     NewBox(),
		labelColor:                              Styles.SecondaryTextColor,
		fieldBackgroundColor:                    Styles.MoreContrastBackgroundColor,
		fieldBackgroundColorFocused:             Styles.ContrastBackgroundColor,
		fieldTextColor:                          Styles.PrimaryTextColor,
		fieldTextColorFocused:                   Styles.PrimaryTextColor,
		placeholderTextColor:                    Styles.ContrastSecondaryTextColor,
		autocompleteListTextColor:               Styles.PrimitiveBackgroundColor,
		autocompleteListBackgroundColor:         Styles.MoreContrastBackgroundColor,
		autocompleteListSelectedTextColor:       Styles.PrimitiveBackgroundColor,
		autocompleteListSelectedBackgroundColor: Styles.PrimaryTextColor,
		autocompleteSuggestionTextColor:         Styles.ContrastSecondaryTextColor,
		fieldNoteTextColor:                      Styles.SecondaryTextColor,
		labelColorFocused:                       ColorUnset,
		placeholderTextColorFocused:             ColorUnset,
	}
}

///////////////////////////////////// <MUTEX> ///////////////////////////////////

func (i *InputField) set(setter func(i *InputField)) *InputField {
	i.mu.Lock()
	setter(i)
	i.mu.Unlock()
	return i
}

func (i *InputField) get(getter func(i *InputField)) {
	i.mu.RLock()
	getter(i)
	i.mu.RUnlock()
}

///////////////////////////////////// <BOX> ////////////////////////////////////

// GetTitle returns the title of this InputField.
func (i *InputField) GetTitle() string {
	return i.box.GetTitle()
}

// SetTitle sets the title of this InputField.
func (i *InputField) SetTitle(title string) *InputField {
	i.box.SetTitle(title)
	return i
}

// GetTitleAlign returns the title alignment of this InputField.
func (i *InputField) GetTitleAlign() int {
	return i.box.GetTitleAlign()
}

// SetTitleAlign sets the title alignment of this InputField.
func (i *InputField) SetTitleAlign(align int) *InputField {
	i.box.SetTitleAlign(align)
	return i
}

// GetBorder returns whether this InputField has a border.
func (i *InputField) GetBorder() bool {
	return i.box.GetBorder()
}

// SetBorder sets whether this InputField has a border.
func (i *InputField) SetBorder(show bool) *InputField {
	i.box.SetBorder(show)
	return i
}

// GetBorderColor returns the border color of this InputField.
func (i *InputField) GetBorderColor() tcell.Color {
	return i.box.GetBorderColor()
}

// SetBorderColor sets the border color of this InputField.
func (i *InputField) SetBorderColor(color tcell.Color) *InputField {
	i.box.SetBorderColor(color)
	return i
}

// GetBorderAttributes returns the border attributes of this InputField.
func (i *InputField) GetBorderAttributes() tcell.AttrMask {
	return i.box.GetBorderAttributes()
}

// SetBorderAttributes sets the border attributes of this InputField.
func (i *InputField) SetBorderAttributes(attr tcell.AttrMask) *InputField {
	i.box.SetBorderAttributes(attr)
	return i
}

// GetBorderColorFocused returns the border color of this InputField when focusei.
func (i *InputField) GetBorderColorFocused() tcell.Color {
	return i.box.GetBorderColorFocused()
}

// SetBorderColorFocused sets the border color of this InputField when focusei.
func (i *InputField) SetBorderColorFocused(color tcell.Color) *InputField {
	i.box.SetBorderColorFocused(color)
	return i
}

// GetTitleColor returns the title color of this InputField.
func (i *InputField) GetTitleColor() tcell.Color {
	return i.box.GetTitleColor()
}

// SetTitleColor sets the title color of this InputField.
func (i *InputField) SetTitleColor(color tcell.Color) *InputField {
	i.box.SetTitleColor(color)
	return i
}

// GetDrawFunc returns the custom draw function of this InputField.
func (i *InputField) GetDrawFunc() func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
	return i.box.GetDrawFunc()
}

// SetDrawFunc sets a custom draw function for this InputField.
func (i *InputField) SetDrawFunc(handler func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)) *InputField {
	i.box.SetDrawFunc(handler)
	return i
}

// ShowFocus sets whether this InputField should show a focus indicator when focusei.
func (i *InputField) ShowFocus(showFocus bool) *InputField {
	i.box.ShowFocus(showFocus)
	return i
}

// GetMouseCapture returns the mouse capture function of this InputField.
func (i *InputField) GetMouseCapture() func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse) {
	return i.box.GetMouseCapture()
}

// SetMouseCapture sets a mouse capture function for this InputField.
func (i *InputField) SetMouseCapture(capture func(action MouseAction, event *tcell.EventMouse) (MouseAction, *tcell.EventMouse)) *InputField {
	i.box.SetMouseCapture(capture)
	return i
}

// GetBackgroundColor returns the background color of this InputField.
func (i *InputField) GetBackgroundColor() tcell.Color {
	return i.box.GetBackgroundColor()
}

// SetBackgroundColor sets the background color of this InputField.
func (i *InputField) SetBackgroundColor(color tcell.Color) *InputField {
	i.box.SetBackgroundColor(color)
	return i
}

// GetBackgroundTransparent returns whether the background of this InputField is transparent.
func (i *InputField) GetBackgroundTransparent() bool {
	return i.box.GetBackgroundTransparent()
}

// SetBackgroundTransparent sets whether the background of this InputField is transparent.
func (i *InputField) SetBackgroundTransparent(transparent bool) *InputField {
	i.box.SetBackgroundTransparent(transparent)
	return i
}

// GetInputCapture returns the input capture function of this InputField.
func (i *InputField) GetInputCapture() func(event *tcell.EventKey) *tcell.EventKey {
	return i.box.GetInputCapture()
}

// SetInputCapture sets a custom input capture function for this InputField.
func (i *InputField) SetInputCapture(capture func(event *tcell.EventKey) *tcell.EventKey) *InputField {
	i.box.SetInputCapture(capture)
	return i
}

// GetPadding returns the padding of this InputField.
func (i *InputField) GetPadding() (top, bottom, left, right int) {
	return i.box.GetPadding()
}

// SetPadding sets the padding of this InputField.
func (i *InputField) SetPadding(top, bottom, left, right int) *InputField {
	i.box.SetPadding(top, bottom, left, right)
	return i
}

// InRect returns whether the given screen coordinates are within this InputField.
func (i *InputField) InRect(x, y int) bool {
	return i.box.InRect(x, y)
}

// GetInnerRect returns the inner rectangle of this InputField.
func (i *InputField) GetInnerRect() (x, y, width, height int) {
	return i.box.GetInnerRect()
}

// WrapInputHandler wraps the provided input handler function such that
// input capture and other processing of the InputField is preservei.
func (i *InputField) WrapInputHandler(inputHandler func(event *tcell.EventKey, setFocus func(p Widget))) func(event *tcell.EventKey, setFocus func(p Widget)) {
	return i.box.WrapInputHandler(inputHandler)
}

// WrapMouseHandler wraps the provided mouse handler function such that
// mouse capture and other processing of the InputField is preservei.
func (i *InputField) WrapMouseHandler(mouseHandler func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget)) func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return i.box.WrapMouseHandler(mouseHandler)
}

// GetRect returns the rectangle occupied by this InputField.
func (i *InputField) GetRect() (x, y, width, height int) {
	return i.box.GetRect()
}

// SetRect sets the rectangle occupied by this InputField.
func (i *InputField) SetRect(x, y, width, height int) {
	i.box.SetRect(x, y, width, height)
}

// GetVisible returns whether this InputField is visible.
func (i *InputField) GetVisible() bool {
	return i.box.GetVisible()
}

// SetVisible sets whether this InputField is visible.
func (i *InputField) SetVisible(visible bool) {
	i.box.SetVisible(visible)
}

// Focus is called when this InputField receives focus.
func (i *InputField) Focus(delegate func(p Widget)) {
	i.box.Focus(delegate)
}

// HasFocus returns whether this InputField has focus.
func (i *InputField) HasFocus() bool {
	return i.box.HasFocus()
}

// GetFocusable returns the focusable primitive of this InputField.
func (i *InputField) GetFocusable() Focusable {
	return i.box.GetFocusable()
}

// Blur is called when this InputField loses focus.
func (i *InputField) Blur() {
	i.box.Blur()
}

// SetText sets the current text of the input field.
func (i *InputField) SetText(text string) *InputField {
	i.mu.Lock()
	i.text = []byte(text)
	i.cursorPos = len(text)
	if i.changed != nil {
		i.mu.Unlock()
		i.changed(text)
	} else {
		i.mu.Unlock()
	}
	return i
}

// GetText returns the current text of the input field.
func (i *InputField) GetText() (text string) {
	i.get(func(i *InputField) { text = string(i.text) })
	return
}

// SetLabel sets the text to be displayed before the input area.
func (i *InputField) SetLabel(label string) *InputField {
	return i.set(func(i *InputField) { i.label = []byte(label) })
}

// GetLabel returns the text to be displayed before the input area.
func (i *InputField) GetLabel() (label string) {
	i.get(func(i *InputField) { label = string(i.label) })
	return
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (i *InputField) SetLabelWidth(width int) *InputField {
	return i.set(func(i *InputField) { i.labelWidth = width })
}

// SetPlaceholder sets the text to be displayed when the input text is empty.
func (i *InputField) SetPlaceholder(text string) *InputField {
	return i.set(func(i *InputField) { i.placeholder = []byte(text) })
}

// SetLabelColor sets the color of the label.
func (i *InputField) SetLabelColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.labelColor = color })
}

// SetLabelColorFocused sets the color of the label when focused.
func (i *InputField) SetLabelColorFocused(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.labelColorFocused = color })
}

// SetFieldBackgroundColor sets the background color of the input area.
func (i *InputField) SetFieldBackgroundColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.fieldBackgroundColor = color })
}

// SetFieldBackgroundColorFocused sets the background color of the input area
// when focused.
func (i *InputField) SetFieldBackgroundColorFocused(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.fieldBackgroundColorFocused = color })
}

// SetFieldTextColor sets the text color of the input area.
func (i *InputField) SetFieldTextColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.fieldTextColor = color })
}

// SetFieldTextColorFocused sets the text color of the input area when focused.
func (i *InputField) SetFieldTextColorFocused(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.fieldTextColorFocused = color })
}

// SetPlaceholderTextColor sets the text color of placeholder text.
func (i *InputField) SetPlaceholderTextColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.placeholderTextColor = color })
}

// SetPlaceholderTextColorFocused sets the text color of placeholder text when
// focused.
func (i *InputField) SetPlaceholderTextColorFocused(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.placeholderTextColorFocused = color })
}

// SetAutocompleteListTextColor sets the text color of the ListItems.
func (i *InputField) SetAutocompleteListTextColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.autocompleteListTextColor = color })
}

// SetAutocompleteListBackgroundColor sets the background color of the
// autocomplete list.
func (i *InputField) SetAutocompleteListBackgroundColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.autocompleteListBackgroundColor = color })
}

// SetAutocompleteListSelectedTextColor sets the text color of the selected
// ListItem.
func (i *InputField) SetAutocompleteListSelectedTextColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.autocompleteListSelectedTextColor = color })
}

// SetAutocompleteListSelectedBackgroundColor sets the background of the
// selected ListItem.
func (i *InputField) SetAutocompleteListSelectedBackgroundColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.autocompleteListSelectedBackgroundColor = color })
}

// SetAutocompleteSuggestionTextColor sets the text color of the autocomplete
// suggestion in the input field.
func (i *InputField) SetAutocompleteSuggestionTextColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.autocompleteSuggestionTextColor = color })
}

// SetFieldNoteTextColor sets the text color of the note.
func (i *InputField) SetFieldNoteTextColor(color tcell.Color) *InputField {
	return i.set(func(i *InputField) { i.fieldNoteTextColor = color })
}

// SetFieldNote sets the text to show below the input field, e.g. when the
// input is invalid.
func (i *InputField) SetFieldNote(note string) *InputField {
	return i.set(func(i *InputField) { i.fieldNote = []byte(note) })
}

// ResetFieldNote sets the note to an empty string.
func (i *InputField) ResetFieldNote() *InputField {
	return i.set(func(i *InputField) { i.fieldNote = nil })
}

// SetFieldWidth sets the screen width of the input area. A value of 0 means
// extend as much as possible.
func (i *InputField) SetFieldWidth(width int) *InputField {
	return i.set(func(i *InputField) { i.fieldWidth = width })
}

// GetFieldWidth returns this primitive's field width.
func (i *InputField) GetFieldWidth() (width int) {
	i.get(func(i *InputField) { width = i.fieldWidth })
	return
}

// GetFieldHeight returns the height of the field.
func (i *InputField) GetFieldHeight() (height int) {
	i.get(func(i *InputField) {
		if len(i.fieldNote) == 0 {
			height = 1
		} else {
			height = 2
		}
	})
	return
}

// GetCursorPosition returns the cursor position.
func (i *InputField) GetCursorPosition() (cursorPos int) {
	i.get(func(i *InputField) { cursorPos = i.cursorPos })
	return
}

// SetCursorPosition sets the cursor position.
func (i *InputField) SetCursorPosition(cursorPos int) *InputField {
	return i.set(func(i *InputField) { i.cursorPos = cursorPos })
}

// SetMaskCharacter sets a character that masks user input on a screen. A value
// of 0 disables masking.
func (i *InputField) SetMaskCharacter(mask rune) *InputField {
	return i.set(func(i *InputField) { i.maskCharacter = mask })
}

// SetAutocompleteFunc sets an autocomplete callback function which may return
// ListItems to be selected from a drop-down based on the current text of the
// input field. The drop-down appears only if len(entries) > 0. The callback is
// invoked in this function and whenever the current text changes or when
// Autocomplete() is called. Entries are cleared when the user selects an entry
// or presses Escape.
func (i *InputField) SetAutocompleteFunc(callback func(currentText string) (entries []*ListItem)) *InputField {
	i.mu.Lock()
	i.autocomplete = callback
	i.mu.Unlock()
	return i.Autocomplete()
}

// Autocomplete invokes the autocomplete callback (if there is one). If the
// length of the returned autocomplete entries slice is greater than 0, the
// input field will present the user with a corresponding drop-down list the
// next time the input field is drawn.
//
// It is safe to call this function from any goroutine. Note that the input
// field is not redrawn automatically unless called from the main goroutine
// (e.g. in response to events).
func (i *InputField) Autocomplete() *InputField {
	i.mu.Lock()
	if i.autocomplete == nil {
		i.mu.Unlock()
		return i
	}
	i.mu.Unlock()

	// Do we have any autocomplete entries?
	entries := i.autocomplete(string(i.text))
	if len(entries) == 0 {
		// No entries, no list.
		i.mu.Lock()
		i.autocompleteList = nil
		i.autocompleteListSuggestion = nil
		i.mu.Unlock()
		return i
	}

	i.mu.Lock()

	// Make a list if we have none.
	if i.autocompleteList == nil {
		l := NewList()
		l.SetChangedFunc(i.autocompleteChanged)
		l.ShowSecondaryText(false)
		l.SetMainTextColor(i.autocompleteListTextColor)
		l.SetSelectedTextColor(i.autocompleteListSelectedTextColor)
		l.SetSelectedBackgroundColor(i.autocompleteListSelectedBackgroundColor)
		l.SetHighlightFullLine(true)
		l.SetBackgroundColor(i.autocompleteListBackgroundColor)

		i.autocompleteList = l
	}

	// Fill it with the entries.
	currentEntry := -1
	i.autocompleteList.Clear()
	for index, entry := range entries {
		i.autocompleteList.AddItem(entry)
		if currentEntry < 0 && entry.GetMainText() == string(i.text) {
			currentEntry = index
		}
	}

	// Set the selection if we have one.
	if currentEntry >= 0 {
		i.autocompleteList.SetCurrentItem(currentEntry)
	}

	i.mu.Unlock()

	return i
}

// autocompleteChanged gets called when another item in the
// autocomplete list has been selected.
func (i *InputField) autocompleteChanged(_ int, item *ListItem) {
	mainText := item.GetMainBytes()
	secondaryText := item.GetSecondaryBytes()
	if len(i.text) < len(secondaryText) {
		i.autocompleteListSuggestion = secondaryText[len(i.text):]
	} else if len(i.text) < len(mainText) {
		i.autocompleteListSuggestion = mainText[len(i.text):]
	} else {
		i.autocompleteListSuggestion = nil
	}
}

// SetAcceptanceFunc sets a handler which may reject the last character that was
// entered (by returning false).
//
// This package defines a number of variables prefixed with InputField which may
// be used for common input (e.g. numbers, maximum text length).
func (i *InputField) SetAcceptanceFunc(handler func(textToCheck string, lastChar rune) bool) *InputField {
	return i.set(func(i *InputField) { i.accept = handler })
}

// SetChangedFunc sets a handler which is called whenever the text of the input
// field has changed. It receives the current text (after the change).
func (i *InputField) SetChangedFunc(handler func(text string)) *InputField {
	return i.set(func(i *InputField) { i.changed = handler })
}

// SetDoneFunc setFocus(s a handler which is called when the user is done entering
// text. The callback function is provided with the key that was pressed, which
// is one of the following:
//
//   - KeyEnter: Done entering text.
//   - KeyEscape: Abort text input.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (i *InputField) SetDoneFunc(handler func(key tcell.Key)) *InputField {
	return i.set(func(i *InputField) { i.done = handler })
}

// SetFinishedFunc sets a callback invoked when the user leaves this form item.
func (i *InputField) SetFinishedFunc(handler func(key tcell.Key)) *InputField {
	return i.set(func(i *InputField) { i.finished = handler })
}

// Draw draws this primitive onto the screen.
func (i *InputField) Draw(screen tcell.Screen) {
	if !i.GetVisible() {
		return
	}

	i.box.Draw(screen)

	i.mu.Lock()
	defer i.mu.Unlock()

	// Select colors
	labelColor := i.labelColor
	fieldBackgroundColor := i.fieldBackgroundColor
	fieldTextColor := i.fieldTextColor
	if i.GetFocusable().HasFocus() {
		if i.labelColorFocused != ColorUnset {
			labelColor = i.labelColorFocused
		}
		if i.fieldBackgroundColorFocused != ColorUnset {
			fieldBackgroundColor = i.fieldBackgroundColorFocused
		}
		if i.fieldTextColorFocused != ColorUnset {
			fieldTextColor = i.fieldTextColorFocused
		}
	}

	// Prepare
	x, y, width, height := i.GetInnerRect()
	rightLimit := x + width
	if height < 1 || rightLimit <= x {
		return
	}

	// Draw label.
	if i.labelWidth > 0 {
		labelWidth := i.labelWidth
		if labelWidth > rightLimit-x {
			labelWidth = rightLimit - x
		}
		Print(screen, i.label, x, y, labelWidth, AlignLeft, labelColor)
		x += labelWidth
	} else {
		_, drawnWidth := Print(screen, i.label, x, y, rightLimit-x, AlignLeft, labelColor)
		x += drawnWidth
	}

	// Draw input area.
	i.fieldX = x
	fieldWidth := i.fieldWidth
	if fieldWidth == 0 {
		fieldWidth = math.MaxInt32
	}
	if rightLimit-x < fieldWidth {
		fieldWidth = rightLimit - x
	}
	fieldStyle := tcell.StyleDefault.Background(fieldBackgroundColor)
	for index := 0; index < fieldWidth; index++ {
		screen.SetContent(x+index, y, ' ', nil, fieldStyle)
	}

	// Text.
	var cursorScreenPos int
	text := i.text
	if len(text) == 0 && len(i.placeholder) > 0 {
		// Draw placeholder text.
		placeholderTextColor := i.placeholderTextColor
		if i.GetFocusable().HasFocus() && i.placeholderTextColorFocused != ColorUnset {
			placeholderTextColor = i.placeholderTextColorFocused
		}
		Print(screen, EscapeBytes(i.placeholder), x, y, fieldWidth, AlignLeft, placeholderTextColor)
		i.offset = 0
	} else {
		// Draw entered text.
		if i.maskCharacter > 0 {
			text = bytes.Repeat([]byte(string(i.maskCharacter)), utf8.RuneCount(i.text))
		}
		var drawnText []byte
		if fieldWidth > runewidth.StringWidth(string(text)) {
			// We have enough space for the full text.
			drawnText = EscapeBytes(text)
			Print(screen, drawnText, x, y, fieldWidth, AlignLeft, fieldTextColor)
			i.offset = 0
			iterateString(string(text), func(main rune, comb []rune, textPos, textWidth, screenPos, screenWidth int) bool {
				if textPos >= i.cursorPos {
					return true
				}
				cursorScreenPos += screenWidth
				return false
			})
		} else {
			// The text doesn't fit. Where is the cursor?
			if i.cursorPos < 0 {
				i.cursorPos = 0
			} else if i.cursorPos > len(text) {
				i.cursorPos = len(text)
			}
			// Shift the text so the cursor is inside the field.
			var shiftLeft int
			if i.offset > i.cursorPos {
				i.offset = i.cursorPos
			} else if subWidth := runewidth.StringWidth(string(text[i.offset:i.cursorPos])); subWidth > fieldWidth-1 {
				shiftLeft = subWidth - fieldWidth + 1
			}
			currentOffset := i.offset
			iterateString(string(text), func(main rune, comb []rune, textPos, textWidth, screenPos, screenWidth int) bool {
				if textPos >= currentOffset {
					if shiftLeft > 0 {
						i.offset = textPos + textWidth
						shiftLeft -= screenWidth
					} else {
						if textPos+textWidth > i.cursorPos {
							return true
						}
						cursorScreenPos += screenWidth
					}
				}
				return false
			})
			drawnText = EscapeBytes(text[i.offset:])
			Print(screen, drawnText, x, y, fieldWidth, AlignLeft, fieldTextColor)
		}
		// Draw suggestion
		if i.maskCharacter == 0 && len(i.autocompleteListSuggestion) > 0 {
			Print(screen, i.autocompleteListSuggestion, x+runewidth.StringWidth(string(drawnText)), y, fieldWidth-runewidth.StringWidth(string(drawnText)), AlignLeft, i.autocompleteSuggestionTextColor)
		}
	}

	// Draw field note
	if len(i.fieldNote) > 0 {
		Print(screen, i.fieldNote, x, y+1, fieldWidth, AlignLeft, i.fieldNoteTextColor)
	}

	// Draw autocomplete list.
	if i.autocompleteList != nil {
		// How much space do we need?
		lheight := i.autocompleteList.GetItemCount()
		lwidth := 0
		for index := 0; index < lheight; index++ {
			entry, _ := i.autocompleteList.GetItemText(index)
			width := TaggedStringWidth(entry)
			if width > lwidth {
				lwidth = width
			}
		}

		// We prefer to drop down but if there is no space, maybe drop up?
		lx := x
		ly := y + 1
		_, sheight := screen.Size()
		if ly+lheight >= sheight && ly-2 > lheight-ly {
			ly = y - lheight
			if ly < 0 {
				ly = 0
			}
		}
		if ly+lheight >= sheight {
			lheight = sheight - ly
		}
		if i.autocompleteList.scrollBarVisibility == ScrollBarAlways || (i.autocompleteList.scrollBarVisibility == ScrollBarAuto && i.autocompleteList.GetItemCount() > lheight) {
			lwidth++ // Add space for scroll bar
		}
		i.autocompleteList.SetRect(lx, ly, lwidth, lheight)
		i.autocompleteList.Draw(screen)
	}

	// Set cursor.
	if i.box.focus.HasFocus() {
		screen.ShowCursor(x+cursorScreenPos, y)
	}
}

// InputHandler returns the handler for this primitive.
func (i *InputField) InputHandler() func(event *tcell.EventKey, setFocus func(p Widget)) {
	return i.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Widget)) {
		i.mu.Lock()

		// Trigger changed events.
		currentText := i.text
		defer func() {
			i.mu.Lock()
			newText := i.text
			i.mu.Unlock()

			if !bytes.Equal(newText, currentText) {
				i.Autocomplete()
				if i.changed != nil {
					i.changed(string(i.text))
				}
			}
		}()

		// Movement functions.
		home := func() { i.cursorPos = 0 }
		end := func() { i.cursorPos = len(i.text) }
		moveLeft := func() {
			iterateStringReverse(string(i.text[:i.cursorPos]), func(main rune, comb []rune, textPos, textWidth, screenPos, screenWidth int) bool {
				i.cursorPos -= textWidth
				return true
			})
		}
		moveRight := func() {
			iterateString(string(i.text[i.cursorPos:]), func(main rune, comb []rune, textPos, textWidth, screenPos, screenWidth int) bool {
				i.cursorPos += textWidth
				return true
			})
		}
		moveWordLeft := func() {
			i.cursorPos = len(regexRightWord.ReplaceAll(i.text[:i.cursorPos], nil))
		}
		moveWordRight := func() {
			i.cursorPos = len(i.text) - len(regexLeftWord.ReplaceAll(i.text[i.cursorPos:], nil))
		}

		// Add character function. Returns whether the rune character is
		// accepted.
		add := func(r rune) bool {
			newText := append(append(i.text[:i.cursorPos], []byte(string(r))...), i.text[i.cursorPos:]...)
			if i.accept != nil && !i.accept(string(newText), r) {
				return false
			}
			i.text = newText
			i.cursorPos += len(string(r))
			return true
		}

		// Finish up.
		finish := func(key tcell.Key) {
			if i.done != nil {
				i.done(key)
			}
			if i.finished != nil {
				i.finished(key)
			}
		}

		// Process key event.
		switch key := event.Key(); key {
		case tcell.KeyRune: // Regular character.
			if event.Modifiers()&tcell.ModAlt > 0 {
				// We accept some Alt-key combinations.
				switch event.Rune() {
				case 'a': // Home.
					home()
				case 'e': // End.
					end()
				case 'b': // Move word left.
					moveWordLeft()
				case 'f': // Move word right.
					moveWordRight()
				default:
					if !add(event.Rune()) {
						i.mu.Unlock()
						return
					}
				}
			} else {
				// Other keys are simply accepted as regular characters.
				if !add(event.Rune()) {
					i.mu.Unlock()
					return
				}
			}
		case tcell.KeyCtrlU: // Delete all.
			i.text = nil
			i.cursorPos = 0
		case tcell.KeyCtrlK: // Delete until the end of the line.
			i.text = i.text[:i.cursorPos]
		case tcell.KeyCtrlW: // Delete last word.
			newText := append(regexRightWord.ReplaceAll(i.text[:i.cursorPos], nil), i.text[i.cursorPos:]...)
			i.cursorPos -= len(i.text) - len(newText)
			i.text = newText
		case tcell.KeyBackspace, tcell.KeyBackspace2: // Delete character before the cursor.
			iterateStringReverse(string(i.text[:i.cursorPos]), func(main rune, comb []rune, textPos, textWidth, screenPos, screenWidth int) bool {
				i.text = append(i.text[:textPos], i.text[textPos+textWidth:]...)
				i.cursorPos -= textWidth
				return true
			})
			if i.offset >= i.cursorPos {
				i.offset = 0
			}
		case tcell.KeyDelete: // Delete character after the cursor.
			iterateString(string(i.text[i.cursorPos:]), func(main rune, comb []rune, textPos, textWidth, screenPos, screenWidth int) bool {
				i.text = append(i.text[:i.cursorPos], i.text[i.cursorPos+textWidth:]...)
				return true
			})
		case tcell.KeyLeft:
			if event.Modifiers()&tcell.ModAlt > 0 {
				moveWordLeft()
			} else {
				moveLeft()
			}
		case tcell.KeyRight:
			if event.Modifiers()&tcell.ModAlt > 0 {
				moveWordRight()
			} else {
				moveRight()
			}
		case tcell.KeyHome, tcell.KeyCtrlA:
			home()
		case tcell.KeyEnd, tcell.KeyCtrlE:
			end()
		case tcell.KeyEnter: // We might be done.
			if i.autocompleteList != nil {
				currentItem := i.autocompleteList.GetCurrentItem()
				selectionText := currentItem.GetMainText()
				if currentItem.GetSecondaryText() != "" {
					selectionText = currentItem.GetSecondaryText()
				}
				i.mu.Unlock()
				i.SetText(selectionText)
				i.mu.Lock()
				i.autocompleteList = nil
				i.autocompleteListSuggestion = nil
				i.mu.Unlock()
			} else {
				i.mu.Unlock()
				finish(key)
			}
			return
		case tcell.KeyEscape:
			if i.autocompleteList != nil {
				i.autocompleteList = nil
				i.autocompleteListSuggestion = nil
				i.mu.Unlock()
			} else {
				i.mu.Unlock()
				finish(key)
			}
			return
		case tcell.KeyDown, tcell.KeyTab: // Autocomplete selection.
			if i.autocompleteList != nil {
				count := i.autocompleteList.GetItemCount()
				newEntry := i.autocompleteList.GetCurrentItemIndex() + 1
				if newEntry >= count {
					newEntry = 0
				}
				i.autocompleteList.SetCurrentItem(newEntry)
				i.mu.Unlock()
			} else {
				i.mu.Unlock()
				finish(key)
			}
			return
		case tcell.KeyUp, tcell.KeyBacktab: // Autocomplete selection.
			if i.autocompleteList != nil {
				newEntry := i.autocompleteList.GetCurrentItemIndex() - 1
				if newEntry < 0 {
					newEntry = i.autocompleteList.GetItemCount() - 1
				}
				i.autocompleteList.SetCurrentItem(newEntry)
				i.mu.Unlock()
			} else {
				i.mu.Unlock()
				finish(key)
			}
			return
		}

		i.mu.Unlock()
	})
}

// MouseHandler returns the mouse handler for this primitive.
func (i *InputField) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
	return i.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Widget)) (consumed bool, capture Widget) {
		x, y := event.Position()
		_, rectY, _, _ := i.GetInnerRect()
		if !i.InRect(x, y) {
			return false, nil
		}

		// Process mouse event.
		if action == MouseLeftClick && y == rectY {
			// Determine where to place the cursor.
			if x >= i.fieldX {
				if !iterateString(string(i.text), func(main rune, comb []rune, textPos int, textWidth int, screenPos int, screenWidth int) bool {
					if x-i.fieldX < screenPos+screenWidth {
						i.cursorPos = textPos
						return true
					}
					return false
				}) {
					i.cursorPos = len(i.text)
				}
			}
			setFocus(i)
			consumed = true
		}

		return
	})
}

var (
	regexRightWord = regexp.MustCompile(`(\w*|\W)$`)
	regexLeftWord  = regexp.MustCompile(`^(\W|\w*)`)
)
