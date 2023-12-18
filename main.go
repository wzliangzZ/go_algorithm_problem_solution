package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n, k int
	fmt.Fscan(in, &n, &k)
	arr := make([]int64, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(in, &arr[i])
	}
	var last int64 = 0x3f3f3f3f * 2
	for i := k - 2; i >= 0; i-- {
		if arr[i + 1] - arr[i] > last{
			fmt.Fprintln(out, "NO")
			return
		}
		last = arr[i + 1] - arr[i]
	}
	n -= k - 1
	if int64(n) * last < arr[0]{
		fmt.Fprintln(out, "NO")
	}else{
		fmt.Fprintln(out, "YES")
	}
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	//fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		solve(in, out)
	}

}
