package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	if strings.Contains(str, "H") || strings.Contains(str, "Q") || strings.Contains(str, "9") {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
