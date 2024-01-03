package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1849C(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	var s string
	fmt.Fscan(in, &s)
	l := make([]int, n+1)
	r := make([]int, n+2)
	for i := 0; i < n; i++ {
		if s[i] == '0' {
			l[i+1] = i + 1
		} else {
			l[i+1] = l[i]
		}
	}
	r[n+1] = n + 1
	for i := n; i > 0; i-- {
		if s[i-1] == '1' {
			r[i] = i
		} else {
			r[i] = r[i+1]
		}
	}
	st := make(map[[2]int]struct{})
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		a, b = r[a], l[b]
		if a > b {
			a, b = 0, 0
		}
		st[[2]int{a, b}] = struct{}{}
	}
	fmt.Fprintln(out, len(st))
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1849C(in, out)
	}
}
*/
