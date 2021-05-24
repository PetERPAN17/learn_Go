package accounts

// Acount struct
type account struct {
	owner   string
	balance int
}

// Creates account
func NewAccount(owner string) *account {
	newAccount := account{owner: owner, balance: 0}
	return &newAccount
}
