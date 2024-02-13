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

func Test_Contains_BinarySearchTree_Returns_True(t *testing.T) {
	tree := &datastructures.BinarySearchTree{}

	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(4)

	expected := true
	result := tree.Contains(6)

	if result != expected {
		t.Errorf(TestErrorResponse(expected, result))
	} else {
		t.Logf(TestSuccessResponse(expected, result))
	}
}

func Test_Contains_BinarySearchTree_Returns_False(t *testing.T) {
	tree := &datastructures.BinarySearchTree{}

	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(4)

	expected := false
	result := tree.Contains(9)

	if result != expected {
		t.Errorf(TestErrorResponse(expected, result))
	} else {
		t.Logf(TestSuccessResponse(expected, result))
	}
}
