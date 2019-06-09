package main

import (
	"fmt"
)

func main() {
	var n, i int
	fmt.Scan(&n)
	var arr, r [100]int
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
		r[arr[i]-1] = i + 1
	}
	for i = 0; i < n; i++ {
		fmt.Print(r[i])
		fmt.Print(" ")
	}
}
