package main

import "fmt"

func main() {
	var num1 int = 10
	fmt.Println(num1)
	// 等号右侧的值运算清楚后，在赋值给=左侧
	var num2 int = (10+20)%3 + 3 - 7
	fmt.Println(num2)

	var num3 int = 10
	num3 += 20
	fmt.Println(num3)

	//交换两个数的值 并输出结果
	var a int = 8
	var b int = 4
	fmt.Printf("a = %v, b = %v", a, b)
	// 交换 引入一个中间变量
	var t int
	t = a
	a = b
	b = t
	fmt.Printf("a = %v, b = %v", a, b)
}
