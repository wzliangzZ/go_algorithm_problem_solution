package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1921E(in *bufio.Reader, out *bufio.Writer) {
	var h, w, xa, ya, xb, yb int
	fmt.Fscan(in, &h, &w, &xa, &ya, &xb, &yb)
	if xa >= xb {
		fmt.Fprintln(out, "Draw")
		return
	}
	var gap int = xb - xa - 1
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	la, ra := max(1, ya-(gap+1)/2), min(w, ya+(gap+1)/2)
	lb, rb := max(1, yb-gap/2), min(w, yb+gap/2)
	if gap%2 == 0 {
		if la-1 <= lb && ra+1 >= rb {
			fmt.Fprintln(out, "Alice")
		} else {
			fmt.Fprintln(out, "Draw")
		}
	} else {
		if lb-1 <= la && rb+1 >= ra {
			fmt.Fprintln(out, "Bob")
		} else {
			fmt.Fprintln(out, "Draw")
		}
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
		CF1921E(in, out)
	}
}
*/
