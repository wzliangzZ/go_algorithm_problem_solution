package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1790E(in *bufio.Reader, out *bufio.Writer) {
	var x int
	fmt.Fscan(in, &x)
	var a, b int
	a = x
	for i := 32; i >= 0; i-- {
		if x&(1<<i) > 0 {
			continue
		}
		if 2*x-a-b >= (2 << i) {
			a |= (1 << i)
			b |= (1 << i)
		}
	}
	if a^b == x && a+b == 2*x {
		fmt.Fprintln(out, a, b)
	} else {
		fmt.Fprintln(out, -1)
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
		CF1790E(in, out)
	}
}
*/
