package main

func mergeTwo(arr1, arr2 []int) []int {
	i, j := 0, 0
	res := []int{}
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	for i < len(arr1) {
		res = append(res, arr1[i])
		i++
	}
	for j < len(arr2) {
		res = append(res, arr2[j])
		j++
	}
	return res
}
