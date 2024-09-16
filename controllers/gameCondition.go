package controllers

import (
	"fmt"
	"strings"
)

func ValidatehumanVal(gamePositions []string, humanVal int) bool {

	fmt.Println("arr", gamePositions[humanVal-1])

	if strings.Contains(gamePositions[humanVal], "X") || strings.Contains(gamePositions[humanVal], "O") {
		return true
	}

	return false
}

func ValidatecompVal(gamePositions []string, compval int) bool {

	if strings.Contains(gamePositions[compval], "X") || strings.Contains(gamePositions[compval], "O") {
		return true
	}

	return false

}
