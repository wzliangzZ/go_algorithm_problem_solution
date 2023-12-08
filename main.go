package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const (
	INT int   = 0x3f3f3f3f
	MOD int   = 1_000_000_007
	I64 int64 = (1<<bits.UintSize)/2 - 1
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)

}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		solve(in, out)
	}
}
