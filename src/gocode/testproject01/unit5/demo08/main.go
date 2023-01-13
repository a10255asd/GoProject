package main

import "fmt"

// 定义一个函数：
func test(num int) {
	fmt.Println(num)
}

// 定义一个函数，把另一个函数作为形参
func test02(num1 int, num2 float32, testFunc func(int)) {

}

type myFunc func(int)

func test03(num1 int, num2 float32, testFunc myFunc) {
	fmt.Println("--------test03")
}

func test04(num1 int, num2 int) (int, int) {
	// 定义一个函数 求两个数的和和差
	fmt.Println("--------test04")
	resul01 := num1 + num2
	resul02 := num1 - num2
	return resul01, resul02
}

func test05(num1 int, num2 int) (sum int, sub int) {
	// 定义一个函数 求两个数的和和差
	fmt.Println("--------test04")
	sum = num1 + num2
	sub = num1 - num2
	return
}

func main() {
	a := test
	//a 对应的类型是：func(int) ,test函数对应的类型是： func(int)
	fmt.Printf("a 对应的类型是：%T ,test函数对应的类型是： %T", a, test)
	// 通过该变量可以对函数调用
	a(10) //等价于 test（10）

	test02(10, 3.19, test)
	test02(10, 3.19, a)

	// 自定义数据类型： 给int类型起了别名，myInt类型
	type myInt int
	var num1 myInt = 30
	fmt.Println("num1", num1)

	// 虽然是别名，但是在go中变异识别的时候还是认为myInt和int不是一种数据类型
	var num2 int = 30
	num2 = int(num1)
	fmt.Println(num2)
}
