package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度，按字节进行统计
	// 在golang中，汉字是utf-8字符集，一个汉字3个字节，空格是一个字节。
	str := "golang 你好"
	fmt.Println(len(str))

	// 对字符串进行遍历：
	// 方式一：利用键值循环： for-range
	for i, value := range str {
		fmt.Printf("索引为：%d，具体值为： %c \n", i, value)
	}
	// 方式二：利用切片
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("%c \n", r[i])
	}

	//字符串转整数 func Atoi
	num1, err := strconv.Atoi("666")
	fmt.Println(num1, err)

	//整数转字符串 func Itoa
	str1 := strconv.Itoa(88)
	fmt.Println(str1)

	// 统计一个字符串有几个指定的子串：
	count := strings.Count("golangandjava", "ga")
	fmt.Println(count)

	// 不区分大小写的字符串比较
	flag := strings.EqualFold("hello", "HELLO")
	fmt.Println(flag)

	// 区分大小写进行字符串比较
	fmt.Println("hello" == "HELLO")

	// 返回子串在字符串中第一次出现的索引数
	fmt.Println(strings.Index("golangandjava", "ga"))
}
