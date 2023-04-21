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

// 闭包：闭包是一个函数值，它来自函数体的外部的变量引用。该函数可以访问并赋予其引用的变量的值，换句话说该函数被这些变量“绑定”在一起。
// 例如，函数adder返回一个闭包。每个闭包都被绑定在其各自的sum变量上。
// 闭包的好处是可以访问一个函数内部的变量，即使该函数已经执行完毕。·1
