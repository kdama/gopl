#!/usr/bin/env bash

find . -name "*.txt" | xargs rm -f
go run main.go http://pi.karmona.com
diff -r out/1 out/2
