// -----------------------------------------------------------------------------
// balarabe@protonmail.com                                          License: MIT
// :v: 2018-02-24 01:44:00 4A0F67                    [zr_fs/dir_watcher_test.go]
// -----------------------------------------------------------------------------

package fs

/*
to test all items in dir_watcher.go use:
    go test --run Test_dirw_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import "os"      // standard
import "time"    // standard
import "testing" // standard

import "github.com/balacode/zr" // Zirconium

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
