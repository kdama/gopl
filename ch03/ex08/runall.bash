#!/usr/bin/env bash

go run main.go -x -2 -zoom 16 -type complex64  > out/complex64_z16.png
go run main.go -x -2 -zoom 16 -type complex128 > out/complex128_z16.png

go run main.go -x -2 -zoom 45 -type complex128 > out/complex128_z45.png
go run main.go -x -2 -zoom 45 -type bigfloat   > out/bigfloat_z45.png

go run main.go -x -2 -zoom 53 -iterations 32 -type bigfloat > out/bigfloat_z53.png
go run main.go -x -2 -zoom 53 -iterations 32 -type bigrat   > out/bigrat_z53.png
