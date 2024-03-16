package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1941B(in *bufio.Reader, out *bufio.Writer) {
	var T int
	for fmt.Fscan(in, &T); T > 0; T-- {
		var n int
		fmt.Fscan(in, &n)
		arr := make([]int, n)
		res := true
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}
		for i := 1; i < n-1; i++ {
			arr[i] -= 2 * arr[i-1]
			arr[i+1] -= arr[i-1]
			if arr[i] < 0 {
				res = false
				break
			}
		}
		if arr[n-1] == 0 && arr[n-2] == 0 && res {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF1941B(in, out)
// }
