# testall.bash

次の表は、ワードの大きさを 32 バイトに設定した IntSet において掛かった時間を 1.00 として、ワードの大きさを 64 バイトに設定した IntSet と、マップに基づくセット MapSet において掛かった時間を示している。
DifferenceWith のような、2 つの集合に関する計算については、IntSet64 の性能が最も良く、MapSet の性能が著しく悪かった。

|                     | IntSet32 | IntSet64 | MapSet   |
|---------------------|----------|----------|----------|
| Add                 |     1.00 |     0.98 |    15.45 |
| AddAll              |     1.00 |     0.99 |    28.70 |
| Clear               |     1.00 |     0.95 |    16.51 |
| Copy                |     1.00 |     0.49 |   885.69 |
| DifferenceWith      |     1.00 |     0.56 |   689.74 |
| Elems               |     1.00 |     0.94 |    27.76 |
| Has                 |     1.00 |     0.96 |    29.99 |
| InteresectWith      |     1.00 |     0.68 |   885.65 |
| Len                 |     1.00 |     0.84 |     0.00 |
| Remove              |     1.00 |     0.96 |     0.94 |
| SymmetricDifference |     1.00 |     0.59 |   462.57 |
| UnionWith           |     1.00 |     0.47 |   188.47 |

ベンチマークの出力は次の通り。

```
BenchmarkIntSet32Add-4                              5000            277449 ns/op
BenchmarkIntSet32AddAll-4                           2000            648814 ns/op
BenchmarkIntSet32Clear-4                        200000000                6.42 ns/op
BenchmarkIntSet32Copy-4                           200000             10364 ns/op
BenchmarkIntSet32DifferenceWith-4                 500000              2496 ns/op
BenchmarkIntSet32Elems-4                            3000            478009 ns/op
BenchmarkIntSet32Has-4                             20000             59068 ns/op
BenchmarkIntSet32IntersectWith-4                  500000              2295 ns/op
BenchmarkIntSet32Len-4                            100000             19554 ns/op
BenchmarkIntSet32Remove-4                          10000            159785 ns/op
BenchmarkIntSet32SymmetricDifference-4            500000              2400 ns/op
BenchmarkIntSet32UnionWith-4                      300000              3874 ns/op
BenchmarkIntSet64Add-4                              5000            272224 ns/op
BenchmarkIntSet64AddAll-4                           2000            641635 ns/op
BenchmarkIntSet64Clear-4                        200000000                6.10 ns/op
BenchmarkIntSet64Copy-4                           300000              5097 ns/op
BenchmarkIntSet64DifferenceWith-4                1000000              1394 ns/op
BenchmarkIntSet64Elems-4                            3000            450034 ns/op
BenchmarkIntSet64Has-4                             30000             57000 ns/op
BenchmarkIntSet64IntersectWith-4                 1000000              1565 ns/op
BenchmarkIntSet64Len-4                            100000             16419 ns/op
BenchmarkIntSet64Remove-4                          10000            154045 ns/op
BenchmarkIntSet64SymmetricDifference-4           1000000              1410 ns/op
BenchmarkIntSet64UnionWith-4                     1000000              1814 ns/op
BenchmarkMapSetAdd-4                                 300           4285646 ns/op
BenchmarkMapSetAddAll-4                              100          18618812 ns/op
BenchmarkMapSetClear-4                          20000000               106 ns/op
BenchmarkMapSetCopy-4                                200           9179242 ns/op
BenchmarkMapSetDifferenceWith-4                     1000           1721590 ns/op
BenchmarkMapSetElems-4                               100          13267311 ns/op
BenchmarkMapSetHas-4                                1000           1771228 ns/op
BenchmarkMapSetIntersectWith-4                       500           2032573 ns/op
BenchmarkMapSetLen-4                            2000000000               0.60 ns/op
BenchmarkMapSetRemove-4                            10000            150888 ns/op
BenchmarkMapSetSymmetricDifference-4                 300           5910172 ns/op
BenchmarkMapSetUnionWith-4                           300           4604139 ns/op
```
