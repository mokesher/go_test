package func_go

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var rmutex sync.RWMutex

//var count = 0
var m = make(map[int]int, 0)

//func sync_test1() {
//	mutex.Lock()
//	count++
//	fmt.Println("the count is:", count)
//	time.Sleep(time.Millisecond)
//	mutex.Unlock()
//	wg.Done()
//}
// 互斥锁
func sync_test(num int) {
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	fmt.Printf("key=%v value=%v\n", num, sum)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()

}

func read() {
	rmutex.RLock()
	fmt.Println("-----执行读操作")
	time.Sleep(time.Second * 2)
	rmutex.RUnlock()
	wg.Done()
}

func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

func sync_main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go write()
		//go sync_test(i)
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}
