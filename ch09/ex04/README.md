# testall.bash

RAM 1 GB の環境では、メモリ不足になることなく作成できるパイプラインの最大段数は 262144 段であった。
また、262144 段のパイプラインにおいて、一つの値が全体を伝わるのに要した時間は 1.7 秒であった。

`./testall.bash` の出力を次に示す。

```
BenchmarkPipeline1-2         	 1000000	      1073 ns/op	       8 B/op	       1 allocs/op
BenchmarkPipeline2-2         	 1000000	      1282 ns/op	       8 B/op	       1 allocs/op
BenchmarkPipeline4-2         	 1000000	      2770 ns/op	       8 B/op	       1 allocs/op
BenchmarkPipeline8-2         	  500000	      6068 ns/op	       8 B/op	       1 allocs/op
BenchmarkPipeline16-2        	  200000	      7625 ns/op	       8 B/op	       1 allocs/op
BenchmarkPipeline32-2        	  100000	     15997 ns/op	       8 B/op	       1 allocs/op
BenchmarkPipeline64-2        	   50000	     50878 ns/op	       8 B/op	       1 allocs/op
BenchmarkPipeline128-2       	   10000	    145881 ns/op	      13 B/op	       1 allocs/op
BenchmarkPipeline256-2       	   10000	    155463 ns/op	      22 B/op	       1 allocs/op
BenchmarkPipeline512-2       	   10000	    318187 ns/op	      30 B/op	       1 allocs/op
BenchmarkPipeline1024-2      	    2000	    586682 ns/op	     295 B/op	       2 allocs/op
BenchmarkPipeline2048-2      	    1000	   1575388 ns/op	     216 B/op	       3 allocs/op
BenchmarkPipeline4096-2      	     500	   3571316 ns/op	    1115 B/op	       9 allocs/op
BenchmarkPipeline8192-2      	     200	   7351650 ns/op	    8632 B/op	      51 allocs/op
BenchmarkPipeline16384-2     	     100	  20255402 ns/op	   19516 B/op	     177 allocs/op
BenchmarkPipeline32768-2     	      50	  36272431 ns/op	   64015 B/op	     670 allocs/op
BenchmarkPipeline65536-2     	      20	 103393327 ns/op	  390858 B/op	    3592 allocs/op
BenchmarkPipeline131072-2    	      10	 148866744 ns/op	 1403854 B/op	   13773 allocs/op
BenchmarkPipeline262144-2    	       1	1755310811 ns/op	98547424 B/op	  518884 allocs/op
```
