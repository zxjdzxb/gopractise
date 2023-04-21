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

	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // 如果发现因子，则不是素数
			}
		}
		if j > (i / j) {
			fmt.Printf("%d  是素数\n", i)
		}
	}
	// for range
	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。
	// 格式如下：
	// for key, value := range oldMap {
	//     newMap[key] = value
	//}
}
