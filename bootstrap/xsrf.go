package bootstrap

import "github.com/astaxie/beego"

// SetupXSRF initializes XSRF protection against Cross-site request forgery
func SetupXSRF()  {
	enabled := beego.AppConfig.DefaultBool("xsrf::enabled", true)
	if enabled {
		beego.BConfig.WebConfig.EnableXSRF = enabled
		key := beego.AppConfig.String("xsrf::xsrfkey")
		beego.BConfig.WebConfig.XSRFKey = key
		beego.BConfig.WebConfig.XSRFExpire = beego.AppConfig.DefaultInt("xsrf::xsrfexpire", 3600)
	}
}
