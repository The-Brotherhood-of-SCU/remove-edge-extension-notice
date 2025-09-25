package main

import (
	"fmt"
	"remove-edge-extension/runner"
)

func main() {
	r := runner.Runner{}
	for {
		fmt.Println("输入操作编号：")
		fmt.Println("1.备份Edge配置")
		fmt.Println("2.恢复Edge配置")
		fmt.Println("3.写入配置")
		fmt.Println("4.退出")
		op := 0
		fmt.Scan(&op)
		switch op {
		case 1:
			r.Backup()
		case 2:
			r.Recovery()
		case 3:
			r.WriteConfig()
		case 4:
			return
		}
	}
}
