package main

import (
	"fmt"
	"strings"
)

func isLucky(n int64) bool {
	var d int64
	for n > 0 {
		d = n % 10
		if d != 4 && d != 7 {
			return false
		}
		n = n / 10
	}
	return true
}

func main() {
	var s string
	var num, four, seven int64
	fmt.Scan(&s)
	four = int64(strings.Count(s, "4"))
	seven = int64(strings.Count(s, "7"))
	num = four + seven
	if num >= 4 && isLucky(num) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
