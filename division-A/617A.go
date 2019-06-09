package main

import (
	"fmt"
)

func main() {
	var n, i, ctr int
	fmt.Scan(&n)
	for i = 5; i > 0; i-- {
		ctr += n / i
		n = n % i
	}
	fmt.Println(ctr)
}
