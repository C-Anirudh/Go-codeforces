package main

import (
	"fmt"
)

func main() {
	var t int // size of input
	var n, k int64
	fmt.Scan(&t)
	for t > 0 {
		var ctr int64
		fmt.Scan(&n, &k)
		for n >= k {
			r := n % k
			if r == 0 {
				n = n / k
				ctr++
			} else {
				n = n - r
				ctr += r
			}
		}
		ctr += n
		fmt.Println(ctr)
		t = t - 1
	}
}
