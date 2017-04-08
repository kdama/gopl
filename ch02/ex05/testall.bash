#!/usr/bin/env bash

go test
go test -bench=. -args a b c
