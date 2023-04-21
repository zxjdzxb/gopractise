package main

import "fmt"

func main() {
	a, b := 0, 0

	fmt.Printf("====== 初始化 ====== \n")
	fmt.Printf("地址n a: %p, b: %p\n", &a, &b)
	fmt.Printf("值 a: %d, b: %d\n", a, b) // 0 0

	Add(a)

	AddPtr(&b)

	fmt.Printf("\n ======  final ====== \n")
	fmt.Printf("地址n a: %p, b: %p\n", &a, &b)
	fmt.Printf("值 a: %d, b: %d\n", a, b) // 0 1
}

// 通过值传递
func Add(x int) {
	fmt.Printf("\n======   'Add' 函数 ====== \n")
	fmt.Printf("Before Add, 地址n: %p, 值: %d\n", &x, x)
	x++
	fmt.Printf("After Add, 地址n: %p, 值: %d\n", &x, x)
}

// 通过指针传递
func AddPtr(x *int) {
	fmt.Printf("\n ======  'AddPtr' 函数 ====== \n")
	fmt.Printf("Before AddPtr, 地址n: %p, 值: %d\n", x, *x)
	*x++ // We add * in front of the variable because it is a pointer, * will call 值 of a pointer
	fmt.Printf("After AddPtr, 地址n: %p, 值: %d\n", x, *x)
}
