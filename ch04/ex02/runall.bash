#!/usr/bin/env bash

go run main.go < in/in.txt
go run main.go -mode sha256 < in/in.txt
go run main.go -mode sha384 < in/in.txt
go run main.go -mode sha512 < in/in.txt
