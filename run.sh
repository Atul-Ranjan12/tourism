#!/bin/bash
set GOOS=windows
set GOARCH=amd64

go build -o tourism.exe cmd/web/*.go
./tourism.exe