package main

import (
	"testing"

	"github.com/Joepolymath/go-dsa/algos"
)

func Test_BinarySearch_NotFound(t *testing.T) {
	testArray := []int{1, 3, 5, 6, 7, 9, 12}
	expected := false
	result := algos.BinarySearch(testArray, 2)

	if result != expected {
		t.Errorf("FAILED: Expected %v, got %v", expected, result)
	} else {
		t.Logf("SUCCESS. Expected %v, got %v", expected, result)
	}
}
func Test_BinarySearch_Found(t *testing.T) {
	testArray := []int{1, 3, 5, 6, 7, 9, 12}
	expected := true
	result := algos.BinarySearch(testArray, 5)

	if result != expected {
		t.Errorf("FAILED: Expected %v, got %v", expected, result)
	} else {
		t.Logf("SUCCESS. Expected %v, got %v", expected, result)
	}
}