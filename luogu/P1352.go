package main

import (
	"bufio"
	"fmt"
	//"os"
)

func P1352(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	f := make([][]int, n+1)
	g := make([][]int, n+1)
	is_son := make([]bool, n+1)

	for i := 0; i < n; i++ {
		f[i+1] = []int{0, 0}
		fmt.Fscan(in, &f[i+1][1])
	}
	for i := 0; i < n-1; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		g[b] = append(g[b], a)
		is_son[a] = true
	}
	var dfs func(int)
	var res int
	dfs = func(r int) {
		for _, v := range g[r] {
			dfs(v)
			f[r][0] += max(f[v][0], f[v][1])
			f[r][1] += f[v][0]
		}
	}
	for i := 1; i <= n; i++ {
		if !is_son[i] {
			dfs(i)
			res = max(f[i][0], f[i][1])
		}
	}
	fmt.Fprintln(out, res)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		P1352(in, out)
	}
}
*/
