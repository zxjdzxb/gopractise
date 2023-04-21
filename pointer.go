package main

import "fmt"

const MAX int = 3

func main() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */
	var ptr *int

	ip = &a /* 指针变量的存储地址 */

	if ptr != nil {
		fmt.Println("ptr 不是空指针")
	} /* ptr 不是空指针 */
	if ptr == nil {
		fmt.Println("ptr 是空指针")
	} /* ptr 是空指针 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	fmt.Println("======指针数组======")

	a2 := []int{10, 100, 200}
	var i int
	var ptr2 [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr2[i] = &a2[i] /* 整数地址赋值给指针数组 */
	}

	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr2[i])
	}
	fmt.Println("======指向指针的指针======")
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Println("======指针作为函数参数======")

	/* 定义局部变量 */
	var a3 int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a3)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	/* 调用函数用于交换值
	 * &a 指向 a 变量的地址
	 * &b 指向 b 变量的地址
	 */
	swap(&a3, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a3)
	fmt.Printf("交换后 b 的值 : %d\n", b)

}

func swap(x *int, y *int) {
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
}
