package stack

import "testing"

func TestPop(t *testing.T) {
	stack := Stack{}

	stack.Push("p")
	stack.Push("hdasda")

	result, error := stack.Pop()
	if result != "hdasda" || error != nil {
		t.Errorf("test failed")
	}

	result, error = stack.Pop()
	if result != "p" || error != nil {
		t.Errorf("test failed")
	}

	result, error = stack.Pop()
	if result != "" || error == nil {
		t.Errorf("test failed")
	}
}
