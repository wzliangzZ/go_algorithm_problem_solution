package DataStruct

import (
	"fmt"
	"math"
)
func ST() {
    N := 100000 + 10
    n := 100
    m := 10

    // 初始化 2^0
    f := make([][]int, N)
    for i := range f {
        f[i] = make([]int, 20)
    }
    for i := 0; i < n; i++ {
        f[i][0] = i
    }

    // 预处理
    for j := 1; j < 20; j++ {
        // 枚举每个长度为(1<<j)长度的区间左端点，若下标从1开始为+2
        for i := 0; i < n - (1 << j) + 1; i++ {
            f[i][j] = max(f[i][j-1], f[i+(1<<uint(j-1))][j-1])
        }
    }

    // 查询
    for i := 0; i < m; i++ {
        l, r := 0, 0
        k := int(math.Log2(float64(r-l+1)))
        res := max(f[l][k], f[r-1<<uint(k)+1][k])
		fmt.Println(res)
    }
}