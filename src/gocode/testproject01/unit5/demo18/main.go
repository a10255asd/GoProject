package main

import "fmt"

func main() {
	// new 分配内存，new函数的实参是一个类型，而不是具体数值，new函数返回值是对应类型的指针 num：*int
	num := new(int)
	fmt.Printf("num 的类型是 ：%T，num的值是：%v，num的地址是：%v,num指针指向的值是：%v", num, num, &num, *num)
}
