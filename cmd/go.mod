module github.com/govice/golinks/cmd

go 1.12

require (
	github.com/google/uuid v1.1.1
	github.com/govice/golinks/block v0.0.0-20190804205723-8c69e5636931
	github.com/govice/golinks/blockmap v0.0.0-20190804205723-8c69e5636931
	github.com/govice/golinks/walker v0.0.0-20190804205723-8c69e5636931
	github.com/mitchellh/go-homedir v1.1.0
	github.com/pierrre/archivefile v0.0.0-20170218184037-e2d100bc74f5
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.4.0
	github.com/urfave/cli v1.22.1
)

replace github.com/govice/golinks/archivemap => ../archivemap

replace github.com/govice/golinks/blockmap => ../blockmap

replace github.com/govice/golinks/block => ../block

replace github.com/govice/golinks/fs => ../fs

replace github.com/govice/golinks/walker => ../walker
