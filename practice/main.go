package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(11)
	var guess int
	attempts := 0
	fmt.Println("Guess a number between 0 and 10")
	for {
		fmt.Scan(&guess)
		attempts++
		if guess < x {
			fmt.Println("Too low!")
		} else if guess > x {
			fmt.Println("Too high!")
		} else {
			fmt.Println("Correct!", "Attempts:", attempts)
			break

		}
	}
}
