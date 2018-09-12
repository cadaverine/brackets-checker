package stack

import "testing"

func TestPop(t *testing.T) {
	stack := Stack{}

	stack.Push('a')
	stack.Push('b')

	result, error := stack.Pop()
	if result != 'a' || error != nil {
		t.Errorf("test failed")
	}

	result, error = stack.Pop()
	if result != 'b' || error != nil {
		t.Errorf("test failed")
	}

	result, error = stack.Pop()
	if result != 0 || error == nil {
		t.Errorf("test failed")
	}
}
