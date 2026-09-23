package main

import "fmt"

func main() {

	students := make(map[string]int)
	students["Shakil"] = 85
	students["Rahim"] = 72
	students["Karim"] = 91
	students["Farabi"] = 65
	students["Opu"] = 88

	fmt.Println("----TOP SCORERS----")
	for name, marks := range students {
		if marks > 80 {
			fmt.Printf("\n%s: %d marks", name, marks)

		}
	}
	fmt.Println()
	fmt.Println("\n----Avegrage----")
	var total int
	result := 0
	for _, marks := range students {
		total += marks

	}
	result = total / len(students)
	fmt.Println(result)
	fmt.Println()
	fmt.Println("----Highest Mark----")

	highest := 0
	var topscorer string
	for name, marks := range students {

		if marks > highest {
			highest = marks
			topscorer = name
		}

	}
	fmt.Printf("%s got hight mark: %d", topscorer, highest)

}
