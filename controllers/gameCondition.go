package controllers

import (
	"fmt"
	"strings"
)

func ValidatehumanVal(gamePositions []string, humanVal int) bool {

	fmt.Println("arr", gamePositions[humanVal-1])

	if strings.Contains(gamePositions[humanVal-1], "X") || strings.Contains(gamePositions[humanVal-1], "O") {
		return true
	}

	return false
}

func ValidatecompVal(gamePositions []string, compval int) bool {

	if strings.Contains(gamePositions[compval-1], "X") || strings.Contains(gamePositions[compval-1], "O") {
		return true
	}

	return false

}
