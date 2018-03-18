// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-18 23:19:27 53A1FC              [zr/fs/dir_watcher_nonwindows.go]
// -----------------------------------------------------------------------------
// +build !windows

package fs

import "github.com/balacode/zr" // Zircon-Go

// waitForDirChange __
func waitForDirChange(c chan string, dir string) {
	zr.IMPLEMENT()
} //                                                            waitForDirChange

//end
