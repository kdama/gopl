#!/usr/bin/env bash

# almost the same
go run main.go -x -2 -zoom 15 -type complex64  > out/complex64_z15.png
go run main.go -x -2 -zoom 15 -type complex128 > out/complex128_z15.png

# different
go run main.go -x -2 -zoom 16 -type complex64  > out/complex64_z16.png
go run main.go -x -2 -zoom 16 -type complex128 > out/complex128_z16.png

# almost the same
go run main.go -x -2 -zoom 44 -type complex128 > out/complex128_z44.png
go run main.go -x -2 -zoom 44 -type bigfloat   > out/bigfloat_z44.png

# different
go run main.go -x -2 -zoom 45 -type complex128 > out/complex128_z45.png
go run main.go -x -2 -zoom 45 -type bigfloat   > out/bigfloat_z45.png

# almost the same
go run main.go -x -2 -zoom 52 -iterations 1 -type bigfloat > out/bigfloat_z52.png
go run main.go -x -2 -zoom 52 -iterations 1 -type bigrat   > out/bigrat_z52.png

# different
go run main.go -x -2 -zoom 53 -iterations 1 -type bigfloat > out/bigfloat_z53.png
go run main.go -x -2 -zoom 53 -iterations 1 -type bigrat   > out/bigrat_z53.png
