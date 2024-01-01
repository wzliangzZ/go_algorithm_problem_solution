package main

import (
	"bufio"
	"fmt"
	//"os"
)

func P1855(in *bufio.Reader, out *bufio.Writer) {
	var n, m, t int
	fmt.Fscan(in, &n, &m, &t)
	f := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]int, t+1)
	}
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		for j := m; j >= a; j-- {
			for k := t; k >= b; k-- {
				f[j][k] = max(f[j][k], f[j-a][k-b]+1)
			}
		}
	}
	fmt.Fprintln(out, f[m][t])

}
/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		P1855(in, out)
	}
}
*/