package main

import "fmt"

type Account struct {
	Owner   string
	Balance int
}

func findAccount(users []Account, name string) int {
	for i := 0; i < len(users); i++ {
		if users[i].Owner == name {
			return i
		}
	}
	return -1
}

func main() {

	users := []Account{
		{Owner: "Shakil", Balance: 4000},
		{Owner: "Fahad", Balance: 5000},
		{Owner: "Rahim", Balance: 7000},
	}

	user := findAccount(users, "Shakil")
	if user != -1 {
		users[user].Balance += 500
		fmt.Printf("\n%s balance: %d", users[user].Owner, users[user].Balance)
	}

	user = findAccount(users, "Rahim")
	ammount := 60000
	if user != -1 {
		if users[user].Balance >= ammount {
			users[user].Balance -= ammount
			fmt.Printf("\n%s balance: %d", users[user].Owner, users[user].Balance)
		} else {
			fmt.Println(" \ninffsfficient money")
		}
	}

	for _, value := range users {
		fmt.Printf("\n%s users balance is: %d", value.Owner, value.Balance)
	}

}
