from bisect import bisect_left
from queue import Queue   # 队列，FIFO
from queue import PriorityQueue  #优先级队列，优先级高的先输出
from collections import defaultdict
from sortedcontainers import SortedList     #有序列表
from sortedcontainers import SortedDict     #有序字典
from sortedcontainers import SortedSet      #有序集合


#队列，FIFO
q = Queue(100) 	#构建一个队列对象，maxsize初始化默认为零，此时队列无穷大
q.empty()		#判断队列是否为空(取数据之前要判断一下)
q.full()		#判断队列是否满了
#q.put(item)		#向队列存放数据
q.get()			#从队列取数据,并且删除
q.qsize()		#获取队列大小，可用于判断是否满/空

#优先队列
q = PriorityQueue()
q.put((2, "Lisa"))

'''
优先队列继承队列
队列的变体，按优先级顺序（最低优先）检索打开的条目。
条目通常是以下格式的元组：
插入格式：q.put((priority number, data))
特点：priority number 越小，优先级越高
其他的操作和队列相同
'''

import heapq        #封装了操作堆的函数（默认小根堆，大根堆将数字取负数）

array = [1,3,5,2,4]
heap = []

for num in array:
    heapq.heappush(heap,num)    #将元素(num)插入堆(heap)中

heapq.heapify(array)            #在线性时间内将列表转换成堆。

x = heapq.heappop(heap)         #将堆顶元素删除，如果堆为空，则会引发IndexError。返回删除的值。
x = heapq.heappushpop(heap,6)   #将item添加到heap中去，然后将堆顶的元素删除，并返回这个被删除的值。

heapq.heapify(array)            #在线性时间内将列表转换成堆。

x = heapq.heapreplace(heap,7)   #删除掉堆顶的元素，然后将新元素添加到堆中去。如果堆是空的，则会引发IndexError。

#nlargest(n,iterable,key=None)
heapq.nlargest(3,heap)          #返回数据集中最大的n个元素组成的列表。
heapq.nsmallest(3,heap)         #返回数据集中最大的n个元素组成的列表。





sl = SortedList([0,1,2,3])
sl.add(8)       #添加单个元素
sl.update([5, 1, 3, 4, 2])      #添加多个元素
sl.discard(1)
sl.remove(2)         #按照值删除,不存在会报错
sl.pop(0)        #按照索引删除，无参数默认最右边
sl.clear()      #删除所有数据
sl.count(1)     #计数

#sl.index(2)     #返回下标
#del sl[0]
#支持 in []运算和二分查找的bisect_
iter(sl)   #迭代
reversed(sl)    #翻转
#五个方法没有实现：直接赋值、reverse()、append()、extend()、insert()



sd = SortedDict()  #sd = SortedDict(neg, enumerate('abc', start=1))  # key-function使用neg函数
sd['e'] = 5
sd['b'] = 2     # SortedDict({'b': 2, 'e': 5})
sd.update({'d': 4, 'c': 3})     # SortedDict({'b': 2, 'c': 3, 'd': 4, 'e': 5})
sd.setdefault('a', 1)   # 1 如果键位于排序后的结果中，则返回其值。如果键不在排序的结果中，则插入带有默认值的键并返回默认
# SortedDict({'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5})
del sd['d']
sd.pop('c')  # 3
sd.popitem(index=-1)  # ('e', 5)
# SortedDict({'a': 1, 'b': 2})
sd['b']  # 2
'c' in sd  # False
sd.get('z') is None  # True
sd.peekitem(index=-1)  # ('b', 2)
sd.bisect_right('b')  # 2
sd.index('a')  # 0
list(sd.irange('a', 'z'))  # ['a', 'b']
# SortedDict({'a': 1, 'b': 2})
keys = sd.keys()
keys[0]  # 'a'
items = sd.items()
items[-1]  # ('b', 2)
values = sd.values()
values[:]  # [1, 2]

#同上
ss = SortedSet()  
#ss = SortedSet([1, 2, 3], key=neg)
# SortedSet([3, 2, 1], key=<built-in function neg>)



