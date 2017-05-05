package main

import (
	"math/big"
	"testing"

	"github.com/kdama/gopl/ch03/ex08/complexfloat"
	"github.com/kdama/gopl/ch03/ex08/complexrat"
)

func BenchmarkMandelbrotComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := float32(-2); y <= 2; y++ {
			for x := float32(-2); x <= 2; x++ {
				mandelbrotComplex64(complex(x, y))
			}
		}
	}
}

func BenchmarkMandelbrotComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := float64(-2); y <= 2; y++ {
			for x := float64(-2); x <= 2; x++ {
				mandelbrotComplex128(complex(x, y))
			}
		}
	}
}

func BenchmarkMandelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := float64(-2); y <= 2; y++ {
			for x := float64(-2); x <= 2; x++ {
				mandelbrotBigFloat(&complexfloat.ComplexFloat{big.NewFloat(x), big.NewFloat(y), 1024})
			}
		}
	}
}

func BenchmarkMandelbrotBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := int64(-2); y <= 2; y++ {
			for x := int64(-2); x <= 2; x++ {
				mandelbrotBigRat(&complexrat.ComplexRat{big.NewRat(x, 1), big.NewRat(y, 1)})
			}
		}
	}
}
