package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Quiz Game!")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello.. %v, welcome to the game! \n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yes...you can play!")
	} else {
		fmt.Println("Nopee.. you can't!")
		return
	}

	score := 0
	num_questions := 2

	fmt.Printf("What is my favorite? Chili's 100 or Onion's 55? ")
	
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	// Clean whitespace and convert everything to lowercase automatically
	userAnswer := strings.ToLower(strings.TrimSpace(answer + " " + answer2))

	// Switch statement replaces all the if/else if blocks cleanly
	switch userAnswer {
	case "chili's 100", "chilis 100":
		fmt.Println("Correct...!")
		score++
	default:
		// This handles "banana", wrong answers, or anything else not on the list
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How old am I? 35 or 38? ")

	var guess int
	fmt.Scan(&guess)

	// Switch statement checks the integer directly
	switch guess {
	case 35: // Change this to whichever number is actually your real age
		fmt.Println("Correct...!")
		score++
	default:
		// This catches 38, 99, or any other number/wrong input
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You got %v out of %v", score, num_questions)
}