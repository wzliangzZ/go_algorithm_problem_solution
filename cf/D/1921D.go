package main

import (
	"bufio"
	"fmt"

	//"os"
	"sort"
)

func CF1921D(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &n, &m)
	arra := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arra[i])
	}
	arrb := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &arrb[i])
	}
	sort.Ints(arra)
	sort.Slice(arrb, func(i, j int) bool {
		return arrb[i] > arrb[j]
	})
	abs := func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}
	max := func(a, b int) int {
		if a < b {
			return b
		} else {
			return a
		}
	}
	var res int
	for k := 0; k < n; k++ {
		res += max(abs(arra[k]-arrb[k]), abs(arra[k]-arrb[m-n+k]))
	}
	fmt.Fprintln(out, res)

}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1921D(in, out)
	}
}
*/
