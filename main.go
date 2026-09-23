package main

import "fmt"

type Student struct {
	Name  string
	Marks int
	Grade string
}

func main() {
	students := []Student{
		{Name: "shakil", Marks: 5},
		{Name: "Fahad", Marks: 72},
		{Name: "Uzzal", Marks: 91},
		{Name: "Sumaiya", Marks: 65},
		{Name: "Prantu", Marks: 88},
	}

	for i := 0; i < len(students); i++ {
		if students[i].Marks >= 80 {
			students[i].Grade = "A+"
		} else if students[i].Marks >= 70 {
			students[i].Grade = "A"
		} else if students[i].Marks >= 60 {
			students[i].Grade = "B"
		} else {
			students[i].Grade = "F"
		}
		fmt.Printf("\n %s: %d  ---> %s", students[i].Name, students[i].Marks, students[i].Grade)

	}

}
