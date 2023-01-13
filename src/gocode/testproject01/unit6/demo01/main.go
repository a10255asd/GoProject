package main

import "fmt"

func test() {
	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}

func main() {
	test()
	fmt.Println("上面的除法执行成功")
	fmt.Println("正常执行下面的逻辑")
}
