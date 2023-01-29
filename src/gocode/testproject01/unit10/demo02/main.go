package main

import "fmt"

// 定义 老师的结构体

type Teacher struct {
	// 变量名字大写，外界可以访问
	Name   string
	Age    int
	School string
}

// func main() {
// 	// 创建老师结构体的实例、对象、变量
// 	var ma Teacher  //
// 	fmt.Println(ma) // { 0 }
// 	ma.Name = "马"
// 	ma.Age = 33
// 	ma.School = "school"
// 	fmt.Println(ma)
// }

// func main() {
// 	var t Teacher = Teacher{}
// 	fmt.Println(t)
// 	t.Age = 31
// 	t.Name = "ss"
// 	t.School = "aa"
// 	fmt.Println(t)
// }
// func main() {
// 	var t *Teacher = new(Teacher)
// 	// t 是指针 ，指向的就是地址，应该给地址指向的字段进行复制
// 	(*t).Name = "bb"
// 	(*t).Age = 43
// 	// 为了符合程序员的编程习惯，go提供了简化的赋值方式
// 	t.School = "school"
// 	fmt.Println(*t)
// }
func main() {
	var t *Teacher = &Teacher{}
	t.Age = 33
	t.Name = "cc"
	fmt.Println(t)
}
