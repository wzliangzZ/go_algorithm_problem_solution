package main

import (
	"bufio"
	"fmt"
)

func CF1902A(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	var a, b int
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			a++
		} else {
			b++
		}
	}
	if a > b || (a > 0 && b > 0) {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
	}
}
