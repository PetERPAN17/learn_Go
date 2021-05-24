package main

import (
	"fmt"

	"github.com/learn_Go/packages/accounts"
)

func main() {
	account := accounts.NewAccount("TestUser")
	fmt.Println(account)
}
