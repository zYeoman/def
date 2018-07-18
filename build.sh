#! /bin/sh
#
# build.sh
# Copyright (C) 2018 Yongwen Zhuang <zeoman@163.com>
#
# Distributed under terms of the MIT license.
#

GOOS=windows GOARCH=amd64 go build -o defd-win.exe defd.go data.go
GOOS=darwin  GOARCH=amd64 go build -o defd-osx defd.go data.go
go build -o defd-linux defd.go data.go

GOOS=windows GOARCH=amd64 go build -o def-win.exe def.go
GOOS=darwin  GOARCH=amd64 go build -o def-osx def.go
go build -o def-linux def.go

zip def-win.zip def-win.exe defd-win.exe
zip def-osx.zip def-osx defd-osx
zip def-linux.zip def-linux defd-linux

rm def-win.exe defd-win.exe def-osx defd-osx def-linux defd-linux
