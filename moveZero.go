package main

import "fmt"

// moveZero function takes in an array of integers and moves all the zeros to the end of the array
// while maintaining the order of the non-zero elements
//Ask in facebook

func moveZero(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		if arr[left] == 0 && arr[right] != 0 {
			temp := arr[left]
			arr[left] = arr[right]
			arr[right] = temp
			// arr[left], arr[right] = arr[right], arr[left] // this is a more concise way of swapping values
			left++
			right--
		} else if arr[left] != 0 {
			left++
		} else if arr[right] == 0 {
			right--
		}
	}
	return arr

}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			fmt.Println(mid, "ss")
			return mid
		} else if arr[mid] < target {
			left++
		} else {
			right--
		}

	}
	return -1

}
