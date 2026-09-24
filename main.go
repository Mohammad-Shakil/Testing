package main

import "fmt"

func main() {

	fruits := []string{"mango", "apple", "tometo", "pineapple", "lemon", "jackfruit"}

	fmt.Println("Before:", fruits)
	index := -1
	target := "tometo"
	for i := range fruits {

		if fruits[i] == target {
			index = i
		}
	}
	if index != -1 {
		for i := index; i < len(fruits)-1; i++ {
			fruits[i] = fruits[i+1]

		}
	} else {
		fmt.Printf("ERROR NO: %s ", target)
		return
	}
	var newfruit []string
	for i := 0; i < len(fruits)-1; i++ {
		newfruit = append(newfruit, fruits[i])
	}
	fmt.Println("After: ", newfruit)

}
