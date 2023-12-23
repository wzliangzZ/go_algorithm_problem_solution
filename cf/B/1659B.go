package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1659B(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)
	cnt := make([]int, n)
	tk := k
	for i := 0; i < n && tk > 0; i++ {
		if k%2 == int(s[i]-'0') {
			cnt[i]++
			tk--
		}
	}
	cnt[n-1] += tk
	sb := []byte(s)
	for i, v := range cnt {
		if (k-v)%2 == 1 {
			sb[i] = '1' - (s[i] - '0')
		}
	}
	fmt.Fprintln(out, string(sb))
	for _, v := range cnt {
		fmt.Fprint(out, v, " ")
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
		CF1659B(in, out)
	}
}
*/
