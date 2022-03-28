//beginanswer
package account

//endanswer
type Account struct {
	Name    string
	Balance int
}

func (a Account) GetBalance() int {
	//beginanswer
	return a.Balance
	//endanswer
}

func (a *Account) Deposit(amount int) {
	//beginanswer
	a.Balance += amount
	//endanswer
}
