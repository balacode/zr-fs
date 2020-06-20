// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2020-06-20 09:43:21 3FED84                    zr-fs/[read_file_chunks.go]
// -----------------------------------------------------------------------------

package fs

import (
	"os"

	"github.com/balacode/zr"
)

// ReadFileChunks reads a file in chunks and repeatedly calls
// the supplied 'reader' function with each read chunk.
// This function is useful for processing large files to
// avoid  having to load the entire file into memory.
//
// filename:  Name of the file to read from.
//
// chunkSize: Size of chunks to read from the file, in bytes.
//
// reader:    Function to call with each read chunk, which is passed
//            in the 'chunk' parameter. The function should return
//            int64(len(chunk)) to continue reading from the next
//            position, or 0 if further reading should stop.
//            It also accepts a negative value to skip
//            the reading position back.
//
// Returns an error if the file can't be opened or a read fails.
func ReadFileChunks(
	filename string,
	chunkSize int64,
	reader func(chunk []byte) int64,
) error {
	// sanity checks
	if filename == "" {
		return zr.Error(zr.EInvalidArg, "^filename", "zero-length")
	}
	if chunkSize < 1 {
		return zr.Error(zr.EInvalidArg, "^chunkSize", chunkSize)
	}
	if reader == nil {
		return zr.Error(zr.EInvalidArg, "^reader", "is nil")
	}
	// open the file
	var file *os.File
	var err error
	file, err = os.Open(filename)
	if file == nil || err != nil {
		return zr.Error("Failed opening", filename, "due to:", err)
	}
	defer file.Close()
	//
	// repeatedly read chunks from the file and call reader()
	chunk := make([]byte, chunkSize)
	pos := int64(0)
	for {
		// advance the reading position
		_, err = file.Seek(pos, 0)
		if err != nil {
			return zr.Error("Failed seeking", filename, "due to:", err)
		}
		// read the next chunk
		var bytesRead int
		bytesRead, err = file.Read(chunk[:])
		if bytesRead == 0 {
			break
		}
		if err != nil {
			return zr.Error("Failed reading", filename, "due to:", err)
		}
		// call the reader function with the chunk of dat
		offset := reader(chunk[:bytesRead])
		if offset == 0 {
			break
		}
		pos += offset
	}
	return nil
} //                                                              ReadFileChunks

//end
