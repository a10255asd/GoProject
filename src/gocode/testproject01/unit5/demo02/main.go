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
	sum3, result := cal2(10, 10)
	fmt.Println(sum3, result)

	sum4, _ := cal2(10, 50)
	fmt.Println(sum4)

}

func cal(num1 int, num2 int) int {
	// 如果返回值类型就一个 （）可以不写
	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

func cal2(num1 int, num2 int) (int, int) {
	// 计算两个数的和 和 两个数的 差
	var sum int = 0
	var result int = num1 - num2
	sum += num1
	sum += num2
	return sum, result

}
