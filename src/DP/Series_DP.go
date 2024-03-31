package DP

/*
划分型dp:将序列分成（恰好/至多）k 个连续区间，求解与这些区间有关的最优值。

f[i][j]表示将前 i 个元素划分为 j 个区间时的最优解，考虑最后一个区间，用 k 表示划分点，
前 k 个元素表示前 j-1 个元素，从 k+1 个元素到第 i 个元素表示第 j 个区间。

f[i][j] = max(f[i][j], f[k][j-1]+pf[i]-pf[k])

i >= j
j = min(i, m)
j - 1 <= k < i

*/

func SeriesDP(nums []int, m int) {
	//预处理
	n := len(nums)
	f := make([][]int, n+1)
	pf := make([]int, n+1)
	for i, x := range nums {
		pf[i+1] = pf[i] + x
	}
	for i := 0; i < n+1; i++ {
		f[i] = make([]int, m+1)
	}

	//核心代码，注意 j k 的取值范围
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, m); j++ {
			for k := j - 1; k < i; k++ {
				f[i][j] = max(f[i][j], f[k][j-1]+pf[i]-pf[k])
			}
		}
	}
}
