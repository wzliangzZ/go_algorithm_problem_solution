package DP

/*
划分型dp:将序列分成（恰好/至多）kkk 个连续区间，求解与这些区间有关的最优值。
f[i][j]表示将前i个元素划分为j个区间时的最优解，考虑最后一个区间，用k表示划分点，
前k个元素表示前j-1个元素，从k+1个元素到第i个元素表示第j个区间
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
