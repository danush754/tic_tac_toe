package controllers

import (
	"fmt"
	"strconv"
)

var positionArr = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func GreetUser() {

	fmt.Println("-------Welcome to Tic-Tac-Toe---------")
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")

}

func IntialBoard() {

	fmt.Println("")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", positionArr[0], positionArr[1], positionArr[2])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", positionArr[3], positionArr[4], positionArr[5])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", positionArr[6], positionArr[7], positionArr[8])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")

	fmt.Println("")
}

func InstructUser() {

	fmt.Println("You have to choose a place in the box where you will put your x or o")
}

func GameBoard(humanValInt, compValInt int) {
	isAlreadyGvnHumanInput := ValidatehumanVal(positionArr, humanValInt)

	isAlreadyGvnCompInput := ValidatecompVal(positionArr, compValInt)
	fmt.Println("check", isAlreadyGvnHumanInput, isAlreadyGvnCompInput)

	if isAlreadyGvnHumanInput {
		GetUserInput()
	}

	if isAlreadyGvnCompInput {
		GenerateCompInput()
	}

	humanVal := strconv.Itoa(humanValInt)
	compVal := strconv.Itoa(compValInt)

	// 	positionArr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	switch humanVal {
	case "1":
		positionArr[0] = "X"
	case "2":
		positionArr[1] = "X"
	case "3":
		positionArr[2] = "X"
	case "4":
		positionArr[3] = "X"
	case "5":
		positionArr[4] = "X"
	case "6":
		positionArr[5] = "X"
	case "7":
		positionArr[6] = "X"
	case "8":
		positionArr[7] = "X"
	case "9":
		positionArr[8] = "X"
	}

	switch compVal {
	case "1":
		positionArr[0] = "O"
	case "2":
		positionArr[1] = "O"
	case "3":
		positionArr[2] = "O"
	case "4":
		positionArr[3] = "O"
	case "5":
		positionArr[4] = "O"
	case "6":
		positionArr[5] = "O"
	case "7":
		positionArr[6] = "O"
	case "8":
		positionArr[7] = "O"
	case "9":
		positionArr[8] = "O"
	}

	fmt.Println("")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", positionArr[0], positionArr[1], positionArr[2])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", positionArr[3], positionArr[4], positionArr[5])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", positionArr[6], positionArr[7], positionArr[8])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")

	fmt.Println("")

}
