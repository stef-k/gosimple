package routers

import (
	"github.com/astaxie/beego"
	"github.com/stef-k/gosimple/controllers"
)

// SetupV1Api sets a namespace /api with child namespace /v1
func SetupV1Api() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/v1",
			//beego.NSBefore(controllers.Auth),
			beego.NSInclude(
				&controllers.ApiV1Controller{},
			),
		),
	)
	beego.AddNamespace(ns)
}

