package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误结果是：", err)
		panic(err)
	}
	fmt.Println("上面的除法执行成功")
	fmt.Println("正常执行下面的逻辑")
}

func test() (err error) {
	// 利用 defer 和 recover 来捕获错误，defer后加上匿名函数的调用  “()”
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出自定义错误；
		return errors.New("除数不能为0哦~~~~~")
	} else {
		// 如果除数不为 0 ，那么正常执行就可以了
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
