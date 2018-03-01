// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-01 17:06:35 6BF41A              [zr/fs/dir_watcher_nonwindows.go]
// -----------------------------------------------------------------------------
// +build !windows

package fs

import "ase/zr" // Zircon-Go

// waitForDirChange __
func waitForDirChange(c chan string, dir string) {
	zr.IMPLEMENT()
} //                                                            waitForDirChange

//end
