package main

import (
	"fmt"

	"./stack"
)

func some(slice []interface{}, predicat func() bool) bool {
	for _, value := range slice {
		if predicat(value) {
			return true
		}
	}
	return false
}

func checkBrackets(text string) bool {
	for i, char := range text {
		if some(['{', '(', '['], func (value rune) { return value == char }) {

		}
	}

	return false
}

func main() {

	store := stack.Stack{}

	store.Push("1")
	store.Push("2")
	store.Push("3")

	fmt.Println(store.Data)
}
