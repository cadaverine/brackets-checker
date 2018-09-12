package stack

import "errors"

type IStack interface {
	Push(interface{})
	Pop() interface{}
}

type Stack struct {
	Data []string
}

func (stack *Stack) Push(value string) {
	stack.Data = append(stack.Data, value)
}

func (stack *Stack) Pop() (string, error) {
	length := len(stack.Data)

	if length > 0 {
		last := stack.Data[length-1]
		stack.Data = stack.Data[:length-1]
		return last, nil
	}

	return "", errors.New("stack is empty")
}

func push(stack IStack, value interface{}) {
	stack.Push(value)
}

func pop(stack IStack) interface{} {
	return stack.Pop()
}
