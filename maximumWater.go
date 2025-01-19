package main

func maximumWater(arr []int) int {
	left, right := 0, len(arr)-1
	max := 0
	for left < right {
		//	max = Max(max, Min(arr[left], arr[right])*(right-left))
		l := Min(arr[left], arr[right])
		b := right - left
		max = Max(max, l*b)
		if arr[left] < arr[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
