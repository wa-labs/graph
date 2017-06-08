# Gonum Graph [![Build Status](https://travis-ci.org/wa-labs/graph.svg?branch=master)](https://travis-ci.org/wa-labs/graph) [![Coverage Status](https://coveralls.io/repos/wa-labs/graph/badge.svg?branch=master&service=github)](https://coveralls.io/github/wa-labs/graph?branch=master) [![GoDoc](https://godoc.org/github.com/wa-labs/graph?status.svg)](https://godoc.org/github.com/wa-labs/graph)

This is a generalized graph package for the Go language. It aims to provide a clean, transparent API for common algorithms on arbitrary graphs such as finding the graph's strongly connected components, dominators, or searces.

The package is currently in testing, and the API is "semi-stable". The signatures of any functions like AStar are unlikely to change much, but the Graph, Node, and Edge interfaces may change a bit.

## Issues

If you find any bugs, feel free to file an issue on the github issue tracker. Discussions on API changes, added features, code review, or similar requests are preferred on the Gonum-dev Google Group.

https://groups.google.com/forum/#!forum/gonum-dev

## License

Please see github.com/gonum/license for general license information, contributors, authors, etc on the Gonum suite of packages.
