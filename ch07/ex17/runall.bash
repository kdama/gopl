#!/usr/bin/env bash

# go get gopl.io/ch1/fetch
# go build gopl.io/ch1/fetch
# ./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | go run main.go div div h2

go run main.go div class=d1 p id=p2 < examples/in.html
