import datetime


''' 判断极值（元素无重复) '''
a, i = [], 1
((a[i] > a[i - 1]) == (a[i] > a[i + 1]))

''' 两日期间隔多少天 '''
# 年 月 日 ---> 天数
s = datetime.date(2023, 7, 4) - datetime.date(2022, 6, 3)
# print(s.days)

'''24小时制相差多久'''
# 现将时间转化成秒
get_secends = lambda h, m, s: s + m * 60 + h * 3600
# 将相差的秒转化为24小时制
get_time = lambda s: (s // 3600, (s % 3600) // 60, (s % 3600) % 60)

