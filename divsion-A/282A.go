package main

import (
	"fmt"
)

func main() {
	var x int
	var p []byte
	ctr := 0
	fmt.Scan(&x)
	for x > 0 {
		fmt.Scan(&p)
		if string(p[0]) == "+" || string(p[2]) == "+" {
			ctr++
		} else {
			ctr--
		}
		x--
	}
	fmt.Print(ctr)
}
