package main

import (
	"fmt"
)

func isPalindrome(a1, a2 string) bool {
	var i, n int
	n = len(a1)
	for i = 0; i < n; i++ {
		if a1[i] != a2[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var s1, s2 string
	fmt.Scan(&s1, &s2)
	if len(s1) == len(s2) && isPalindrome(s1, s2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
