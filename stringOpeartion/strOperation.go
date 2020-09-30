package main

import (
	"fmt"
	"strings"
)

// go语言使用strings包来完成字符串的主要操作

func main() {
	str := "hello,world"
	// 判断字符串是否含有特定前缀
	fmt.Println(strings.HasPrefix(str, "hel"))
	// 判断字符串是否含有特定后缀
	fmt.Println(strings.HasSuffix(str, "ld"))

	// 字符串包含关系 Contains 函数
	fmt.Println(strings.Contains(str, "lo,"))

	// 判断子串或者单个字符在原字符串中出现的位置
	// Index函数会返回子串第一个字符在原字符串中出现的位置
	fmt.Println(strings.Index(str, "el"))

	// LastIndex()函数返回子串的第一个字符在原字符串中最后一次出现的位置.
	// 如果子串首先不在原字符串中,则返回 -1
	fmt.Println(strings.LastIndex(str, "ld"))

	// 当字符串为非 ASCII 码字符,则应使用 strings.IndexRune判断
	s := "hello,世界"
	fmt.Println(strings.IndexRune(s, '界'))

	// 字符串替换
	// Replace函数将原串中的前n 个 old 子串替换为 new 子串, n 为人工设定具体需要替换几个
	// n=0 时不替换,n<0 时替换所有.n>0 时按具体数字进行替换
	fmt.Println(strings.Replace(s, "世界", "new world", 1))

	// 统计子串在原串中出现的次数
	fmt.Println(strings.Count("HHYYQQyxyxuHioH", "HH"))

	//重复参数中的字符串 s n 次, 包括字符中所含的空格也一并会被重复
	fmt.Println(strings.Repeat("Yes ", 3))

	// 修改字符串大小写
	fmt.Println(strings.ToLower("HAPPY BIRTHDAY TO YOU"))
	fmt.Println(strings.ToUpper("i am legend"))
	// 一般字符的 title 格式就是其大写格式
	fmt.Println(strings.ToTitle("fighting for free"))

	// 剔除字符串开头和结尾的空格
	fmt.Println(strings.TrimSpace(" Hello, I'm Chen "))
	// 剔除字符串开头和结尾的指定字符串
	// 如果开头或者结尾没有则只删除开头或者结尾处的指定字符串,如果有重叠则同样删除
	fmt.Println(strings.Trim("HHHello", "HH")) // 此处会删除三个 H,只剩下 ello
	// 如果仅想删除开头或者结尾的字符串,则需使用 TrimLeft or TrimRight
	fmt.Println(strings.TrimLeft("Hello, Bingo", "He"))
	fmt.Println(strings.TrimRight("Hello,Bingo", "go"))

	// 切割字符串函数 Fields 和 Split函数都会返回一个切片
	// Fileds()会使用空格符去动态分割字符串,返回一个 slice. 如果字符串只有空格符,则返回一个空 slice
	// Splice 需要指定一个分隔符去分割,如果分隔符不在字符串中,则将原字符串当做返回的字符串切片中的第一个元素
	// 如果分割符存在于字符串中,则根据分割符将字符串分成 n 份.并将每一个分割后的子串当做[]string 中的元素.
	// 如果分割符是空串,则 split 函数会将每一个字符分隔开作为[]string 中的元素,包括标点符号
	// 因为这两个函数都会返回 slice,所以一般用 for-range 函数去进行处理.
	sli := strings.Split("我爱你,中国", "")
	fmt.Printf("%#v\n", sli)
	fmt.Println(len(sli))
	sli2 := strings.Fields("你 好 啊 我 的 宝 贝")
	fmt.Printf("%#v\n", sli2)
	fmt.Println(len(sli2))

	// 字符串拼接
	// Join() 函数用于将元素类型为 string 的切片使用分割符来组成一个字符串
	var sli3 = []string{"h", "e", "l", "l", "o"}
	fmt.Println(strings.Join(sli3, ""))

	// 从字符串中读取内容:
	// strings.NewReader(str)函数会生成一个用于读取 str 的 Reader 对象(结构体实例),
	// 该实例对象拥有 Len(), Size() 和 Read()方法.
	// Len()方法用于返回剩余未读取的字符数, Size()返回 str 的字符数.
	// Read([]byte)方法用来读取字符信息到参数指定的byte 切片中. 该方法返回具体读到的数据长度
	reader := strings.NewReader(s)
	buf := make([]byte, 3)           // 创建 3 个字节大小的缓冲区
	fmt.Println(reader.Len())        //未读取字符时,reader中剩余字节长度
	fmt.Println(reader.Size())       // 字符串字节总长度
	readLen, err := reader.Read(buf) // 用 reader 向 buf 中读取数据
	if err != nil {
		fmt.Println("错误:", err)
	}
	fmt.Println(readLen)       // 读取数据后,返回读到的长度
	fmt.Println(reader.Len())  // 读取数据后,剩余字符串的字节长度
	fmt.Println(reader.Size()) // 字符串的字节总长度
}
