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

	beego.GlobalControllerRouter["github.com/stef-k/gosimple/controllers:RegistrationController"] = append(beego.GlobalControllerRouter["github.com/stef-k/gosimple/controllers:RegistrationController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/stef-k/gosimple/controllers:TokenController"] = append(beego.GlobalControllerRouter["github.com/stef-k/gosimple/controllers:TokenController"],
		beego.ControllerComments{
			Method: "GetToken",
			Router: `/api/get-token/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
