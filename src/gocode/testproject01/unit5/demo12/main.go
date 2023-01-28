package main

import "fmt"

// 函数的名字是 getSum 参数为空  返回值是一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}
}

// 闭包： 返回的匿名函数 +  匿名函数以外的变量 num
func main() {
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(2))
}
