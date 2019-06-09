package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	if (int(str[0]) >= 97 && int(str[0]) <= 122) && strings.Compare(str[1:], strings.ToUpper(str[1:])) == 0 {
		fmt.Println(strings.Title(strings.ToLower(str)))
	} else if strings.Compare(str, strings.ToUpper(str)) == 0 {
		fmt.Println(strings.ToLower(str))
	} else {
		fmt.Println(str)
	}
}
