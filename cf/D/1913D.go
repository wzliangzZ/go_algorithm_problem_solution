package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1913D(in *bufio.Reader, out *bufio.Writer) {
	const mod = 998244353
	var n, j int

	fmt.Fscan(in, &n)
	arr := make([]int, n)
	sum := make([]int, n+1)
	var stk []int
	f := make([]int, n)
	var st_sum int64
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		for len(stk) != 0 && arr[stk[len(stk)-1]] > arr[i] {
			st_sum -= int64(f[stk[len(stk)-1]])
			//st_sum -= f[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			j = 0
			f[i] += 1
		} else {
			j = stk[len(stk)-1] + 1
		}
		f[i] = (f[i] + int(st_sum%mod) + sum[i] - sum[j]) % mod
		sum[i+1] = (sum[i] + f[i]) % mod
		st_sum = (st_sum + int64(f[i])) % mod
		stk = append(stk, i)
	}
	fmt.Fprintln(out, (st_sum+mod)%mod)
}

/*
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF1913D(in, out)
	}

}
*/
