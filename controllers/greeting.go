package controllers

import (
	"fmt"
	"strconv"
)

// var positionArr = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func GreetUser() {

	fmt.Println("-------Welcome to Tic-Tac-Toe---------")
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")

}

func IntialBoard() {

	fmt.Println("")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", PositionArr[0], PositionArr[1], PositionArr[2])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", PositionArr[3], PositionArr[4], PositionArr[5])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", PositionArr[6], PositionArr[7], PositionArr[8])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")

	fmt.Println("")
}

func InstructUser() {

	fmt.Println("You have to choose a place in the box where you will put your x or o")
}

func GameBoard(humanValInt, compValInt int) {

	humanVal := strconv.Itoa(humanValInt)
	compVal := strconv.Itoa(compValInt)

	// 	positionArr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	switch humanVal {
	case "1":
		PositionArr[0] = "X"
	case "2":
		PositionArr[1] = "X"
	case "3":
		PositionArr[2] = "X"
	case "4":
		PositionArr[3] = "X"
	case "5":
		PositionArr[4] = "X"
	case "6":
		PositionArr[5] = "X"
	case "7":
		PositionArr[6] = "X"
	case "8":
		PositionArr[7] = "X"
	case "9":
		PositionArr[8] = "X"
	}

	switch compVal {
	case "1":
		PositionArr[0] = "O"
	case "2":
		PositionArr[1] = "O"
	case "3":
		PositionArr[2] = "O"
	case "4":
		PositionArr[3] = "O"
	case "5":
		PositionArr[4] = "O"
	case "6":
		PositionArr[5] = "O"
	case "7":
		PositionArr[6] = "O"
	case "8":
		PositionArr[7] = "O"
	case "9":
		PositionArr[8] = "O"
	}

	fmt.Println("")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", PositionArr[0], PositionArr[1], PositionArr[2])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", PositionArr[3], PositionArr[4], PositionArr[5])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", PositionArr[6], PositionArr[7], PositionArr[8])
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")

	fmt.Println("")

	// GenerateCompInput()

}
