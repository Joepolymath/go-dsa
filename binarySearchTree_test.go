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
