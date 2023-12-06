package api

import (
	"fmt"
	"sort"
	"testing"
	"slices"
)

// go1.21
func TestSlicesFunc(t *testing.T){
	var arr = []int{1,2,6,5,20}
	// func 中的值为数组里的元素
	slices.SortFunc[[]int](arr, func(a, b int) int{
		if a < b { return -1}
		if a > b { return 1 }
		return 0
	})
	// slices.SortStableFunc  相等元素保留相对位置
	fmt.Println(arr)

}

// 自定义排序:类型要实现 sort.Interface 接口的 Len、Less 和 Swap 方法
// 按字符串长度大小排序
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func TestCmp(*testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println("自定义按字符串长度排序：", fruits)

}
func TestSort(*testing.T) {
	fmt.Println("--")
	var arr = []int{3, 2, 1, 10, 10}
	sort.Ints(arr) // 原地排序
	// sort.Float64s([]float64{23.1, 10.2, -10.3})
	// sort.Strings([]string{"dhf", "djfl", "dkjfh"})

	//	自定义排序规则， 实现 less func(i int, j int) bool
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] <= arr[j] {
			return true
		} else {
			return false
		}
	})

	// 有序的arr 中返回第一个出现的目标元素下标，否则len(arr)
	fmt.Println(sort.SearchInts(arr, 10))
	// sort.SearchFloat64s()
	// sort.SearchStrings()

	b_arr := []int{7, 2, 3, 10, 1, 2, 5, 0}
	sort.Ints(b_arr)
	// 2 分查找，返回第一个满足 如下条件的下标，否则 len(r_arr)
	idx := sort.Search(len(b_arr), func(i int) bool {
		if b_arr[i] >= 6 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(idx, b_arr)

}
