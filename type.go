package main

import "fmt"

// 结构体
type Person struct {
	name string
	age  int
}

func main() {
	// bool 默认值是false
	var isFlag bool = true
	var isFlag2 bool
	fmt.Println(isFlag)
	fmt.Printf("%T,%t\n", isFlag, isFlag)
	fmt.Println(isFlag2)
	fmt.Printf("%T,%t\n", isFlag2, isFlag2)
	// number
	//int uint(无符号)
	var i int = 100
	fmt.Printf("%T,%d\n", i, i)
	//float 32 64  complex(复数)
	var f float32 = 3.14
	fmt.Println(i, f)
	fmt.Printf("%T,%.2f\n", f, f)
	//string
	s := "Hello, World!"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	fmt.Println(s[:5])      // "hello"
	fmt.Println(s[7:])      // "world"
	fmt.Println(s[:])       // "hello, world"
	fmt.Printf("%T,%s\n", s, s)
	fmt.Println(s[:5] + s[7:])
	// fmt.Printf("%T,%c\n", s[0], s[0])

	u := Person{"Tom", 18}
	fmt.Println(u)

	//引用类型(指针类型)

}
