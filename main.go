package main

import (
	"bufio"
	"fmt"
	"os"

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

func isPaired(first rune, second rune) bool {
	switch true {
	case first == '{' && second == '}':
	case first == '}' && second == '{':
	case first == '[' && second == ']':
	case first == ']' && second == '[':
	case first == '(' && second == ')':
	case first == ')' && second == '(':
		return true
	}
	return false
}

func findWrongBrackets(text string) int {
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
			stack.Push(char)
		} else if isRight {
			item, err := stack.Pop()
			if err != nil || !isPaired(item, char) {
				return i
			}
		} else {
			return i
		}
	}

	if stack.IsEmpty() {
		return -1
	}

	return stack.Size() - 1
}

func checkBrackets(out *os.File, text string) {
	result := findWrongBrackets(text)

	if result != -1 {
		fmt.Println(result)
	} else {
		fmt.Println("Success")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error)
	}

	checkBrackets(os.Stdout, text)
}
