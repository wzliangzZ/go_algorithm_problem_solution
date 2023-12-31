package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1915E(in *bufio.Reader, out *bufio.Writer) {
	var n, t, tot int
	fmt.Fscan(in, &n)
	hm := make(map[int]struct{})
	var res bool
	hm[0] = struct{}{}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		if i%2 == 0 {
			t = -t
		}
		tot += t
		if _, ok := hm[tot]; ok {
			res = true
		}
		hm[tot] = struct{}{}
	}
	if res {
		fmt.Fprintln(out, "YES")
	} else {
		fmt.Fprintln(out, "NO")
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
		CF1915E(in, out)
	}
}
*/
