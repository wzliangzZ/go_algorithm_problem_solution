package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF923B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n+1)
	t := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 1; i <= n; i++ {
		fmt.Fscan(in, &t[i])
	}
	f := make([]int, n+2)
	extr := make([]int, n+2)
	s := make([]int64, n+2)
	for i := 1; i <= n; i++ {
		s[i] += s[i-1] + int64(t[i])
	}
	find := func(target, start int) int {
		l, r := start, n
		for l < r {
			mid := (l + r + 1) / 2
			if s[mid]-s[start-1] <= int64(target) {
				l = mid
			} else {
				r = mid - 1
			}
		}
		return l
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	for i := 1; i <= n; i++ {
		idx := find(a[i], i)
		extr[idx+1] += max(0, int(int64(a[i])-(s[idx]-s[i-1])))
		if i == idx {
			extr[i] += -max(-a[i], -t[i])
		} else {
			f[i]++
			f[idx+1]--
		}
	}
	var pre int
	for i := 1; i <= n; i++ {
		pre += f[i]
		fmt.Fprint(out, extr[i]+pre*t[i], " ")
	}
	fmt.Fprintln(out)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF923B(in, out)
	}
}
*/
