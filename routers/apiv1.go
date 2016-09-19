package routers

import (
	"github.com/astaxie/beego"
	"github.com/stef-k/gosimple/controllers"
)

// SetupV1Api sets a namespace /api/v1/websocket
func SetupV1Api() {
	beego.Router("/api/v1/websocket", &controllers.ApiV1Controller{})
}
// for some reason Beego's namespace function does not work for me :/
// this should work
//func SetupV1Api() {
//	ns := beego.NewNamespace("/api",
//		beego.NSNamespace("/v1",
//			//beego.NSBefore(controllers.Auth),
//			beego.NSInclude(
//				&controllers.ApiV1Controller{},
//			),
//		),
//	)
//	beego.AddNamespace(ns)
//}

