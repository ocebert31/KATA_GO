package main

import (
	"fmt"
	"sort"
)

func countingDuplicates() {
	ascNumbers := sortAscendingNumbers()
	resultCountDuplicate := calculateDuplicateNumber(ascNumbers);
	result := fmt.Sprintf("Vous avez %d doublons.", resultCountDuplicate)
	fmt.Println(result)
}

func sortAscendingNumbers() []int {
	allNumbers := []int{1, 5, -2, 3, 5, 5, -2, 6, 9, 6}
	sort.Ints(allNumbers)
	return allNumbers
}

func calculateDuplicateNumber(ascNumbers []int) int {
	var countDuplicate = 0
	for i := 0; i < len(ascNumbers)-1 ;i++  {
		countDuplicate = countDuplicatesIfFirstOccurrence(ascNumbers, countDuplicate, i)
    }
	return countDuplicate
}

func countDuplicatesIfFirstOccurrence(ascNumbers []int, countDuplicate int, i int)int {
	isDuplicate := ascNumbers[i] == ascNumbers[i+1]
	isFirstOccurrence  := i == 0 || ascNumbers[i-1] != ascNumbers[i+1]
	if isDuplicate  && isFirstOccurrence  {
		countDuplicate = countDuplicate + 1
	}
	return countDuplicate;
}