package main

import (
	"bufio"
	"fmt"
	"sort"
	//"os"
)

func CF1902C(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	if n == 1 {
		fmt.Fprintln(out, 1)
		return
	}
	sort.Ints(arr)
	var gcd func(a, b int) int
	gcd = func(a, b int) int {
		if b == 0 {
			return a
		}
		return gcd(b, a%b)
	}
	var k, num int
	for i := 1; i < n; i++ {
		num = arr[i] - arr[i-1]
		if i == 1 {
			k = num
		} else {
			k = gcd(num, k)
		}
		if k == 1 {
			break
		}
	}

	other := arr[n-1] - k
	var res int = 1
	var flag bool = false
	for i := n - 2; i >= 0; i-- {
		res += (arr[n-1] - arr[i]) / k
		if !flag {
			if arr[i] == other {
				res++
				other -= k
			} else {
				flag = true
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
		CF1902C(in, out)
	}

}
*/
