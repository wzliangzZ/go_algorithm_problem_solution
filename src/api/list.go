package api

import (
	"container/list"
	"fmt"
)

/*
创建	New()
清空	Init()

取值	Back()		Front()		Len()
插入	PushBack()	PushFront()
删除	Remove(e *Element)
*/

func DoubleList() {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// 反向迭代
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
	// 正向迭代
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
