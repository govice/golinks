# golinks
[![CircleCI](https://circleci.com/gh/govice/golinks/tree/master.svg?style=svg)](https://circleci.com/gh/govice/golinks/tree/master)
[![codecov](https://codecov.io/gh/govice/golinks/branch/master/graph/badge.svg)](https://codecov.io/gh/govice/golinks)
[![GoDoc](https://godoc.org/github.com/govice/golinks?status.svg)](https://godoc.org/github.com/govice/golinks)



Golinks is a command line tool and library used to deep-hash file systems. This project is still under early development and subject to frequent change.

**Author**:     Kevin Gentile

**Contact**:    kevin@govice.org


# Installation
```
go get -u github.com/govice/golinks
go install github.com/govice/golinks
golinks -h
```

# Usage

Create a link file for an archive located at directory `archive`

## Generation
Generate a `.link` file used as a reference when validating archives.
```
golinks link ~/[pathToArchive]/archive
```

## Validation
Determine if a linked archive is valid
```
golinks validate ~/[pathToArchive]/archive
```


# Contributing
Contributions are welcome. We use a [forking workflow](https://www.atlassian.com/git/tutorials/comparing-workflows/forking-workflow) for all contributions.
 Check out this article about [working with forked repositories in Go](https://blog.sgmansfield.com/2016/06/working-with-forks-in-go/).
Be sure to open an issue for any new work. 
Branch names should be discriptive and append the issue number (ex. `update-readme-123`).
All pull requests should be "squashed" into a single commit and resolve the commit issue (ex. `[Resolves #123] update readme`). 
Pull requests should be similarly named and close the issue (ex. `[Closes #123] update readme`).

Happy coding :)


# Testing

The default resource folder used by this tool is located at `~/.golinks`

The default test root is located at `~/.golinks/test`

## Test Archive

To generate a test archive in the test root:
```bash
golinks buildtest -s [small|medium|large]
```
Which creates 10 folders within the test archive each containing 10 files of the following sizes:
```
 small:     1 B
 medium:    1 KB
 large:     1 MB
```
To delete a test archive:
```
golinks buildtest clean
```

## Environment

It can be useful to specify enviornment variables for testing

* Windows:

    ```
    TEST_ROOT : "%userprofile%\.golinks\test"
    ```

* Linux:
    ```
    TEST_ROOT=~/.golinks/test
    ```


## GoDoc

* [block](https://godoc.org/github.com/govice/goLinks/block)

* [blockchain](https://godoc.org/github.com/govice/goLinks/blockchain)

* [blockmap](https://godoc.org/github.com/govice/goLinks/blockmap)

* [fs](https://godoc.org/github.com/govice/goLinks/fs)

* [walker](https://godoc.org/github.com/govice/goLinks/walker)


## License
Copyright (C) 2019-2020 Kevin Gentile & other contributors (see AUTHORS.md)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
