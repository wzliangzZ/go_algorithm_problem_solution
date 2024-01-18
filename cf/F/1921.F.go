package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1921F(in *bufio.Reader, out *bufio.Writer) {
	var T int
	const B = 300
	var pre, sum [B][1e5 + B]int
	for fmt.Fscan(in, &T); T > 0; T-- {
		var n, q int
		fmt.Fscan(in, &n, &q)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		for d := 1; d < B; d++ {
			for i := 0; i < n; i++ {
				pre[d][i+d] = pre[d][i] + arr[i]
				sum[d][i+d] = sum[d][i] + arr[i]*(int(i/d)+1)
			}
		}

		for m := 0; m < q; m++ {
			var s, d, k int
			fmt.Fscan(in, &s, &d, &k)
			s--
			if d < B {
				var l, r = s, s + k*d
				fmt.Fprint(out, sum[d][r]-sum[d][l]-(l/d)*(pre[d][r]-pre[d][l]), " ")
			} else {
				var cur int
				for i := 0; i < k; i++ {
					cur += (i + 1) * arr[s+i*d]
				}
				fmt.Fprint(out, cur, " ")
			}
		}
		fmt.Fprintln(out, " ")
	}

}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	CF1921F(in, out)
}
*/
