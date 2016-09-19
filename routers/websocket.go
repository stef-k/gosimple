package routers

import (
	"github.com/astaxie/beego"
	"github.com/stef-k/gosimple/controllers"
)

// Setup websocket routes
func SetupWebsocketRoutes() {
	beego.Router("/websocket/", &controllers.WebsocketController{})
}
