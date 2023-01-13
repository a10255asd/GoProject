package main

import "fmt"

// 自定义函数： 功能交换两个输
func swap(num1 int, num2 int) (int, int) {
	var t int
	t = num1
	num1 = num2
	num2 = t
	return num1, num2
}

func main() {
	// 嗲用函数： 交换10 和20
	var num1 int = 10
	var num2 int = 20
	fmt.Printf("交换前的两个数 ： num1 = %v,num2 = %v \n", num1, num2)
	num1, num2 = swap(num1, num2)
	fmt.Printf("交换后的两个数 ： num1 = %v,num2 = %v \n", num1, num2)
}
