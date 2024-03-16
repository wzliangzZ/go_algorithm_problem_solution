package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1922B(in *bufio.Reader, out *bufio.Writer) {
	var T int

	C := func(n, k int) int {
		if k == 2 {
			return n * (n - 1) / 2
		}
		return n * (n - 1) * (n - 2) / 6
	}

	for fmt.Fscan(in, &T); T > 0; T-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n+1)

		for i := 0; i < n; i++ {
			var t int
			fmt.Fscan(in, &t)
			arr[t]++
		}
		var res, cnt int
		for i := 0; i <= n; i++ {
			if arr[i] >= 2 {
				res += C(arr[i], 2) * cnt
			}
			if arr[i] >= 3 {
				res += C(arr[i], 3)
			}
			cnt += arr[i]
		}
		fmt.Fprintln(out, res)
	}

}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF1922B(in, out)
// }
