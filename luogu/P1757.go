package main

import (
	"bufio"
	"fmt"
	//"os"
)

func P1757(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &m, &n)
	v := make([][]int, 101)
	w := make([][]int, 101)
	f := make([]int, m+1)
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		w[c] = append(w[c], a)
		v[c] = append(v[c], b)
	}
	// 现枚举顺序： 组别、体积、组中物品
	for k := 1; k <= 100; k++ {
		for j := m; j > 0; j-- {
			for i := 0; i < len(v[k]); i++ {
				if j >= w[k][i] {
					//每组物品选或不选，相当于01背包问题每个物品选货不选
					//f[k][j] = max(f[k][j], f[k - 1][j - w[k][i]] + v[k][i])	
					f[j] = max(f[j], f[j-w[k][i]]+v[k][i])
				}
			}
		}
	}
	fmt.Fprintln(out, f[m])
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		P1757(in, out)
	}
}
*/
