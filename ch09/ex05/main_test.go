package main

import (
	"testing"
)

func BenchmarkPingpong(b *testing.B) {
	alice := make(chan interface{})
	bob := make(chan interface{})
	done := make(chan struct{})
	go func() {
		for i := 0; i < b.N; i++ {
			bob <- 42
			<-alice
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
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
