package main

import (
	"bufio"
	"fmt"
)

func CF1521B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	pos := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] < arr[pos] {
			pos = i
		}
	}

	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	fmt.Fprintln(out, n-1)
	for i := 0; i < n; i++ {
		if i != pos {
			fmt.Fprintln(out, pos+1, i+1, arr[pos], arr[pos]+abs(pos-i))
		}
	}

}
