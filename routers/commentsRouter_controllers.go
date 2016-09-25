package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/stef-k/gosimple/controllers:ApiV1Controller"] = append(beego.GlobalControllerRouter["github.com/stef-k/gosimple/controllers:ApiV1Controller"],
		beego.ControllerComments{
			Method: "Websocket",
			Router: `/websocket-status/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
