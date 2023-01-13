package main

import "fmt"

func main() {
	// 实现功能：根据给出的学生分数，判断学生的等级
	var score int = 97
	// 根据分数判断对应的等级
	switch score / 10 {
	case 10:
		fmt.Println("A")
	case 9:
		fmt.Println("B")
		fallthrough
	case 8:
		fmt.Println("C")
	case 7:
		fmt.Println("D")
	case 6:
		fmt.Println("E")
	case 5:
		fmt.Println("F")
	case 4:
		fmt.Println("F")
	default:
		fmt.Println("您的程序有误")
	}
}
