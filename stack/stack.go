package stack

import "errors"

// IStack base interface
type IStack interface {
	Push(interface{})
	Pop() interface{}
	IsEmpty() bool
	Size() int
}

// Stack simple implementation
type Stack struct {
	Data []rune
}

// Push - method of Stack
func (stack *Stack) Push(value rune) {
	stack.Data = append(stack.Data, value)
}

// Pop - method of Stack
func (stack *Stack) Pop() (rune, error) {
	length := len(stack.Data)

	if length > 0 {
		last := stack.Data[length-1]
		stack.Data = stack.Data[:length-1]
		return last, nil
	}

	return 0, errors.New("stack is empty")
}

// IsEmpty - method of Stack
func (stack *Stack) IsEmpty() bool {
	return len(stack.Data) == 0
}

// Size - method of Stack
func (stack *Stack) Size() int {
	return len(stack.Data)
}

func push(stack IStack, value interface{}) {
	stack.Push(value)
}

func pop(stack IStack) interface{} {
	return stack.Pop()
}
