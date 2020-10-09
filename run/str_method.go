package main

import "fmt"

// Str is func
func Str() {
	//name := "a网络"
	//fmt.Println(name)
	//for i:=0; i<len(name); i++ {
	//	fmt.Println(i, name[i])
	//}
	//for index, item := range name {
	//	fmt.Println(index, item, string(item))
	//}

	//dataList := []rune(name)
	//fmt.Println(dataList[0], string(dataList[0]))
	//fmt.Println(dataList[1], string(dataList[1]))

	var num [3]int
	num[0] = 1
	num[1] = 2
	num[2] = 3

	var name = [2]string{"12","moke"}
	var ages = [3]int{0:0,1:1,2:2}

	var na1 = [...]string{"1","2"}

	fmt.Println(num ,name, ages, na1)




}

