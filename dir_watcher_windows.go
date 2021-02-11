// -----------------------------------------------------------------------------
// ZR Library - File System Package               zr-fs/[dir_watcher_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------
// +build windows

package fs

import (
	"time"

	"github.com/balacode/zr"
	win "github.com/balacode/zr-win"
)

// waitForDirChange _ _
func waitForDirChange(c chan string, dir string) {
	// start watching the folder (and check that handle value is correct)
	var handles [2]win.HANDLE
	{
		NOTIFY := win.FILE_NOTIFY_CHANGE_CREATION |
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
	prev := time.Now()
	//
	// begin loop that waits for a change to occur
	for {
		// wait for notification
		status := win.WaitForMultipleObjects(
			1, &handles[0], win.TRUE, win.INFINITE,
		)
		if status != win.WAIT_OBJECT_0 {
			zr.Error("Unhandled wait status", status)
			return
		}
		now := time.Now()
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
