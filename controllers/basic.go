package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Address"] = beego.AppConfig.String("httpaddr")
	c.Data["Port"] = beego.AppConfig.String("httpport")
	c.TplName = "index.html"
}
