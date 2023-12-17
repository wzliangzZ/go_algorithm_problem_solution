package main

import (
	"bufio"
	"fmt"
)

func CF1867B(in *bufio.Reader, out *bufio.Writer) {
	var n, k, gap int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	var i, j int = 0, n - 1
	for i < j {
		if s[i] != s[j] {
			k++
		}
		i++
		j--
	}
	var res []byte
	var rem int = n - 2*k
	for i := 0; i <= n; i++ {
		if i < k {
			res = append(res, '0')
		} else {
			gap = i - k
			if rem < gap {
				res = append(res, '0')
			} else if rem%2 != 0 {
				res = append(res, '1')
			} else if gap%2 != 0 {
				res = append(res, '0')
			} else {
				res = append(res, '1')
			}
		}
	}
	fmt.Fprintln(out, string(res))
}
