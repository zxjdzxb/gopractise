package main

import "fmt"

//通道（channel）是用来传递数据的一个数据结构。
/*
ch <- v    // 把 v 发送到通道 ch
v := <-ch  // 从 ch 接收数据
           // 并把值赋给 v

//声明
ch := make(chan int)
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 把 sum 发送到通道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(c)
	fmt.Println(s[:len(s)/2])
	fmt.Println(s[len(s)/2:])
	go sum(s[:len(s)/2], c) //-9
	go sum(s[len(s)/2:], c) //17
	x, y := <-c, <-c        // 从通道 c 中接收
	fmt.Println()
	fmt.Println(x, y, x+y)
	fmt.Println("======通道缓冲区=========")
	ch := make(chan int, 2)

	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 2

	// 获取这两个数据
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
