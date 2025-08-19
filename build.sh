#!/bin/zsh
export GOOS=windows
export GOARCH=amd64
go build -o health.exe cmd/cmd.go  