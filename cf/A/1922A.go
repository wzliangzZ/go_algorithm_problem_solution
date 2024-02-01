package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1922A(in *bufio.Reader, out *bufio.Writer) {
	var T int
	for fmt.Fscan(in, &T); T > 0; T-- {
		var n int
		fmt.Fscan(in, &n)
		var a, b, c string
		fmt.Fscan(in, &a, &b, &c)
		var res = "NO"
		for i := 0; i < n; i++ {
			if c[i] != a[i] && c[i] != b[i] {
				res = "YES"
				break
			}
		}
		fmt.Fprintln(out, res)
	}
}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF1922A(in, out)
// }
