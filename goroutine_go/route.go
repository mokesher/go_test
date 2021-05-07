package goroutine_go

import (
	"fmt"
	"go_test/common"
)

func GorouteMethod() {
	fmt.Printf("in the %v.\n", common.GetFuncName())

}
