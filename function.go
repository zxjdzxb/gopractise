package main

import "fmt"

func main() {
	/* 定义局部变量 */
	var a int = 100
	var b int = 200
	var ret int

	/* 调用函数并返回最大值 */
	ret = max(a, b)
	sum(1, 2, 3, 4, 5)
	fmt.Printf("交换前 a 的值为 : %d\n", a)
	fmt.Printf("交换前 b 的值为 : %d\n", b)
	swap(a, b)

	fmt.Printf("a is %d\n", a)
	fmt.Printf("b is %d\n", b)

	fmt.Printf("最大值是 : %d\n", ret)
	fmt.Printf("sum is %d\n", sum(1, 2, 3, 4, 5))

	fmt.Printf("交换前，a 的值 : %d\n", a)
	fmt.Printf("交换前，b 的值 : %d\n", b)

	/* 调用 swap() 函数
	* &a 指向 a 指针，a 变量的地址
	* &b 指向 b 指针，b 变量的地址
	 */
	swap2(&a, &b)

	fmt.Printf("交换后，a 的值 : %d\n", a)
	fmt.Printf("交换后，b 的值 : %d\n", b)
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 定义局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

// 可变参数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// 值传递: 值传递是指在调用函数时将实际参数复制一份传递到函数中，对于函数内的操作不影响实际参数。

func swap(x, y int) int {
	var temp int

	temp = x /* 保存 x 的值 */
	x = y    /* 将 y 值赋给 x */
	y = temp /* 将 temp 值赋给 y*/

	return temp
}

// 引用传递
// 引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的操作就是对其本身进行的操作，而不是参数的拷贝。
func swap2(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}
