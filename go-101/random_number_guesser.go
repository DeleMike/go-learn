package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("I'm thinking of a number...")
	// generate a random number between 0 to 100
	randomNumber := rand.Intn(100) + 1
	//fmt.Println(randomNumber)
	for numOfTrials := 10; numOfTrials > 0; numOfTrials-- {
		// ask user to enter their guessed number
		fmt.Print("Enter the guessed number: ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if err != nil {
			log.Fatal(err)
		}

		guessedNumber, _ := strconv.Atoi(text)
		//fmt.Println(guessedNumber)

		if guessedNumber == randomNumber {
			fmt.Println("Good job! You guessed it!")
			break
		}

		if guessedNumber < randomNumber {
			fmt.Println("Oops! Your guess was LOW.")
		} else {
			fmt.Println("Oops! Your guess was HIGH.")
		}

		if numOfTrials == 1 {
			fmt.Println("Sorry, you didn't guess my number. It was: ", randomNumber)
		} else {
			fmt.Println("You have", numOfTrials-1, "trials left...")

		}

	}

}
