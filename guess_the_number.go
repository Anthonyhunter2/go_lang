package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

//rand.Seed(time.Now().Nanosecond())
//var number rand.Intn(1,100)

func mainb() {
	number := rand.Intn(100)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("I'm thinking of a number between 1 and 100, can you guess my number?\n")
	guess, _ := reader.Reader
	if guess == number {
		fmt.Println("OMG YOU guessed my number")
	} else {
		too_high(guess)
		too_low(guess)
	}
}

func too_high(num int) int {
	if num > number {
		fmt.Println("Sorry thats to high to be my number")
	}
}
func too_low(num int) int {
	if num < number {
		fmt.Println("Sorry but thats too low to be my number")
	}
}
