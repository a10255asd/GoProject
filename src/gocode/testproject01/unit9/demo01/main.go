package main

import "fmt"

func main() {
	// 方式一
	// 定义 map 变量
	var a map[int]string
	// 只声明map 内存是没有分配空间的
	// 必须通过make 函数初始化
	a = make(map[int]string, 10)
	a[200095452] = "张三"
	a[200095123] = "李四"
	a[202125452] = "王五"
	a[200095123] = "朱六"
	fmt.Println(a)
	// 方式二
	b := make(map[int]string)
	b[200095452] = "张三"
	b[200095123] = "李四"
	fmt.Println(b)
	// 方式三
	c := map[int]string{200095452: "张三", 1: "二二"}
	fmt.Println(c)
	// 查找
	value, flag := b[200095123]
	fmt.Println(value, flag)
}
