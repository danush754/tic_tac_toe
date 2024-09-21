package main

import "tic_tac_toe/controllers"

func main() {

	controllers.GreetUser()
	controllers.IntialBoard()
	controllers.InstructUser()

	for {
		controllers.GetUserInput()
		if controllers.ValidateUserWin(controllers.PositionArr) {
			break
		}
		controllers.GenerateCompInput()
		if controllers.ValidateCompWin(controllers.PositionArr) {
			break
		}

	}

}
