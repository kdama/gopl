#!/usr/bin/env bash

# go get gopl.io/ch1/fetch
# go build gopl.io/ch1/fetch
# ./fetch https://golang.org | go run main.go

go run main.go < in/in.html > out/out.txt
diff expected/out.txt out/out.txt
