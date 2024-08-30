package main

import (
	"fmt"
)

func findSmallest(array []int) int {
	var smallest int = array[0]
	var smallest_index int = 0
	var length_of_array int = len(array) 
	for i := 1; i < length_of_array; i++ {
		if array[i] < smallest {
			smallest = array[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func selectionSort(array []int) []int {
	var new_array []int
	var length_of_array int = len(array)
	for i := 0; i < length_of_array; i++ {
		var smallest int = findSmallest(array)
		new_array = append(new_array, array[smallest])
		array = append(array[:smallest], array[smallest+1:]...)
	}
	return new_array
}

func main() {
	var array []int = []int{9, 2, 4, 6, 10, 3, 8, 5}
	fmt.Println("Original array:", array)
	var sorted_array []int = selectionSort(array)
	fmt.Println("Sorted array:", sorted_array)
}