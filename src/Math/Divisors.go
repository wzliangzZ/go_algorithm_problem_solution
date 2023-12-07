package Math

func Divisors[T int | int8 | int16 | int32 | int64](x T) []T {
	var arr []T
	var i T = 1
	for ; i <= x/i; i++ {
		if x%i == 0 {
			arr = append(arr, i)
			// 约数成对出现
			if i != x/i {
				arr = append(arr, x/i)
			}
		}
	}
	return arr
}
