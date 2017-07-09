// ch09/ex05 は、二つのバッファなしチャネルを介してメッセージを送り合うプログラムです。
package main

const n = 10

func main() {
	alice := make(chan interface{})
	bob := make(chan interface{})
	done := make(chan struct{})
	go func() {
		for i := 0; i < n; i++ {
			bob <- 42
			<-alice
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < n; i++ {
			<-bob
			alice <- 42
		}
		done <- struct{}{}
	}()
	<-done
	<-done
	close(alice)
	close(bob)
	close(done)
}
