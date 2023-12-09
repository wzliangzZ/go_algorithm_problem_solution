package main

import (
	"bufio"
	"fmt"
)

func CF1811C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &arr[i])
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	fmt.Fprint(out, arr[0], " ")
	for i := 1; i < n-1; i++ {
		fmt.Fprint(out, min(arr[i], arr[i-1]), " ")
	}
	if n >= 2 {
		fmt.Fprint(out, arr[n-2])
	}
	fmt.Fprintln(out)

}
