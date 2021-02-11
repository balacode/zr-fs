// -----------------------------------------------------------------------------
// ZR Library - File System Package               zr-fs/[dir_watcher_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package fs

import (
	"os"

	"github.com/balacode/zr"
)

// A DirWatcher holds a channel that delivers a folder's
// path when any file in the folder changes.
type DirWatcher struct {
	Chan <-chan string // the channel on which path changes are sent
	dir  string
} //                                                                  DirWatcher

// NewDirWatcher returns a new DirWatcher that contains
// a channel that be sent the name of a file every time
// a file in the folder or one of its subfolders changes.
func NewDirWatcher(dir string) (*DirWatcher, error) {
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		return nil, zr.Error("Folder^", dir, "does not exist")
	}
	c := make(chan string, 1)
	ret := &DirWatcher{Chan: c, dir: dir}
	go waitForDirChange(c, dir)
	return ret, nil
} //                                                               NewDirWatcher

//end
