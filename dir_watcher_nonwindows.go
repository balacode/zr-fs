// -----------------------------------------------------------------------------
// ZR Library - File System Package            zr-fs/[dir_watcher_nonwindows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------
// +build !windows

package fs

import (
	"github.com/balacode/zr"
)

// waitForDirChange _ _
func waitForDirChange(c chan string, dir string) {
	zr.IMPLEMENT()
} //                                                            waitForDirChange

//end
