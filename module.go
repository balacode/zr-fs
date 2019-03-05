// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-03-04 20:18:51 491D56                              zr-fs/[module.go]
// -----------------------------------------------------------------------------

// Package fs implements file-system related functions
package fs

import (
	"fmt"

	"github.com/balacode/zr"
)

// LB specifies a line break string.
// On Windows it is a pair of CR and LF.
// CR is decimal 13, hex 0D.
// LF is decimal 10, hex 0A.
const LB = "\r\n"

// LF specifies a line feed string ("\n").
const LF = "\n"

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// VL is zr.VerboseLog() but is used only for debugging.
var VL = zr.VerboseLog

// Callers returns a human-friendly string showing the call stack with each
// callers calling method or function's name and line number. The most
// immediate are shown first, followed by their callers, and so on.
// For brevity, 'runtime.*', 'syscall.*' and 'testing.*'
// top-level callers are not included.
//
// func Callers(options ...interface{}) string
var callers = zr.Callers

// Error outputs an error message to the standard output and to a
// log file named 'run.log' saved in the program's current directory,
// It also outputs the call stack (names and line numbers of callers.)
// Returns an error value initialized with the message.
var myError = zr.Error

//end
