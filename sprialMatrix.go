package main

import "fmt"

func spiralMatrixs(arr [][]int) {
	if len(arr) == 0 {
		fmt.Print("Empty matrix")
	}

	rows := len(arr)
	cols := len(arr[0])
	top := 0
	bottom := rows - 1
	left := 0
	right := cols - 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			fmt.Print(arr[top][i], " ")
		}
		top++
		for i := top; i <= bottom; i++ {
			fmt.Print(arr[i][right], " ")
		}
		right--
		if top <= bottom {
			for i := right; i >= left; i-- {
				fmt.Print(arr[bottom][i], " ")
			}

		}
		bottom--

		if left <= right {
			for i := bottom; i >= top; i-- {
				fmt.Print(arr[i][left], " ")
			}
		}
		left++

	}

	fmt.Println()

}
