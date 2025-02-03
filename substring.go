package main

func subString(s string, exact string) {
	s1 := []rune(s)
	s2 := []rune(exact)
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[0] {
			for j := 0; j < len(s2); j++ {
				if s1[i+j] != s2[j] {
					break
				}
				if j == len(s2)-1 {
					println("substring found")
					return
				}
			}
		}
	}
	println("substring not found")

}

func flatanAr(arr [][]int) []int {
	var result []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			result = append(result, arr[i][j])
		}
	}
	return result

}
