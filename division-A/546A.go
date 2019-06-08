package main

import (
	"fmt"
)

func main() {
	var k, n, w int
	fmt.Scan(&k, &n, &w)
	cost := ((w * (w + 1)) / 2) * k
	if cost < n {
		fmt.Println(0)
	} else {
		fmt.Println(cost - n)
	}
}
