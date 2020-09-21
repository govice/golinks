/*
 *Copyright 2018-2019 Kevin Gentile
 *
 *Licensed under the Apache License, Version 2.0 (the "License");
 *you may not use this file except in compliance with the License.
 *You may obtain a copy of the License at
 *
 *http://www.apache.org/licenses/LICENSE-2.0
 *
 *Unless required by applicable law or agreed to in writing, software
 *distributed under the License is distributed on an "AS IS" BASIS,
 *WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *See the License for the specific language governing permissions and
 *limitations under the License.
 */

package fs

import (
	"crypto/sha512"
	"io/ioutil"
	"os"

	"github.com/govice/golinks/walker"

	"path/filepath"

	"archive/zip"

	"io"

	"github.com/pkg/errors"
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

//HashFile returns a sha512 hash of the file at the provided path
func HashFile(path string) ([]byte, error) {
	//If path is null return
	if path == "" {
		return nil, ErrNullPath
	}
	//Open open and verify file in path
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, &FsErr{
			Path: path,
			Err:  err,
		}
	}

	fileHash := sha512.New()
	if _, err := fileHash.Write(fileBytes); err != nil {
		return nil, &FsErr{
			Path: path,
			Err:  err,
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
