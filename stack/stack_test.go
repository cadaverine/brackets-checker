package stack

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	stack := Stack{}

	a := 'a'
	b := 'b'

	stack.Push(a)
	stack.Push(b)

	result, error := stack.Pop()
	if result != 'a' || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = stack.Pop()
	if result != 'b' || error != nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}

	result, error = stack.Pop()
	if result != 0 || error == nil {
		t.Errorf("Test failed.\nResult: %v\nType: %v\nError: %v", result, reflect.TypeOf(result), error)
	}
}
