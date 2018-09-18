package stack

import "errors"

// Stack simple implementation
type Stack struct {
	Data []interface{}
}

// Push - method of Stack
func (stack *Stack) Push(value interface{}) {
	stack.Data = append(stack.Data, value)
}

// Pop - method of Stack
func (stack *Stack) Pop() (interface{}, error) {
	length := stack.Size()

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
