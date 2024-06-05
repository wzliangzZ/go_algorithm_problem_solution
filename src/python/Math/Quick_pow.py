#快速幂
def ksm(m: int, k: int, p: int) -> int:
    res = 1 % p
    while k:
        if k & 1: res = res * m % p
        m = m * m % p
        k >>= 1
    return res