package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/stef-k/gosimple/models"
	"encoding/json"
)

type TokenController struct {
	beego.Controller
}

type Parameters struct {
	Username string
	Password string
}

// @Title GetToken
// @Description issues a new JWT token upon client authentication
// @Param username body string true "user's name"
// @Param password body string true "user's password"
// @Success 200 {string}
// @Failure 400 body is empty
// @router /api/get-token/ [post]
func (tc *TokenController) GetToken() {

	var params Parameters
	json.Unmarshal(tc.Ctx.Input.RequestBody, &params)

	if params.Username == "" || params.Password == "" {
		tc.Abort("400")
	}

	// authenticate user
	if models.AuthenticateUser(params.Username, params.Password) {
		expiresIn := beego.AppConfig.DefaultInt("jwt::expiresIn", 48)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username" : params.Username,
			"expires" : time.Now().Add(time.Hour * time.Duration(expiresIn)).Unix(),
		})

		// signing key
		key := beego.AppConfig.String("jwt::key")

		signedString, _ := token.SignedString([]byte(key))
		tc.Data["json"] = signedString
	} else {
		tc.Data["json"] = "Authentication Failed"
	}

	tc.ServeJSON()
}
