package main

import "fmt"

func main() {
	// var (
	// 	name   string
	// 	age    int
	// 	addess string
	// )
	// name = "Hello, World!"
	// age = 18
	// addess = "China"
	// var nkie string = "heu!"
	// := 自动推断类型
	name := "Hello, World!"
	age := 18
	addess := "China"
	fmt.Println(name, age, addess)
	fmt.Printf("age: %d,内存地址: %p", age, &age)
	//常量
	const pi = 3.1415926
	//iota常量计数器
	fmt.Println(pi)
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday = 100
		Friday   = iota
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
