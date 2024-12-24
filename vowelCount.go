package main

import (
	"fmt"
	"strings"
)

func vowelCount() {
	valueVowelResult := getUserValue()
	vowelTotal := vowelCountInWord(valueVowelResult)
	result := fmt.Sprintf("Vous avez %d voyelles dans votre mot.", vowelTotal)
	fmt.Println(result)
}

func getUserValue() string{
	var valueUser string
    fmt.Print("Entrez une valeur : ")
    fmt.Scanln(&valueUser)
	return valueUser
}

func vowelCountInWord(valueVowelResult string) int {
	result := strings.Split(valueVowelResult, "")
	var countVowel = 0
	vowels := "aeiouAEIOU"
	for _, char := range result {
		countVowel += countWhenIsVowel(vowels, char)
	}
	return countVowel
}

func countWhenIsVowel(vowels string, char string) int {
	if strings.Contains(vowels, char) {
		return 1
	}
	return 0
}