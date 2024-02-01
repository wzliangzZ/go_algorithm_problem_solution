package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1922C(in *bufio.Reader, out *bufio.Writer) {
	var T int

	for fmt.Fscan(in, &T); T > 0; T-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n+2)
		closet := make([]int, n+2)
		left := make([]int, n+2)
		right := make([]int, n+2)
		arr[0], arr[n+1] = -0x3f3f3f3f, 0x3f3f3f3f3f3f
		for i := 1; i <= n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		choose := func(i, k, j int) int {
			abs := func(x int) int {
				if x < 0 {
					return -x
				} else {
					return x
				}
			}
			if abs(arr[k]-arr[i]) <= abs(arr[k]-arr[j]) {
				return i
			}
			return j
		}

		for i := 1; i <= n; i++ {
			closet[i] = choose(i-1, i, i+1)
			if closet[i] > i {
				left[i+1] += left[i] + 1
			} else {
				left[i+1] += left[i] + arr[i+1] - arr[i]
			}
		}
		for i := n; i >= 1; i-- {
			if closet[i] < i {
				right[i-1] += right[i] + 1
			} else {
				right[i-1] += right[i] + arr[i] - arr[i-1]
			}
		}
		var m int
		fmt.Fscan(in, &m)
		for i := 0; i < m; i++ {
			var a, b int
			fmt.Fscan(in, &a, &b)
			if a < b {
				fmt.Fprintln(out, left[b]-left[a])
			} else {
				fmt.Fprintln(out, right[b]-right[a])
			}
		}
	}

}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF1922C(in, out)
// }
