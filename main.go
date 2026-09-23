package main

import "fmt"

func stringreverse(p string) string {

	runes := []int32(p)

	low := 0
	high := len(runes) - 1

	for low < high {
		runes[low], runes[high] = runes[high], runes[low]
		low++
		high--
	}
	return string(runes)
}

func main() {

	fmt.Println(stringreverse("shakil"))
}
