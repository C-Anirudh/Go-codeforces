package main

import (
	"fmt"
)

func main() {
	var n, i, ctr, c1, c2, c3, c4 int
	fmt.Scan(&n)
	s := make([]int, n)
	for i = 0; i < n; i++ {
		fmt.Scan(&s[i])
		if s[i] == 1 {
			c1++
		} else if s[i] == 2 {
			c2++
		} else if s[i] == 3 {
			c3++
		} else {
			c4++
		}
	}

	ctr += c4 + c2/2
	c2 = c2 % 2
	if c3 > 0 && c1 > 0 {
		if c3 > c1 {
			ctr += c1
			c3 = c3 - c1
			c1 = 0
		} else {
			ctr += c3
			c1 = c1 - c3
			c3 = 0
		}
	}
	ctr += c3

	if c2 == 1 && c1 >= 1 {
		ctr++
		if c1 >= 2 {
			c2 = 0
			c1 = c1 - 2
		} else {
			c2 = 0
			c1 = c1 - 1
		}
	} else if c2 == 1 {
		c2 = 0
		ctr++
	}

	if c2 > 0 {
		ctr++
	}

	if c1 >= 4 {
		ctr += c1 / 4
		c1 = c1 % 4
	}
	if c1 > 0 {
		ctr++
	}

	fmt.Println(ctr)
}
