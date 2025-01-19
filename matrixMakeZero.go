package main

import "fmt"

func convertCol(arr [][]int, row int, col int, i int) {
	for j := 0; j < col; j++ {
		if arr[i][j] != 0 {
			arr[i][j] = -1
		}
	}
}

func convertRow(arr [][]int, row int, col int, j int) {
	for i := 0; i < row; i++ {
		if arr[i][j] != 0 {
			arr[i][j] = -1
		}
	}
}

func matrixMakeZero(arr [][]int) {
	row := len(arr)
	col := len(arr[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if arr[i][j] == 0 {
				convertCol(arr, row, col, i)
				convertRow(arr, row, col, j)
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if arr[i][j] == -1 {
				arr[i][j] = 0
			}
		}
	}
	fmt.Println(arr)
}
