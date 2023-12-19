package main

import (
	"bufio"
	"fmt"
)

func CF1913C(in *bufio.Reader, out *bufio.Writer) {
	var n, a, b int
	fmt.Fscan(in, &n)
	cnt := make([]int, 30)

	check := func(x int) string {
		arr := make([]int, 30)
		copy(arr, cnt)
		for i := 0; i < 30; i++ {
			if (x>>i)&1 == 1 {
				if arr[i] > 0 {
					arr[i]--
				} else {
					var k int = 1 << i
					for j := i - 1; j >= 0 && k > 0; j-- {
						if (1<<j)*arr[j] <= k {
							k -= (1 << j) * arr[j]
							arr[j] = 0
						} else {
							arr[j] -= k / (1 << j)
							k %= (1 << j)
						}
						// for arr[j] > 0 && k >= (1 << j){
						// 	arr[j]--
						// 	k -= 1 << j
						// }
					}
					if k != 0 {
						return "NO"
					}
				}
			}
		}
		return "YES"
	}

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b)
		if a == 1 {
			cnt[b]++
		} else {
			fmt.Fprintln(out, check(b))
		}
	}

}
