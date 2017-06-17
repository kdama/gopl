# testall.bash

次は、プロセッサ 8 個のコンピュータでベンチマークを実行した結果です。
BenchmarkRender*N*-8 で始まる行は、*N* 個のゴルーチンを使用していることを表しています。

最も性能が良かったのは、64～16384 個のゴルーチンを使用した場合でした。
64 個よりもゴルーチンの数が少ないと性能は悪くなりました。
また、ゴルーチンの数を増やしすぎても性能は悪くなりました。

```txt
BenchmarkRender1-8         	       5	 307461480 ns/op
BenchmarkRender2-8         	      10	 153530700 ns/op
BenchmarkRender4-8         	      10	 133126620 ns/op
BenchmarkRender8-8         	      10	 110922180 ns/op
BenchmarkRender16-8        	      20	  86292255 ns/op
BenchmarkRender32-8        	      20	  78940785 ns/op
BenchmarkRender64-8        	      20	  71139225 ns/op
BenchmarkRender128-8       	      20	  69063810 ns/op
BenchmarkRender256-8       	      20	  74364870 ns/op
BenchmarkRender512-8       	      20	  74364870 ns/op
BenchmarkRender1024-8      	      20	  67338465 ns/op
BenchmarkRender2048-8      	      20	  67888575 ns/op
BenchmarkRender4096-8      	      20	  68738745 ns/op
BenchmarkRender8192-8      	      20	  69763950 ns/op
BenchmarkRender16384-8     	      20	  71639325 ns/op
BenchmarkRender32768-8     	      20	  75090015 ns/op
BenchmarkRender65536-8     	      20	  79690935 ns/op
BenchmarkRender131072-8    	      20	  89767950 ns/op
BenchmarkRender262144-8    	      10	 119073810 ns/op
BenchmarkRender524288-8    	      10	 177685530 ns/op
BenchmarkRender1048576-8   	       5	 316363260 ns/op
PASS
ok  	github.com/kdama/gopl/ch08/ex05	35.321s
```
