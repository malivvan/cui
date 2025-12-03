//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && !cgo

package nixpty

import (
	"os"
	"syscall"

	"github.com/malivvan/cui/terminal/pty/nixpty/native"
)

// Open returns a control pty(ptm) and the linked process tty(pts).
func open() (ptm *os.File, pts *os.File, err error) {
	ptm, err = native.Openpt(syscall.O_RDWR)
	if err != nil {
		return
	}

	err = native.Grantpt(ptm)
	if err != nil {
		ptm.Close()
		return
	}

	err = native.Unlockpt(ptm)
	if err != nil {
		ptm.Close()
		return
	}

	ptsname, err := native.Ptsname(ptm)
	if err != nil {
		ptm.Close()
		return
	}

	pts, err = os.OpenFile(ptsname, syscall.O_RDWR|syscall.O_NOCTTY|syscall.O_CLOEXEC, 0)
	return
}
