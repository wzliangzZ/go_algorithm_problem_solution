
# 模式串P,主串S,求模式串P在主串S中出现的位置(下标从1开始)
# 时间复杂度O(n+m)

# p,s下标从1开始
def kmp(p, s):
    ne = [0] * len(p)
    #求 ne 数组
    ne[1] = j = 0
    for i in range(2, len(p)):
        #判断j+1
        while j and p[i] != p[j + 1]: j = ne[j]
        if p[i] == p[j + 1]: j += 1
        ne[i] = j
    j = 0
    for i in range(1, len(s)):
        if j and s[i] != p[j + 1]: j = ne[j]
        if s[i] == p[j + 1]: j += 1
        if j == len(p) - 1: return i - j + 1
    return -1

# 取 最长的 相等前后缀
# 对模式串前后缀 自我匹配 ,计算next数组
# next[i] 表示模式串p[1,i)中 相等前后缀的 最长 长度

#求ne数组
# 双指针： i扫描模式串，j扫描前缀
# 初始化：ne[1]=0, i=2, j=0
# 每轮for循环，i向右走一步
# 1)若p[i]!=p[j+1],让j回眺到能匹配位置，若没有，则j=0
# 2)若p[i]==p[j+1],则j+=1,指向匹配前缀的末尾
# 3)ne[i]=j



#扩展KMP

# 下标从1开始
def get_z(s):
    z = [0] * len(s)
    z[1] = len(s)
    l = r = 0
    for i in range(len(s)):
        if i <= r: z[i] = min(z[i - l + 1], r - l + 1)
        while (i + z[i] < len(s) and s[1 + z[i]] == s[i + z[i]]): z[i] += 1
        if i + z[i] - 1 > r: l, r = i, i + z[i] - 1
    return z

# Z函数：对于一个长度为n的字符串S。z[i]表示S与其后缀s[i,n]的最长公共前缀(LCP)的长度
#   i   1  2  3  4  5  6  7  8  9
#   s   a  a  a  b  a  a  a  b  c
# z[i]  9  2  1  0  4  2  1  0  0

# Z-box
# 对于i，区间 [i,i+z[i]-1] 是i的匹配段，即：z-box
# 维护右端点的最靠右的匹配段，记作[l,r],S[l,r]一定是S的前缀
# 计算z[i]时，保证 l<=i

#   i   1  2  3  4  5  6  7  8  9
#   s   a  a  a  b |a  a  a  b| c
# z[i]  9  2  1  0 |4  2  1  0| 0

#算法流程
# 2) 计算完前i-1的z值，盒子长度为[l,r], S[l,r]==S[1, r-l+1]
# 1. 如果 i<=r (在盒内),则有 S[i,r]== S[i-l+1, r-l+1]
#    i.若z[i-l+1]<r-i+1,则z[i]=z[i-l+1]
#   ii.若z[i-l+1]>=r-i+1, 则z[i]=z[r-i+1]，从r+1往后暴力枚举
# 2. 如果i>r(在盒外),则从i开始枚举
# 3. 求出z[i]后，如果i+z[i]-1>r,则更新盒子，l=i, r=i+z[i]-1
