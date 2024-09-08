package main

import (
	"fmt"
	"tic_tac_toe/controllers"
)

func main() {

	controllers.GreetUser()

	var i int

	fmt.Print("Enter a number: ")

	fmt.Scan(&i)

	fmt.Printf("You have entered: %d ", i)

}
