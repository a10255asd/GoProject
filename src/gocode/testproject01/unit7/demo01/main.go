package main

import "fmt"

func main() {
	// 给出 5 个学生的成绩，求出成绩的总和，平均数
	score1 := 95
	score2 := 85
	score3 := 75
	score4 := 65
	score5 := 55

	// sum
	sum := score1 + score2 + score3 + score4 + score5
	// avg
	avg := sum / 5
	fmt.Printf("成绩的总和为：%v，成绩的平均数为：%v", sum, avg)
}
