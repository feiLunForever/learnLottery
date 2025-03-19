package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"learnLottery/bootstap"
	"learnLottery/services"
	"learnLottery/web/controllers"
)

func Configure(b *bootstap.Bootstrapper) {
	giftService := services.NewGiftService()

	index := mvc.New(b.Party("/"))
	index.Register(giftService)
	index.Handle(new(controllers.IndexController))
}
