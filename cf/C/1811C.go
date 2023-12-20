package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1811C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(in, &arr[i])
	}

	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	fmt.Fprint(out, arr[0], " ")
	for i := 1; i < n-1; i++ {
		fmt.Fprint(out, min(arr[i], arr[i-1]), " ")
	}
	if n >= 2 {
		fmt.Fprint(out, arr[n-2])
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
		CF1811C(in, out)
	}
}
*/
