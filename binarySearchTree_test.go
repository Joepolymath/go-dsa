package main

import (
	"testing"

	datastructures "github.com/Joepolymath/go-dsa/dataStructures"
)

func Test_Insert_BinarySearchTree(t *testing.T) {
	
	tree := &datastructures.BinarySearchTree{}

	expected := true
	result := tree.Insert(5)
	

	if result != expected {
		t.Errorf(TestErrorResponse(expected, result))
	} else {
		t.Logf(TestSuccessResponse(expected, result))
	}
}

// func Test_Pop_StackWithLinkedList(t *testing.T) {
// 	stack := &datastructures.Stack2{
// 		MaxSize: 5,
// 	}
// 	stack.Push(55)
// 	expected := 55
// 	result, err := stack.Pop()
// 	if err != nil {
// 		t.Errorf("FAILED: Push method threw an error %v", err)
// 	}

// 	if result != expected {
// 		t.Errorf(TestErrorResponse(expected, result))
// 	} else {
// 		t.Logf(TestSuccessResponse(expected, result))
// 	}
// }