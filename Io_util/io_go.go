package Io_util

import (
	"bufio"
	"fmt"
	"go_test/common"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func Io_method() {
	fmt.Printf("in the %v.\n", common.GetFuncName())
	var fileDir = "E:\\github\\go_test\\testfile.txt"

	file, err := os.Open(fileDir)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		// n个字节
		n, err := file.Read(tempSlice)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("读取失败：", err)
		}
		strSlice = append(strSlice, tempSlice[:n]...)

	}
	fmt.Println("读取文件:", string(strSlice))

	//bufio 读取文件
	file, err = os.Open(fileDir)
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取失败：", err)
	}
	fmt.Println("bufio:", str)

	// ioutil
	byteStr, err := ioutil.ReadFile(fileDir)
	if err != nil {
		fmt.Println("读取失败：", err)
	}
	fmt.Println("ioutil:", string(byteStr))

	file, err = os.OpenFile(fileDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("读取失败：", err)
	}

	var st = "mmmmmmm111"
	file.Write([]byte(st))

	writer := bufio.NewWriter(file)
	writer.WriteString("写入" + strconv.Itoa(1) + "\n")
	writer.Flush()

	str = "Hello Golang"
	ex := ioutil.WriteFile(fileDir, []byte(str), 0666)
	if err != nil {
		fmt.Println("ioutil写入失败：", ex)
	}

}
