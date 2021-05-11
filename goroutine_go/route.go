package goroutine_go

import (
	"fmt"
	"go_test/common"
	"sync"
	"time"
)

func RouteTest1() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("go-route-", i)
			time.Sleep(time.Millisecond * 80)
		}
		wg.Done() // 计数器-1
	}()

	for i := 1; i < 10; i++ {
		fmt.Println("main-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Wait()

}

func RouteTest2() {
	// 管道，引用数据类型
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30

	a := <-ch
	b := <-ch
	fmt.Printf("%v-%v 容量：%v 长度：%v\n", a, b, cap(ch), len(ch))

	var ch1 = make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch1 <- i
	}
	close(ch1)

	for v := range ch1 {
		fmt.Println(v)
	}

}

func GorouteMethod() {
	fmt.Printf("in the %v.\n", common.GetFuncName())

	//RouteTest1()
	//RouteTest2()
	start := time.Now().Unix()
	RouteTest3()
	end := time.Now().Unix()
	fmt.Println("耗时：", end-start)

}
