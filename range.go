package main

import "fmt"

//ange 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素
/*
for key, value := range oldMap {
    newMap[key] = value
}
*/
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	//遍历简单的数组，2**%d 的结果为 2 对应的次方数
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println("======key value可省略======")

	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0

	// 读取 key 和 value
	for key, value := range map1 {
		fmt.Printf("key is: %d - value is: %f\n", key, value)
	}

	// 读取 key
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	// 读取 value
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
}
