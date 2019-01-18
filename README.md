# golinks
Golinks is a command line tool and library used to deep-hash file systems. This project is still under early development and subject to frequent change.

## Author 
Kevin Gentile
    
Contact: kevin@govice.org


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


    ### GoDoc

#### [block](https://godoc.org/github.com/govice/goLinks/block)

#### [blockchain](https://godoc.org/github.com/govice/goLinks/blockchain)

#### [blockmap](https://godoc.org/github.com/govice/goLinks/blockmap)

#### [fs](https://godoc.org/github.com/govice/goLinks/fs)

#### [walker](https://godoc.org/github.com/govice/goLinks/walker)
