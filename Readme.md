# 🎯 CLI Number Guessing Game (Go)

Project Page URL - https://roadmap.sh/projects/number-guessing-game

A simple **Command Line Interface (CLI)** number guessing game built in **Go**.  
The computer randomly selects a number between **1 and 100**, and the player has to guess it within a limited number of chances based on the chosen difficulty level.

---

## 📜 Project Requirements

You are required to build a simple number guessing game where:
- The computer randomly selects a number between **1 and 100**.
- The user selects a **difficulty level** (easy, medium, hard), which determines the number of chances.
- The user must guess the number within the given attempts.
- The game ends when the user guesses correctly or runs out of chances.

---

## 🕹️ How the Game Works

1. When the game starts, it displays a **welcome message** and **rules**.
2. The user selects a difficulty level:
   - **Easy** → 10 chances  
   - **Medium** → 7 chances  
   - **Hard** → 5 chances
3. The computer generates a random number between **1 and 100**.
4. The user enters their guesses one by one.
5. After each guess:
   - If correct → 🎉 A congratulatory message appears with the number of attempts.
   - If incorrect → A hint is shown whether the number is higher or lower.
6. The game ends when the user either guesses the number or exhausts all chances.

---
