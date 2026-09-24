package main

import (
	"fmt"
	"mytesting/logics"
)

func main() {

	fmt.Println("\n----String reversal machine----")
	fmt.Println()
	text := []string{"madam", "hello", "lili", "level", "SOS"}

	for i := 0; i < len(text)-1; i++ {

		reversed := logics.Reversal(text[i])
		if reversed == text[i] {
			fmt.Printf("\n%s: is palindrome \n", reversed)
		} else {
			fmt.Printf("\n%s is not plindrome\n", reversed)
		}

	}

}
