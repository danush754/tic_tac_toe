package controllers

import (
	"fmt"
	"strconv"
)

func GreetUser() {

	fmt.Println("-------Welcome to Tic-Tac-Toe---------")
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")

}

func IntialBoard(humanValInt, compValInt int) {

	humanVal := strconv.Itoa(humanValInt)
	compVal := strconv.Itoa(compValInt)

	var firstPosition, secondPosition, thirdPosition = "1", "2", "3"

	var fourthPosition, fifthPosition, sixthPostion = "4", "5", "6"

	var seventhPosition, eigthPosition, ninthPosition = "7", "8", "9"

	switch humanVal {
	case "1":
		firstPosition = "X"
	case "2":
		secondPosition = "X"
	case "3":
		thirdPosition = "X"
	case "4":
		fourthPosition = "X"
	case "5":
		fifthPosition = "X"
	case "6":
		sixthPostion = "X"
	case "7":
		seventhPosition = "X"
	case "8":
		eigthPosition = "X"
	case "9":
		ninthPosition = "X"
	}

	switch compVal {
	case "1":
		firstPosition = "O"
	case "2":
		secondPosition = "O"
	case "3":
		thirdPosition = "O"
	case "4":
		fourthPosition = "O"
	case "5":
		fifthPosition = "O"
	case "6":
		sixthPostion = "O"
	case "7":
		seventhPosition = "O"
	case "8":
		eigthPosition = "O"
	case "9":
		ninthPosition = "O"
	}

	fmt.Println("")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", firstPosition, secondPosition, thirdPosition)
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", fourthPosition, fifthPosition, sixthPostion)
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Printf("   %v	|   %v	|   %v\n", seventhPosition, eigthPosition, ninthPosition)
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")

	fmt.Println("")
}

func InstructUser() {

	fmt.Println("You have to choose a place in the box where you will put your x or o")
}
