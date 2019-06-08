package main

import (
	"fmt"
)

func main() {
	var t, a, b, max, p int
	fmt.Scan(&t)
	for t > 0 {
		fmt.Scan(&a, &b)
		p += b - a
		if p > max {
			max = p
		}
		t--
	}
	fmt.Println(max)
}
