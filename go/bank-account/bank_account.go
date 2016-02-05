package account

import "sync"

type Account struct {
	amt   int64
	vaild bool
	sync.RWMutex
}

func Open(initalDeposit int64) *Account {
	if initalDeposit >= 0 {
		return &Account{
			amt:   initalDeposit,
			vaild: true,
		}
	}
	return nil
}

func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	ok, a.vaild = a.vaild, false
	payout, a.amt = a.amt, 0
	return
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()
	if a.amt+amount < 0 {
		return 0, false
	}
	a.amt += amount
	return a.amt, a.vaild
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()
	return a.amt, a.vaild
}
