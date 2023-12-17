package main

import (
	"bufio"
	"fmt"
)

func CF1905A(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Fprintln(out, max(n, m))

}
