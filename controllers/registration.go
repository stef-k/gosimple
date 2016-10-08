package controllers

import "github.com/astaxie/beego"

func RegisterUser(c *MainController)  {
	registrationIsOpen, _ := beego.AppConfig.Bool("registration::is_open")
	if registrationIsOpen {

	}
}
