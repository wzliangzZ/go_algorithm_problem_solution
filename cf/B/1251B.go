package main

import (
	"bufio"
	"fmt"
	"strings"
	//"os"
)

func CF1251B(in *bufio.Reader, out *bufio.Writer) {
	var n, res, cnt int
	fmt.Fscan(in, &n)
	var s string
	var flag bool
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		if len(s)%2 == 1 {
			flag = true
		}
		cnt += strings.Count(s, "1")
	}
	res = n
	if !flag && cnt%2 > 0 {
		res--
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
		CF1251B(in, out)
	}
}
*/
