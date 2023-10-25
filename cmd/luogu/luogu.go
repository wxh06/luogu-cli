package main

import (
	"github.com/charmbracelet/glamour"
	"github.com/wxh06/luogu-cli"
)

func main() {
	data, err := luogu.Request[luogu.UserData]("GET", "https://www.luogu.com.cn/user/108135", nil)
	if err != nil {
		panic(err)
	}

	println(data.CurrentData.User.Name)
	introduction, err := glamour.Render(data.CurrentData.User.Introduction, "notty")
	if err != nil {
		panic(err)
	}
	println(introduction)
}
