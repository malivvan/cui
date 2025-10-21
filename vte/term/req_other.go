//go:build linux || solaris
// +build linux solaris

package term

import "github.com/malivvan/cui/vte/term/export"

const reqGetTermios = export.TCGETS
const reqSetTermios = export.TCSETS
