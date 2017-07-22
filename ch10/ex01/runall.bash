#!/usr/bin/env bash

go run main.go -format gif < examples/mandelbrot.png > out/out.gif
go run main.go -format jpg < examples/mandelbrot.png > out/out.jpg
go run main.go -format png < examples/mandelbrot.png > out/out.png
