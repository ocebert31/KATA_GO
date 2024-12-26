package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evenOrOdd() {
	number := getValidatedUserInput()
	result := checkImpairOrPair(number)
	fmt.Println(result) 
}

func getValidatedUserInput() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Entrez un nombre : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Erreur : veuillez entrer un nombre entier valide.")
			os.Exit(1)
		}
		return number
	}
}

func checkImpairOrPair(numberUserResult int) string{
	if numberUserResult % 2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}