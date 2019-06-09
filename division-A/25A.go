package main

import (
	"fmt"
)

func main() {
	var n, i, e, ei, oi int
	fmt.Scan(&n)
	var arr [100]int
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i]%2 == 0 {
			e++
			ei = i
		} else {
			oi = i
		}
	}
	if e == 1 {
		fmt.Println(ei + 1)
	} else {
		fmt.Println(oi + 1)
	}
}
