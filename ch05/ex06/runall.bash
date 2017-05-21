#!/usr/bin/env bash

go run main.go > out/out.svg

# Compare output of gopl.io/ch3/surface
diff examples/out.svg out/out.svg
