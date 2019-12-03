#!/bin/bash
echo "starting build.sh"

echo "make dir"


# Stops the process if something fails
# set -xe

# All of the dependencies needed/fetched for your project.
# This is what actually fixes the problem so that EB can find your dependencies. 
# FOR EXAMPLE:

echo "get dependencies"
# go get "github.com/gin-gonic/gin"

# create the application binary that eb uses
go build -o bin/application application.go

chmod +x bin/application

echo "build successful"