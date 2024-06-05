from functools import lru_cache

alloc = lambda *s, dfv: len(s) != 1 and [alloc(*s[1:], dfv=dfv) for _ in range(int(s[0]))] or [dfv] * int(s[0])




#数位DP
"""
提问形式：某一个区间[X,Y]之间满足某一性质的数的个数
f(Y) - f(X - 1)
"""
def digit_dp(n):
    #特判边界
    if n == 0: return 0
    digits = []
    #把n的每一位都存入数组，（倒序存储）
    while n:
        digits.append(n % 10)
        n //= 10
    #存储结果和前面填的数位信息
    res, last = 0, 0
    #倒序枚举从 [len-1, 0]
    for i in range(len(digits) - 1, -1, -1):
        x = digits[i]
        #枚举左边分支 [0,x)
        for j in range(x):
            pass

        #全部枚举完
        if i == 0: res += 1


'''
    返回从i开始填，i前面填的数字是mask, 所有合法方案数
    
    is_limit 表示 前面的数字 是否都是n位上的数字， 如果为true,那么当前位至多为s[i],否则为9
    
    is_num 表示 前面是否填了数字 （是否跳过,false 可以跳过,可以枚举更短位数的数字)。
        如果为true,那么当前位可以从0开始(因为前面有数字)，否则从1开始填数字
    
    如果 is_num 参数存在，即枚举的数字绝不可能以前导0出现，自然数字0也不会统计在内
'''
#记忆化模板
#在一个范围内(限制)，满足某些条件(约束)的数的个数
def digit_dp(n: int):
    s = str(n)  # 把n转换成字符串
    @lru_cache(None)
    def dp(i: int, mask: int, is_limit: bool, is_num: bool) -> int:
        if i == len(s): return int(is_num)
        res = 0
        if not is_num:  #跳过当前位
            res += dp(i + 1, mask, False, False)
        up = int(s[i]) if is_limit else 9

        #枚举要填的数字，枚举范围取决于 is_limt 和  is_num
        for d in range(1 - int(is_num), up + 1):
            if mask >> d & 1 == 0:      #处理约束逻辑
                res += dp(i + 1, mask | (1 << d), is_limit and d == up, True)
        return res

    return dp(0, 0, True, False)


#数组版
def digit_dp(n: int) -> int:
    s = str(n)
    f = [[-1] * 10 for _ in range(len(s))]
    #f = alloc(-1, len(s), 10)
    def dp(i: int, last: int, is_limit: bool, is_num: bool) -> int:
        if i == len(s): return int(is_num)

        #如果 不受限制 且 填了数字 且 计算过   直接返回
        if not is_limit and is_num and f[i][last] != -1: return f[i][last]
        res = 0
        if not is_num: res += dp(i + 1, 0, False, False)
        up = int(s[i]) if is_limit else 9

        for d in range(1 - int(is_num), up + 1):
            if not is_num or abs(last - d) == 1:
                res += dp(i + 1, d, is_limit and d == up, True)

        # 更新答案
        if not is_limit and is_num: f[i][last] = res
        return res


"""
问：记忆化四个状态有点麻烦，能不能只记忆化 i 和 pre 这两个状态？

答：可以的！比如 n=234，第一位填 2，第二位填 3，后面无论怎么递归，
都不会再次递归到第一位填 2，第二位填 3 的情况，所以不需要记录。
第一位不填，第二位也不填，后面无论怎么递归也不会再次递归到这种情况，所以也不需要记录。

根据这个例子，我们可以只记录不受到约束时的状态 (i,mask,false,true)。比如 n=456，
第一位（最高位）填的 3，那么继续递归，后面就可以随便填，所以状态
(1,3,false,true)就表示 i=0 填 3，从 i=1 往后随便填的方案数。
由于后面两个参数恒为 false和true，所以可以不用记忆化，只记忆化 i 和 pre。
"""


# v2.0
def dig_dp(start: str, end: str) -> int:
    low = start
    high = end
    n = len(high)
    low = '0' * (n - len(low)) + low

    #不支持前导0
    @lru_cache()
    def dfs(i:int, limit_low:bool, limit_high:bool) -> int:
        if i == n: return 1
        lo = int(low[i]) if limit_low else 0
        hi = int(high[i]) if limit_high else 9
        res = 0
        for d in range(lo, hi + 1):
            res += dfs(i + 1, limit_low and d == lo, limit_high and d == hi)
        return res
    return dfs(0, True, True)

#要支持前导0， 引入参数 is_num: true -> 前面填了非0数字， dfs(0, True, True, False)
def dig_dp(start: str, end: str) -> int:
    low = start
    high = end
    n = len(high)
    low = '0' * (n - len(low)) + low

    @lru_cache()
    def dfs(i:int, limit_low:bool, limit_high:bool, is_num:bool) -> int:
        if i == n: return int(is_num)
        res = 0
        if not is_num and low[i] == '0':    #前面都是前导0那么limit_low 一定等于true, 因此，此处必须要判断low[i] == '0'，否则，就会忽略下界约束
            res += dfs(i + 1, True, False, False)

        lo = int(low[i]) if limit_low else 0
        hi = int(high[i]) if limit_high else 9

        do = 0 if is_num else 1

        for d in range(max(lo, do), hi + 1):
            res += dfs(i + 1, limit_low and d == lo, limit_high and d == hi, False)
        return res
    return dfs(0, True, True, False)

