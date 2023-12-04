package cubegamevalidator

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateGame(game string) (bool, int) {
	limitMap := map[string]int{
		"blue":  14,
		"red":   12,
		"green": 13,
	}

	gameBagSplit := strings.Split(game, ": ")

	gameNumber, error := strconv.Atoi(strings.Trim(gameBagSplit[0], "Game "))

	if error != nil {
		fmt.Println("Error converting string to int:", error)
		return false, 0
	}

	bagSlice := strings.Split(gameBagSplit[1], "; ")

	for _, bag := range bagSlice {
		gameSlice := strings.Split(bag, ", ")

		for _, cube := range gameSlice {
			if cube == "" {
				continue
			}

			cubeSlice := strings.Split(cube, " ")

			cubeAmount, error := strconv.Atoi(cubeSlice[0])
			if error != nil {
				fmt.Println("Error converting string to int:", error)
				return false, 0
			}

			cubeColor := strings.Trim(cubeSlice[1], ",")

			if cubeAmount > limitMap[cubeColor] {
				return false, gameNumber
			}
		}
	}

	return true, gameNumber
}
