package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1905A(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Fprintln(out, max(n, m))

}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF1905A(in, out)
	}

}
*/
