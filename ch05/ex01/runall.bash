#!/usr/bin/env bash

# go get gopl.io/ch1/fetch
# go build gopl.io/ch1/fetch
# ./fetch https://golang.org | go run main.go

go run main.go < examples/in.html > out/out.txt
diff examples/out.txt out/out.txt
