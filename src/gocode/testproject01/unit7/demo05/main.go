package main

import "fmt"

func main() {

	// 四种定义数组的方式
	// 1
	var arr1 [3]int = [3]int{1, 2, 3}
	fmt.Println(arr1)
	// 2
	var arr2 = [3]int{1, 2, 3}
	fmt.Println(arr2)
	// 3
	var arr3 = [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)
	//4
	var arr4 = [...]int{2: 99, 1: 33, 3: 22}
	fmt.Println(arr4)
}
