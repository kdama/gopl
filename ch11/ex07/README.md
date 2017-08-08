# testall.bash

次の表は、ワードの大きさを 16 ビットに設定したセット (BitSet16) において掛かった時間を 1 として、ワードの大きさを 32 ビット、64 ビットに設定した BitSet と、マップに基づくセット MapSet において掛かった時間を示している。

|                     | BitSet16 | BitSet32 | BitSet64 | MapSet   |
| ------------------- | -------- | -------- | -------- | -------- |
| Add                 |        1 |     0.99 |     0.99 |    25.36 |
| AddAll              |        1 |     0.99 |     0.98 |    24.23 |
| Clear               |        1 |     1.01 |     1.00 |    16.79 |
| Copy                |        1 |     0.73 |     0.61 |   446.15 |
| DifferenceWith      |        1 |     0.51 |     0.28 |   838.21 |
| Elems               |        1 |     0.94 |     0.95 |    14.96 |
| Has                 |        1 |     1.00 |     1.00 |    82.03 |
| IntersectWith       |        1 |     0.55 |     0.26 |   814.29 |
| Len                 |        1 |     0.67 |     0.51 |     0.00 |
| Remove              |        1 |     0.99 |     0.95 |    25.05 |
| SymmetricDifference |        1 |     0.55 |     0.26 | 1,812.87 |
| UnionWith           |        1 |     0.52 |     0.27 | 1,059.46 |

表から、次の性質が分かる。

- BitSet については、どのような演算においても、BitSet16 よりも BitSet32 のほうが性能が良く、BitSet32 よりも BitSet64 のほうが性能が良い。
- Len を除くすべての演算で、MapSet の性能は悪い。特に、DifferenceWith のような 2 つの集合に関する計算の性能が著しく悪い。

ベンチマークの出力は次の通り。

```
BenchmarkBitSet16Add-4                                30          48574741 ns/op
BenchmarkBitSet16AddAll-4                             10         116247776 ns/op
BenchmarkBitSet16Clear-4                        300000000                5.67 ns/op
BenchmarkBitSet16Copy-4                              500           2748983 ns/op
BenchmarkBitSet16DifferenceWith-4                   2000            853575 ns/op
BenchmarkBitSet16Elems-4                              20         101473959 ns/op
BenchmarkBitSet16Has-4                               200           7139585 ns/op
BenchmarkBitSet16IntersectWith-4                    2000            785878 ns/op
BenchmarkBitSet16Len-4                               100          10493501 ns/op
BenchmarkBitSet16Remove-4                             20          81804526 ns/op
BenchmarkBitSet16SymmetricDifference-4              2000            781388 ns/op
BenchmarkBitSet16UnionWith-4                        1000           1175321 ns/op
BenchmarkBitSet32Add-4                                30          48199604 ns/op
BenchmarkBitSet32AddAll-4                             10         114729479 ns/op
BenchmarkBitSet32Clear-4                        300000000                5.72 ns/op
BenchmarkBitSet32Copy-4                             1000           1999194 ns/op
BenchmarkBitSet32DifferenceWith-4                   3000            431406 ns/op
BenchmarkBitSet32Elems-4                              20          95525818 ns/op
BenchmarkBitSet32Has-4                               200           7122982 ns/op
BenchmarkBitSet32IntersectWith-4                    3000            429872 ns/op
BenchmarkBitSet32Len-4                               200           7034466 ns/op
BenchmarkBitSet32Remove-4                             20          80948434 ns/op
BenchmarkBitSet32SymmetricDifference-4              3000            430235 ns/op
BenchmarkBitSet32UnionWith-4                        2000            605338 ns/op
BenchmarkBitSet64Add-4                                30          47928553 ns/op
BenchmarkBitSet64AddAll-4                             10         114323131 ns/op
BenchmarkBitSet64Clear-4                        300000000                5.67 ns/op
BenchmarkBitSet64Copy-4                             1000           1689454 ns/op
BenchmarkBitSet64DifferenceWith-4                   5000            241746 ns/op
BenchmarkBitSet64Elems-4                              20          96458791 ns/op
BenchmarkBitSet64Has-4                               200           7115726 ns/op
BenchmarkBitSet64IntersectWith-4                    5000            205683 ns/op
BenchmarkBitSet64Len-4                               300           5356103 ns/op
BenchmarkBitSet64Remove-4                             20          79429189 ns/op
BenchmarkBitSet64SymmetricDifference-4              5000            202963 ns/op
BenchmarkBitSet64UnionWith-4                        5000            320187 ns/op
BenchmarkMapSetAdd-4                                   1        1231970556 ns/op
BenchmarkMapSetAddAll-4                                1        2816510242 ns/op
BenchmarkMapSetClear-4                          20000000                95.2 ns/op
BenchmarkMapSetCopy-4                                  1        1226467119 ns/op
BenchmarkMapSetDifferenceWith-4                        2         715476271 ns/op
BenchmarkMapSetElems-4                                 1        1517646265 ns/op
BenchmarkMapSetHas-4                                   2         585687783 ns/op
BenchmarkMapSetIntersectWith-4                         2         639930156 ns/op
BenchmarkMapSetLen-4                            2000000000               0.54 ns/op
BenchmarkMapSetRemove-4                                1        2064251811 ns/op
BenchmarkMapSetSymmetricDifference-4                   1        1416554046 ns/op
BenchmarkMapSetUnionWith-4                             1        1245203210 ns/op
```
