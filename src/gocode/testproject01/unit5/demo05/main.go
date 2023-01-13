package main

import "fmt"

// 定义int类型的可变参数，传入0ge 。。。 n个可变数据
func test(args ...int) {
	// 在函数内部处理可变参数的时候，将可变参数当作切片来处理
	// 遍历可变参数
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}

func main() {
	test()
	fmt.Println("-------")
	test(1)
	test(1, 2, 3, 4, 5)
}
