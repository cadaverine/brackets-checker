package main

import (
	"fmt"

	"./stack"
)

func checkBrackets(text string) bool {

	return false
}

func main() {

	store := stack.Stack{}

	store.Push("1")
	store.Push("2")
	store.Push("3")

	fmt.Println(store.Data)
}
