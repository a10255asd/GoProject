package main

import "fmt"

func main() {
	// 定义切片 ：三个参数，切片对应类型，切片长度，切片的容量
	slice := make([]int, 4, 20)
	fmt.Println(slice) //[0 0 0 0]
	fmt.Println("切片的长度", len(slice))
	fmt.Println("切片的容量", cap(slice))
	slice[0] = 66
	slice[1] = 88
	fmt.Println(slice) //[66 88 0 0]

	slice2 := []int{1, 4, 7}
	fmt.Println(slice2) //
	fmt.Println("切片的长度", len(slice2))
	fmt.Println("切片的容量", cap(slice2))
}
