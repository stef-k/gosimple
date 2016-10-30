package routers

import (
	"github.com/stef-k/gosimple/controllers"
	"github.com/astaxie/beego"
)

// Sets some basic routes

func SetupBasicRoutes () {
	beego.Router("/", &controllers.MainController{})
	// include all registration related controllers
	beego.Include(&controllers.RegistrationController{})
	// enable classic login
	if beego.AppConfig.DefaultBool("login::classicLoginEnabled", true) {
		beego.Include(&controllers.LoginController{})
	}
	// JWT route & controller setup if disabled the route does not exist
	if beego.AppConfig.DefaultBool("jwt::enabled", true) {
		beego.Include(&controllers.TokenController{})
	}
}
