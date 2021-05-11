package goroutine_go

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func PrimeNum(intChan chan int, PrimeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			PrimeChan <- num
		}
	}
	exitChan <- true
	wg.Done()
}

func PutNum(intChan chan int, totalNum int) {
	for i := 2; i < totalNum; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

func PrintPrime(PrimeChan chan int) {
	for v := range PrimeChan {
		fmt.Println(v)
	}
	wg.Done()
}

func RouteTest3() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 100000)
	exitChan := make(chan bool, 16)

	// 存放数字的协程
	wg.Add(1)
	go PutNum(intChan, 200000)

	//统计素数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go PrimeNum(intChan, primeChan, exitChan)
	}

	// 打印素数的协程
	wg.Add(1)
	go PrintPrime(primeChan)

	//判断exitChan是否存满值
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("执行完毕.")

}
