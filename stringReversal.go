package main

import (
	"fmt"
	"strings"
)

func stringReversal() {
	valueUserResult := getUserValue()
	inverseResult := reverseCharacters(valueUserResult)
	fmt.Println("Chaîne inversée :", inverseResult) 
}

func getUserValue() string{
	var valueUser string
    fmt.Print("Entrez une valeur : ")
    fmt.Scanln(&valueUser)
	return valueUser
}

func reverseCharacters(valueUserResult string) string{
	result := strings.Split(valueUserResult, "")
	var inverse []string
	for i := len(result)-1; i >= 0 ;i--  {
		inverse = append(inverse, result[i])
    }
	return strings.Join(inverse, "")
}