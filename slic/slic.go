package slic

import "fmt"

func Slice_method() {
	fmt.Println("in the Slice")
	var arr1 []int
	fmt.Println(arr1)

	var arr2 = []int{1,2,3,4}
	fmt.Println(arr2)

	var arr3 = []int{1:1,2:2}
	fmt.Println(arr3)


}
