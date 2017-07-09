# testall.bash

次は、プロセッサ 2 個のコンピュータでベンチマークを実行した結果である。

`GOMAXPROCS=1` のとき、最も性能が良かったのは 16384 個のゴルーチンを使用した場合で、0.633 秒であった。
`GOMAXPROCS=2` のとき、最も性能が良かったのは 2 個のゴルーチンを使用した場合で、0.393 秒であった。
`GOMAXPROCS=4` のとき、最も性能が良かったのは 512 個のゴルーチンを使用した場合で、0.423 秒であった。
`GOMAXPROCS=8` のとき、最も性能が良かったのは 128 個のゴルーチンを使用した場合で、0.417 秒であった。

## GOMAXPROCS=1

```txt
BenchmarkRender1       	       1	1031848790 ns/op
BenchmarkRender2       	       1	1069674571 ns/op
BenchmarkRender4       	       1	1130412337 ns/op
BenchmarkRender8       	       2	 702081394 ns/op
BenchmarkRender16      	       2	1359770949 ns/op
BenchmarkRender32      	       1	1328384510 ns/op
BenchmarkRender64      	       2	 663993200 ns/op
BenchmarkRender128     	       2	 838262406 ns/op
BenchmarkRender256     	       2	 741063365 ns/op
BenchmarkRender512     	       2	 676512901 ns/op
BenchmarkRender1024    	       2	 737423574 ns/op
BenchmarkRender2048    	       2	 804697270 ns/op
BenchmarkRender4096    	       2	 759873746 ns/op
BenchmarkRender8192    	       2	 891444459 ns/op
BenchmarkRender16384   	       2	 633448360 ns/op
BenchmarkRender32768   	       2	 643524170 ns/op
BenchmarkRender65536   	       2	1333691770 ns/op
BenchmarkRender131072  	       1	1475505820 ns/op
BenchmarkRender262144  	       1	1373965478 ns/op
BenchmarkRender524288  	       1	1560809903 ns/op
BenchmarkRender1048576 	       1	2608771388 ns/op
PASS
ok  	github.com/kdama/gopl/ch09/ex06	43.381s
```

## GOMAXPROCS=2

```txt
BenchmarkRender1-2         	       2	1112117313 ns/op
BenchmarkRender2-2         	       3	 393419435 ns/op
BenchmarkRender4-2         	       2	 736368213 ns/op
BenchmarkRender8-2         	       2	1352373516 ns/op
BenchmarkRender16-2        	       1	1653094619 ns/op
BenchmarkRender32-2        	       3	 522031045 ns/op
BenchmarkRender64-2        	       3	 525613106 ns/op
BenchmarkRender128-2       	       3	 768987224 ns/op
BenchmarkRender256-2       	       2	 590673321 ns/op
BenchmarkRender512-2       	       3	 704121911 ns/op
BenchmarkRender1024-2      	       2	 510017329 ns/op
BenchmarkRender2048-2      	       2	1016409982 ns/op
BenchmarkRender4096-2      	       1	2678278632 ns/op
BenchmarkRender8192-2      	       1	1058428164 ns/op
BenchmarkRender16384-2     	       3	 439379795 ns/op
BenchmarkRender32768-2     	       2	 541122391 ns/op
BenchmarkRender65536-2     	       2	 517492745 ns/op
BenchmarkRender131072-2    	       2	 888025610 ns/op
BenchmarkRender262144-2    	       1	2868075135 ns/op
BenchmarkRender524288-2    	       1	1427782443 ns/op
BenchmarkRender1048576-2   	       1	1873798345 ns/op
PASS
ok  	github.com/kdama/gopl/ch09/ex06	50.297s
```

## GOMAXPROCS=4

```txt
BenchmarkRender1-4         	       2	1058569137 ns/op
BenchmarkRender2-4         	       1	1101955595 ns/op
BenchmarkRender4-4         	       2	 610677812 ns/op
BenchmarkRender8-4         	       2	 587184872 ns/op
BenchmarkRender16-4        	       2	 560835417 ns/op
BenchmarkRender32-4        	       2	 631555096 ns/op
BenchmarkRender64-4        	       2	1297750673 ns/op
BenchmarkRender128-4       	       3	 424851357 ns/op
BenchmarkRender256-4       	       3	 432132119 ns/op
BenchmarkRender512-4       	       3	 423621461 ns/op
BenchmarkRender1024-4      	       3	 473986820 ns/op
BenchmarkRender2048-4      	       3	 463048065 ns/op
BenchmarkRender4096-4      	       3	 536973122 ns/op
BenchmarkRender8192-4      	       3	 469872715 ns/op
BenchmarkRender16384-4     	       2	 685576088 ns/op
BenchmarkRender32768-4     	       3	 466562243 ns/op
BenchmarkRender65536-4     	       2	 537890473 ns/op
BenchmarkRender131072-4    	       2	1010229439 ns/op
BenchmarkRender262144-4    	       1	1664215512 ns/op
BenchmarkRender524288-4    	       1	1000732538 ns/op
BenchmarkRender1048576-4   	       1	1369003732 ns/op
PASS
ok  	github.com/kdama/gopl/ch09/ex06	48.261s
```

## GOMAXPROCS=8

```txt
BenchmarkRender1-8         	       2	1037608338 ns/op
BenchmarkRender2-8         	       1	1677544075 ns/op
BenchmarkRender4-8         	       2	 581534481 ns/op
BenchmarkRender8-8         	       3	 485808520 ns/op
BenchmarkRender16-8        	       3	 453192931 ns/op
BenchmarkRender32-8        	       3	 448486249 ns/op
BenchmarkRender64-8        	       3	 425978336 ns/op
BenchmarkRender128-8       	       3	 417876620 ns/op
BenchmarkRender256-8       	       3	 532778174 ns/op
BenchmarkRender512-8       	       3	 419296751 ns/op
BenchmarkRender1024-8      	       3	 423441118 ns/op
BenchmarkRender2048-8      	       2	 698810034 ns/op
BenchmarkRender4096-8      	       3	 466537922 ns/op
BenchmarkRender8192-8      	       2	 526384314 ns/op
BenchmarkRender16384-8     	       2	 511459012 ns/op
BenchmarkRender32768-8     	       2	 537766316 ns/op
BenchmarkRender65536-8     	       2	 536562305 ns/op
BenchmarkRender131072-8    	       2	1026339336 ns/op
BenchmarkRender262144-8    	       2	 688765860 ns/op
BenchmarkRender524288-8    	       2	1124110787 ns/op
BenchmarkRender1048576-8   	       1	1584840678 ns/op
PASS
ok  	github.com/kdama/gopl/ch09/ex06	49.059s
```
