package Math

const mod = 1_000_000_007
const mx = 100_000

var fac, invFac [mx]int

func init() {
	fac[0] = 1
	for i := 1; i < mx; i++ {
		fac[i] = fac[i-1] * i % mod
	}
	invFac[mx-1] = Pow(fac[mx-1], mod-2)
	for i := mx - 1; i > 0; i-- {
		invFac[i-1] = invFac[i] * i % mod
	}
}

func Comb(n, k int) int {
	return fac[n] * invFac[k] % mod * invFac[n-k] % mod
}

func Pow(x, n int) (res int) {
	res = 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return
}
