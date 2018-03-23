// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-23 11:40:22 A40992                           [zr-fs/walk_path.go]
// -----------------------------------------------------------------------------

package fs

import "fmt"           // standard
import "os"            // standard
import "path/filepath" // standard
import "sync"          // standard
import "time"          // standard
import str "strings"   // standard

// WalkPathOptions __
type WalkPathOptions struct {
	FileExts     []string
	MinSize      int64
	MaxSize      int64
	ExcludeFunc  func(path string, info os.FileInfo)
	ProgressFunc func(scanned, listed int, size int64)
	WalkFunc     func(path string, info os.FileInfo, err error) error
} //                                                             WalkPathOptions

// WalkPath __
func WalkPath(path string, opts WalkPathOptions) []string {
	var ret = []string{}
	var scanCount = 0
	var listCount = 0
	var size = int64(0)
	var mutex = &sync.Mutex{}
	var update = true
	var runProgressFunc = func() {
		if opts.ProgressFunc == nil {
			return
		}
		mutex.Lock()
		opts.ProgressFunc(scanCount, listCount, size)
		mutex.Unlock()
	}
	var appendFile = func(path string, info os.FileInfo, err error) error {
		scanCount++
		// ignore directories and files in system folders
		if err != nil {
			if !str.Contains(str.ToLower(fmt.Sprintf("%s", err)),
				"$recycle.bin") {
				myError(err)
			}
			return nil
		}
		if info.IsDir() ||
			str.Contains(path, "\\System Volume Information\\") ||
			str.Contains(path, "$RECYCLE.BIN") {
			return nil
		}
		if len(opts.FileExts) > 0 && !IsFileExt(path, opts.FileExts) {
			return nil
		}
		if info.Size() > opts.MaxSize {
			if opts.ExcludeFunc != nil {
				mutex.Lock()
				opts.ExcludeFunc(path, info)
				mutex.Unlock()
			}
			return nil
		}
		ret = append(ret, path)
		listCount++
		size += info.Size()
		if opts.WalkFunc != nil {
			mutex.Lock()
			defer mutex.Unlock()
			return opts.WalkFunc(path, info, err)
		}
		return nil
	}
	if opts.ProgressFunc != nil {
		go func() {
			for update {
				time.Sleep(100 * time.Millisecond)
				runProgressFunc()
			}
		}()
	}
	var err = filepath.Walk(path, appendFile)
	if err != nil {
		myError(err)
		ret = []string{}
	}
	update = false
	runProgressFunc()
	return ret
} //                                                                    WalkPath

//end
