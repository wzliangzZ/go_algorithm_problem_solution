package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1921C(in *bufio.Reader, out *bufio.Writer) {
	var n, f, a, b int
	fmt.Fscan(in, &n, &f, &a, &b)
	var pre int
	min := func(a, b int) int {
		if a > b {
			return b
		} else {
			return a
		}
	}
	for i := 0; i < n; i++ {
		var cur int
		fmt.Fscan(in, &cur)
		f -= min((cur-pre)*a, b)
		pre = cur
	}
	if f <= 0 {
		fmt.Fprintln(out, "NO")
	} else {
		fmt.Fprintln(out, "YES")
	}

}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1921C(in, out)
	}
}
*/
