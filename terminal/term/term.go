package term

import (
	"fmt"
	"io"
	"os"
)

type Term struct {
	fields
	stdin        *os.File
	stdout       *os.File
	wrapStdin    io.Reader
	wrapStdout   io.Writer
	onSizeChange func(cols, rows uint16)
	onExit       func()
}

// Open a terminal connected to the given stdin & stdout.
// onSizeChange is called when the terminal window size changes.
func Open(stdin, stdout *os.File, ignore_error bool, onSizeChange func(cols, rows uint16)) (t *Term, err error) {
	if stdin == nil {
		err = fmt.Errorf("invalid stdin")
		return
	}
	if stdout == nil {
		err = fmt.Errorf("invalid stdout")
		return
	}

	t = &Term{
		stdin:        stdin,
		stdout:       stdout,
		onSizeChange: onSizeChange,
	}
	if !ignore_error {
		if !t.isTerminal() {
			err = fmt.Errorf("not terminal")
			return
		}
	}

	err = t.wrapStdInOut()
	if !ignore_error {
		if err != nil {
			return
		}
	}
	if t.wrapStdout == nil {
		t.wrapStdout = stdout
		err = nil
	}

	t.onExit = t.captureSizeChangeEvent(t.onSizeChange)
	return
}

// Close the terminal and restore the original status.
func (t *Term) Close() {
	t.restore()
	if t.onExit != nil {
		t.onExit()
	}
}

// Read from Term.
func (t *Term) Read(b []byte) (n int, err error) {
	return t.wrapStdin.Read(b)
}

// Write to Term.
func (t *Term) Write(b []byte) (n int, err error) {
	return t.wrapStdout.Write(b)
}
