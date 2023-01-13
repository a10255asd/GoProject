package main

import "fmt"

func main() {
	// 实现功能：键盘录入学生的年龄，姓名，成绩，是否是VIP
	// 方式1 ：Scanln
	var age int
	fmt.Println("请输入学生的年龄：")
	// 传入age的地址的目的：在Scanln函数中，对地址中的值进行改变的时候，实际外面的age被影响了。
	// 录入数据的时候，类型一定要匹配，因为底层会自动判断类型。
	fmt.Scanln(&age)

	var name string
	fmt.Println("请输入学生的姓名：")
	fmt.Scanln(&name)

	var score float32
	fmt.Println("请输入学生的成绩：")
	fmt.Scanln(&score)

	var isVip bool
	fmt.Println("请输入学生是否是VIP：")
	fmt.Scanln(&isVip)

	// 将上述数据在控制台打印输出
	fmt.Printf("学生的年龄为：%v 姓名为：%v成绩为：%v是否为VIP：%v", age, name, score, isVip)

	// 方式2 ：Scanf
	fmt.Println("请录入学生的姓名，年龄，成绩，是否为VIP，使用空格进行分隔")
	fmt.Scanf("%s %d %f %t", &age, &name, &score, &isVip)

}
