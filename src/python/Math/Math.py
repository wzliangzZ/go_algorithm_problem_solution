alloc = lambda dfv, *s: len(s) != 1 and [alloc(dfv, *s[1:]) for _ in range(int(s[0]))] or [dfv] * int(s[0])
#求a,b的最大公约数
def gcd(a, b):
    return gcd(b, a % b) if b else a

gcd = lambda a, b: gcd(b, a % b) if b else a

def multi_gcd(array: []) -> []:
    l = len(array)
    if l == 1: return array[0]
    elif l == 2: return gcd(array[0], array[1])
    else: return gcd(multi_gcd(array[:l // 2]), multi_gcd(array[l // 2:]))


#求最小公倍数
lcm = lambda a, b: a * b // gcd(a, b)

def multi_lcm(array: []) -> []:
    l = len(array)
    if l == 1: return array[0]
    elif l == 2: return lcm(array[0], array[1])
    else: return lcm(multi_lcm(array[:l // 2]), multi_lcm(array[l // 2:]))


#试除法求约数
def get_divisors(x: int) -> int:
    res, i = [], 1
    while i <= x // i:
        if x % i == 0:
            res.append(i)
            if i != x // i: res.append(x // i)
        i += 1
    res.sort()
    return res






#r点在向量pq左侧(上方)时大于0，等于0三点共线
def cross(p: [], q: [], r: []) -> int:
    return (q[0] - p[0]) * (r[1] - q[1]) - (q[1] - p[1]) * (r[0] - q[0])

