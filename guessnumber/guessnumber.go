package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const totalAttemptsNumber = 3
const lowerBound = 1
const upperBound = 20

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	fmt.Printf("The number on the segment [%d, %d] is guessed.\n", lowerBound, upperBound)
	fmt.Println("Try to guess it!")
	guessedNum := rand.Intn(upperBound) + lowerBound

	isGuessed := false
	for i := 0; i < totalAttemptsNumber; i++ {
		fmt.Println()
		showAttemptsNumber(totalAttemptsNumber - i)
		userNum := getUserNumber()
		if userNum > guessedNum {
			fmt.Println("Your number is bigger.")
		} else if userNum < guessedNum {
			fmt.Println("Your number is less")
		} else {
			isGuessed = true
			break
		}
	}

	fmt.Println()
	if isGuessed {
		fmt.Printf("Congratulations! It's really %d\n", guessedNum)
	} else {
		fmt.Printf("The number was %d\n", guessedNum)
		fmt.Println("Next time it will definitely succeed!")
	}
}

func showAttemptsNumber(attemptsNumber int) {
	fmt.Printf("You have %d ", attemptsNumber)
	if attemptsNumber == 1 {
		fmt.Println("attempt")
	} else {
		fmt.Println("attempts")
	}
}

func getUserNumber() int {
	reader := bufio.NewReader(os.Stdin)
	var userNum int
	for {
		fmt.Print("Please enter your number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		userNum, err = strconv.Atoi(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, "You need to enter an integer.")
			continue
		}
		break
	}
	return userNum
}
