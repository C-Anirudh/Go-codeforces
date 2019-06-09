package main

import (
	"fmt"
	"sort"
)

func sum(a []int) int {
	var i, sum int
	for i = 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	var n, i int
	fmt.Scan(&n)
	s := make([]int, n)
	for i = 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	suml := s[0]
	sumr := sum(s) - s[0]
	if suml > sumr {
		fmt.Println(1)
	} else {
		for i = 1; i < n; i++ {
			suml += s[i]
			sumr -= s[i]
			if suml > sumr {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
	