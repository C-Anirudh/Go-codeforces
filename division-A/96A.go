package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	if strings.Contains(str, "1111111") || strings.Contains(str, "0000000") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}