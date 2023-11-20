// using two-pointer approach

package main

import "fmt"

func findMiddleElements(arr []int) []int {
	middleIndex := len(arr) / 2

	if len(arr)%2 == 0 {
		// Even number of elements, return both middle elements
		return []int{arr[middleIndex-1], arr[middleIndex]}
	} else {
		// Odd number of elements, return the middle element
		return []int{arr[middleIndex]}
	}
}

func mergeSortedArrays(arr1 []int, arr2 []int) []int {
	merged := make([]int, 0, len(arr1)+len(arr2))
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	// Append remaining elements from arr1, if any
	for i < len(arr1) {
		merged = append(merged, arr1[i])
		i++
	}

	// Append remaining elements from arr2, if any
	for j < len(arr2) {
		merged = append(merged, arr2[j])
		j++
	}

	return merged

}

func main() {
	a := []int{1, 5, 9}
	b := []int{2, 6, 7}

	merged := mergeSortedArrays(a, b)
	middleElements := findMiddleElements(merged)
	fmt.Println("Middle Element(s):", middleElements)
}
