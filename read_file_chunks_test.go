// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 EFF386               [zr-fs/read_file_chunks_test.go]
// -----------------------------------------------------------------------------

package fs

/*
to test all items in read_file_chunks.go use:
    go test --run Test_rdfc_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	"os"
	str "strings"
	"testing"

	"github.com/balacode/zr" // Zircon-Go
)

// go test --run Test_rdfc_ReadFileChunks_
func Test_rdfc_ReadFileChunks_(t *testing.T) {
	const ChunkSize = 1024
	const SamplePath = `x:\test`
	const SampleFile = SamplePath + "ReadFileChunks.tmp"
	// -------------------------------------------------------------------------
	zr.DisableErrors()
	// -------------------------------------------------------------------------
	// return an error if 'filename' is blank
	{
		var reader = func(chunk []byte) int64 {
			t.Error("Called reader() when 'filename' is blank.")
			return 0
		}
		var err = ReadFileChunks("", ChunkSize, reader)
		if err == nil {
			t.Error("Did not return an error when 'filename' is blank.")
		}
	}
	// -------------------------------------------------------------------------
	// return an error if 'filename' is not blank, but doesn't exist
	{
		var reader = func(chunk []byte) int64 {
			t.Error("Called reder() when 'filename' doesn't exist.")
			return 0
		}
		var err = ReadFileChunks("NOFILE.TMP", ChunkSize, reader)
		if err == nil {
			t.Error("Did not return an error when 'filename' doesn't exist.")
		}
	}
	// -------------------------------------------------------------------------
	// return an error if 'chunkSize' is zero
	{
		var reader = func(chunk []byte) int64 {
			t.Error("Called reader() when 'chunkSize' is zero.")
			return 0
		}
		var err = ReadFileChunks(SampleFile, 0, reader)
		if err == nil {
			t.Error("Did not return an error when 'chunkSize' is zero.")
		}
	}
	// -------------------------------------------------------------------------
	// return an error if 'chunkSize' is negative
	{
		var reader = func(chunk []byte) int64 {
			t.Error("Called reader() when 'chunkSize' is negative.")
			return 0
		}
		var err = ReadFileChunks(SampleFile, -1, reader)
		if err == nil {
			t.Error("Did not return an error when 'chunkSize' is negative.")
		}
	}
	// -------------------------------------------------------------------------
	// return an error if 'reader' is nil
	{
		var err = ReadFileChunks(SampleFile, ChunkSize, nil)
		if err == nil {
			t.Error("Did not return an error when 'reader' is nil.")
		}
	}
	var fillers = []string{"1", "2", "3", "4", "5", "6", "7"}
	var createSampleFile = func() {
		// create a file and fill it with some data
		os.Remove(SampleFile)
		for _, filler := range fillers {
			zr.AppendToTextFile(SampleFile, str.Repeat(filler, ChunkSize))
		}
	}
	// -------------------------------------------------------------------------
	zr.EnableErrors()
	// -------------------------------------------------------------------------
	// is the read data consistent?
	{
		// reader() function that will check for consistency
		var i = 0
		var reader = func(chunk []byte) int64 {
			var expect = str.Repeat(fillers[i], ChunkSize)
			if string(chunk) != expect {
				t.Error("Read chunk doesn't match expected data.")
			}
			i++
			return int64(len(chunk))
		}
		createSampleFile()
		var err = ReadFileChunks(SampleFile, ChunkSize, reader)
		if err != nil {
			t.Error("Expected to return nil, but returned error:", err)
		}
	}
	// -------------------------------------------------------------------------
	// when 'reader' returns 0, further reading must stop
	{
		// reader() function must be called only once
		var i = 0
		var reader = func(chunk []byte) int64 {
			if i > 0 {
				t.Error("reader() returned false, but reading did not stop.")
			}
			var expect = str.Repeat(fillers[0], ChunkSize)
			if string(chunk) != expect {
				t.Error("Read chunk doesn't match expected data.")
			}
			i++
			return 0
		}
		createSampleFile()
		var err = ReadFileChunks(SampleFile, ChunkSize, reader)
		if err != nil {
			t.Error("Expected to return nil, but returned error:", err)
		}
	}
	os.Remove(SampleFile)
} //                                                   Test_rdfc_ReadFileChunks_

//end
