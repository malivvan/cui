//go:build windows

package terminal

import "syscall"

var sysProcAttr = &syscall.SysProcAttr{}
