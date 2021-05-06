package interface_go

import (
	"fmt"
	"go_test/common"
)

type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func Interface_method() {
	fmt.Printf("in the %v.\n", common.GetFuncName())
	p := Phone{
		Name: "mi",
	}
	p.start()

}
