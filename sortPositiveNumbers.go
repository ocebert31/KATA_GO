package main

import "fmt"

func sortPositiveNumbers() {
	allNumbers := []int{1, 2, 3, -5, 0, -1, 8}
    positiveNumbers := []int{}
    for i := 0; i < len(allNumbers); i++ {
        positiveNumbers = filterPositiveNumbers(allNumbers, positiveNumbers, i)
    }
    fmt.Println(positiveNumbers)
}

func filterPositiveNumbers(allNumbers []int, positiveNumbers []int, i int) []int {
    if allNumbers[i] >= 0 {
        //Elle prend en premier argument le slice auquel vous voulez ajouter des éléments, et ensuite les éléments à ajouter.
        // nouveauTableau = append(nouveauTableau, monTableau[i]) 
        positiveNumbers = append(positiveNumbers, allNumbers[i]) 
    }
    return positiveNumbers;
}