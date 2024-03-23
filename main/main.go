package main

import (
	"fmt"
	"loging"
)

var log *loging.Loging

func initF() {
	fmt.Println("初始化函数开始工作了")
	log.Debug("初始化函数开始工作了")
}

func main() {
	log = loging.Default()
	log.Debug("hell")
	log.Debug("你好！")
	initF()
}
