package main

import (
	"fmt"
	"sort"
)

func getSmallestNumber() {
	allNumbers := []int{1, 2, 3, -5, 0, -1, 8}
	sort.Ints(allNumbers)
	fmt.Println(allNumbers[0])
}