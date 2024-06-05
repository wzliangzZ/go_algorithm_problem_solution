from functools import lru_cache

alloc = lambda *s, dfv: len(s) != 1 and [alloc(*s[1:], dfv=dfv) for _ in range(int(s[0]))] or [dfv] * int(s[0])



#区间DP
def interval_dp():
    n = 100
    f = alloc(n+1,n+1,dfv=0)
    prefix = []

    for len in range(1, n + 1):
        for l in range(1, n - len + 2):
            r = l + len - 1
            #特判长度为1
            if len == 1:
                f[l][r] = 0
                continue
            #枚举断点k, f(l,r) = f(l,k) + f(k+1,r)
            for k in range(l, r):
                pass
