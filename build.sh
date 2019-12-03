#!/bin/bash
echo "starting build.sh"

echo "ls before .."
ls -l


# Stops the process if something fails
# set -xe

# All of the dependencies needed/fetched for your project.
# This is what actually fixes the problem so that EB can find your dependencies. 
# FOR EXAMPLE:

echo "get dependencies"
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/static
go get github.com/sirupsen/logrus
go get github.com/mazeForGit/WordlistExtractor

# create the application binary that eb uses
go build -o bin/application server.go

chmod +x bin/application

echo "ls after .."
cd ./bin
ls -l

echo "build successful"