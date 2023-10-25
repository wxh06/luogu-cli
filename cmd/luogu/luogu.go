package main

import "github.com/wxh06/luogu-cli"

func main() {
	data, _ := luogu.Request[luogu.UserData]("GET", "https://www.luogu.com.cn/user/108135", nil)
	println(data.CurrentData.User.Name)
}
