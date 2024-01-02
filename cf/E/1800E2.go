package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1800E2(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s, t string
	fmt.Fscan(in, &s, &t)
	cnt := [26]int{}
	for i := 0; i < n; i++ {
		if i >= n-k && i < k && s[i] != t[i] {
			fmt.Fprintln(out, "NO")
			return
		}
		cnt[s[i]-'a']++
		cnt[t[i]-'a']--
	}
	if cnt == [26]int{} {
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
		CF1800E2(in, out)
	}
}
*/
