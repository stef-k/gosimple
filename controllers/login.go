package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/stef-k/gosimple/models"
	"time"
	"github.com/stef-k/gosimple/limiter"
	"fmt"
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
	limiterEnabled := beego.AppConfig.DefaultBool("limiter::loginLimiterEnabled", true)
	incomingIP := lc.Ctx.Input.IP()
	lockMinutes, _ := beego.AppConfig.Int("limiter::loginLockMinutes")
	// if limiter is enabled and login limit reached, show login locked...
	if limiter.LoginLimitReached(incomingIP) && limiterEnabled {
		lc.Data["json"] = "Login limit has been reached, please wait " + fmt.Sprintf("%v", lockMinutes) + " minutes to retry."
		lc.ServeJSON()
	} else {
		lc.Data["json"] = "You must imlement the login form in this controller"
		lc.ServeJSON()
	}
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
		limiterEnabled := beego.AppConfig.DefaultBool("limiter::loginLimiterEnabled", true)
		incomingIP := lc.Ctx.Input.IP()
		lockMinutes, _ := beego.AppConfig.Int("limiter::loginLockMinutes")
		if limiter.LoginLimitReached(incomingIP) && limiterEnabled {
			lc.Data["json"] = "Login limit has been reached, please wait " + fmt.Sprintf("%v", lockMinutes) + " minutes to retry."
			lc.ServeJSON()
		}
		var params loginParameters
		json.Unmarshal(lc.Ctx.Input.RequestBody, &params)
		isAuthenticated, user, error := models.AuthenticateUser(params.UsernameOrEmail, params.Password)
		if isAuthenticated {
			userInfo := make(map[string]interface{})
			userInfo["username"] = user.Username
			userInfo["role"] = user.Role
			userInfo["timestamp"] = time.Now()
			lc.SetSession(sessionName, userInfo)
			lc.Data["json"] = "Welcome " + user.Username
		} else {
			// with rate limiter enabled
			if limiterEnabled {
				_, attempts, _ := limiter.RecordLoginAttempt(params.UsernameOrEmail, incomingIP)
				lc.Data["json"] = "Login failed, please try again. Reason: " + error + " Remaining attempts: " + fmt.Sprintf("%v", attempts)
			} else {
				// with rate limiter disabled
				lc.Data["json"] = "Login failed, please try again. Reason: " + error
			}

		}
	}
	lc.ServeJSON()
}
