package cui

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/gdamore/tcell/v2"
)

type testCase struct {
	mod     tcell.ModMask
	key     tcell.Key
	ch      rune
	encoded string
}

var testCases = []testCase{
	{mod: tcell.ModNone, key: tcell.KeyRune, ch: 'a', encoded: "a"},
	{mod: tcell.ModNone, key: tcell.KeyRune, ch: '+', encoded: "+"},
	{mod: tcell.ModNone, key: tcell.KeyRune, ch: ';', encoded: ";"},
	{mod: tcell.ModNone, key: tcell.KeyTab, ch: rune(tcell.KeyTab), encoded: "Tab"},
	{mod: tcell.ModNone, key: tcell.KeyEnter, ch: rune(tcell.KeyEnter), encoded: "Enter"},
	{mod: tcell.ModNone, key: tcell.KeyPgDn, ch: 0, encoded: "PageDown"},
	{mod: tcell.ModAlt, key: tcell.KeyRune, ch: 'a', encoded: "Alt+a"},
	{mod: tcell.ModAlt, key: tcell.KeyRune, ch: '+', encoded: "Alt++"},
	{mod: tcell.ModAlt, key: tcell.KeyRune, ch: ';', encoded: "Alt+;"},
	{mod: tcell.ModAlt, key: tcell.KeyRune, ch: ' ', encoded: "Alt+Space"},
	{mod: tcell.ModAlt, key: tcell.KeyRune, ch: '1', encoded: "Alt+1"},
	{mod: tcell.ModAlt, key: tcell.KeyTab, ch: rune(tcell.KeyTab), encoded: "Alt+Tab"},
	{mod: tcell.ModAlt, key: tcell.KeyEnter, ch: rune(tcell.KeyEnter), encoded: "Alt+Enter"},
	{mod: tcell.ModAlt, key: tcell.KeyBackspace2, ch: rune(tcell.KeyBackspace2), encoded: "Alt+Backspace"},
	{mod: tcell.ModCtrl, key: tcell.KeyCtrlC, ch: rune(tcell.KeyCtrlC), encoded: "Ctrl+C"},
	{mod: tcell.ModCtrl, key: tcell.KeyCtrlD, ch: rune(tcell.KeyCtrlD), encoded: "Ctrl+D"},
	{mod: tcell.ModCtrl, key: tcell.KeyCtrlSpace, ch: rune(tcell.KeyCtrlSpace), encoded: "Ctrl+Space"},
	{mod: tcell.ModCtrl, key: tcell.KeyCtrlRightSq, ch: rune(tcell.KeyCtrlRightSq), encoded: "Ctrl+]"},
	{mod: tcell.ModCtrl | tcell.ModAlt, key: tcell.KeyRune, ch: '+', encoded: "Ctrl+Alt++"},
	{mod: tcell.ModCtrl | tcell.ModShift, key: tcell.KeyRune, ch: '+', encoded: "Ctrl+Shift++"},
}

func TestEncode(t *testing.T) {
	t.Parallel()

	for _, c := range testCases {
		encoded, err := BindEncode(c.mod, c.key, c.ch)
		if err != nil {
			t.Errorf("failed to encode key %d %d %d: %s", c.mod, c.key, c.ch, err)
		}
		if encoded != c.encoded {
			t.Errorf("failed to encode key %d %d %d: got %s, want %s", c.mod, c.key, c.ch, encoded, c.encoded)
		}
	}
}

func TestDecode(t *testing.T) {
	t.Parallel()

	for _, c := range testCases {
		mod, key, ch, err := BindDecode(c.encoded)
		if err != nil {
			t.Errorf("failed to decode key %s: %s", c.encoded, err)
		}
		if mod != c.mod {
			t.Errorf("failed to decode key %s: invalid modifiers: got %d, want %d", c.encoded, mod, c.mod)
		}
		if key != c.key {
			t.Errorf("failed to decode key %s: invalid key: got %d, want %d", c.encoded, key, c.key)
		}
		if ch != c.ch {
			t.Errorf("failed to decode key %s: invalid rune: got %d, want %d", c.encoded, ch, c.ch)
		}
	}
}

const pressTimes = 7

func TestConfiguration(t *testing.T) {
	t.Parallel()

	wg := make([]*sync.WaitGroup, len(testCases))

	config := NewBindConfig()
	for i, c := range testCases {
		wg[i] = new(sync.WaitGroup)
		wg[i].Add(pressTimes)

		i := i // Capture
		if c.key != tcell.KeyRune {
			config.SetKey(c.mod, c.key, func(ev *tcell.EventKey) *tcell.EventKey {
				wg[i].Done()
				return nil
			})
		} else {
			config.SetRune(c.mod, c.ch, func(ev *tcell.EventKey) *tcell.EventKey {
				wg[i].Done()
				return nil
			})
		}

	}

	done := make(chan struct{})
	timeout := time.After(5 * time.Second)

	go func() {
		for i := range testCases {
			wg[i].Wait()
		}

		done <- struct{}{}
	}()

	errs := make(chan error)
	for j := 0; j < pressTimes; j++ {
		for i, c := range testCases {
			i, c := i, c // Capture
			go func() {
				k := tcell.NewEventKey(c.key, c.ch, c.mod)
				if k.Key() != c.key {
					errs <- fmt.Errorf("failed to test capturing keybinds: tcell modified EventKey.Key: expected %d, got %d", c.key, k.Key())
					return
				} else if k.Rune() != c.ch {
					errs <- fmt.Errorf("failed to test capturing keybinds: tcell modified EventKey.Rune: expected %d, got %d", c.ch, k.Rune())
					return
				} else if k.Modifiers() != c.mod {
					errs <- fmt.Errorf("failed to test capturing keybinds: tcell modified EventKey.Modifiers: expected %d, got %d", c.mod, k.Modifiers())
					return
				}

				ev := config.Capture(tcell.NewEventKey(c.key, c.ch, c.mod))
				if ev != nil {
					errs <- fmt.Errorf("failed to test capturing keybinds: failed to register case %d event %d %d %d", i, c.mod, c.key, c.ch)
				}
			}()
		}
	}

	select {
	case err := <-errs:
		t.Fatal(err)
	case <-timeout:
		t.Fatal("timeout")
	case <-done:
	}
}

// Example of creating and using an input configuration.
func ExampleNewConfiguration() {
	// Create a new input configuration to store the key bindings.
	c := NewBindConfig()

	handleSave := func(ev *tcell.EventKey) *tcell.EventKey {
		// Save
		return nil
	}

	handleOpen := func(ev *tcell.EventKey) *tcell.EventKey {
		// Open
		return nil
	}

	handleExit := func(ev *tcell.EventKey) *tcell.EventKey {
		// Exit
		return nil
	}

	// Bind Alt+s.
	if err := c.Set("Alt+s", handleSave); err != nil {
		log.Fatalf("failed to set keybind: %s", err)
	}

	// Bind Alt+o.
	c.SetRune(tcell.ModAlt, 'o', handleOpen)

	// Bind Escape.
	c.SetKey(tcell.ModNone, tcell.KeyEscape, handleExit)

	// Capture input. This will differ based on the framework in use (if any).
	// When using tview or cview, call Application.SetInputCapture before calling
	// Application.Run.
	// app.SetInputCapture(c.Capture)
}

// Example of capturing key events.
func ExampleConfiguration_Capture() {
	// See the end of the NewBindConfig example.
}
