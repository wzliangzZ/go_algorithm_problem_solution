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

}

func main() {
    in := bufio.NewReader(os.Stdin) // 创建一个读取标准输入的读取器
    out := bufio.NewWriter(os.Stdout) // 创建一个向标准输出写入数据的写入器
    defer out.Flush() // 在函数返回之前，将缓冲区中的数据写入标准输出

    var tot int = 1 // 定义一个整数变量tot并初始化为1
    fmt.Fscan(in, &tot) // 从标准输入中读取一个整数，并将其赋值给tot

    for i := 0; i < tot; i++ { // 循环tot次
        solve(in, out) // 调用solve函数，传入in和out作为参数
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
