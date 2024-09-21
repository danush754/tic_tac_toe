package controllers

import "fmt"

func ValidateUserWin(PositionArr []string) bool {
	if (PositionArr[0] == "X" && PositionArr[1] == "X" && PositionArr[2] == "X") || (PositionArr[3] == "X" && PositionArr[4] == "X" && PositionArr[5] == "X") || (PositionArr[6] == "X" && PositionArr[7] == "X" && PositionArr[8] == "X") || (PositionArr[0] == "X" && PositionArr[3] == "X" && PositionArr[6] == "X") || (PositionArr[1] == "X" && PositionArr[4] == "X" && PositionArr[7] == "X") || (PositionArr[2] == "X" && PositionArr[5] == "X" && PositionArr[8] == "X") || (PositionArr[2] == "X" && PositionArr[4] == "X" && PositionArr[6] == "X") || (PositionArr[0] == "X" && PositionArr[4] == "X" && PositionArr[8] == "X") {
		fmt.Println("********************You have won the match********************")
		return true
	}

	return false
}

func ValidateCompWin(PositionArr []string) bool {
	if (PositionArr[0] == "O" && PositionArr[1] == "O" && PositionArr[2] == "O") || (PositionArr[3] == "O" && PositionArr[4] == "O" && PositionArr[5] == "O") || (PositionArr[6] == "O" && PositionArr[7] == "O" && PositionArr[8] == "O") || (PositionArr[0] == "O" && PositionArr[3] == "O" && PositionArr[6] == "O") || (PositionArr[1] == "O" && PositionArr[4] == "O" && PositionArr[8] == "O") || (PositionArr[2] == "O" && PositionArr[5] == "O" && PositionArr[8] == "O") || (PositionArr[2] == "O" && PositionArr[4] == "O" && PositionArr[6] == "O") || (PositionArr[0] == "O" && PositionArr[4] == "O" && PositionArr[8] == "O") {
		fmt.Println("********************Comp have won the match********************")

		return true
	}

	return false
}
