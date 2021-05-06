package Jsonp

import (
	"encoding/json"
	"fmt"
	"go_test/common"
)

// 结构体标签
type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"` //首字母需大写，私有属性json无法访问

}

type Class struct {
	Title    string
	Students []Student
}

func Json_method() {
	fmt.Printf("in the %v.\n", common.GetFuncName())
	var s1 = Student{
		ID:   17,
		Name: "moke",
	}
	// 结构体对象转换为json字符串
	fmt.Println(s1)
	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("json: %v\n", jsonStr)

	// json字符串转为结构体对象
	var str = `{"ID":17,"Name":"moke"}`
	var s2 Student
	err := json.Unmarshal([]byte(str), &s2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", s2)
		fmt.Println(s2.Name)
	}

	c := Class{
		Title:    "07班",
		Students: make([]Student, 0),
	}
	for i := 1; i <= 10; i++ {
		s := Student{
			ID:   i,
			Name: fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}
	fmt.Println(c)

	strByte, _ := json.Marshal(c)
	fmt.Println(string(strByte))

	var cc = &Class{}
	err = json.Unmarshal(strByte, cc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", c)
	}

}
