from functools import lru_cache


"""
求组合数的3种方式(模MOD)
1)  递推     C(a, b) = C(a - 1, b) + C(a - 1, b - 1)    时间复杂度: O(n^2)
2)  预处理
        C(a, b) = a! / (a - b)! * b!
        fact[a]     表示a的阶乘
        infact[a]   表示a的阶乘的逆元
        fact[0] = infact[0] = 1
        C(a, b) = fact[a] * infact[a - b] * infact[b]
    时间复杂度：O(nlogn)

3)  Lucas (MOD值较小)
        C(a, b) 同余 C(a % MOD, b % MOD) * lucas(a // MOD, b // MOD)
        相当于把n变成MOD进制
    时间复杂度:O(log(MOD, n) * MOD * log(MOD))


"""
# O(nlogn)


def C(N: int = 0, MOD: int = 1000000007) -> []:
    c = [[0] * N for _ in range(N)]
    for i in range(N):
        for j in range(i + 1):
            c[i][j] = (1 if j == 0 else (c[i - 1][j] + c[i - 1][j - 1]) % MOD)
    return c

@lru_cache(None)
def C(a: int = 0, b: int = 0, MOD: int = 1000000007) -> int:
    if a < 0: return 0
    if b == 0: return 1
    if b == 1: return a
    return (C(a - 1, b) + C(a - 1, b - 1)) % MOD

# 预处理 + lucas
class Comb(object):

    def __init__(self, n: int = 0, MOD: int = 1000000007):
        self.L = n + 1
        self.MOD = MOD
        self.__fac = [0] * self.L
        self.__invfac = [0] * self.L
        self.__init()

    def __init(self) -> None:
        self.__fac[0] = self.__invfac[0] = 1
        for i in range(1, self.L):
            self.__fac[i] = self.__fac[i - 1] * i % self.MOD
            self.__invfac[i] = self.__invfac[i - 1] * self.__ksm(i, self.MOD - 2) % self.MOD

    def __ksm(self, m: int, k: int) -> int:
        res = 1 % self.MOD
        while k:
            if k & 1: res = res * m % self.MOD
            m = m * m % self.MOD
            k >>= 1
        return res

    def __C_by_def(self, a: int = 0, b: int = 0) -> int:
        if a < b: return 0
        res = 1 % self.MOD
        # a! / (b!(a-b)!) = (a - b + 1) * ... * a / b! 分子有b项
        i, j = 1, a
        while i <= b:
            res = res * j % self.MOD
            res = res * self.__ksm(i, self.MOD - 2) % self.MOD
            i += 1
            j -= 1
        return res

    def fac(self, x: int = 0) -> int: return self.__fac[x]
    def invfac(self, x: int = 0) -> int: return self.__invfac[x]

    def updateMOD(self, mod: int = 1000000007): self.MOD = mod

    #预处理方式  求组合数
    def C(self, a: int = 0, b: int = 0) -> int: return self.__fac[a] * self.__invfac[a - b] * self.__invfac[b] % self.MOD

    #lucas  求组合数
    def lucas(self, a: int = 0, b: int = 0) -> int:
        if a < self.MOD and b < self.MOD: return self.__C_by_def(a, b)
        return self.__C_by_def(a % self.MOD, b % self.MOD) * self.lucas(a // self.MOD, b // self.MOD) % self.MOD

