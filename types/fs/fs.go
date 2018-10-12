/*
 *Copyright 2017 Kevin Gentile
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
	"encoding/gob"
	"github.com/laughingcabbage/golinks/types/walker"
	"os"

	"path/filepath"

	"archive/zip"

	"io"

	"github.com/pkg/errors"
)

//HashFile returns a sha512 hash of the file at the provided path
func HashFile(path string) ([]byte, error) {
	//If path is null return
	if path == "" {
		return nil, errors.New("fs: failed to hash null path")
	}
	//Open open and verify file in path
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "fs: failed to read "+path)
	}

	//Get the file size for reading
	fileStat, err := f.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "fs: failed read size of "+path)
	}
	//Construct a buffer based on the file size
	buffer := make([]byte, fileStat.Size())
	//Read the file and verify bytes read equal the stat size
	bytesRead, err := f.Read(buffer)
	if err != nil {
		return nil, errors.Wrap(err, "fs: failed to to buffer in file "+path)
	}
	if bytesRead != int(fileStat.Size()) {
		return nil, errors.New("fs: bytes read not equal to stat size of file" + path)
	}

	if err := f.Close(); err != nil {
		return nil, errors.Wrap(err, "fs: failed to close file "+path)
	}

	fileHash := sha512.New()
	if _, err := fileHash.Write(buffer); err != nil {
		return nil, err
	}
	return fileHash.Sum(nil), nil
}

//Compress stores a zip file of in the provided path
func Compress(path, target string) error {
	//TODO this is...dense. cyclomatic complexity >10
	//TODOAY we refactor this to use go channels

	//Verify directory exists
	s, err := os.Stat(path)
	if err != nil {
		return errors.Wrap(err, "fs: compress failed to detect archive from path")
	}

	//TODO i'm not convinced this will work on individual files sooooo...
	if !s.IsDir() {
		return errors.New("fs: compress operation requires path to a directory")
	}

	//Get the archives parent for a default storage location
	parentPath, err := filepath.Abs(target)
	if err != nil {
		return errors.Wrap(err, "fs: compress failed to extract absolute path of archive")
	}

	//Open the zip archive for writing
	archiveBuffer, err := os.OpenFile(parentPath+string(os.PathSeparator)+s.Name()+".zip", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return errors.Wrap(err, "fs: compress failed to open "+archiveBuffer.Name()+".zip")
	}

	//Initialize compression writer
	z := zip.NewWriter(archiveBuffer)

	//walk the provided directory
	w := walker.New(path)
	if err := w.Walk(); err != nil {
		return errors.Wrap(err, "fs: compress failed to walk archive")
	}

	for _, file := range w.Archive() {
		//Get the files relative path in the archive
		relPath, err := filepath.Rel(path, file)
		if err != nil {
			return errors.Wrap(err, "fs: compress failed to extract relative file path in archive")
		}
		//Create zip file buffer for compression storage
		zipFile, err := z.Create(s.Name() + string(os.PathSeparator) + relPath)
		if err != nil {
			return errors.Wrap(err, "fs: compress failed to create zip file "+relPath)
		}
		//Open file for copying
		f, err := os.Open(file)
		if err != nil {
			return errors.Wrap(err, "fs: compress failed to open "+file)
		}

		//copy file into zip archive
		if _, err := io.Copy(zipFile, f); err != nil {
			return errors.Wrap(err, "fs: failed to write file to zip folder")
		}
		//close file opened in iteration
		if err := f.Close(); err != nil {
			return errors.Wrap(err, "fs: compress failed to close "+file)
		}
	}

	//close zip archive
	if err := z.Close(); err != nil {
		return errors.Wrap(err, "fs: compress failed to close zip writer")
	}
	//close archive buffer
	if err := archiveBuffer.Close(); err != nil {
		return errors.Wrap(err, "fs: compress failed to close archive buffer")
	}

	return nil

}

//Decompress extracts a compressed archive at path to target
func Decompress(path, target string) error {

	//create zip reader
	r, err := zip.OpenReader(path)
	if err != nil {
		return errors.Wrap(err, "fs: decompress failed to open zip reader")
	}

	//iterate and extract all files in the zip archive to target
	for _, file := range r.File {
		//Get absolute location from archive relative path
		destFile := filepath.Join(target, file.Name)

		//Get files directory name and create the folder hierarchy
		destDir, _ := filepath.Split(destFile)
		if err := os.MkdirAll(destDir, file.Mode()); err != nil {
			return errors.Wrap(err, "fs: decompress failed to create destination")
		}

		//open zipped file for decompression
		zippedFile, err := file.Open()
		if err != nil {
			return errors.Wrap(err, "fs: decompress failed to open zipped file")
		}

		//open destination file
		dest, err := os.OpenFile(destFile, os.O_RDWR|os.O_CREATE, file.Mode())
		if err != nil {
			return errors.Wrap(err, "fs: decompress failed to open destination file")
		}

		//Write unzipped file to destination
		if _, err := io.Copy(dest, zippedFile); err != nil {
			return errors.Wrap(err, "fs: decompress failed to write unzip file to destination")
		}

		//close opened zip file
		if err := zippedFile.Close(); err != nil {
			return errors.Wrap(err, "fs: decompress failed to close zipped file")
		}

		//close opened destination file
		if err := dest.Close(); err != nil {
			return errors.Wrap(err, "fs: decompress failed to close destination file")
		}

	}

	//close zip reader
	if err := r.Close(); err != nil {
		return errors.Wrap(err, "fs: decompress failed to close zip reader")
	}

	return nil
}

//SaveGob writes a gob file to path
func SaveGob(path string, object interface{}) error {
	file, err := os.Create(path)
	defer file.Close()
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	return err
}

//ReadGob reads a gob file from path
func ReadGob(path string, object interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	if err = decoder.Decode(object); err != nil {
		return err
	}

	return nil
}
