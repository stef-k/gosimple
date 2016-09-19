package routers

import (
	"github.com/stef-k/gosimple/controllers"
	"github.com/astaxie/beego"
)

// Sets some basic routes

func SetupBasicRoutes () {
	beego.Router("/", &controllers.MainController{})
}
