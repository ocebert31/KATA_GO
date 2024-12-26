package main

import (
	"fmt"
	"strings"
)

func palindromeCheck() {
	valueUserResult := getUserValue()
	inverseResult := verifyIsPalindrome(valueUserResult)
	fmt.Println(inverseResult) 
}

func getUserValue() string{
	var valueUser string
    fmt.Print("Entrez une valeur : ")
    fmt.Scanln(&valueUser)
	return valueUser
}

func verifyIsPalindrome(valueUserResult string) string {
	initialValue := strings.Split(valueUserResult, "")
	var inverse []string
	for i := len(initialValue)-1; i >= 0 ;i--  {
		inverse = append(inverse, initialValue[i])
    }
	inverseInitialValue := strings.Join(inverse, "")
	if valueUserResult == inverseInitialValue {
		return "C'est un palindrome"
	} else {
		return "Ce n'est pas un palindrome"
	}
}