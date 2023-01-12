package main

import "fmt"

func main() {
	// 功能 10 + 20  两数之和
	var num1 int = 10
	var num2 int = 20
	var sum int = 0
	sum += num1
	sum += num2
	fmt.Println(sum)
	// 调用函数
	sum2 := cal(10, 23)
	fmt.Println(sum2)
}

func cal(num1 int, num2 int) int {
	// 如果返回值类型就一个 （）可以不写
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}
