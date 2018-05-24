// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-24 03:05:23 E0B25C              [zr-fs/dir_watcher_nonwindows.go]
// -----------------------------------------------------------------------------
// +build !windows

package fs

import (
	"github.com/balacode/zr"
)

// waitForDirChange __
func waitForDirChange(c chan string, dir string) {
	zr.IMPLEMENT()
} //                                                            waitForDirChange

//end
