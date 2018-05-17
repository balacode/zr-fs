// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-17 10:44:18 E96906              [zr-fs/dir_watcher_nonwindows.go]
// -----------------------------------------------------------------------------
// +build !windows

package fs

import (
	"github.com/balacode/zr" // Zircon-Go
)

// waitForDirChange __
func waitForDirChange(c chan string, dir string) {
	zr.IMPLEMENT()
} //                                                            waitForDirChange

//end
