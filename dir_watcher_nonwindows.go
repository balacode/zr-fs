// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 8E649D              [zr/fs/dir_watcher_nonwindows.go]
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
