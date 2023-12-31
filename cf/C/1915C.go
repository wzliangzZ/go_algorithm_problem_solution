package main

import (
	"bufio"
	"fmt"
	"math"
	//"os"
)

func CF1915C(in *bufio.Reader, out *bufio.Writer) {
	var n, s, t int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		s += t
	}
	if int(math.Sqrt(float64(s)))*int(math.Sqrt(float64(s))) == s {
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
		CF1915C(in, out)
	}
}
*/
