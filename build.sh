#!/bin/bash
echo "starting build.sh"

# Stops the process if something fails
# set -xe

# All of the dependencies needed/fetched for your project.
# This is what actually fixes the problem so that EB can find your dependencies. 
# FOR EXAMPLE:

echo "get dependencies"
# go get "github.com/gin-gonic/gin"

# create the application binary that eb uses
GOOS=linux GOARCH=amd64 go build -o bin/application -ldflags="-s -w"

chmod +x bin/application

echo "build successful"