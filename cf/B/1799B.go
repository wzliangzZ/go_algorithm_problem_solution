package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1799B(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	var cnt int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] == 1 {
			cnt++
		}
	}
	if cnt > 0 && cnt != n {
		fmt.Fprintln(out, -1)
		return
	}
	var res [][]int
	get := func(k int) int {
		var cur, idx int
		if k > 0 {
			cur = -1
		} else {
			cur = 1e9 + 7
		}
		for i := 0; i < n; i++ {
			if arr[i]*k > cur*k {
				cur = arr[i]
				idx = i
			}
		}
		return idx
	}
	for {
		a, b := get(1), get(-1)
		if a == b {
			break
		}
		res = append(res, []int{a, b})
		arr[a] = (arr[a] + arr[b] - 1) / arr[b]
	}
	fmt.Fprintln(out, len(res))
	for _, v := range res {
		fmt.Fprintln(out, v[0]+1, v[1]+1)
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
		CF1799B(in, out)
	}
}
*/
