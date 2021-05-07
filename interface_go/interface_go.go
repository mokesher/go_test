package interface_go

import (
	"fmt"
	"go_test/common"
)

// 接口是一个规范
type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

type Computer struct {
}

type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}

func (p Camera) stop() {
	fmt.Println("相机关机")
}

// 电话与相机都实现了usb接口
func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

// 空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("%v-%T\n", a, a)
}

func MyPrint(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("...default")
	}
}

func EmptyInterface() {
	var a interface{}
	a = 20
	v, ok := a.(int)
	if ok {
		fmt.Println("a 为int", v)
	} else {
		fmt.Println("断言失败")
	}
	show(a)
	a = true
	show(a)
	a = "sss"
	show(a)

	// map表示任意类型
	var m1 = make(map[string]interface{})
	m1["name"] = "moke"
	m1["age"] = 22
	fmt.Println(m1)
	var s1 = []interface{}{1, 2, "23", "tre"}
	fmt.Println(s1)

}

// 接口定义参数与返回值
type Animaler interface {
	SetName(string)
	GetName() string
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func Interface_method() {
	fmt.Printf("in the %v.\n", common.GetFuncName())
	phone := Phone{
		Name: "mi",
	}
	var camera = Camera{}
	var com = Computer{}
	com.work(phone)
	com.work(camera)

	EmptyInterface()
	d := &Dog{
		Name: "xx",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("xxx")
	fmt.Println(d1.GetName())

}
