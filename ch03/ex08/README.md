# ch03/ex08

## 性能とメモリ使用量

性能とメモリ使用量は、complex64 を 1 として、次のようになった。
小さいほど良い。

```txt
            Performance     Memory
complex64          1.00       1.00
complex128         0.89       1.00
big.Float         95.22   33040.00
big.Rat          168.39   55872.03
```

```txt
$ go test -benchmem -bench=.
BenchmarkMandelbrotComplex64-4            100000             18338 ns/op              30 B/op         25 allocs/op
BenchmarkMandelbrotComplex128-4           100000             16400 ns/op              30 B/op         25 allocs/op
BenchmarkMandelbrotBigFloat-4               1000           1746090 ns/op          991200 B/op      11827 allocs/op
BenchmarkMandelbrotBigRat-4                  500           3087937 ns/op         1676161 B/op      36621 allocs/op
PASS
ok      github.com/kdama/gopl/ch03/ex08 7.639s
```

## レンダリング結果

2^15 倍の拡大では、complex64 と complex128 のレンダリング結果に違いは確認できなかった。(`examples/complex64_z15.png`, `examples/complex128_z15.png`)
2^16 倍の拡大では、 complex64 と complex128 のレンダリング結果に違いを確認できた。(`examples/complex64_z16.png`, `examples/complex128_z16.png`)

2^44 倍の拡大では、complex128 と big.Float のレンダリング結果に違いは確認できなかった。(`examples/complex128_z44.png`, `examples/bigfloat_z44.png`)
2^45 倍の拡大では、complex128 と big.Float のレンダリング結果に違いを確認できた。(`examples/complex128_z45.png`, `examples/bigfloat_z45.png`)

big.Float と big.Rat のレンダリング結果の差は、詳細には得られなかった。
2^52 倍の拡大では、big.Float と big.Rat のレンダリング結果に違いは確認できなかった。(`examples/bigfloat_z52.png`, `examples/bigrat_z52.png`)
2^53 倍の拡大では、big.Float と big.Rat のレンダリング結果に違いを確認できた。(`examples/bigfloat_z53.png`, `examples/bigrat_z53.png`)
ただし、
2^53 倍の拡大においては、big.Float, big.Rat ともに精度の限界を越えたような明らかに間違ったレンダリング結果が得られている。
