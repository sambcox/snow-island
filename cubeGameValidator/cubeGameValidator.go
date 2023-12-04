package cubegamevalidator

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidateGame(game string) int {
	maxCubeMap := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	gameBagSplit := strings.Split(game, ": ")

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
				return 0
			}

			cubeColor := strings.Trim(cubeSlice[1], ",")

			if maxCubeMap[cubeColor] < cubeAmount {
				maxCubeMap[cubeColor] = cubeAmount
			}
		}
	}

	gamePower := maxCubeMap["blue"] * maxCubeMap["red"] * maxCubeMap["green"]
	return gamePower
}
