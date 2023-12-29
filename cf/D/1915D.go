package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1915D(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	V := func(c byte) bool {
		if c == 'a' || c == 'e' {
			return true
		}
		return false
	}
	var res []byte
	for i := 0; i < n; i++ {
		if i != 0 && i != n-1 && !V(s[i]) && ((V(s[i-1]) && V(s[i+1])) || !V(s[i-1])) {
			res = append(res, '.')
		}
		res = append(res, s[i])
	}
	fmt.Fprintln(out, string(res))
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1915D(in, out)
	}
}
*/
