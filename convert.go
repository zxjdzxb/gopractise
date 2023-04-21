package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)

	// 初始化一个32位整型值
	var a int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", a, a)

	// 将a变量数值转换为十六进制, 发生数值截断
	b := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", b, b)
	fmt.Println("=======数值类型转换=========")
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)
	fmt.Println("=======字符串转换=========")
	str := "123"
	//strconv.Atoi 函数返回两个值，第一个是转换后的整型值，第二个是可能发生的错误 , _ 来忽略错误
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("转换错误:", err)
	} else {
		fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
	}
	num2 := 123
	str2 := strconv.Itoa(num)
	fmt.Printf("整数 %d  转换为字符串为：'%s'\n", num2, str2)

	str3 := "3.14"
	num3, err3 := strconv.ParseFloat(str3, 64)
	if err3 != nil {
		fmt.Println("转换错误:", err3)
	} else {
		fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str3, num3)
	}

	num4 := 3.14
	str4 := strconv.FormatFloat(num4, 'f', 2, 64)
	fmt.Printf("浮点数 %f 转为字符串为：'%s'\n", num4, str4)
	fmt.Println("=======接口类型转换=========")
	var i interface{} = "Hello, World"
	str, ok := i.(string)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}
