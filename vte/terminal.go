package vte

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/malivvan/cui"
	"github.com/malivvan/cui/vte/pty"
)

type Terminal struct {
	*cui.Box

	term    *VT
	running bool
	opt     pty.Options
	app     *cui.App
	w       int
	h       int

	sync.RWMutex
}

func NewTerminal(app *cui.App, opt pty.Options) *Terminal {
	t := &Terminal{
		Box:  cui.NewBox(),
		term: New(),
		app:  app,
		opt:  opt,
	}
	return t
}

func (t *Terminal) Draw(s tcell.Screen) {
	if !t.GetVisible() {
		return
	}
	t.Box.Draw(s)
	t.Lock()
	defer t.Unlock()

	x, y, w, h := t.GetInnerRect()
	view := views.NewViewPort(s, x, y, w, h)
	t.term.SetSurface(view)
	if w != t.w || h != t.h {
		t.w = w
		t.h = h
		t.term.Resize(w, h)
	}

	if !t.running {
		err := t.term.Start(t.opt)
		if err != nil {
			panic(err)
		}
		t.term.Attach(t.HandleEvent)
		t.running = true
	}
	if t.HasFocus() {
		cy, cx, style, vis := t.term.Cursor()
		if vis {
			s.ShowCursor(cx+x, cy+y)
			s.SetCursorStyle(style)
		} else {
			s.HideCursor()
		}
	}
	t.term.Draw()
}

func (t *Terminal) HandleEvent(ev tcell.Event) {
	switch ev.(type) {
	case *EventRedraw:
		go func() {
			t.app.QueueUpdateDraw(func() {})
		}()
	}
}

func (t *Terminal) InputHandler() func(event *tcell.EventKey, setFocus func(p cui.Widget)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p cui.Widget)) {
		t.term.HandleEvent(event)
	})
}

func (t *Terminal) MouseHandler() func(action cui.MouseAction, event *tcell.EventMouse, setFocus func(p cui.Widget)) (consumed bool, capture cui.Widget) {
	return t.WrapMouseHandler(func(action cui.MouseAction, event *tcell.EventMouse, setFocus func(p cui.Widget)) (consumed bool, capture cui.Widget) {
		if action == cui.MouseLeftClick && t.InRect(event.Position()) {
			setFocus(t)
			return t.term.HandleEvent(event), nil
		}
		return false, nil
	})
}
