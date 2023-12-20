package main

import (
	"bufio"
	"fmt"
	"math"
	//"os"
)

func CF962C(in *bufio.Reader, out *bufio.Writer) {
	var num int
	fmt.Fscan(in, &num)
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	var res int = 15
	for s := 0; s < 1<<len(arr); s++ {
		var sum, cnt int
		var k int = 1
		for i := 0; i < len(arr); i++ {
			if s>>i&1 == 1 {
				cnt++
				sum += k * arr[i]
				k *= 10
			}
		}
		if k/10-sum > 0 || sum == 0 {
			continue
		}

		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
		check := func(sum int) bool {
			var n int = int(math.Sqrt(float64(sum)))
			return n*n == sum
		}
		if check(sum) {
			res = min(res, len(arr)-cnt)
		}

	}
	if res == 15 {
		fmt.Fprintln(out, -1)
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
	//fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		CF962C(in, out)
	}

}
*/
