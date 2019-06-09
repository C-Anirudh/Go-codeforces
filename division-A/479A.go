package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a == 1 && b == 1 && c == 1 {
		fmt.Println(a + b + c)
	} else if a == 1 && b == 1 {
		fmt.Println((a + b) * c)
	} else if b == 1 && c == 1 {
		fmt.Println((c + b) * a)
	} else if a == 1 && c == 1 {
		fmt.Println(a + b + c)
	} else if a == 1 {
		fmt.Println((a + b) * c)
	} else if b == 1 {
		if a > c {
			fmt.Println((c + b) * a)
		} else {
			fmt.Println(c * (b + a))
		}
	} else if c == 1 {
		fmt.Println((c + b) * a)
	} else {
		fmt.Println(a * b * c)
	}
}
