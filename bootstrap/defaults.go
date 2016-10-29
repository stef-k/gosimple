package bootstrap

import "github.com/astaxie/beego"

// SetupDefaults sets some defaults that may or maynot need user's input
// This is an opinionated setup
func SetupDefaults()  {
	// flash name equals appname + _FLASH
	appname := beego.AppConfig.DefaultString("appname", "GoSimple")
	beego.BConfig.WebConfig.FlashName = appname + "_FLASH"

	// SERVER NAME defaults to appname
	beego.BConfig.ServerName = appname
}
