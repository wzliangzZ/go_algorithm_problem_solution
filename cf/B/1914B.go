package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1914B(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	var p int = n - k
	for i := p; i <= n; i++ {
		fmt.Fprint(out, i, " ")
	}
	for i := p - 1; i > 0; i-- {
		fmt.Fprint(out, i, " ")
	}
	fmt.Fprintln(out)

}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF1914B(in, out)
	}

}
*/
