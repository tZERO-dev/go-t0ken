#!/bin/sh


COMMIT=`git rev-parse --short=10 HEAD`

xgo \
	-v \
	--dest=dist \
	--targets="darwin/amd64,linux/amd64,windows/amd64" \
	-ldflags="-X github.com/tzero-dev/go-t0ken/commands.GitCommit=$COMMIT" \
	github.com/tzero-dev/go-t0ken/cmd/t0ken
