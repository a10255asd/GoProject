package main

import "fmt"

func main() {
	// 实现功能 如果口罩的库存小于30个，提示：库存不足
	var count int = 20
	if count < 30 {
		fmt.Println("对不起，口罩的数量小于30")
	}
	// go语言里，if后面可以并列的加入变量的定义：
	if count1 := 20; count1 < 30 {
		fmt.Println("if后面可以并列的加入变量的定义")
	}
}
