package slic

import "fmt"

func Slice_method() {
	fmt.Println("in the Slice")
	//空切片声明1，长度为0，需要append扩容
	var arr1 []int
	fmt.Println(arr1, arr1 == nil)
	arr1 = append(arr1, 10, 22)
	fmt.Printf("%v - %v - %v\n", arr1, len(arr1), cap(arr1))

	// 声明2
	var arr2 = []int{1, 2, 3, 4}
	fmt.Printf("%v - %T- 长度:%v\n", arr2, arr2, len(arr2))

	// 声明3
	var arr3 = []int{1: 11, 2: 22}
	fmt.Println(arr3)

	// make只用于 slice、dict、channel
	// 声明4，make，切片复制好处
	var sli = make([]int, 4, 8)
	sli[0] = 10
	sli[1] = 11
	fmt.Printf("make %v %d--%d\n", sli, len(sli), cap(sli))

	//切片的指针
	var v1 = new([]int)
	fmt.Println(v1)
	// 指针类型
	var v2 *[]int
	fmt.Println(v2)

	//循环遍历切片1
	var strSlice = []string{"python", "java", "golang"}
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}
	//循环遍历切片2
	for k, v := range strSlice {
		fmt.Println(k, v)
	}

	//基于数组定义切片
	a := [5]int{1, 2, 3, 4, 5}
	b := a[:]
	fmt.Printf("%v-%T\n", b, b)

	//切片
	//长度: 包含元素的个数
	//容量：从第一个元素到底层数组元素末尾的个数
	s := []int{2, 3, 4, 5, 1, 3}
	fmt.Printf("长度：%d - 容量：%d\n", len(s), cap(s))

	s1 := s[:3]
	fmt.Printf("长度：%d - 容量：%d\n", len(s1), cap(s1))

	// 切片合并
	ss1 := []string{"1", "2"}
	ss2 := []string{"1", "2"}
	ss1 = append(ss1, ss2...)
	fmt.Println(ss1)

	// 值类型：改变变量副本值，不会改变本身的值
	// 引用类型：改变变量副本值，会改变本身的值

	//切片属于引用数据类型，可以使用copy
	ss3 := []int{1, 2, 3}
	fmt.Printf("ss3: %v\n", ss3)
	ss4 := make([]int, 4, 4)
	copy(ss3, ss4) // dst目标 src要复制的

	fmt.Printf("ss3: %v ss4: %v\n", ss3, ss4)

	ss3[0] = 333
	fmt.Printf("ss3: %v ss4: %v\n", ss3, ss4)

	// 切片删除元素，append与切片
	// append ,删除索引为2的元素
	ss5 := []int{2, 3, 4, 5, 6, 7}
	ss6 := append(ss5[:2], ss5[3:]...)
	fmt.Println(ss6)

}
