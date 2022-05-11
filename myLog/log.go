package myLog

import (
	"fmt"
	"io"
	"os"
)

func checkFile(Filename string) bool {
	var exist = true
	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		exist = false
		if err != nil {
			fmt.Println("检查文件失败")
		}
	}
	return exist
}

//写入文件
func Logfile(Log string) {
	var f1 *os.File
	var err1 error

	Filenames := "count.log" //也可将name作为参数传进来

	if checkFile(Filenames) { //如果文件存在
		f1, err1 = os.OpenFile(Filenames, os.O_APPEND|os.O_WRONLY, 0666) //打开文件,第二个参数是写入方式和权限
		if err1 != nil {
			fmt.Println("文件存在，已打开")
		}
	} else {
		f1, err1 = os.Create(Filenames) //创建文件
		if err1 != nil {
			fmt.Println("创建文件失败")
		}
	}
	_, err1 = io.WriteString(f1, Log+"\n") //写入文件(字符串)
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Printf("写入 %d 个字节\n", n)

	return
}
