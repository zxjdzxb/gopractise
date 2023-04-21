package main

import "fmt"

func main() {
	// var n [10]int /* n 是一个长度为 10 的数组 */
	// var i, j int

	// fmt.Println(n)
	// /* 为数组 n 初始化元素 */
	// for i = 0; i < 10; i++ {
	// 	n[i] = i + 100 /* 设置元素为 i + 100 */
	// }

	// /* 输出每个数组元素的值 */
	// for j = 0; j < 10; j++ {
	// 	fmt.Printf("Element[%d] = %d\n", j, n[j])
	// }
	/* [0 0 0 0 0 0 0 0 0 0]
	Element[0] = 100
	Element[1] = 101
	Element[2] = 102
	Element[3] = 103
	Element[4] = 104
	Element[5] = 105
	Element[6] = 106
	Element[7] = 107
	Element[8] = 108
	Element[9] = 109 */
	var i, j, k int
	// 声明数组的同时快速初始化数组
	balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50}

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		fmt.Printf("balance[%d] = %f\n", i, balance[i])
	}
	fmt.Println("======")
	balance2 := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0, 100.0, 200.0}
	/* 输出每个数组元素的值 */
	for j = 0; j < 7; j++ {
		fmt.Printf("balance2[%d] = %f\n", j, balance2[j])
	}
	fmt.Println("======")

	//  将索引为 1 和 3 的元素初始化
	balance3 := [5]float32{1: 2.0, 3: 7.0}
	for k = 0; k < 5; k++ {
		fmt.Printf("balance3[%d] = %f\n", k, balance3[k])
	}
	fmt.Println("多维数组======")
	// Step 1: 创建数组
	values := [][]int{}

	// Step 2: 使用 append() 函数向空的二维数组添加两行一维数组
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)

	// Step 3: 显示两行数据
	fmt.Println("Row 1")
	fmt.Println(values[0])
	fmt.Println("Row 2")
	fmt.Println(values[1])

	// Step 4: 访问第一个元素
	fmt.Println("第一个元素为：")
	fmt.Println(values[0][0])
	fmt.Println("======")

	// 创建二维数组
	sites := [2][2]string{}

	// 向二维数组添加元素
	sites[0][0] = "Google"
	sites[0][1] = "Runoob"
	sites[1][0] = "Taobao"
	sites[1][1] = "Weibo"

	// 显示结果
	fmt.Println(sites)
	fmt.Println("======")

	/* 数组 - 5 行 2 列*/
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	/* 输出数组元素 */
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
	fmt.Println("向量传递数组======")

	//
	var avg float32
	var balance4 = [5]int{1000, 2, 3, 17, 50}

	/* 数组作为参数传递给函数 */
	avg = getAverage(balance4, 5)

	/* 输出返回的平均值 */
	fmt.Printf("平均值为: %f\n ", avg)
	fmt.Println("======")
	//转整型来设置精度
	a2 := 1690                         // 表示1.69
	b2 := 1700                         // 表示1.70
	c2 := a2 * b2                      // 结果应该是2873000表示 2.873
	fmt.Println(c2)                    // 内部编码
	fmt.Println(float64(c2) / 1000000) // 显示

}

func getAverage(arr [5]int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg
}
