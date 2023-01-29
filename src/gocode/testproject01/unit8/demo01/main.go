package main

import "fmt"

func main() {
	// 定义数组
	var intarr [6]int = [6]int{3, 6, 9, 1, 4, 7}
	// 切片构建在数组之上：
	// 定义一个切片 名字为slice ,[]动态变化的数组长度不写，int类型
	// 切出的一段片段- [1:3] 从 1 开始 到 3 结束（不包含3）
	var slice []int = intarr[1:3]
	fmt.Println("intarr", intarr)
	fmt.Println("slice", slice)
	fmt.Println("slice的元素个数：", len(slice))
	fmt.Println("slice的容量：", cap(slice))
	fmt.Println("数组中下标为1位置的地址：%p", &intarr[1])
	fmt.Println("切片中下标为0位置的地址：%p", &slice[0])
	slice[0] = 16
	fmt.Println("intarr", intarr)
	fmt.Println("slice", slice)
}
