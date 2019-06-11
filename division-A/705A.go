package main

import (
	"fmt"
)

func main() {
	var n, i int
	fmt.Scan(&n)
	hIt := "I hate it"
	lIt := "I love it"
	hThat := "I hate that"
	lThat := "I love that"
	if n == 1 {
		fmt.Println(hIt)
	} else if n == 2 {
		fmt.Println(hThat + " " + lIt)
	} else {
		if n%2 == 1 {
			for i = 1; i < n-1; i += 2 {
				fmt.Print(hThat)
				fmt.Print(" " + lThat + " ")
			}
			fmt.Print(hIt)
		} else {
			fmt.Print(hThat)
			for i = 1; i < n-2; i += 2 {
				fmt.Print(" " + lThat + " ")
				fmt.Print(hThat)
			}
			fmt.Print(" " + lIt)
		}
	}
}
