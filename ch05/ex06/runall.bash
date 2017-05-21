#!/usr/bin/env bash

go run main.go > out/out.svg
diff examples/out.svg out/out.svg
