package reflect_go

import (
	"fmt"
	"go_test/common"
	"reflect"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (p Person) GetInfo() string {
	var st = fmt.Sprintf("姓名：%v 年龄：%v 成绩：%v", p.Name, p.Age, p.Score)
	return st
}

func (p *Person) SetInfo(name string, age int, score int) {
	p.Name = name
	p.Age = age
	p.Score = score
}

func (p Person) Print() {
	fmt.Println("这是一个打印方法")
}

type myInt int32

func PrintStruct(p interface{}) {
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}
	// 1、通过类型变量里面的Field可以获取结构体字段
	field0 := t.Field(0)
	fmt.Println(field0)
	fmt.Println(field0.Name)
	fmt.Println(field0.Type)
	fmt.Println(field0.Tag.Get("json"))

	// 2、通过类型变量名称获取
	field1, ok := t.FieldByName("Age")
	if ok {
		fmt.Println(field1.Name)
		fmt.Println(field1.Type)
		fmt.Println(field1.Tag.Get("json"))

	}

	// 3、通过NumField
	var fieldCount = t.NumField()
	fmt.Println("结构体有", fieldCount, "个属性")

	// 4、通过值变量获取结构体属性对应值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println(v.FieldByName("Score"))

	//
	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称：%v 属性值：%v 属性类型：%v 属性Tag：%v \n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}

}

func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Struct && t.Elem().Kind() != reflect.Struct {
		fmt.Println("传入的参数不是一个结构体")
		return
	}

	// 1、通过类型变量的method可以获取结构体的方法，与结构体方法的Ascii有关
	method0 := t.Method(0)
	fmt.Println(method0.Name)
	fmt.Println(method0.Type)

	// 2、通过类型变量获取这个结构体有多少个方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println("Print method:")
		fmt.Println(method1.Name)
		fmt.Println(method1.Type)
	}

	//3、通过值变量执行方法，
	v.Method(1).Call(nil)
	v.MethodByName("Print").Call(nil)
	// 切片
	info := v.MethodByName("GetInfo").Call(nil)
	fmt.Println("info:", info)

	//4、执行方法传入参数，修改
	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的参数不是结构体指针类型")
		return
	}
	var params []reflect.Value
	params = append(params, reflect.ValueOf("fake"))
	params = append(params, reflect.ValueOf(22))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params)

	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println("info2:", info2)

	// 获取方法的数量
	fmt.Println("方法数量：", t.NumMethod())
}

func reflectFn(x interface{}) {
	// 获取变量的类型
	v := reflect.TypeOf(x)
	// 类型名称 ：Person
	// 类型种类 ：底层类型
	fmt.Printf("类型：%v 类型名称：%v 类型种类：%v \n", v, v.Name(), v.Kind())

}

func reflectValue(x interface{}) {
	//类型断言
	b, _ := x.(int)
	var num = 10 + b
	fmt.Printf("reflect num: %v\n", num)
	//返回reflect.Value 类型，包含原始值的值信息，可以相互转换
	v := reflect.ValueOf(x)
	fmt.Printf("v:%v\n", v)

	kind := v.Kind()
	switch kind {
	case reflect.Int:
		var num1 = v.Int() + 20
		fmt.Printf("reflect num1: %v\n", num1)
	case reflect.Float32:
		fmt.Printf("reflect float32: %v\n", v.Float()+10.1)
	default:
		fmt.Println("nonononon")
	}
}

func reflectSetValue(x interface{}) {
	// 反射设置变量值
	v := reflect.ValueOf(x)
	//fmt.Println(v.Elem().Kind())
	kind := v.Elem().Kind()
	switch kind {
	case reflect.Int64:
		v.Elem().SetInt(123)
	case reflect.String:
		v.Elem().SetString("gogogo")
	}

}

func Test_reflect() {
	fmt.Printf("in the %v.\n", common.GetFuncName())

	a := 10
	b := 23.4
	c := true
	d := "你好Go"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	var f = Person{
		Name: "moke",
		Age:  18,
	}

	reflectFn(e)
	reflectFn(f)

	var h = 25
	reflectFn(&h)

	var i = [3]int{1, 2, 3}
	reflectFn(i)

	var j = []int{11, 22, 33}
	reflectFn(j)

	reflectValue(h)

	var k int64 = 22
	reflectSetValue(&k)
	fmt.Println(k)
	var l = "你好 golang"
	reflectSetValue(&l)
	fmt.Println(l)

	P := Person{
		Name:  "moke",
		Age:   18,
		Score: 100,
	}
	PrintStruct(P)
	PrintStructFn(&P)

}
