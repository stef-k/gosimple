package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/stef-k/gosimple/models"
	"time"
	"github.com/stef-k/gosimple/limiter"
)

type LoginController struct {
	beego.Controller
}

type loginParameters struct {
	UsernameOrEmail string
	Password        string
}


// @Title ShowLoginTemplate
// @Description showl the login template/form
// @Success 200 {string}
// @Failure 400 bad request
// @router /login/ [get]
func (lc *LoginController) ShowLoginTemplate()  {
	lc.Data["json"] = "You must imlement the login form in this controller"
	lc.ServeJSON()
}

// @Title login
// @Description logs in a user
// @Success 200 {string}
// @Failure 400 bad request
// @router /login/ [post]
func (lc *LoginController) Login() {
	sessionName := beego.AppConfig.String("session::sessionName")
	isLoggedin := lc.GetSession(sessionName) != nil

	if !isLoggedin {
		var params loginParameters
		json.Unmarshal(lc.Ctx.Input.RequestBody, &params)
		isAuthenticated, user := models.AuthenticateUser(params.UsernameOrEmail, params.Password)
		if isAuthenticated {
			userInfo := make(map[string]interface{})
			userInfo["username"] = user.Username
			userInfo["role"] = user.Role
			userInfo["timestamp"] = time.Now()
			lc.SetSession(sessionName, userInfo)
		} else {
			//flash := beego.NewFlash()
			//flash.Error("Your username/email and/or password are incorrect")
			limitReached, attempts, timestamp := limiter.RecordLoginAttempt(params.UsernameOrEmail, lc.Ctx.Input.IP())
			beego.Warning("Login Limit reached? ", limitReached, " remaining attempts: ", attempts, " time: ", timestamp)
		}
	}
	lc.ServeJSON()
}
