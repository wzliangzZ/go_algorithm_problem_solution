package main

import (
	"bufio"
	"fmt"
)

func CF1861D(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	suf := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		suf[i] = suf[i+1]
		if arr[i] >= arr[i+1] {
			suf[i]++
		}
	}
	var res int = suf[0]
	pre := 1
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i := 1; i < n; i++ {
		res = min(res, pre+suf[i])
		if arr[i] >= arr[i-1] {
			pre++
		}
	}
	fmt.Fprintln(out, res)
}
