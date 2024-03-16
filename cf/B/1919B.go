package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1919B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	var a, b int
	for i := 0; i < n; i++ {
		if s[i] == '+' {
			a++
		} else {
			b++
		}
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}
	fmt.Fprintln(out, abs(a-b))
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1919B(in, out)
	}
}
*/
