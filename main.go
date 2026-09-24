package main

import (
	"fmt"
	"mytesting/logics"
)

func main() {
	fruits := []string{"apple", "banana", "watermelon", "mango", "pineapple"}

	fmt.Println("Final list: ", logics.Elcheck(fruits))

}
