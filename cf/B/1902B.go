package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1902B(in *bufio.Reader, out *bufio.Writer) {
	var n, p, l, t int
	fmt.Fscan(in, &n, &p, &l, &t)
	tasks := n / 7
	if n%7 != 0 {
		tasks++
	}
	ats := (tasks / 2) * (t*2 + l)
	if ats >= p {
		k := (2*t + l)
		res := (p + k - 1) / k
		fmt.Fprintln(out, n-res)
	} else {
		res := tasks / 2
		p -= res * (t*2 + l)
		if tasks%2 != 0 {
			p -= t + l
			res++
		}
		if p > 0 {
			res += (p + l - 1) / l
		}
		fmt.Fprintln(out, n-res)
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
		CF1902B(in, out)
	}

}
*/
