package main

import (
	"fmt"

	"github.com/Joepolymath/go-dsa/algos"
)

func main() {
	testArray := []int{1, 3, 5, 6, 7, 9, 12}
	fmt.Println(algos.LinearSearch(testArray, 5)) // should return true
	fmt.Println(algos.BinarySearch(testArray, 2)) // should return false
}

// test error
func TestErrorResponse(expected, result interface{}) string {
	return fmt.Sprintf("FAILED: Expected %v, got %v", expected, result)
}

func TestSuccessResponse(expected, result interface{}) string {
	return fmt.Sprintf("SUCCESS. Expected %v, got %v", expected, result)
}