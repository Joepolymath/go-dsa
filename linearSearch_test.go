package main

import (
	"testing"

	"github.com/Joepolymath/go-dsa/algos"
)

func Test_LinearSearch(t *testing.T) {
	testArray := []int{1, 3, 5, 6, 7, 9, 12}
	expected := true
	result := algos.LinearSearch(testArray, 5)

	if result != expected {
		t.Errorf("FAILED: Expected %v, got %v", expected, result)
	} else {
		t.Logf("SUCCESS. Expected %v, got %v", expected, result)
	}
}