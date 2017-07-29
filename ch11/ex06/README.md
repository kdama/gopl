# testall.bash

ベンチマークの結果は次の通り。2^0 から 2^16 の間のどのサイズでも、表に基づく方法が最も高速で、次に LSB をクリアする練習問題 2.5 の方法が速く、練習問題 2.4 の方法が最も遅かった。

```
BenchmarkTablePopCount0-4               2000000000               1.16 ns/op
BenchmarkTablePopCount1-4               1000000000               2.25 ns/op
BenchmarkTablePopCount2-4               300000000                4.24 ns/op
BenchmarkTablePopCount4-4               100000000               16.6 ns/op
BenchmarkTablePopCount8-4                5000000               265 ns/op
BenchmarkTablePopCount16-4                 20000             67701 ns/op
BenchmarkBitShiftPopCount0-4            2000000000               1.43 ns/op
BenchmarkBitShiftPopCount1-4            20000000                94.1 ns/op
BenchmarkBitShiftPopCount2-4             5000000               295 ns/op
BenchmarkBitShiftPopCount4-4             1000000              1605 ns/op
BenchmarkBitShiftPopCount8-4               50000             31116 ns/op
BenchmarkBitShiftPopCount16-4                200           8342895 ns/op
BenchmarkLSBPopCount0-4                 2000000000               1.47 ns/op
BenchmarkLSBPopCount1-4                 500000000                3.81 ns/op
BenchmarkLSBPopCount2-4                 200000000                8.08 ns/op
BenchmarkLSBPopCount4-4                 30000000                41.5 ns/op
BenchmarkLSBPopCount8-4                  1000000              1043 ns/op
BenchmarkLSBPopCount16-4                    3000            429576 ns/op
```