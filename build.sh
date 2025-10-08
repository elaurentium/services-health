#!/bin/zsh
export GOOS=windows
export GOARCH=amd64
rsrc -manifest healthservice.manifest -icon icon.ico -o health.syso
go build -o health.exe .