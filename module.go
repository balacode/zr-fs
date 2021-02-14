// -----------------------------------------------------------------------------
// ZR Library - File System Package                            zr-fs/[module.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

// Package fs implements file-system related functions
package fs

import (
	"fmt"

	"github.com/balacode/zr"
)

var (
	// PL is fmt.Println() but is used only for debugging.
	PL = fmt.Println

	// VL is zr.VerboseLog() but is used only for debugging.
	VL = zr.VerboseLog

	// Callers returns a human-friendly string showing the call stack with
	// each calling method or function's name and line number. The most
	// immediate are shown first, followed by their callers, and so on.
	// For brevity, 'runtime.*', 'syscall.*' and 'testing.*'
	// top-level callers are not included.
	//
	// func Callers(options ...interface{}) string
	callers = zr.Callers

	// Error outputs an error message to the standard output and to a
	// log file named "<process>.log" in the program's current directory,
	// It also outputs the call stack (names and line numbers of callers.)
	// Returns an error value initialized with the message.
	myError = zr.Error
)

// end
