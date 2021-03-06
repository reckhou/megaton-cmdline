#!/bin/sh

CURDIR=`pwd`
EXPORTPATH=$CURDIR/bin/mt
SRCPATH=$CURDIR/mtcmd.go

echo "=== Running govet tools to check code validity ==="
go tool vet ./
echo "=== govet ends ==="

mkdir -p $CURDIR/bin
mkdir -p $CURDIR/bin/config
cp $CURDIR/doc/config.json $CURDIR/bin/config/
gofmt -w=true -tabwidth=2 -tabs=false $CURDIR
go build -v -o $EXPORTPATH $SRCPATH

EXPORTPATH=$CURDIR/bin/mt_linux_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o $EXPORTPATH $SRCPATH
