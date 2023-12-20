package main

import (
	"bufio"
	"fmt"
)

func CF1914B(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	var p int = n - k
	for i := p; i <= n; i++ {
		fmt.Fprint(out, i, " ")
	}
	for i := p - 1; i > 0; i-- {
		fmt.Fprint(out, i, " ")
	}
	fmt.Fprintln(out)

}
