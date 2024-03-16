package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1919A(in *bufio.Reader, out *bufio.Writer) {
	var n, m int
	fmt.Fscan(in, &m, &n)
	if (n+m)%2 == 1 {
		fmt.Fprintln(out, "Alice")
	} else {
		fmt.Fprintln(out, "Bob")
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
		CF1919A(in, out)
	}
}
*/
