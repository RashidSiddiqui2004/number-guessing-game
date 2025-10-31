package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🎯 Welcome to the Number Guessing Game!")
	fmt.Println("----------------------------------------")

	bestScore := -1 // stores minimum attempts (lower is better)

	for {
		// Select difficulty
		fmt.Println("\nSelect difficulty level:")
		fmt.Println("1. Easy   (10 chances)")
		fmt.Println("2. Medium (7 chances)")
		fmt.Println("3. Hard   (5 chances)")
		fmt.Print("Enter your choice (1-3): ")

		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)
		choice, err := strconv.Atoi(choiceInput)
		if err != nil || choice < 1 || choice > 3 {
			fmt.Println("❌ Invalid choice. Defaulting to Medium difficulty.")
			choice = 2
		}

		maxAttempts := 7
		switch choice {
		case 1:
			maxAttempts = 10
		case 2:
			maxAttempts = 7
		case 3:
			maxAttempts = 5
		}

		target := rand.Intn(100) + 1
		attempts := 0
		startTime := time.Now()

		fmt.Printf("\nI'm thinking of a number between 1 and 100.\nYou have %d chances to guess it!\n\n", maxAttempts)

		for attempts < maxAttempts {
			fmt.Print("Enter your guess (or type 'exit' to quit this round): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if input == "exit" {
				fmt.Println("👋 Round exited. The number was:", target)
				break
			}

			guess, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("❌ Please enter a valid number.")
				continue
			}

			attempts++

			if guess < target {
				fmt.Println("📉 Too low!")
			} else if guess > target {
				fmt.Println("📈 Too high!")
			} else {
				duration := time.Since(startTime).Seconds()
				fmt.Printf("\n🎉 Correct! The number was %d.\n", target)
				fmt.Printf("✅ You guessed it in %d attempts and %.2f seconds.\n", attempts, duration)

				if bestScore == -1 || attempts < bestScore {
					bestScore = attempts
					fmt.Println("🏆 New Best Score!")
				}
				break
			}

			if attempts == maxAttempts {
				fmt.Printf("\n❌ Out of chances! The correct number was %d.\n", target)
			}
		}

		fmt.Println("\nYour best score so far (fewest attempts):", bestScore)
		fmt.Print("\nDo you want to play again? (y/n): ")
		retryInput, _ := reader.ReadString('\n')
		retryInput = strings.TrimSpace(strings.ToLower(retryInput))

		if retryInput != "y" {
			fmt.Println("\n👋 Thanks for playing! Goodbye!")
			break
		}
	}
}
