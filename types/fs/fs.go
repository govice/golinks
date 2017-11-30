package fs

import (
	"crypto/sha512"
	"os"

	"path/filepath"

	"archive/zip"

	"io"

	"github.com/LaughingCabbage/goLinks/types/walker"
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
func Compress(path string) error {
	//TODO this is...dense. cyclomatic complexity >10

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
	parentPath, err := filepath.Abs(path + string(os.PathSeparator) + "..")
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
		return errors.Wrap(err, "fs: failed to close archive buffer")
	}

	return nil

}
