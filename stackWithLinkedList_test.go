package main

import (
	"testing"

	datastructures "github.com/Joepolymath/go-dsa/dataStructures"
)

func Test_Push_StackWithLinkedList(t *testing.T) {
	stack := &datastructures.Stack2{
		MaxSize: 5,
	}
	expected := true
	result, err := stack.Push(55)
	if err != nil {
		t.Errorf("FAILED: Push method threw an error %v", err)
	}

	if result != expected {
		t.Errorf(TestErrorResponse(expected, result))
	} else {
		t.Logf(TestSuccessResponse(expected, result))
	}
}