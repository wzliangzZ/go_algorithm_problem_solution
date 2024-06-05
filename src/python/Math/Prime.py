from math import isqrt


#试除法判断质数
def is_prime(x: int) -> int:
    if x < 2: return False;
    i = 2
    while i <= x // i:
        if x % i == 0:
            return False
        i += 1
    return True

#试除法分解质因数
def divide(x: int) -> int:
    i = 2
    res = []    #(底数，指数)
    while i <= x // i:
        if x % i == 0:
            s = 0
            while x % i == 0:
                x //= i
                s += 1
            res.append((i, s))
        i += 1
    if x > 1: res.append((x, 1))
    return res

#埃氏筛法
def get_primes(n: int) ->[]:
    MXL = 10 ** 5 + 10
    primes = [True] * (MXL + 10)
    primes[1] = False
    for i in range(2, isqrt(MXL) + 1):
        if primes[i]:
            for j in range(i * i, MXL, i):
                primes[j] = False
    return primes


# primes[];     // primes[]存储所有素数
# st[N] --> bool;     // st[x]存储x是否被筛掉
#线性筛法求质数
primes, st = [], []
def get_primes(n: int) -> []:
    for i in range(2, n + 1):
        if not st[i]: primes.append(i)
        j = 0
        while primes[j] <= n // i:
            st[primes[j] * i] = True
            if i % primes[j] == 0: break
            j += 1
    return primes


