package func_go

import (
	"errors"
	"fmt"
	"sort"
)

// 省略前面参数类型
func sumFn(x, y int) int {
	sum := x + y
	return sum
}
func subFn(x, y int) int {
	sum := x + y
	return sum
}

// 函数可变参数
func sumFn1(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}

// return 一次返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub

}

// 返回值重命名,在函数体使用变量
func calc1(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return

}

func SortInt(sli []int) []int {
	sort.Ints(sli)
	return sli
}

// 函数类型
type myInt int
type mycalc func(int, int) int

func calc2(x, y int, cb mycalc) int {
	return cb(x, y)
}

func do(method string) mycalc {
	switch method {
	case "+":
		return sumFn
	case "-":
		return subFn
	default:
		return nil

	}
}

// 0，匿名函数里面操作函数内的变量，返回值不影响
func f1() int {
	var a int // 0
	defer func() {
		a++
	}()
	return a
}

//1
func f2() (a int) {
	defer func() {
		a++
	}()
	return a
}

// defer 执行顺序，从下到上

// recover 与 panic 结合捕获异常
func f3() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	panic("抛出一个异常")
}

func f4(a, b int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	fmt.Println(a / b)
	fmt.Println("********")

}

func readFile(filename string) error {
	if filename == "main.go" {
		fmt.Println("读取成功")
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func funcRead() {
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println("给管理员发送右键")
		}
	}()
	// 有异常做处理，捕获，不影响继续执行
	err := readFile("main.go")
	if err != nil {
		panic(err)
	}
}

func test_defer() {
	fmt.Println("start")

	defer func() {
		fmt.Println("defer...")
	}()

	fmt.Println("end")

	fmt.Printf("f1:%v f2:%v\n", f1(), f2())

	f3()
	f4(10, 0)
	fmt.Println("123333")

	funcRead()
}

func Func_method() {
	sum1 := sumFn(1, 3)
	fmt.Println(sum1)

	sum2 := sumFn1(1, 24, 5)
	fmt.Println(sum2)

	c1, c2 := calc(2, 3)
	fmt.Println(c1, c2)

	c1, c2 = calc1(3, 3)
	fmt.Println(c1, c2)

	sli1 := []int{1, 3, 2, 5, 1}
	fmt.Println(SortInt(sli1))

	fmt.Println(calc2(1, 3, sumFn))
	// 匿名函数
	fmt.Println(calc2(1, 3, func(x, y int) int {
		return x * y
	}))

	add := do("+")
	fmt.Println(add(1, 3))

	test_defer()

}
