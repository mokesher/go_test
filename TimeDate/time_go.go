package TimeDate

import (
	"fmt"
	"time"
)

func Time_method() {
	fmt.Println("in Time_method")

	timeobj := time.Now()
	//
	fmt.Printf("%d-%d-%d-%d-%d-%d \n", timeobj.Year(), timeobj.Month(), timeobj.Day(), timeobj.Hour(), timeobj.Minute(), timeobj.Second())

	fmt.Println(timeobj.Format("2006-01-02 15:04:05"))

	timestamp := timeobj.Unix() // 毫秒
	fmt.Printf("当前时间戳: %v", timestamp)

}
