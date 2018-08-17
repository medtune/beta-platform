#!/bin/bash

# Static global frameworks
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/cobra

# Dependencies
go get -v ./pkg/config 
go get -v ./pkg/session 
go get -v ./pkg/store 
go get -v ./pkg/jsonutil

#
go get -u github.com/anthonynsimon/bild/transform
go get -u github.com/vincent-petithory/dataurl

if [ "$1" = "--hack" ]; then
    echo "installing ./hack/xlsx2pg package dependencies"
    go get -t -u -v github.com/tealeg/xlsx
fi
