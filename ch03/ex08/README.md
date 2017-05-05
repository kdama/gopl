
-z 10000
complex64 と complex128 の違いすら分からないレベル。

-z 100000
complex64 と complex128 の違いが分かるレベル。

```
$ go test -bench=.
BenchmarkMandelbrotComplex64-2             50000             20322 ns/op
BenchmarkMandelbrotComplex128-2           100000             21998 ns/op
BenchmarkMandelbrotBigFloat-2               1000           2454715 ns/op
BenchmarkMandelbrotBigRat-2                  300          10676383 ns/op
```
