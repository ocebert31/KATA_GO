package main

import (
	"fmt"
)

func sumOfArray() {
	allNumbers := []int{1, 5, -2, 3, 5, 5, -2, 6, 9, 6}
	countArray := 0
	for i := 0; i < len(allNumbers); i++ {
		countArray += allNumbers[i]
    }
	fmt.Println(countArray)
}