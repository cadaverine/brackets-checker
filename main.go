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

func isPairedBrackets(first rune, second rune) bool {
	switch true {
	case
		first == '{' && second == '}',
		first == '}' && second == '{',
		first == '[' && second == ']',
		first == ']' && second == '[',
		first == '(' && second == ')',
		first == ')' && second == '(':
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
			if err != nil || !isPairedBrackets(item, char) {
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
		fmt.Fprintln(out, result)
	} else {
		fmt.Fprintln(out, "Success")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err.Error)
	}

	text = text[:len(text)-1]

	checkBrackets(os.Stdout, text)
}
