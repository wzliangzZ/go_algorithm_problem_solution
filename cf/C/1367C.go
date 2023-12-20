package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1367C(in *bufio.Reader, out *bufio.Writer) {
	var n, k, res int
	var s string
	fmt.Fscan(in, &n, &k)
	fmt.Fscan(in, &s)
	var arr = []int{-1}
	for i, v := range s {
		if v == '1' {
			arr = append(arr, i)
		}
	}
	arr = append(arr, n)
	if len(arr) == 2 {
		res = n / (k + 1)
		if n%(k+1) != 0 {
			res++
		}
		fmt.Fprintln(out, res)
		return
	}
	for i := 1; i < len(arr); i++ {
		r := arr[i] - 1
		l := arr[i-1]
		if r > l {
			res += (r - l) / (k + 1)
			if i != 1 && i != len(arr)-1 && (r-l)%(k+1) != k {
				res--
			}
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
	fmt.Fscan(in, &tot)

	for i := 0; i < tot; i++ {
		CF1367C(in, out)
	}
}
*/
