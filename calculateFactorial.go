package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateFactorial() {
	number := getValidatedUserInput()
	factorial := factorial(number)
	fmt.Printf("Le factoriel de %d est : %d\n", number, factorial)
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

func factorial(number int) int {
	if number == 0 {
		return 1
	} else {
		return number * factorial(number-1)
	}
}