package main

import "fmt"

func stringReversal(p string) string {

	runes := []byte(p)

	left := 0
	right := len(runes) - 1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)

}

func main() {

	fmt.Println(stringReversal("Shakil"))
}
