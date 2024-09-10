package main

import (
	"tic_tac_toe/controllers"
)

func main() {

	controllers.GreetUser()
	controllers.IntialBoard(0, 0)
	controllers.InstructUser()

	// var i int

	// fmt.Print("Enter a number: ")

	// fmt.Scan(&i)

	// fmt.Printf("You have entered: %d\n ", i)

}
