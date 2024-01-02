package main

import (
	"bufio"
	"fmt"
	//"os"
)

func P2014(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	m++
	f := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
	}
	g := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(in, &a, &f[i][1])
		g[a] = append(g[a], i)
	}
	var dfs func(int)
	dfs = func(u int) {
		for _, ne := range g[u] {
			dfs(ne)
			for j := m; j > 0; j-- {
				for i := 0; i < j; i++ {
					f[u][j] = max(f[u][j], f[u][j-i]+f[ne][i])
				}
			}
		}
	}
	dfs(0)
	fmt.Fprintln(out, f[0][m])
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		P2014(in, out)
	}
}
*/
