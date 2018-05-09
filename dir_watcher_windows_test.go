// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 F4BB96            [zr-fs/dir_watcher_windows_test.go]
// -----------------------------------------------------------------------------

package fs

/*
to test all items in dir_watcher_windows.go use:
    go test --run Test_dirw_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"os"
	"testing"
	"time"

	"github.com/balacode/zr" // Zircon-Go
)

// go test --run Test_dirw_DirWatcher_
func Test_dirw_DirWatcher_(t *testing.T) {
	zr.TBegin(t)
	const TESTDIR = `X:\TEST`
	const TESTFILE = `X:\TEST\FILE.TMP`
	//
	// this test writes to TESTFILE 5 times every 150ms
	// then checks that the watcher detects a directory change exactly 5 times
	var dirWatchChan = NewDirWatcher(TESTDIR).C
	var intervalChan = time.NewTicker(time.Millisecond * 150).C
	var quitChan = time.NewTimer(time.Second * 1).C
	var intervalCount = 5
	var watchCount = 0
	//
	// all 3 channels should not be nil
	zr.TTrue(t, dirWatchChan != nil)
	zr.TTrue(t, intervalChan != nil)
	zr.TTrue(t, quitChan != nil)
	os.Remove(TESTFILE)
	//
	var run = true
loop:
	for run {
		select {
		case <-intervalChan: // write to the file to cause a directory change
			if intervalCount > 0 {
				zr.AppendToTextFile(TESTFILE, "123")
			}
			intervalCount--
		case <-dirWatchChan: // should occur after every file change
			watchCount++
		case <-quitChan: // end loop when test times out
			run = false
			break loop
		}
	}
	zr.TEqual(t, watchCount, (5))
} //                                                       Test_dirw_DirWatcher_

//end
