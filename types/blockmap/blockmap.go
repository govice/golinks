package blockmap

//BlockMap is a ad-hoc Merkle tree-map
type BlockMap struct {
	archive  map[string][]byte
	rootHash []byte
	root     string
}

//New returns a new BlockMap initialized at the provided root
func New(root string) *BlockMap {
	rootMap := make(map[string][]byte)
	return &BlockMap{archive: rootMap, rootHash: nil, root: root}
}

/*
//HashFile takes a file at path and returns it's hash
func (b *BlockMap) HashFile(path string) {
	//If path is null return
	if path == "" {
		log.Println("BlockMap: failed to hash null path")
		return
	}
	//Open open and verify file in path
	f, err := os.Open(path)
	if err != nil {
		log.Println("BlockMap: failed to read " + path)
		return
	}

	//Get the file size for reading
	fileStat, err := f.Stat()
	if err != nil {
		log.Println("BlockMap: failed read size of " + path)
		return
	}
	//Construct a buffer based on the file size
	buffer := make([]byte, fileStat.Size())
	//Read the file and verify bytes read equal the stat size
	bytesRead, err := f.Read(buffer)
	if err != nil {
		log.Println("BlockMap: failed to to buffer in file " + path)
		return
	}
	if bytesRead != int(fileStat.Size()) {
		panic(errors.New("Blockmap: bytes read not equal to stat size of file" + path))
	}


	b.archive[path] =

}
*/
