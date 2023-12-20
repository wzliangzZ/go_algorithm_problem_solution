package main

import (
	"bufio"
	"fmt"
	"sort"
	//"os"
)

func CF1864B(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	var s string
	fmt.Fscan(in, &s)
	if k%2 == 0 {
		res := []byte(s)
		sort.Slice(res, func(i, j int) bool {
			return res[i] < res[j]
		})
		fmt.Fprintln(out, string(res))
	} else {
		var arr1, arr2 []byte
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr1 = append(arr1, s[i])
			} else {
				arr2 = append(arr2, s[i])
			}
		}
		sort.Slice(arr1, func(i, j int) bool {
			return arr1[i] < arr1[j]
		})
		sort.Slice(arr2, func(i, j int) bool {
			return arr2[i] < arr2[j]
		})
		var i, j int
		for k := 0; k < n; k++ {
			if k%2 == 0 {
				fmt.Fprint(out, string(arr1[i]))
				i++
			} else {
				fmt.Fprint(out, string(arr2[j]))
				j++
			}
		}
		fmt.Fprintln(out)
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
		CF1864B(in, out)
	}
}
*/
