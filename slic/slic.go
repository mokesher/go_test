package slic

import (
	"fmt"
)

func Slice_method() {
	fmt.Println("in the Slice")
	// 切片的声明
	var arr1 []int
	fmt.Println(arr1)

	var arr2 = []int{1,2,3,4}
	fmt.Println(arr2)

	var arr3 = []int{1:1,2:2}
	fmt.Println(arr3)

	var nums []int
	// make只用于 slice、dict、channel
	var users = make([]int,1,2)
	var da = []int{11,22,33}
	fmt.Println(nums, users, da)
	//切片的指针
	var v1 = new([]int)
	fmt.Println(v1)
	// 指针类型
	var v2 *[]int
	fmt.Println(v2)

	//自动扩容
	v3 := make([]int,1,3)
	// 长度 容量
	fmt.Println(len(v3), cap(v3))	//1 3

	data := make([]int,3)	// [0 0 0]
	data2 := append(data, 111)	//append内存地址相同 [0 0 0 111]
	data[0] = 999
	fmt.Println(data, data2)

	s1 := []int{1,2,3,4}
	s2 :=s1
	sc := make([]int, 4,5)
	copy(sc, s1)
	sc[0] = 0

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(sc)




}
