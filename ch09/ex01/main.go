// ex09/ch01 は、引き出しが可能な銀行プログラムです。
package main

import (
	"fmt"

	"github.com/kdama/gopl/ch09/ex01/bank"
)

func main() {
	done := make(chan struct{})

	go func() {
		fmt.Println("Deposit 100")
		bank.Deposit(100)
		done <- struct{}{}
	}()

	go func() {
		fmt.Println("Deposit 200")
		bank.Deposit(200)
		done <- struct{}{}
	}()

	go func() {
		fmt.Printf("Withdraw 100 OK? -> %t\n", bank.Withdraw(200))
		done <- struct{}{}
	}()

	go func() {
		fmt.Printf("Withdraw 200 OK? -> %t\n", bank.Withdraw(200))
		done <- struct{}{}
	}()

	<-done
	<-done
	<-done
	<-done

	fmt.Printf("\nFinal balance: %d\n", bank.Balance())
}
