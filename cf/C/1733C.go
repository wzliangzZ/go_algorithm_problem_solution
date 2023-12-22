package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1733C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	var res [][]int
	calc := func(k int) {
		var t int = -1
		for i := n - 1; i >= 0; i-- {
			if arr[i]%2 == k {
				if t == -1 {
					t = i
					continue
				}
				arr[i] = arr[t]
				res = append(res, []int{i + 1, t + 1})
				t = i
			}

		}
		k ^= 1
		for i := 0; i < n; i++ {
			if arr[i]%2 == k {
				arr[i] = arr[i-1]
				res = append(res, []int{i, i + 1})
			}
		}
	}
	if arr[0]%2 == 0 {
		calc(0)
	} else {
		calc(1)
	}
	fmt.Fprintln(out, len(res))
	for _, v := range res {
		fmt.Fprintln(out, v[0], v[1])
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
		CF1733C(in, out)
	}
}
*/
