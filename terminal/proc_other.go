//go:build !windows

package terminal

import "syscall"

var sysProcAttr = &syscall.SysProcAttr{
	Setsid:  true,
	Setctty: true,
	Ctty:    1,
}
