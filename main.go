package main

import (
	"fmt"
	"mytesting/logics"
)

func main() {
	list := []int{23, 10, 45, 60, 20, 80, 35, 90}

	fmt.Println("Final list: ", logics.ListCompress(list))

}
