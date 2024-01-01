package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	var v, w []int
	f := make([]int, m + 1)
	split := func(a, b, c int) {
		var k int = 1
		for k < c {
			v = append(v, a * k)
			w = append(w, b * k)
			c -= k
			k <<= 1
		}
		v = append(v, a * c)
		w = append(w, b * c)
	}

	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		split(a, b, c)
	}
	for i := 0; i < len(v); i++ {
		for j := m; j >= w[i]; j-- {
			f[j] = max(f[j], f[j - w[i]] + v[i])
		}
	}
	fmt.Fprintln(out, f[m])
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		solve(in, out)
	}
}
