#!/usr/bin/env bash

GOMAXPROCS=1 go test -v ./... -bench=.
GOMAXPROCS=2 go test -v ./... -bench=.
GOMAXPROCS=4 go test -v ./... -bench=.
GOMAXPROCS=8 go test -v ./... -bench=.
