package main

import (
	"dll_inject_to_wechat/src/helper"
)

var (
//核心dll

)

func getWechatInstallExtPath() {

}

func inject() {

}

func unject() {

}

func main() {
	process := helper.GetProcessesByName("WeChat")
	println(process)
	inject()

}
