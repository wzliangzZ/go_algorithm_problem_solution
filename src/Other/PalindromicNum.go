package Other

// 从小到大的顺序生成10^9以内的回文数
func GetPalindromicNum() []int {
	// 枚举一般的范围即可
	var pal = make([]int, 0, 109999)
	// 按顺序从小到大生成所有回文数
	for base := 1; base <= 10000; base *= 10 {
		for i := base; i < base*10; i++ {
			x := i
			for t := i / 10; t > 0; t /= 10 {
				x = x*10 + t%10
			}
			pal = append(pal, x)
		}
		//防止产生的数越界
		if base <= 1000 {
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				pal = append(pal, x)
			}
		}
	}
	pal = append(pal, 1_000_000_001) // 哨兵，防止下标越界
	return pal
}