package main

import (
	"fmt"
	"go_practice/accounts"
	"go_practice/dictionary"
)

func main() {
	// Account Example
	account := accounts.NewAccount("juyoung")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account)

	// Dictionary Example
	dict := dictionary.Dictionary{}
	word := "BTS"
	dict.Add(word,"Best Boy Group")
	dict.Update("t","World Boy Group")
	definition, err := dict.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found", word, " : ", definition)
	}
	dict.Delete(word)
	definition, err = dict.Search(word)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found", word, " : ", definition)
	}
}

