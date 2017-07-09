package main

import (
	"testing"
)

func benchmarkRender(b *testing.B, worker int) {
	for i := 0; i < b.N; i++ {
		render(worker)
	}
}

func BenchmarkRender1(b *testing.B) {
	benchmarkRender(b, 1)
}

func BenchmarkRender2(b *testing.B) {
	benchmarkRender(b, 2)
}

func BenchmarkRender4(b *testing.B) {
	benchmarkRender(b, 4)
}

func BenchmarkRender8(b *testing.B) {
	benchmarkRender(b, 8)
}

func BenchmarkRender16(b *testing.B) {
	benchmarkRender(b, 16)
}

func BenchmarkRender32(b *testing.B) {
	benchmarkRender(b, 32)
}

func BenchmarkRender64(b *testing.B) {
	benchmarkRender(b, 64)
}

func BenchmarkRender128(b *testing.B) {
	benchmarkRender(b, 128)
}

func BenchmarkRender256(b *testing.B) {
	benchmarkRender(b, 256)
}

func BenchmarkRender512(b *testing.B) {
	benchmarkRender(b, 512)
}

func BenchmarkRender1024(b *testing.B) {
	benchmarkRender(b, 1024)
}

func BenchmarkRender2048(b *testing.B) {
	benchmarkRender(b, 2048)
}

func BenchmarkRender4096(b *testing.B) {
	benchmarkRender(b, 4096)
}

func BenchmarkRender8192(b *testing.B) {
	benchmarkRender(b, 8192)
}

func BenchmarkRender16384(b *testing.B) {
	benchmarkRender(b, 16384)
}

func BenchmarkRender32768(b *testing.B) {
	benchmarkRender(b, 32768)
}

func BenchmarkRender65536(b *testing.B) {
	benchmarkRender(b, 65536)
}

func BenchmarkRender131072(b *testing.B) {
	benchmarkRender(b, 131072)
}

func BenchmarkRender262144(b *testing.B) {
	benchmarkRender(b, 262144)
}

func BenchmarkRender524288(b *testing.B) {
	benchmarkRender(b, 524288)
}

func BenchmarkRender1048576(b *testing.B) {
	benchmarkRender(b, 1048576)
}
