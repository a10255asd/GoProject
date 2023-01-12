package main

import "fmt"

func main() {
	var age int = 18
	// & 去出对应的地址
	fmt.Println("age对应的存储空间的地址为：", &age)
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值是：", *ptr)
}
