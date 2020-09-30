package main

import "fmt"

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func containSubstr(s, subStr string) bool {
	/*
		利用已经编写好的判断是否含有前缀的函数,让原字符串从第一个字节开始向后切片,
		每次都用新的原串进行前缀判断,只要O(n)次判断即可得到是否含有子串
	*/
	for i := 0; i < len(s); i++ {
		if hasPrefix(s[i:], subStr) {
			return true
		}
	}
	return false
}

/*
	GO中,for range循环会隐式解码UTF8编码的字符串,因此利用该特性,我们可以快速遍历字符串来统计字符串中rune字符的个数
	在for range 中,由于隐式的解码,所以字符的索引并不一定是连续递增的,中文字符的索引会跳跃3个
*/
func fastCountStr(s string) int {
	n := 0
	for i := range s {
		n++
		fmt.Printf("i:%v\n", i)
	}
	return n
}

func main() {
	/*
		string 是一个值类型的基本数据类型,打印整个字符串时,会打印字面量,但是循环打印字符串内部的值时,会输出字符对应的
		ASCII码值.或者其他整数
		string的底层是字节数组,而byte又是uint8类型的等价类型.所以字符串变量内的单个字符输出时都会按照其ASCII对应的
		字符输出.
		汉字在go中都是以int32类型存储的.
		len()函数会返回字符串中字节的数目而不是字符串中rune字符的数目
	*/
	var s1 string = "123世界"
	fmt.Println(len(s1)) // 9
	// 根据打印的字符串的长度可以知道,len()函数打印的是字符串中每个字符字节数之和
	fmt.Println(hasPrefix(s1, "123"))
	fmt.Println(hasSuffix(s1, "界"))
	fmt.Println(containSubstr(s1, "3世"))
	fmt.Println(fastCountStr(s1))
}
