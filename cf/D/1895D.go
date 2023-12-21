package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1895D(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	cnt := make([]int, 33)
	for i := 1; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		arr[i] ^= arr[i-1]
		for j := 0; j < 31; j++ {
			if (arr[i]>>j)&1 == 1 {
				cnt[j]++
			}
		}
		for j := 0; j < 31; j++ {
			if (i>>j)&1 == 1 {
				cnt[j]--
			}
		}
	}
	var b int
	for i, v := range cnt {
		if v > 0 {
			b |= 1 << i
		}
	}
	fmt.Fprint(out, b, " ")
	for i := 1; i < n; i++ {
		fmt.Fprint(out, arr[i]^b, " ")
	}
	fmt.Fprintln(out)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1895D(in, out)
	}
}
*/
