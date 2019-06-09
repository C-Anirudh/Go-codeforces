package main

import (
	"fmt"
	"strings"
)

func main() {
	var u, l, i int
	var str string
	fmt.Scan(&str)
	for i = 0; i < len(str); i++ {
		if int(str[i]) >= 97 && int(str[i]) <= 122 {
			l++
		} else {
			u++
		}
	}

	if u > l {
		fmt.Println(strings.ToUpper(str))
	} else {
		fmt.Println(strings.ToLower(str))
	}
}
