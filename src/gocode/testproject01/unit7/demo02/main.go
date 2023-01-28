package main

import "fmt"

func main() {
	// 给出 5 个学生的成绩，求出成绩的总和，平均数
	// 定义一个数组
	var scores [5]int
	// 将成绩存入数组

	scores[0] = 95
	scores[1] = 85
	scores[2] = 75
	scores[3] = 65
	scores[4] = 55

	// sum
	sum := 0
	for i := 0; i < len(scores); i++ {
		// 定义一个变量专门接受成绩的和
		sum += scores[i]
	}
	// avg
	avg := sum / len(scores)
	fmt.Printf("成绩的总和为：%v，成绩的平均数为：%v", sum, avg)
}
