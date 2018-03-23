// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-23 11:40:22 C6E28B                              [zr-fs/module.go]
// -----------------------------------------------------------------------------

// Package fs implements file-system related functions
package fs

import "fmt" // standard

import "github.com/balacode/zr" // Zircon-Go

// LB specifies a line break string.
// On Windows it is a pair of CR and LF.
// CR is decimal 13, hex 0D.
// LF is decimal 10, hex 0A.
const LB = "\r\n"

// LF specifies a line feed string ("\n").
const LF = "\n"

// CallTracing specifies if zr.TraceCall() should be called after
// every call to a function in this module. Used for debugging.
const CallTracing = false

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// VL is zr.VerboseLog() but is used only for debugging.
var VL = zr.VerboseLog

// // traceCall is a function that enables call tracing.
// // It not called when set to nil.
// var traceCall = func() func() {
//     if CallTracing {
//         return zr.TraceCall
//     }
//     return nil
// }()

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
