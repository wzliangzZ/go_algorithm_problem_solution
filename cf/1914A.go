package main

import (
	"bufio"
	"fmt"
)

func CF1914A(in *bufio.Reader, out *bufio.Writer) {
	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	cnt := make([]int, 27)
	for i := 0; i < n; i++ {
		cnt[s[i]-'A'+1]++
	}
	var res int
	for i := 1; i < 27; i++ {
		if cnt[i] >= i {
			res++
		}
	}
	fmt.Fprintln(out, res)

}
