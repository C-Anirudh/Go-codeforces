package main

import (
	"fmt"
)

func main() {
	var n, p, q, ctr int
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&p, &q)
		if q-p >= 2 {
			ctr++
		}
		n--
	}
	fmt.Println(ctr)
}
