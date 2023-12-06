package Math

func Ksm[T int | int8 | int16 | int32 | int64](m, k, p T) T {
	res, t := 1%p, m
	for k > 0 {
		if k&1 == 1 {
			res = res * t % p
		}
		t = t * t % p
		k >>= 1
	}
	return res
}
