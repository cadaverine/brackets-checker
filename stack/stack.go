package stack

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

func (stack *Stack) Pop() string {
	last := stack.Data[len(stack.Data)-1]

	stack.Data = stack.Data[:len(stack.Data)-2]

	return last
}

func push(stack IStack, value interface{}) {
	stack.Push(value)
}

func pop(stack IStack) interface{} {
	return stack.Pop()
}
