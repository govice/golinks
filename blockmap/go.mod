module github.com/govice/golinks/blockmap

go 1.12

require (
	github.com/govice/golinks/archivemap v0.0.0-20190730011947-94e806ef1fbb
	github.com/govice/golinks/fs v0.0.0-20190730011947-94e806ef1fbb
	github.com/govice/golinks/walker v0.0.0-20190730011947-94e806ef1fbb
	github.com/pkg/errors v0.8.1
)

replace github.com/govice/golinks/archivemap => ../archivemap

replace github.com/govice/golinks/fs => ../fs

replace github.com/govice/golinks/walker => ../walker
