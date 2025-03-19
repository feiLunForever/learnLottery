package main

import (
	"fmt"
	"learnLottery/bootstap"
	"learnLottery/conf"
	"learnLottery/web/routes"
)

var port = 8081

func newApp() *bootstap.Bootstrapper {
	// 初始化应用
	app := bootstap.New("Go抽奖系统", "一凡Sir")
	app.Bootstrap()
	app.Configure(routes.Configure)

	return app
}

func main() {
	// 服务器集群的时候才需要区分这项设置
	// 比如：根据服务器的IP、名称、端口号等，或者运行的参数
	if port == 8081 {
		conf.RunningCrontabService = true
	}
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
