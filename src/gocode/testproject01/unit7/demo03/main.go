package main

import "fmt"

func main() {
	var arr [3]int16
	// 获取数组的长度
	fmt.Println(len(arr)) // [0 0 0 ]
	fmt.Printf("arr 的地址为：%p \n", &arr)
	// 第一个空间的地址
	fmt.Printf("arr[0] 的地址为：%p \n", &arr[0])
	fmt.Printf("arr[1] 的地址为：%p \n", &arr[1])
	fmt.Printf("arr[2] 的地址为：%p \n", &arr[2])

	// 赋值
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
}
