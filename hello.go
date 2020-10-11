package main

import "fmt"

func main() {

	// for {
	// 	fmt.Print("请输入用户名：")
	// 	var name string
	// 	var age int
	// 	_, err := fmt.Scanln(&name, &age)

	// 	if err == nil {
	// 		result := fmt.Sprintf("用户名：%s 年龄：%d", name, age)
	// 		fmt.Println(result)
	// 	} else {
	// 		fmt.Println("用户输入错误：", err)
	// 	}
	// 	println("end")

	// }
	// name := "网络"
	// fmt.Println(name[0])
	// fmt.Println(name[1])
	// fmt.Println(name[2])
	// fmt.Println(name[3])
	// fmt.Println(name[4])

	// var v1 int8 = 10
	// var v2 int16 = 20
	// v3 := int16(v1) + v2
	// fmt.Println(v3)

	// var v4 int16 = 129
	// v5 := int8(v4)

	// fmt.Println(v5)

	//v6 := 10
	//result1 := strconv.Itoa(v6)
	//fmt.Println(result1, reflect.TypeOf(result1))
	//
	//v7 := "111"
	//result2, err := strconv.Atoi(v7)
	//if err == nil {
	//	fmt.Println("sucdess", result2, reflect.TypeOf(result2))
	//} else {
	//	fmt.Println("error", result2, reflect.TypeOf(result2), err)
	//}
	//name1 := [2]string{"111", "222"}
	//name2 := name1
	//// name2不会发生改变，变量赋值时重新拷贝了一份
	//name1[1] = "moke"
	//fmt.Printf("name1: %v , name2: %v", name1, name2)

	//name3 := [4]string{"111", "222", "333", "444"}
	////长度
	//fmt.Println(len(name3))
	////索引
	//data := name3[1]
	//fmt.Println(data)
	////切片
	//fmt.Println(name3[1:3])

	// slice
	//var nums []int
	//// make只用于 slice、dict、channel
	//var users = make([]int,1,2)
	//var data = []int{11,22,33}
	//fmt.Println(nums, users, data)
	////切片的指针
	//var v1 = new([]int)
	//// 指针类型
	//var v2 *[]int

	//自动扩容
	//v1 := make([]int,1,3)
	//fmt.Println(len(v1), cap(v1))	//1 3

	data := make([]int,3)	// [0 0 0]
	data2 := append(data, 111)	//append内存地址相同 [0 0 0 111]
	data[0] = 999
	fmt.Println(data, data2)









}
