#!/bin/bash

#
#
#go get -u -d gocv.io/x/gocv
#cd $GOPATH/src/gocv.io/x/gocv
#make install
#make deps
#make download
#make build
#make sudo_install

# Static global frameworks
go get -u -v ./{internal,cmd}/...

if [ "$1" = "--hack" ]; then
    echo "installing ./hack/xlsx2pg package dependencies"
    go get -t -u -v github.com/tealeg/xlsx
fi
