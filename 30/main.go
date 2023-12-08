package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
	lock    sync.Mutex
}

var print func(a ...any) (n int, err error) = fmt.Println
var printf func(format string, a ...any) (n int, err error) = fmt.Printf

func (acc *Account) GetBalance() int {
	acc.lock.Lock()
	defer acc.lock.Unlock()
	return acc.balance
}

func (acc *Account) Withdraw(amount int) {
	acc.lock.Lock()
	defer acc.lock.Unlock()
	if amount > acc.balance {
		print("Insufficient Balance.")
	} else {
		acc.balance -= amount
		printf("Withdranwn: $%d\nBalance: $%d\n", amount, acc.balance)
	}

}

func main() {
	account := Account{balance: 250}

	for i := 1; i <= 15; i++ {
		go account.Withdraw(i)
	}

	time.Sleep(2 * time.Second)
}
