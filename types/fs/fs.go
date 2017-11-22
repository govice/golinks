package fs

import (
	"os"

	"crypto/sha512"

	"github.com/pkg/errors"
)

func HashFile(path string) ([]byte, error) {
	//If path is null return
	if path == "" {
		return nil, errors.New("BlockMap: failed to hash null path")
	}
	//Open open and verify file in path
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "BlockMap: failed to read "+path)
	}

	//Get the file size for reading
	fileStat, err := f.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "BlockMap: failed read size of "+path)
	}
	//Construct a buffer based on the file size
	buffer := make([]byte, fileStat.Size())
	//Read the file and verify bytes read equal the stat size
	bytesRead, err := f.Read(buffer)
	if err != nil {
		return nil, errors.Wrap(err, "BlockMap: failed to to buffer in file "+path)
	}
	if bytesRead != int(fileStat.Size()) {
		return nil, errors.New("Blockmap: bytes read not equal to stat size of file" + path)
	}

	fileHash := sha512.New()
	if _, err := fileHash.Write(buffer); err != nil {
		return nil, err
	}
	return fileHash.Sum(nil), nil
}
