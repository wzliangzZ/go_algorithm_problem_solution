package main

import (
	"bufio"
	"fmt"
	//"os"
)

func P1616(in *bufio.Reader, out *bufio.Writer) {
	var t, m int
	fmt.Fscan(in, &t, &m)
	f := make([]int, t+1)
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		for j := a; j <= t; j++ {
			// f[i][j] = max(f[i-1][j], f[i][j-a]+b)
			f[j] = max(f[j], f[j-a]+b)
		}
	}
	fmt.Fprintln(out, f[t])
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	// fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		P1616(in, out)
	}
}
*/
