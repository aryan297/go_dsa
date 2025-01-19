package main

import (
	"fmt"
	"sort"
)

// boats function takes in an array of integers and an integer limit
// and returns the number of boats needed to carry all the people in the array

func boats(arr []int, limit int) {
	left, right := 0, len(arr)-1
	count := 0
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for left <= right {
		if arr[left]+arr[right] <= limit {
			left++
			count++
			right--
		} else {
			right--
			count++
		}
	}
	fmt.Println(count)

}
