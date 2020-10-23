#!/bin/bash

# Reset the tree to the current commit to handle
# any writes during the build.
git reset --hard

go generate github.com/ArkaGPL/gonum/blas
go generate github.com/ArkaGPL/gonum/blas/gonum
go generate github.com/ArkaGPL/gonum/unit
go generate github.com/ArkaGPL/gonum/unit/constant
go generate github.com/ArkaGPL/gonum/graph/formats/dot
go generate github.com/ArkaGPL/gonum/stat/card

if [ -n "$(git diff)" ]; then	
	git diff
	exit 1
fi
