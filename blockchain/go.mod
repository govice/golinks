module github.com/govice/golinks/blockchain

go 1.12

require (
	github.com/govice/golinks/block v0.0.0-20190730011947-94e806ef1fbb
	github.com/pkg/errors v0.8.1
)

replace github.com/govice/golinks/block => ../block
