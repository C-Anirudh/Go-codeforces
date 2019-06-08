package main

import (
	"fmt"
	"math"
)

func main() {
	var mat [5][5]int
	var i, j, p, q int
	for i = 0; i < 5; i++ {
		for j = 0; j < 5; j++ {
			fmt.Scan(&mat[i][j])
			if mat[i][j] == 1 {
				p, q = i, j
			}
		}
	}
	r := math.Abs(float64(p-3)) + math.Abs(float64(q-3))
	fmt.Println(r)
}
