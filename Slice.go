package main

import (
	"fmt"
)

// Slice 解决的就是对不定长数组的需求
func main() {
	// var identifier []type
	/*
		var slice1 []type = make([]type, len)
		也可以简写为
		slice1 := make([]type, len)
	*/
	var s1 []string           //[] 声明一个字符串类型切片
	var s2 = []int{1, 2, 3}   //[]  声明一个整数类型切片并初始化
	s3 := []bool{true, false} //[true false]
	// var s4 = []bool{true, false} //[true false]
	fmt.Println(s1 == nil) //true 未初始化等于nil
	fmt.Println(s2 == nil) //false 初始化未赋值也不等于nil
	fmt.Println(s3 == nil) //false
	// fmt.Println(s3 == s4) //切片是引用类型，不支持直接比较，只能和nil比较
	fmt.Println(s3) //[]string
	fmt.Println(s2) //[1 2 3]
	fmt.Println("======make======")
	//s :=make([]int,len,cap)
	//len()获取长度 和 cap() 函数获取切片的容量
	var numbers = make([]int, 3, 5)

	printSlice(numbers)
	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)
	fmt.Println("======append() 和 copy()======")

	/* append() 和 copy() 函数 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
