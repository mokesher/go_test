package pointer

import "fmt"

func po1(x int) {
	x = 10
}

func po2(x *int) {
	*x = 40
}

func Pointer_method() {
	fmt.Println("in Pointer_method")

	a := 10
	p := &a            // p指针变量 类型*int ， & 取地址
	fmt.Println(p, *p) // * 地址取值

	*p = 20
	fmt.Println(&a, a)
	fmt.Println(p, *p)

	b := 1
	po1(b)
	fmt.Println(b)
	//像引用类型，直接改变内存地址
	po2(&b)
	fmt.Println(b)

	var c = new(int)
	fmt.Printf("%v---%T---%v\n", c, c, *c)
	// new 返回指针类型
	var d *int
	d = new(int)
	*d = 100
	fmt.Println(*d)

}
