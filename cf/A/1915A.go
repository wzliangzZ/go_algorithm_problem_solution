package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1915A(in *bufio.Reader, out *bufio.Writer) {
	var a, b, c int
	fmt.Fscan(in, &a, &b, &c)
	if a == b {
		fmt.Fprintln(out, c)
	} else if a == c {
		fmt.Fprintln(out, b)
	} else {
		fmt.Fprintln(out, a)
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
		CF1915A(in, out)
	}
}
*/
