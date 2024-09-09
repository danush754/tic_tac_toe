package controllers

import "fmt"

func GreetUser() {

	fmt.Println("-------Welcome to Tic-Tac-Toe---------")
	fmt.Println("-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*")

}

func IntialBoard() {

	fmt.Println("")
	fmt.Println("	|	|	")
	fmt.Println("   1	|   2	|   3")
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Println("   4	|   5	|   6")
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")
	fmt.Println("   7	|   8	|   9")
	fmt.Println("-----------------------")
	fmt.Println("	|	|	")

	fmt.Println("")
}

func InstructUser() {

	fmt.Println("You have to choose a place in the box where you will put your x or o")
}
