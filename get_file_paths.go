// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-06-05 23:50:24 31E0D5                      zr-fs/[get_file_paths.go]
// -----------------------------------------------------------------------------

package fs

// # Types
//   PathAndSize struct
//   FilesMap map[int64][]*PathAndSize
//   Options struct
//
// # Function
//   GetFilePaths(dir string, exts ...string) []string
//   getFilesMap(dir, filter string) Files

import (
	"fmt"
	"os"
	"path/filepath"
	str "strings"
)

// -----------------------------------------------------------------------------
// # Function

// GetFilePaths returns a list of file names (with full path) contained
// in folder 'dir' that match the given file extensions.
// Extensions should be specified as: "ext", or ".ext", not "*.ext"
// If you don't specify 'exts', returns all files in 'dir'
func GetFilePaths(dir string, exts ...string) []string {
	if dir == "" {
		fmt.Println("GetFilePaths(): 'dir' arg is blank.", callers())
		return nil
	}
	var ret []string
	filepath.Walk(
		dir, func(path string, info os.FileInfo, err error) error {
			// skip directory entries (Walk takes care of reading subfolders)
			if str.Contains(path, "$RECYCLE.BIN") {
				return nil
			}
			if err != nil {
				fmt.Printf("in path %s: %s"+LF, path, err)
				return nil
			}
			if info.IsDir() {
				return nil
			}
			// skip files that don't match needed extension(s)
			var match = len(exts) == 0
			for _, ext := range exts {
				ext = str.ToLower(ext)
				if !str.HasPrefix(ext, ".") {
					ext = "." + ext
				}
				if str.HasSuffix(str.ToLower(path), ext) {
					match = true
					break
				}
			}
			if !match {
				return nil
			}
			// append the file name to returned list
			ret = append(ret, path)
			return nil
		},
	)
	return ret
} //                                                                GetFilePaths

//TODO: global: find return.*\{\}$ and replace with return nil

//end
