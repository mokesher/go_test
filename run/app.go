package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//Add()
	//api.Baidu()
	//var name string
	//var age int
	//name = "12"
	//for {
	//	fmt.Print("请输入用户名：")
	//	_,err := fmt.Scanln(&age)
	//	if err == nil {
	//		fmt.Printf("用户名：%s 年龄: %d\n",name, age)
	//		//fmt.Print(age)
	//	} else {
	//		fmt.Println("用户输入错误：", err)
	//	}
	//	//time.Sleep(time.Second * 2)
	//}

	// }
	//name := "网络"
	//fmt.Println(name[0], strconv.FormatInt(int64(name[0]), 2))
	//fmt.Println(name[1])
	//fmt.Println(name[2])
	//
	//fmt.Println(name[3])
	//fmt.Println(name[4])
	//fmt.Println(name[5])
	//fmt.Println(len(name))

	//n2 := "abc"
	//fmt.Println(len(n2))
	//字符串转换为字节集合
	//byteSet := []byte(n2)
	//fmt.Println(byteSet)
	//字节集合转换为字符串
	//byteList := []byte{97, 98, 99}
	//targetString := string(byteList)
	//fmt.Println(targetString)

	//tempSet := []rune(name)
	//fmt.Println(tempSet)
	//fmt.Println(tempSet[0], strconv.FormatInt(int64(tempSet[0]), 16))
	//字符的长度
	//runelength := utf8.RuneCountInString(name)
	//fmt.Println(runelength)

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

	// 字符串转换为整数，--> int
	//v7 := "111"
	//result2, err := strconv.Atoi(v7)
	//if err == nil {
	//	fmt.Println("success", result2, reflect.TypeOf(result2))
	//} else {
	//	fmt.Println("error", result2, reflect.TypeOf(result2), err)
	//}

	// 整型 (int64)--> 2进制 (string)
	//v8 := 10
	//r1 := strconv.FormatInt(int64(v8), 2)
	//r2 := strconv.FormatInt(int64(v8), 16)
	//fmt.Println(r1, reflect.TypeOf(r1))
	//fmt.Println(r2, reflect.TypeOf(r2))
	//
	//// 字符串转换为int64, 本质与Atoi相同
	//data := "100100101"
	//result, err := strconv.ParseInt(data, 2, 10)
	//fmt.Println(result, reflect.TypeOf(result), err)

	// 超大整型
	//var v1 big.Int
	//v1.SetInt64(111111)
	//v1.SetString("11111", 10)
	//fmt.Println(v1)

	//var v2 *big.Int  // 指针 无法直接赋值，一般复制的时候使用，
	//v3 := new(big.Int)
	//v3.SetInt64(22222)
	//v3.SetString("22222", 10)
	//fmt.Println(v3)

	//超大整型 运算
	//n1 := new(big.Int)
	//n1.SetInt64(100)
	//n2 := new(big.Int)
	//n2.SetInt64(200)
	//result := new(big.Int)
	//result.Add(n1,n2)
	//fmt.Println(result, reflect.TypeOf(result))
	//
	//n3 := big.NewInt(111)
	//n4 := big.NewInt(100)
	//result1 := new(big.Int)
	//result1.Add(n3,n4)
	//fmt.Println(result1.String(), reflect.TypeOf(result1.String()))
	//
	//var v1 big.Int
	//v1.SetString("111",10)
	//var v2 big.Int
	//v2.SetString("222",10)
	//
	//var result big.Int
	//result.Add(&v1, &v2)
	//fmt.Println(result.String())

	//var v1 float32 = 3.14
	//v2 := 10.1
	//v3 :=  float64(v1) + v2
	//fmt.Println(v3)
	//
	//v4 := 0.1
	//v5 := 0.2
	//v6 := v4 + v5
	//fmt.Println(v6)

	//var d1 = decimal.NewFromFloat(0.111)
	//var d2 = decimal.NewFromFloat(0.222)
	//var d3 = d1.Add(d2)
	//fmt.Println(d3)

	//result, err := strconv.ParseBool("t1")
	//fmt.Println(result, err)

	Str()

}

