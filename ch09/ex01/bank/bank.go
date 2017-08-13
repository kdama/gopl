// Package bank は、一つの口座を持つ並行的に安全な銀行を提供します。
package bank

type withdrawMessage struct {
	amount     int
	sufficient chan<- bool
}

var deposits = make(chan int)              // send amount to deposit
var balances = make(chan int)              // receive balance
var withdraws = make(chan withdrawMessage) // send amount to withdraw

// Deposit は、指定された金額を口座に預け入れます。
func Deposit(amount int) {
	deposits <- amount
}

// Withdraw は、指定された金額を口座から引き出します。
func Withdraw(amount int) bool {
	sufficient := make(chan bool)
	withdraws <- withdrawMessage{amount, sufficient}
	return <-sufficient
}

// Balance は、口座の残高を返します。
func Balance() int {
	return <-balances
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case msg := <-withdraws:
			ok := balance >= msg.amount
			if ok {
				balance -= msg.amount
			}
			msg.sufficient <- ok
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
