module github.com/govice/golinks

go 1.12

require (
	github.com/govice/golinks/block v0.0.0-20190510053004-bd4f7003a03c // indirect
	github.com/govice/golinks/blockmap v0.0.0-20190510053004-bd4f7003a03c // indirect
	github.com/govice/golinks/cmd v0.0.0-20190510053004-bd4f7003a03c
	github.com/pierrre/archivefile v0.0.0-20170218184037-e2d100bc74f5 // indirect
	github.com/spf13/cobra v0.0.5 // indirect
	github.com/spf13/viper v1.4.0 // indirect
	github.com/urfave/cli v1.20.0 // indirect
)

replace github.com/govice/golinks/archivemap => ./archivemap

replace github.com/govice/golinks/block => ./block

replace github.com/govice/golinks/blockmap => ./blockmap

replace github.com/govice/golinks/cmd => ./cmd

replace github.com/govice/golinks/fs => ./fs

replace github.com/govice/golinks/walker => ./walker
