package main

import (
	"bufio"
	"fmt"
	//"os"
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

	var i, res, cur int
	for i < n {
		arr = append(arr, idx[i])
		i = idx[i] + 1
	}
	sb := []byte(s)

	for i := len(arr) - 1; i >= 0; i-- {
		sb[arr[i]] = s[arr[cur]]
		cur++
	}
	res = len(arr)
	for i := 1; i < len(arr); i++ {
		if s[arr[i]] != s[arr[i-1]] {
			res -= i
			break
		}
	}
	if res == len(arr) {
		res = 0
	}

	for i := 1; i < len(sb); i++ {
		if sb[i-1] > sb[i] {
			res = -1
			break
		}
	}

	if res > 0 {
		fmt.Fprintln(out, res)
	} else {
		fmt.Fprintln(out, res)
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
		CF1905C(in, out)
	}

}
*/
