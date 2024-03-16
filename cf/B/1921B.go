package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1921B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s, f string
	var a, b int
	fmt.Fscan(in, &s, &f)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' && f[i] == '0' {
			a++
		}
		if s[i] == '0' && f[i] == '1' {
			b++
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}
	fmt.Fprintln(out, max(a, b))
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1921B(in, out)
	}
}
*/
