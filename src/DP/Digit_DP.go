package DP

import (
	"strconv"
	"strings"
)

// 不包含前导零
func DigitDP(start int64, end int64) int64 {
	low, high := strconv.FormatInt(start, 10), strconv.FormatInt(end, 10)
	n := len(high)
	// low 与 high 对齐
	low = strings.Repeat("0", n - len(low)) + low 
	
	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}
	var dfs func(int, bool, bool) int64

	dfs = func(i int, low_limit, high_limit bool) (res int64) {
		if i == n { return 1 }

		if !low_limit && !high_limit {
			p := &memo[i]
			if *p >= 0 {
				return *p
			}
			defer func(){ *p = res }()
		}

		lo, hi := 0, 9
		if low_limit { lo = int(low[i] - '0')}
		if high_limit { hi = int(high[i] - '0')}
		for d := lo; d <= hi; d++ {
			res += dfs(i + 1, low_limit && d == lo, high_limit && d == hi)
		}
		return res
	}
	return dfs(0, true, true)
}

// 包含前导0, 每一位都不相同的方案数
func DigitDP_zero(start int64, end int64) int64{
	low, high := strconv.FormatInt(start, 10), strconv.FormatInt(end, 10)
	n := len(high)
	// low 与 high 对齐
	low = strings.Repeat("0", n - len(low)) + low 
	
	memo := make([][]int64, n)
	for i := range memo {
		memo[i] = make([]int64, 1 << 10)
        for j := range memo[i]{
            memo[i][j] = -1
        }
	}
	var dfs func(int, int, bool, bool, bool) int64

	dfs = func(i, mask int, low_limit, high_limit, is_num bool) (res int64) {
		if i == n { if is_num { return 1} else { return 0 } }	// 0 不算入答案
		// if i == n { return 1 } // 0 也算入答案
		if !low_limit && !high_limit {
			p := &memo[i][mask]
			if *p >= 0 {
				return *p
			}
			defer func(){ *p = res }()
		}
		if !is_num && low[i] == '0' {
			res += dfs(i + 1, mask, true, false, false)
		}
		lo, hi, d0 := 0, 9, 1
		if low_limit { lo = int(low[i] - '0')}
		if high_limit { hi = int(high[i] - '0')}
		if is_num { d0 = 0 }
		
		for d := max(lo, d0); d <= hi; d++ {
            if (mask >> d) & 1 == 0{
                res += dfs(i + 1, mask | (1 << d), low_limit && d == lo, high_limit && d == hi, true)
            }
		}
		return 
	}
	return dfs(0, 0, true, true, false)
}