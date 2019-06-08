package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	slice := strings.Split(s, "+")
	sort.Strings(slice)
	str := strings.Join(slice, "+")
	fmt.Println(str)
}
