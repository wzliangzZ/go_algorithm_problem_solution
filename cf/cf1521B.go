package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
)

const (
	INT int   = 0x3f3f3f3f
	MOD int   = 1000000007
	I64 int64 = (1<<bits.UintSize)/2 - 1
)

func solve(in *bufio.Reader, out *bufio.Writer) {
	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	pos := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
		if arr[i] < arr[pos] {
			pos = i
		}
	}
	fmt.Fprintln(out, n-1)
	for i := 0; i < n; i++ {
		if i != pos {
			fmt.Fprintln(out, pos+1, i+1, arr[pos], arr[pos]+Abs(pos-i))
		}
	}

}

func CF1521B() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var tot int = 1
	fmt.Fscan(in, &tot)
	for i := 0; i < tot; i++ {
		solve(in, out)
	}
}

func Abs[T byte | int | int8 | int16 | int32 | int64 | float32 | float64](data T) T {
	if data < 0 {
		return -data
	} else {
		return data
	}
}

func Max[T byte | int | int16 | int32 | int64 | int8 | float32 | float64 | string](data ...T) T {
	res := data[0]
	for i := 1; i < len(data); i++ {
		if res < data[i] {
			res = data[i]
		}
	}
	return res
}

func Min[T byte | int | int16 | int32 | int64 | int8 | float32 | float64 | string](data ...T) T {
	res := data[0]
	for i := 1; i < len(data); i++ {
		if res > data[i] {
			res = data[i]
		}
	}
	return res
}
