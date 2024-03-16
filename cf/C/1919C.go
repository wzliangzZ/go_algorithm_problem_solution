package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1919C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var a, b, res, cur int = 0x3f3f3f3f, 0x3f3f3f3f, 0, 0
	exc := func(a, b, c *int) {
		if *a > *b {
			*b = *c
		} else {
			*a = *c
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &cur)
		if cur > a && cur > b {
			res++
			exc(&a, &b, &cur)
		} else if cur <= a && cur <= b {
			exc(&a, &b, &cur)
		} else if a > b {
			a = cur
		} else {
			b = cur
		}
	}

	fmt.Fprintln(out, res)

}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1919C(in, out)
	}
}
*/
