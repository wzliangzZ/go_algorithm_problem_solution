package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1730B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arrx := make([]int, n)
	var t int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arrx[i])
	}
	var mx, mn int = -0x3f3f3f3f, 0x3f3f3f3f
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		a, b := arrx[i]-t, arrx[i]+t
		if a < mn {
			mn = a
		}
		if b > mx {
			mx = b
		}
	}

	fmt.Fprintln(out, float32(mx+mn)/2)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF1730B(in, out)
	}

}
*/
