package main

import (
	"bufio"
	"fmt"

	//"os"
	"sort"
)

func CF1941A(in *bufio.Reader, out *bufio.Writer) {
	var T int
	for fmt.Fscan(in, &T); T > 0; T-- {
		var n, m, k int
		fmt.Fscan(in, &n, &m, &k)
		arrn := make([]int, n)
		arrm := make([]int, m)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arrn[i])
		}
		for i := 0; i < m; i++ {
			fmt.Fscan(in, &arrm[i])
		}
		sort.Ints(arrn)
		sort.Ints(arrm)

		f := func(t int) int {
			l, r := 0, len(arrm)-1
			for l < r {
				mid := (l + r + 1) >> 1
				if arrm[mid] <= t {
					l = mid
				} else {
					r = mid - 1
				}
			}
			if arrm[l] <= t {
				return l + 1
			}
			return 0
		}

		var res int
		for i := 0; i < n; i++ {
			target := k - arrn[i]
			res += f(target)
		}
		fmt.Fprintln(out, res)
	}
}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF1941A(in, out)
// }
