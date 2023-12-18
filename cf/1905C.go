package main

import (
	"bufio"
	"fmt"
)

func CF1905C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	var s string
	fmt.Fscan(in, &s)
	var arr []int
	idx := make([]int, n)
	idx[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		if s[i] >= s[idx[i+1]] {
			idx[i] = i
		} else {
			idx[i] = idx[i+1]
		}
	}

	var i, res int
	for i < n {
		arr = append(arr, idx[i])
		i = idx[i] + 1
	}

	sb := []byte(s)

	cnt := make([]int, 26)
	for _, v := range arr {
		if cnt[sb[v]-'a'] == 0 {
			res++
		}
		cnt[sb[v]-'a']++
	}
	var cur int
	for i := len(arr) - 1; i >= 0; i-- {
		sb[arr[i]] = s[arr[cur]]
		cur++
	}

	for i := 1; i < len(sb); i++ {
		if sb[i-1] > sb[i] {
			res = -1
			break
		}
	}

	if res > 0 {
		fmt.Fprintln(out, res-1)
	} else {
		fmt.Fprintln(out, res)
	}

}
