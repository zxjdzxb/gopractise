package main

import "fmt"

// func main() {
// 	sum := 0
// 	for i := 0; i <= 10; i++ {
// 		sum += i
// 	}
// 	fmt.Println(sum)
// }
func main() {
	sum := 1
	// ; 可以省略，但是不能省略条件表达式的括号
	// 这样写也可以，更像 While 语句形式
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
	// 打印5*5方阵
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
	fmt.Println()
	// 打印九九乘法表
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}
