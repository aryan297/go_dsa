package main

import "fmt"

func twoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		nums := arr[i]
		if val, ok := m[target-nums]; ok {
			return []int{val, i}
		}
		m[nums] = i

	}
	return []int{}

}

func singleCount(arr []int) int {
	count := make(map[int]int)
	for _, value := range arr {
		count[value]++
	}
	fmt.Println(count)
	for key, v := range count {
		if v == 1 {
			return key
		}
	}
	return -1
}

func findPrime(n int) int {
	count := 0
	for i := 2; i < n; i++ {
		isPrime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			count++
		}
	}
	return count

}

func findMissingNum(arr []int) int {
	presentArray := make(map[int]bool)

	for _, j := range arr {
		presentArray[j] = true
	}

	for i := 1; i < len(arr)-1; i++ {
		if !presentArray[i] {
			return i
		}
	}

	return -1

}

func findIndex(arr []int, target int) {
	firstIndex := 0
	lastIndex := 0
	newArr := []int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == target && firstIndex == 0 {
			firstIndex = i
			lastIndex = i
			// newArr=append(newArr,firstIndex);
			//  newArr=append(newArr,lastIndex)
		} else if arr[i] == target {
			lastIndex = i
			// newArr=append(newArr,lastIndex);
		}
	}
	fmt.Println(newArr)
	fmt.Println(firstIndex, lastIndex)
}

func twoSum1(arr []int, target int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}

}

func duplicates(arr []int) bool {
	m := make(map[int]int)

	for _, num := range arr {
		m[num]++
	}

	for _, value := range m {
		fmt.Println(value, "ss")
		if value > 1 {
			return true
		}
	}

	return false
}

func duplicate2(arr []int) bool {
	m := make(map[int]bool)
	for _, num := range arr {
		if m[num] {
			return true
		}
		m[num] = true
	}
	return false
}
