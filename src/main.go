package main

import (
	"dll_inject_to_wechat/src/helper"
)

var (
//核心dll

)

func main() {
	dll := "C:\\Users\\Godtoy\\source\\repos\\WechatHookDemo1\\Debug\\GetWxInfo.dll"
	var wx = "WeChat.exe"
	err := helper.Inject(wx, dll)
	println(err)
}
