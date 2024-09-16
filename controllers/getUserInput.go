package controllers

import (
	"fmt"
	"math/rand"
)

// var gamePositions = []string{}

func GetUserInput() {

	var userPosition int
	fmt.Printf("Please enter the position where you want to put X: ")
	fmt.Scan(&userPosition)
	fmt.Println("")

	GameBoard(userPosition, 0)

}

func GenerateCompInput() {

	randomVal := rand.Intn(10)

	fmt.Printf("Computer has chosen the position: %d", randomVal)

	if randomVal != 0 {
		GameBoard(0, randomVal)
	} else {
		GenerateCompInput()
	}

}
