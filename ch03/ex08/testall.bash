#!/usr/bin/env bash

go test -v ./...
go test -benchmem -bench=.
