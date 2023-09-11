package main

import "fmt"

func main() {
	testArray := []int{1, 3, 5, 6, 7, 9, 12}
	fmt.Println(LinearSearch(testArray, 5)) // should return true
	fmt.Println(BinarySearch(testArray, 2)) // should return false
}