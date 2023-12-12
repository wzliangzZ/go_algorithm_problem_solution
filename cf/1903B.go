package main

import (
	"bufio"
	"fmt"
)

func CF1903B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1<<30 - 1
	}
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &arr[i][j])
			if i != j {
				res[i] &= arr[i][j]
				res[j] &= arr[i][j]
			}

		}
	}
	ok := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && res[i]|res[j] != arr[i][j] {
				ok = false
			}
		}
	}
	if ok {
		fmt.Fprintln(out, "Yes")
		for i := 0; i < n; i++ {
			fmt.Fprint(out, res[i], " ")
		}
		fmt.Fprintln(out)
	} else {
		fmt.Fprintln(out, "No")
	}

}
