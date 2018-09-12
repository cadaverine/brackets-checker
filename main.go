package main

import (
	"fmt"

	"./stack"
)

func some(slice []rune, predicat func(rune) bool) bool {
	for _, value := range slice {
		if predicat(value) {
			return true
		}
	}
	return false
}

func checkBrackets(text string) (bool, int) {
	stack := stack.Stack{}
	leftBrackets := []rune{'{', '[', '('}
	rightBrackets := []rune{'}', ']', ')'}

	isEqualCarry := func(char rune) func(rune) bool {
		return func(value rune) bool {
			return char == value
		}
	}

	for i, char := range text {
		isEqual := isEqualCarry(char)

		isLeft := some(leftBrackets, isEqual)
		isRight := some(rightBrackets, isEqual)

		if isLeft {
			// stack.Push(char)
		} else if isRight {

		} else {
			return false, i
		}
	}
}

func main() {
	leftBrackets := []rune{'{', '[', '('}
	rightBrackets := []rune{'}', ']', ')'}

	flag := some(leftBrackets, func(value rune) bool { return value == '}' })
	flag2 := some(rightBrackets, func(value rune) bool { return value == '}' })

	fmt.Println(flag)
	fmt.Println(flag2)

	store := stack.Stack{}

	store.Push("1")
	store.Push("2")
	store.Push("3")

	fmt.Println(store.Data)
}
