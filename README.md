# golinks
goLinks was created to realize a blockchain implemented in Go. This project has grown beyond it's initial scope into a
blockchain and filesystem library as well as a command line tool. Updates to come.
## Author 
Kevin Gentile
    
Contact: kevin@kevingentile.com


# Installation
```
go get -u github.com/laughingcabbage/golinks
go install github.com/laughingcabbage/golinks
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

#### [block](https://godoc.org/github.com/laughingcabbage/goLinks/block)

#### [blockchain](https://godoc.org/github.com/laughingcabbage/goLinks/blockchain)

#### [blockmap](https://godoc.org/github.com/laughingcabbage/goLinks/blockmap)

#### [fs](https://godoc.org/github.com/laughingcabbage/goLinks/fs)

#### [walker](https://godoc.org/github.com/laughingcabbage/goLinks/walker)