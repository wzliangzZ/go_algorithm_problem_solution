#快速幂
def ksm(m: int, k: int, p: int) -> int:
    res, t = 1 % p, m
    while k:
        if k & 1: res = res * t % p
        t = t * t % p
        k >>= 1
    return res

