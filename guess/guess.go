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

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("你能猜猜看嗎?")
	// fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was Low.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was High.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
		// fmt.Println(guess)
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
