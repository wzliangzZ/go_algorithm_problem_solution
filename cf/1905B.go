package main

import (
	"bufio"
	"fmt"
)

func CF1905B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	cnt := make([]int, n+1)
	for i := 0; i < n-1; i++ {
		var x, y int
		fmt.Fscan(in, &x, &y)
		cnt[x]++
		cnt[y]++
	}
	var m int
	for i := 1; i <= n; i++ {
		if cnt[i] == 1 {
			m++
		}
	}
	fmt.Fprintln(out, (m+1)/2)

}
