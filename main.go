package main

import "fmt"

func main() {
	fmt.Println(validMountain([]int{0, 3, 2, 4}))
	fmt.Println(maximumWater([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	boats([]int{3, 2, 2, 1}, 3)
	fmt.Println(moveZero([]int{0, 1, 0, 3, 12}))
	lengthOfLongestSubsString("abcabcbb")
	lengthOfLongestSubsString2("abcabcbb")
	lengthOfLongestSubsString3("abvcabcbb")
	arr := []int{10, 11, 11, 11, 17, 18}
	arr1 := []int{1, 2, 2, 3, 4}
	arr4 := []int{1, 2, 3, 4, 5}
	findIndex(arr, 11)
	fmt.Println(findPrime(5))
	fmt.Println(findMissingNum(arr1))
	fmt.Println(singleCount(arr1))
	fmt.Println(twoSum(arr1, 4))
	fmt.Println(twoSum1(arr1, 4), "check")
	fmt.Println(duplicates(arr1))
	fmt.Println(duplicate2(arr1))
	fmt.Println(binarySearch(arr4, 3))

	st := "abbbac"
	fmt.Println("Palindrome longest", longestPalindrome(st))

	matrix := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	spiralMatrixs(matrix)

	arr6 := []int{1, 2, 5, 6, 2, 3}
	low := 0
	high := 5
	mergeSort(arr6, low, high)
	fmt.Println("merge sort", arr6)
	//make zero
	mat := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	matrixMakeZero(mat)

}
