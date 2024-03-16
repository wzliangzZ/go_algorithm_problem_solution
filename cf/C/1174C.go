package main

import (
	"bufio"
	"fmt"
	//"os"
)

func CF1174C(in *bufio.Reader, out *bufio.Writer) {
	// var T int
	// for fmt.Fscan(in, &T); T > 0; T-- {

	// }
	var n int
	fmt.Fscan(in, &n)
	res := make([]int, n + 1)
	getPrime := func() {
		var primes []int
		st := make([]bool, 1_000_010)
		for i := 2; i <= n; i++ {
			if !st[i] {
				primes = append(primes, i)
			}
			for j := 0; primes[j] <= n/i; j++ {
				st[primes[j]*i] = true
				res[primes[j] * i] = j + 1
				if i%primes[j] == 0 {
					break
				}
			}
		}
	}
	getPrime()
	var cnt int = 1
	for i := 2; i <= n; i++ {
		if res[i] == 0 {
			fmt.Fprint(out, cnt, " ")
			cnt++
		}else{
			fmt.Fprint(out, res[i], " ")
		}
	}
	fmt.Fprintln(out)

}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	CF1174C(in, out)
// }
