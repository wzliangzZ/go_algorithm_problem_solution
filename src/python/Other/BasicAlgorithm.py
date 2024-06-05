import bisect
import datetime


def check(*s):
    pass


#区间[l, r]被划分成[l, mid]和[mid + 1, r]时使用：
def bsearch_1(l: int, r: int, target: int) -> int:
    while (l < r):
        mid = l + r >> 1
        if check(mid): r = mid    # check()判断mid是否满足性质
        else: l = mid + 1
    return l


# 区间[l, r]被划分成[l, mid - 1]和[mid, r]时使用：
def bsearch_2(l: int, r: int, target: int) -> int:
    while (l < r):
        mid = l + r + 1 >> 1
        if check(mid): l = mid
        else: r = mid - 1
    return l


#浮点数二分
def bsearch_3(l: int, r: int, target: int) -> float:
    eps = 1e-6;   # eps 表示精度，取决于题目对精度的要求
    while (r - l > eps):
        mid = (l + r) / 2
        if (check(mid)): r = mid
        else: l = mid
    return l



#获得下标从1开始的一维前缀和
def prefix_sum1(arr: []) -> int:
    res = [0]
    for num in arr: res.append(res[-1] + num)
    return res


#获得下标从1开始的二维前缀和
def prefix_sum2(arr: []) -> int:
    n, m = len(arr) + 1, len(arr[0]) + 1
    res = [[0] * m for _ in range(n)]
    for i in range(1, n):
        for j in range(1, m):
            res[i][j] = res[i - 1][j] + res[i][j - 1] - res[i - 1][j - 1] + arr[i - 1][j - 1]
    return res


#给区间[l, r]中的每个数加上c：A[l] += c, A[r + 1] -= c, 最后求A的前缀和
#一维差分数组，长度为len(arr) + 2
def difference1(arr: []):
    res = [0]
    for i in range(len(arr)):
        if i == 0: res.append(arr[0])
        else: res.append(arr[i] - arr[i - 1])
    res.append(0)
    return res


# 给以(x1, y1)为左上角，(x2, y2)为右下角的子矩阵中的所有元素加上c：
# S[x1, y1] += c, S[x2 + 1, y1] -= c, S[x1, y2 + 1] -= c, S[x2 + 1, y2 + 1] += c 最后求S的前缀和
#二维差分在原数组进行
def difference2_insert(x1: int, y1: int, x2: int, y2: int, arr: [], c: int):
    arr[x1][y1] += c
    arr[x2 + 1][y2 + 1] += c
    arr[x2 + 1][y1] -= c
    arr[x1][y2 + 1] -= c




#设arr是排好序的列表
#greater = n - bisect.bisect_right(arr, x)      # arr中有多少个数严格大于x
#lesser = bisect.bisect_left(arr, x)            # arr中有多少个数严格小于x

'''
说明
1. bisect.bisect_left返回大于等于x的第一个下标(相当于cpp的lower_bound)
2. bisect.bisect_right返回大于x的第一个下标(相当于cpp的upper_bound),x最大就返回len(arr)
3. n - greater 即为小于等于x的数量
4. n - lesser 即为大于等于x的数量
'''


'''
date1 = datetime.date(2022, 12, 30)
delta = datetime.timedelta(days=1)
date1 += delta # 下一天
year, month, day = date1.timetuple()[:3]
weekday = date1.isoweekday() # 星期，从1到7
passdays = date1.timetuple().tm_yday # 当年的第几天，从1到365
'''
# d1 = datetime.date(2021, 1, 12)
# d2 = datetime.date(2021,1,14)
# s = d2 - d1
# print(s.days)

def get(h, m, mi):
    return mi + m * 60 + h * 3600

res = get(22, 45,53) - get(18, 55, 24)
print(res)
print(res // 3600, (res % 3600) // 60, (res % 3600) % 60)
