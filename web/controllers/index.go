package controllers

import (
	"github.com/kataras/iris/v12"
	"learnLottery/models"
	"learnLottery/services"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceGift services.GiftService
}

// http://localhost:8080/
//
//	func (c *IndexController) Get() string {
//		c.Ctx.Header("Content-Type", "text/html; charset=utf-8")
//		c.Ctx.WriteString("")
//		return "welcome to Go抽奖系统，<a href='/public/index.html'>开始抽奖</a>"
//	}
func (c *IndexController) Get() {
	c.Ctx.HTML("welcome to Go抽奖系统，<a href='/public/index.html'>开始抽奖</a>")
}

// http://localhost:8080/gifts
func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""

	datalist := c.ServiceGift.GetAll(true)
	list := make([]models.LtGift, 0)
	for _, data := range datalist {
		if data.SysStatus == 0 { // 正常状态的才需要放进来
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}
