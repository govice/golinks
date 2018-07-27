# golinks command line interface
This command line tool is currently under development and will change frequently until the first tagged release. The goal of this project is to produce a data integrity management tool. Updates to come.

## install
Make sure your go bin path is configured.

    go get github.com/LaughingCabbage/golinks/golinks
    go install github.com/LaughingCabbage/golinks/golinks
    golinks -h

## test

    golinks maketest [path] -[small/medium/large]

####Example
    golinks maketest ./ -large
    
## Author
Kevin Gentile

Contact: kevin@kevingentile.com
