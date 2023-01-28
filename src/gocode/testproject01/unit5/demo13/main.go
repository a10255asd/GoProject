package main

import "fmt"

// 函数的名字是 getSum 参数为空  返回值是一个函数，这个函数的参数是一个int类型的参数，返回值也是int类型
func add(num1 int, num2 int) int {
	// 在go语言中，程序遇到defer关键字，不会立即执行defer后的语句，而是将defer压入一个栈中，然后继续执行函数后面的语句
	defer fmt.Println("num1= ", num1)
	defer fmt.Println("num2= ", num2)
	// 栈的特点是先入后出
	// 在函数执行完毕后，从栈中取出语句开始执行
	var sum int = num1 + num2
	defer fmt.Println("sum= ", sum)
	return sum
}

func main() {
	add(30, 60)
}
