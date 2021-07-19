package test

import (
	"fmt"
	"go_test/common"
	"strings"
)

func Test_method() {
	fmt.Printf("in the %v.\n", common.GetFuncName())

	//var n1 float64 = 3.1
	//var n2 float64 = 4.1
	//d1 := decimal.NewFromFloat(n1).Add(decimal.NewFromFloat(n2))
	//fmt.Println(n1 + n2)
	//fmt.Println(d1)

	var volumeId = "11111"
	var path = "/api/v1/block/iscsi/gateways/{volume_id}/num"
	rp := strings.Replace(path, "{volume_id}", fmt.Sprintf("%v", volumeId), -1)
	fmt.Println("replace: ", rp)

}
