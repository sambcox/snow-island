package main

import (
	"bufio"
	"fmt"
	"github.com/sambcox/snow-island/cubeGameValidator"
	"os"
)

func main() {
	file, err := os.Open("snowInput.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalPossible := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameBool, gameNumber := cubegamevalidator.ValidateGame(line)

		if gameBool {
			totalPossible += gameNumber
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println("Sum of total possible games:", totalPossible)
}
