package main

import (
	"fmt"
	"strings"
)

func accum(s string) string {
	res := ""
	for i, r := range s {
		l := string(r)
		res = res + strings.ToUpper(l)
		for j := 0; j < i; j++ {
			res = res + strings.ToLower(l)
		}
		if i+1 != len(s) {
			res = res + "-"
		}
	}
	return res
}

func main() {
	fmt.Println("A-Bb-Ccc-Dddd", accum("abcd"))
	fmt.Println("R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy", accum("RqaEzty"))
	fmt.Println("C-Ww-Aaa-Tttt", accum("cwAt"))
}
