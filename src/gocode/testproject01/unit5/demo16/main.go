package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%v````````对应的类型为 %T", now, now)
	// 2023-01-24 08:26:33.955499 +0800 CST m=+0.000164793
	// Now 函数的返回值是一个结构体 对应的类型是 time.Time

	// 调用结构体中的方法：
	fmt.Printf(" 年： %v \n", now.Year())
	fmt.Printf(" 月： %v \n", now.Month())
	fmt.Printf(" 日： %v \n", now.Day())
	fmt.Printf(" 时： %v \n", now.Hour())
	fmt.Printf(" 分： %v \n", now.Minute())
	fmt.Printf(" 秒： %v \n", now.Second())

	// 将日期以年月日时分秒按照格式
	fmt.Println("------------------------------")
	fmt.Printf("当前年月日是： %d-%d-%d 时分秒：%d-%d-%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

}
