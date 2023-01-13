package main

import "fmt"

var num int = test()

func test() int {
	fmt.Println("test 函数被执行")
	return 10
}

func init() {
	fmt.Println("init 函数被执行了")
}

func main() {
	fmt.Println("main 函数被执行了")
}
