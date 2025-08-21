#!/bin/zsh
export GOOS=windows
export GOARCH=amd64
rsrc -manifest healthservice.manifest -ico favicon.ico -o health.syso
go build -o health.exe .