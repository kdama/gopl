package bank

import (
	"sync"
	"testing"
)

type BankOperation struct {
	op     string // "deposit", "withdraw", "balance"
	amount int    // ignored when op == "balance"
	ok     bool   // expect ok or not. only when op == "withdraw"
}

func TestBank(t *testing.T) {
	tests := []struct {
		operations []BankOperation
		init       int
		want       int
	}{
		{
			[]BankOperation{
				BankOperation{"withdraw", 1, true},
			},
			10,
			9,
		},
		{
			[]BankOperation{
				BankOperation{"withdraw", 1, true},
				BankOperation{"withdraw", 2, true},
				BankOperation{"withdraw", 3, true},
				BankOperation{"withdraw", 4, true},
			},
			10,
			0,
		},
		{
			[]BankOperation{
				BankOperation{"withdraw", 1, true},
				BankOperation{"withdraw", 100, false},
			},
			10,
			9,
		},
		{
			[]BankOperation{
				BankOperation{"deposit", 1, true},
				BankOperation{"deposit", 1, true},
				BankOperation{"withdraw", 1, true},
				BankOperation{"withdraw", 1, true},
				BankOperation{"balance", 1, true},
				BankOperation{"balance", 1, true},
			},
			2,
			2,
		},
	}

	for _, test := range tests {
		// init balance
		Withdraw(Balance() - test.init)

		var wg sync.WaitGroup
		for _, operation := range test.operations {
			operation := operation
			wg.Add(1)
			go func() {
				defer wg.Done()
				switch operation.op {
				case "deposit":
					Deposit(operation.amount)
				case "withdraw":
					ok := Withdraw(operation.amount)
					if ok != operation.ok {
						t.Errorf("Withdraw(%d) in %v = %t, want %t", operation.amount, test.operations, ok, operation.ok)
					}
				case "balance":
					Balance()
				}
			}()
		}
		wg.Wait()
		if got := Balance(); got != test.want {
			t.Errorf("Balance() after %v = %d, want %d", test.operations, got, test.want)
		}
	}
}
