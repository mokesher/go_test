package Arry

import "fmt"

func Arry_method() {
	//数组长度固定
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	fmt.Println(arr1)

	arr2 := [3]string{"1","2","3"}
	fmt.Println(arr2)

	arr3 := [...]string{"php","st","on","12"}
	fmt.Println(arr3)

}