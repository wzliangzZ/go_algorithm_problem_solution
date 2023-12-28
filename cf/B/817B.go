package main

import (
	"bufio"
	"fmt"
	"sort"
	//"os"
)

func CF817B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		cnt[arr[i]]++
	}
	var t []int
	for k, _ := range cnt {
		t = append(t, k)
	}
	sort.Ints(t)
	C := func(n, k int) int {
		var res, b int = 1, 1
		for i := 0; i < k; i++ {
			res *= n - i
			b *= (i + 1)
		}
		return res / b
	}
	var k int = 3
	var res int = 1
	for i := 0; i < 3 && k > 0; i++ {
		if cnt[t[i]] >= k {
			res *= C(cnt[t[i]], k)
			k = 0
		} else {
			k -= cnt[t[i]]
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
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF817B(in, out)
	}
}
*/
