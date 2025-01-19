package main

import "fmt"

func lengthOfLongestSubsString(str string) {
	s := []rune(str)
	maxLength := 0
	seen := make(map[rune]bool)
	for i := 0; i < len(s); i++ {
		currentLength := 0
		for j := i; j < len(s); j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			currentLength++
			maxLength = max(maxLength, currentLength)

		}

	}
	println(maxLength)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubsString2(str string) {
	s := []rune(str)
	maxLength := 0
	seen := make(map[rune]bool)
	for i, j := 0, 0; j < len(s); {
		if !seen[s[j]] {
			seen[s[j]] = true
			j++
			maxLength = max(maxLength, j-i)
		} else {
			delete(seen, s[i])
			i++
		}
	}
	fmt.Println(maxLength, "d")
}

func lengthOfLongestSubsString3(str string) {
	s := []rune(str)
	maxLength := 0
	seen := make(map[rune]int)
	l := 0
	r := 0
	for l < r {
		character := s[r]
		fmt.Println(character, "sss")
		if seen[character] > 0 {
			l = max(l, seen[character])
			fmt.Println(seen)
		}
		seen[s[r]] = r + 1
		maxLength = max(maxLength, r-l+1)
		r++

	}
	fmt.Println(maxLength)
}
