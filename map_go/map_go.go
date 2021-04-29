package map_go

import "fmt"

func Map_method() {
	fmt.Println("in the Map")
	// make
	var m1 = make(map[string]string, 20)
	m1["name"] = "moke"
	m1["sex"] = "男"
	fmt.Println(m1)

	// var声明，map属于引用数据类型
	var m2 = map[string]string{
		"name": "momo",
		"age":  "18",
	}
	fmt.Println(m2)
	// 推到
	m3 := map[string]string{
		"name": "momo",
		"age":  "18",
	}
	fmt.Println(m3)

	for k, v := range m3 {
		fmt.Printf("key: %v value: %v\n", k, v)
	}
	// 判断某个键存在
	v, ok := m3["name"]
	fmt.Println(v, ok)
	v, ok = m3["1name"]
	fmt.Println(v, ok)

	// delte 删除map key value
	delete(m3, "name")
	fmt.Println(m3)

	var m4 = make([]map[string]string, 3, 3)
	fmt.Println(m4[0] == nil)

	// map切片
	var m5 = make(map[string][]string)
	m5["job"] = []string{
		"aaa",
		"bbb",
	}
	m5["ho"] = []string{
		"ccc",
		"ddd",
	}

	for _, v := range m5 {
		for _, value := range v {
			fmt.Println(value)
		}
	}
	fmt.Println(m5)

}
