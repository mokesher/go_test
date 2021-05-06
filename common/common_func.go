package common

import (
	"runtime"
	"strings"
)

func GetFuncName() string {
	// 获取函数名称
	//fn := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	sli := strings.Split(f.Name(), ".")
	//fmt.Println(sli)
	return sli[len(sli)-1]

}
