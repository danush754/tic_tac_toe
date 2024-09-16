package controllers

import (
	"fmt"
	"math/rand"
)

var PositionArr = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func GetUserInput() {

	var userPosition int
	fmt.Printf("Please enter the position where you want to put X: ")
	fmt.Scan(&userPosition)
	fmt.Println("")

	isAlreadyGvnHumanInput := ValidatehumanVal(PositionArr, userPosition)

	if isAlreadyGvnHumanInput {
		GetUserInput()
	} else {
		GameBoard(userPosition, 0)
	}

}

func GenerateCompInput() {

	randomVal := rand.Intn(10)

	isAlreadyGvnCompInput := ValidatecompVal(PositionArr, randomVal)

	if randomVal != 0 && !isAlreadyGvnCompInput {
		fmt.Printf("Computer has chosen the position: %d", randomVal)
		GameBoard(0, randomVal)
	} else {
		GenerateCompInput()
	}

}
