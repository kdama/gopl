// ch09/ex04 は、任意の数のゴルーチンをチャネルで接続するパイプラインです。
package main

import "fmt"

func main() {
	in, out := pipeline(10)
	in <- 42
	fmt.Println(<-out)
}

func pipeline(n int) (chan<- interface{}, <-chan interface{}) {
	if n < 1 {
		return nil, nil
	}
	in := make(chan interface{})
	out := in
	for i := 1; i < n; i++ {
		prev := out
		next := make(chan interface{})
		go func() {
			for val := range prev {
				next <- val
			}
			close(next)
		}()
		out = next
	}
	return in, out
}
