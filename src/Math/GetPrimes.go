package Math

const N = 1_000_010

func GetPrimes(n int) []int {
	var primes []int
	st := make([]bool, N)
	for i := 2; i <= n; i++ {
		if !st[i] {
			primes = append(primes, i)
		}
		for j := 0; primes[j] <= n/i; j++ {
			st[primes[j]*i] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
	return primes
}
