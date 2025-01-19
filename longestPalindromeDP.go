package main

func longestPalindrome(st string) string {
	ans := ""
	n := len(st)
	isPal := make([][]bool, n)

	for i := range isPal {
		isPal[i] = make([]bool, n)
	}

	for i := 0; i < n; i++ {
		isPal[i][i] = true
	}
	for e := 0; e < n; e++ {
		for s := 0; s < e; s++ {
			medium := e - s - 1 // medium is the length of the string
			sameChar := st[s] == st[e]

			if sameChar && (medium <= 1 || isPal[s+1][e-1]) { // if the characters are the same and the length of the string is less than or equal to 1 or the string is a palindrome
				isPal[s][e] = true
				if e-s > len(ans) {
					ans = st[s : e+1]
				}

			}
		}
	}
	return ans
}
