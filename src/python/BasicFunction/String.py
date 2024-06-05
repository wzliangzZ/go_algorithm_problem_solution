# python 的 string 为不可变对象

string = "abc"
pattern = "a"
str1, str2 = "", ""

#返回 str 在 string 里面出现的次数，如果 beg 或者 end 指定则返回指定范围内 str 出现的次数
string.count(pattern, beg=0, end=len(string))


#检测 str 是否包含在 string 中，如果 beg 和 end 指定范围，则检查是否包含在指定范围内，如果是返回开始的索引值，否则返回-1
string.find(pattern, beg=0, end=len(string))
#同上，不存在则报异常
string.index(pattern, beg=0, end=len(string))

#把 string 中的 str1 替换成 str2,如果 num 指定，则替换不超过 num 次.
string.replace(str1, str2,  num=string.count(str1))

#检查字符串是否以 obj 结束，如果beg 或者 end 指定则检查指定的范围内是否以 obj 结束，如果是，返回 True,否则返回 False.
string.endswith(pattern, beg=0, end=len(string))

#截掉 string 左边的空格
string.lstrip()
#截掉 string 右边的空格
string.rstrip()

#以 str 为分隔符切片 string，如果 num 有指定值，则仅分隔 num+1 个子字符串
string.split(pattern="", num=string.count(pattern))

#翻转string中的大小写
string.swapcase()

