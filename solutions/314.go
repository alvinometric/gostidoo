package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main(){
	numToGuess := rand.Intn(100)
	guesses := 0
	var input string
	inputAsNumber := -1

	fmt.Println("👀 Guess a number between 0 and 100");

	for inputAsNumber != numToGuess{
		fmt.Print("Your number: ")
  	fmt.Scanln(&input)
		trimmed := strings.TrimSpace(input)
		inputAsNumber, err := strconv.Atoi(trimmed)
		guesses++

		if err != nil{
			fmt.Println("Are you sure you entered a number and nothing else?");
		} else if inputAsNumber > numToGuess{
			fmt.Println("Lower!")
		} else if inputAsNumber < numToGuess {
			fmt.Println("Higher!")
		}else{
			fmt.Println("🥳 You won in", guesses, "guesses")
			os.Exit(0)
		}
	}
}