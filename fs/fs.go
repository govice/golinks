package fs

import (
	"crypto/sha512"
	"os"

	"github.com/govice/golinks/walker"

	"path/filepath"

	"archive/zip"

	"io"

	"github.com/pkg/errors"
)

const (
	// T_UNBOUND indicates an unbound throttle
	T_UNBOUND int = -1
)

type FsErr struct {
	Path string
	Err  error
}

func (fe *FsErr) Unwrap() error { return fe.Err }

func (fe *FsErr) Error() string {
	if fe.Err != nil {
		return "fs: " + fe.Err.Error() + fe.Path
	}
	return "fs: failed for " + fe.Path
}

//ErrNullPath is returned when fs is given an empty path string
var ErrNullPath = errors.New("fs: failed to hash null path")

var ErrInvalidBufferSize = errors.New("fs: invalid buffer size")

//HashFile returns a sha512 hash of the file at the provided path
func HashFile(path string) ([]byte, error) {
	return throttleHashHelper(path, T_UNBOUND)
}

// ThrottleHashFile returns a hash for a file throttling with bufferSize bytes.
// a neagative or zero size buffer will be evaluated as T_UNBOUND
func ThrottleHashFile(path string, bufferSize int) ([]byte, error) {
	return throttleHashHelper(path, bufferSize)
}

func throttleHashHelper(path string, bufferSize int) ([]byte, error) {
	//If path is null return
	if path == "" {
		return nil, ErrNullPath
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, &FsErr{
			Path: path,
			Err:  err,
		}
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, &FsErr{
			Path: path,
			Err:  err,
		}
	}

	var readBuffer []byte
	// unbound reads will load the file into memory and hash
	if bufferSize <= 0 {
		readBuffer = make([]byte, fi.Size())
	} else {
		readBuffer = make([]byte, bufferSize)
	}

	fileHash := sha512.New()
	for {
		bytesRead, rerr := f.Read(readBuffer)
		if rerr != nil && !errors.Is(rerr, io.EOF) {
			return nil, &FsErr{
				Path: path,
				Err:  rerr,
			}
		}

		if bytesRead > 0 {
			if _, err := fileHash.Write(readBuffer[:bytesRead]); err != nil {
				return nil, &FsErr{
					Path: path,
					Err:  err,
				}
			}
		}

		if errors.Is(rerr, io.EOF) {
			break
		}
	}

	return fileHash.Sum(nil), nil
}

// ErrExpectedDirectory expects a directory path
var ErrExpectedDirectory = errors.New("fs: compress operation requires path to a directory")

//Compress stores a zip file of in the provided path
func Compress(path, target string) error {
	//TODO this is...dense. cyclomatic complexity >10

	//Verify directory exists
	s, err := os.Stat(path)
	if err != nil {
		return &FsErr{
			Err:  err,
			Path: path,
		}
	}

	if !s.IsDir() {
		return &FsErr{
			Err:  ErrExpectedDirectory,
			Path: path,
		}
	}

	//Get the archives parent for a default storage location
	parentPath, err := filepath.Abs(target)
	if err != nil {
		return &FsErr{
			Err:  err,
			Path: path,
		}
	}

	//Open the zip archive for writing
	archiveBuffer, err := os.OpenFile(parentPath+string(os.PathSeparator)+s.Name()+".zip", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return &FsErr{
			Err:  err,
			Path: archiveBuffer.Name() + ".zip",
		}
	}

	//Initialize compression writer
	z := zip.NewWriter(archiveBuffer)

	//walk the provided directory
	w := walker.New(path)
	if err := w.Walk(); err != nil {
		return &FsErr{
			Err:  err,
			Path: path,
		}
	}

	for _, file := range w.Archive() {
		//Get the files relative path in the archive
		relPath, err := filepath.Rel(path, file)
		if err != nil {
			return &FsErr{
				Err:  err,
				Path: path,
			}
		}
		//Create zip file buffer for compression storage
		zipFile, err := z.Create(s.Name() + string(os.PathSeparator) + relPath)
		if err != nil {
			return &FsErr{
				Err:  err,
				Path: relPath,
			}
		}
		//Open file for copying
		f, err := os.Open(file)
		if err != nil {
			return &FsErr{
				Err:  err,
				Path: file,
			}
		}

		//copy file into zip archive
		if _, err := io.Copy(zipFile, f); err != nil {
			return &FsErr{
				Err:  err,
				Path: f.Name(),
			}
		}
		//close file opened in iteration
		if err := f.Close(); err != nil {
			return &FsErr{
				Err:  err,
				Path: f.Name(),
			}
		}
	}

	//close zip archive
	if err := z.Close(); err != nil {
		return &FsErr{
			Err: err,
		}
	}

	//close archive buffer
	if err := archiveBuffer.Close(); err != nil {
		return &FsErr{
			Err: err,
		}
	}

	return nil

}

//Decompress extracts a compressed archive at path to target
func Decompress(path, target string) error {

	//create zip reader
	r, err := zip.OpenReader(path)
	if err != nil {
		return &FsErr{
			Err:  err,
			Path: path,
		}
	}

	//iterate and extract all files in the zip archive to target
	for _, file := range r.File {
		//Get absolute location from archive relative path
		destFile := filepath.Join(target, file.Name)

		//Get files directory name and create the folder hierarchy
		destDir, _ := filepath.Split(destFile)
		if err := os.MkdirAll(destDir, file.Mode()); err != nil {
			return &FsErr{
				Err:  err,
				Path: destDir,
			}
		}

		//open zipped file for decompression
		zippedFile, err := file.Open()
		if err != nil {
			return &FsErr{
				Err:  err,
				Path: file.Name,
			}
		}

		//open destination file
		dest, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, file.Mode())
		if err != nil {
			return &FsErr{
				Err:  err,
				Path: destFile,
			}
		}

		//Write unzipped file to destination
		if _, err := io.Copy(dest, zippedFile); err != nil {
			return &FsErr{
				Err:  err,
				Path: dest.Name(),
			}
		}

		//close opened zip file
		if err := zippedFile.Close(); err != nil {
			return &FsErr{
				Err: err,
			}
		}

		//close opened destination file
		if err := dest.Close(); err != nil {
			return &FsErr{
				Err: err,
			}
		}
	}

	//close zip reader
	if err := r.Close(); err != nil {
		return &FsErr{
			Err: err,
		}
	}

	return nil
}
