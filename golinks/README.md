
# Installation
```
go get -u github.com/laughingcabbage/golinks/golinks
go install github.com/laughingcabbage/golinks/golinks
golinks -h
```

# Testing
The default resource folder used by this tool is located at `~/.golinks`

The default test root is located at `~/.golinks/test`

It can be useful to specify enviornment variables for testing

* Windows:
    ```
    TEST_ROOT : "%userprofile%\.golinks\test"
    ```

* Linux:
    ```
    TEST_ROOT=~/.golinks/test
    ```