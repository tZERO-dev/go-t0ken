#!/bin/sh

xgo \
	-v \
	--targets="darwin/amd64,linux/amd64,windows/amd64" \
	--dest=dist \
	github.com/tzero-dev/go-t0ken/cmd/t0ken
