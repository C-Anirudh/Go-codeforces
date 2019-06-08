package main

import (
	"fmt"
)

func main() {
	var n, i, j, k, p, q, r int
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&i, &j, &k)
		p += i
		q += j
		r += k
		n--
	}

	if p == 0 && q == 0 && r == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
