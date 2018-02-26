// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-26 01:57:01 8F6AD7                 [zr_fs/dir_watcher_windows.go]
// -----------------------------------------------------------------------------

package fs

import "os"   // standard
import "time" // standard

import "github.com/balacode/zr"     // Zirconium
import "github.com/balacode/zr_win" // Zirconium

// A DirWatcher holds a channel that delivers a folder's
// path when any file in the folder changes.
type DirWatcher struct {
	C   <-chan string // The channel on which path changes are sent.
	dir string
} //                                                                  DirWatcher

// NewDirWatcher returns a new DirWatcher that contains
// a channel that be sent the name of a file every time
// a file in the folder or one of its subfolders changes.
func NewDirWatcher(dir string) *DirWatcher {
	var _, err = os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		zr.Error("Folder^", dir, "does not exist")
		return nil
	}
	var c = make(chan string, 1)
	var ret = &DirWatcher{
		C:   c,
		dir: dir,
	}
	go waitForDirChange(c, dir)
	return ret
} //                                                               NewDirWatcher

// waitForDirChange __
func waitForDirChange(c chan string, dir string) {
	// start watching the folder (and check that handle value is correct)
	var handles [2]win.HANDLE
	{
		var NOTIFY = win.FILE_NOTIFY_CHANGE_CREATION |
			win.FILE_NOTIFY_CHANGE_FILE_NAME |
			win.FILE_NOTIFY_CHANGE_LAST_WRITE |
			win.FILE_NOTIFY_CHANGE_SIZE |
			0
			// not relevant:
			// win.FILE_NOTIFY_CHANGE_ATTRIBUTES |
			// win.FILE_NOTIFY_CHANGE_DIR_NAME |
			// win.FILE_NOTIFY_CHANGE_LAST_ACCESS |
			// win.FILE_NOTIFY_CHANGE_SECURITY |
		handles[0] = win.FindFirstChangeNotification(
			dir, win.TRUE, win.DWORD(NOTIFY),
		)
		switch handles[0] {
		case win.INVALID_HANDLE_VALUE:
			zr.Error("FindFirstChangeNotification() failed")
			return
		case win.NULL:
			zr.Error("FindFirstChangeNotification() returned NULL")
			return
		}
	}
	var prev = time.Now()
	//
	// begin loop that waits for a change to occur
	for {
		// wait for notification
		var status = win.WaitForMultipleObjects(
			1, &handles[0], win.TRUE, win.INFINITE,
		)
		if status != win.WAIT_OBJECT_0 {
			zr.Error("Unhandled wait status", status)
			return
		}
		var now = time.Now()
		// only send on channel if more than 0.1s elapsed from last change,
		// if enough time elapsed, wait for 0.1s, then send on channel
		if since := now.Sub(prev).Seconds(); since > 0.1 {
			go func() {
				time.Sleep(100 * time.Millisecond)
				c <- dir
			}()
			prev = now
		}
		if win.FindNextChangeNotification(handles[0]) == win.FALSE {
			zr.Error("FindNextChangeNotification() failed")
			return
		}
	}
} //                                                            waitForDirChange

//end
