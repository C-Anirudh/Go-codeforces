package main

import (
	"fmt"
)

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	r := x - y
	if r == 0 {
		if z == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("?")
		}
	} else if r > z || -r > z {
		if r > 0 {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	} else {
		fmt.Print("?")
	}
}
