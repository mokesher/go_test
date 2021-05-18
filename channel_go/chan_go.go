package channel_go

import (
	"fmt"
	"go_test/common"
)

func selectTest() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "Hello" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println("从intChan读取：", v)
		case v := <-stringChan:
			fmt.Println("从stringChan读取：", v)
		default:
			fmt.Println("数据获取完毕")
			return
		}

	}
}

func ChanMethod() {
	fmt.Printf("in the %v.\n", common.GetFuncName())

	// 1 双向管道
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	m1 := <-ch1
	fmt.Println(m1)

	// 只写
	ch2 := make(chan<- int, 2)
	ch2 <- 30
	fmt.Println(ch2)

	// 只读
	ch3 := make(<-chan int, 2)
	fmt.Println(ch3)

	selectTest()
}