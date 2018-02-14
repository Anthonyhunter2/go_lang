package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	number := rand.Intn(100)
	var input int
	fmt.Println("I'm thinking of a number between 1 and 100, can you guess my number?")
	for input != number {
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			fmt.Println(err)
		} else if input == number {
			fmt.Println("OMG YOU guessed my number")
		} else if tooHigh(input, number) {
			fmt.Println("Thats too high to be my number, try again")
		} else if tooLow(input, number) {
			fmt.Println("Thats too low to my number, try again")
		}
	}
}

func tooHigh(num int, rando int) bool {
	return num > rando
}
func tooLow(num int, rando int) bool {
	return num < rando
}
