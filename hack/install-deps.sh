#!/bin/bash

# Static global frameworks
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/cobra

# Dependencies
go get -v ./pkg/config 
go get -v ./pkg/session 
go get -v ./pkg/store 
go get -v ./pkg/jsonutil

# Can change one day at pkg/service
go get -u github.com/anthonynsimon/bild/transform
go get -u github.com/vincent-petithory/dataurl

