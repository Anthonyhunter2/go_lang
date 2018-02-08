package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
Lab:
	for x := 1; x < 11; x++ {
		rand.Seed(int64(time.Now().Nanosecond()))
		die := rand.Intn(7)
		fmt.Println("For roll number ", x, "your die turned up as:", die)
		reader := bufio.NewReader(os.Stdin)
		if x == 10 {
			fmt.Println("SORRY THAT WAS YOUR LAST TURN")
			break Lab
		}
		fmt.Print("Would you like to continue rolling\n")
		answerb, _ := reader.ReadString('\n')
		switch strings.TrimRight(answerb, "\n") {
		case "yes":
			fmt.Println("Great lets keep rolling")
		default:
			fmt.Println("Well thanks for playing, goodbye")
			break Lab
		}
	}
}
