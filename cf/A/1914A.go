package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1914A(in *bufio.Reader, out *bufio.Writer) {
	var n int
	var s string
	fmt.Fscan(in, &n, &s)
	cnt := make([]int, 27)
	for i := 0; i < n; i++ {
		cnt[s[i]-'A'+1]++
	}
	var res int
	for i := 1; i < 27; i++ {
		if cnt[i] >= i {
			res++
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
		CF1914A(in, out)
	}

}
*/
