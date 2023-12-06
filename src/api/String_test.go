package api

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	fmt.Println("---")

	// 开辟新内存复制字符串
	strings.Clone("hello"[1:]) // ello

	// 比较字符串 a b 大小 a < b -1; a == b 0; a > b 1
	cmp := strings.Compare("abce", "abcd") // 1
	fmt.Println("abce  abcd:", cmp)

	// s 中是否包含 substr
	strings.Contains("abc", "bc") // true

	// s 中是否存在 substr 中的字符
	strings.ContainsAny("abcjxkj", "zzzzza") // true

	// s 中是否存在 c 字符
	strings.ContainsRune("abc", 'd') // false

	// s 中 substr 出现多少次(不重叠)
	ccnt := strings.Count("aaa", "aa") // 1
	fmt.Println("s 中 substr 出现多少次：", ccnt)

	// s 中的 sep 去除（1次），返回前面、后面字符串，以及是否去除了
	strings.Cut("aabbbbcc", "bb") // aa bbcc true

	// 判断2字符串是否相等
	strings.EqualFold("abc", "abc") // true

	// substr 在 s 中第一次出现的下标， 否则-1
	strings.Index("kabcabc", "abc") // 1

	// substr中的字符，第一次出现在 s 的下标，不在 -1
	strings.IndexAny("abcdef", "ca") // 0

	// 字符出现在 s 中的下标，不在 -1
	strings.IndexByte("abc", 'b')
	strings.IndexRune("abc", 'b')

	// 将字符串数组中的字符串 用 sep 字符串连接
	str_join := strings.Join([]string{"aa", "bb", "cc"}, "/") // aa/bb/cc
	fmt.Println("Join:", str_join)

	// 将 s 的前 n 个字符中的全部 old，替换成 new (不重叠)
	str_rep := strings.Replace("aabbbbbcc", "bb", "kk", 10) // aakkkkbcc
	fmt.Println(str_rep)
	// 不设前n个字符的限制
	str_repa := strings.ReplaceAll("aabbbbbcc", "bb", "kk") // aakkkkbcc
	fmt.Println(str_repa)

	// 切分，返回数组
	strings.Split("a/b/c", "/")

	// Tirm(s, substr)	去除左右两边的substr
	// TirmLeft(s, substr) TiemRight(s, substr)
	fmt.Println(strings.Trim("a  b aaa", "a"))

}
