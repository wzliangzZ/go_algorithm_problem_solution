package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1907C(in *bufio.Reader, out *bufio.Writer) {
	var n, mx int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	var cnt = [26]int{}

	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
		if cnt[s[i]-'a'] > mx {
			mx = cnt[s[i]-'a']
		}
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}
	if mx > n/2 {
		fmt.Fprintln(out, max(2*mx-n, 0))
	} else if mx == n/2 {
		fmt.Fprintln(out, n-2*mx)
	} else {
		if n%2 == 0 {
			fmt.Fprintln(out, 0)
		} else {
			fmt.Fprintln(out, 1)
		}
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
		solve(in, out)
	}
}
*/
