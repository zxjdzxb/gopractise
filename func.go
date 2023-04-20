package main

import "fmt"

func test() (int, int) {

	return 1, 2
}
func main() {
	// a, b := test()
	//匿名变量
	a, _ := test()
	fmt.Println(a)
}
