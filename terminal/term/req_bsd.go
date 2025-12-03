//go:build darwin || dragonfly || freebsd || netbsd || openbsd
// +build darwin dragonfly freebsd netbsd openbsd

package term

import "github.com/malivvan/cui/terminal/term/export"

const reqGetTermios = export.TIOCGETA
const reqSetTermios = export.TIOCSETA
