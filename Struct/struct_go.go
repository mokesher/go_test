package Struct

import "fmt"

// 自定义类型
type myInt int

// 类型别名
type myFloat = float64

type Person struct {
	name string
	age  int
	sex  string
}

func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v 年龄：%v 性别：%v \n", p.name, p.age, p.sex)
}

//指针类型的接受者
func (p *Person) setInfo(name string, age int) {
	p.name = name
	p.age = age
}

// MyInt 结构体与自定义类型支持自定义方法
type MyInt int

func (m MyInt) PrintMe() {
	fmt.Println("i am iiiii")
}

// 嵌套匿名结构体
type User struct {
	Username string
	Password string
	Address
}

type Address struct {
	Name string
	City string
}

// 父结构体
type Animal struct {
	name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动\n", a.name)
}

// 子结构体
type Dog struct {
	Age    int
	Animal //结构体嵌套 继承
}

func (d Dog) wang() {
	fmt.Printf("%v 在旺旺\n", d.name)
}

func test_struct() {
	var u User
	u.Username = "moke"
	u.Password = "123"
	u.Address.Name = "mmm"
	u.Address.City = "beijing"
	u.City = "gansu" // 当访问结构体成员时首先在结构体中查找，找不到去匿名结构体中查找
	fmt.Printf("%#v\n", u)

	var d = Dog{
		Age: 20,
		Animal: Animal{
			name: "qiqi",
		},
	}
	d.run()
	d.wang()

}

func Struct_method() {
	fmt.Println("in Struct_method")

	var a myInt = 10
	fmt.Printf("%v--%T \n", a, a)
	// 类型别名，类型显示指定的
	var b myFloat = 10.1
	fmt.Printf("%v--%T \n", b, b)
	// 1
	var p1 Person
	p1.name = "moke"
	p1.age = 18
	p1.sex = "男"
	fmt.Printf("值：%v 类型: %T \n", p1, p1)
	fmt.Printf("值：%#v 类型: %T \n", p1, p1) // Struct.Person

	// 2
	var p2 = new(Person)
	p2.name = "moke2"
	p2.age = 11
	p2.sex = "男"
	(*p2).sex = "fear"
	fmt.Printf("值：%#v 类型: %T \n", p2, p2) // *Struct.Person

	//3，引用地址
	var p3 = &Person{}
	p3.name = "moke3"
	p3.age = 11
	p3.sex = "男"
	fmt.Printf("值：%#v 类型: %T \n", p3, p3) // *Struct.Person

	//4
	var p4 = Person{
		name: "hha",
		age:  18,
		sex:  "aaa",
	}
	fmt.Printf("值：%#v 类型: %T \n", p4, p4) // Struct.Person
	//5
	var p5 = &Person{
		name: "ddd",
		age:  18,
		sex:  "aaa",
	}
	fmt.Printf("值：%#v 类型: %T \n", p5, p5) // *Struct.Person
	//6
	var p6 = &Person{
		name: "tttt",
	}
	fmt.Printf("值：%#v 类型: %T \n", p6, p6) // *Struct.Person
	//7
	var p7 = &Person{
		"tttt",
		20,
		"asd",
	}
	fmt.Printf("值：%#v 类型: %T \n", p7, p7) // *Struct.Person

	p7.setInfo("set", 1)
	p7.PrintInfo()

	var aa MyInt = 10
	aa.PrintMe()

	test_struct()

}
