package bootstrap

import "github.com/astaxie/beego"

// SetupSession Session initilization, reads from config possible setup settings
// and initializes the Session object accordingly. Defaults to memory
func SetupSession() {
	if enabled, ok := beego.AppConfig.Bool("session::sessionEnabled"); ok == nil {
		if enabled {
			// enable session
			beego.BConfig.WebConfig.Session.SessionOn = true
			// storage
			storage := beego.AppConfig.DefaultString("session::sessionStorage", "memory")
			if storage != "memory" {
				sessionStorage := beego.AppConfig.String("session::sessionStorage")
				beego.BConfig.WebConfig.Session.SessionProviderConfig = sessionStorage
			}
			beego.BConfig.WebConfig.Session.SessionProvider = storage

			// session name
			sessionName := beego.AppConfig.String("session::sessionName")
			if sessionName == "" {
				sessionName = beego.AppConfig.String("appname")
			}
			beego.BConfig.WebConfig.Session.SessionName = sessionName

			// session lifetime
			sessionLifetime := beego.AppConfig.DefaultInt64("session::sessionLifetime", 3600)
			beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = sessionLifetime
			// cookie lifetime
			sessionCookieLifetime := beego.AppConfig.DefaultInt("session::sessionCookieLifetime", 0)
			beego.BConfig.WebConfig.Session.SessionCookieLifeTime = sessionCookieLifetime


		} // if enabled

	}

}
