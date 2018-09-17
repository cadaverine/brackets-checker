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
	indexes := stack.Stack{}
	brackets := stack.Stack{}

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

		switch true {
		case isLeft:
			indexes.Push(rune(i))
			brackets.Push(char)
		case isRight:
			item, err := brackets.Pop()
			if err != nil || !isPairedBrackets(item, char) {
				return i
			}
			indexes.Pop()
		}
	}

	if brackets.IsEmpty() {
		return -1
	}

	index, err := indexes.Pop()
	if err != nil {
		panic(err.Error)
	}

	return int(index)
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
