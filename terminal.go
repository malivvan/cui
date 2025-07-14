package cui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/gdamore/tcell/v2/views"
	"github.com/malivvan/cui/vte"
	"os/exec"
	"sync"
)

type Terminal struct {
	*Box

	term    *vte.VT
	running bool
	cmd     *exec.Cmd
	app     *Application
	w       int
	h       int

	sync.RWMutex
}

func NewTerminal(app *Application, cmd *exec.Cmd) *Terminal {
	t := &Terminal{
		Box:  NewBox(),
		term: vte.New(),
		app:  app,
		cmd:  cmd,
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
		err := t.term.Start(t.cmd)
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
	case *vte.EventRedraw:
		go func() {
			t.app.QueueUpdateDraw(func() {})
		}()
	}
}

func (t *Terminal) InputHandler() func(event *tcell.EventKey, setFocus func(p Primitive)) {
	return t.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p Primitive)) {
		t.term.HandleEvent(event)
	})
}

func (t *Terminal) MouseHandler() func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
	return t.WrapMouseHandler(func(action MouseAction, event *tcell.EventMouse, setFocus func(p Primitive)) (consumed bool, capture Primitive) {
		if action == MouseLeftClick && t.InRect(event.Position()) {
			setFocus(t)
			return t.term.HandleEvent(event), nil
		}
		return false, nil
	})
}
