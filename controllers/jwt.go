package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/stef-k/gosimple/models"
)

type TokenController struct {
	beego.Controller
}

// @Title GetToken
// @Description issues a new JWT token upon client authentication
// @Param username body string true "user's name"
// @Param password body string true "user's password"
// @Success 200 {string}
// @Failure 400 body is empty
// @router /get-token/ [post]
func (tc *TokenController) GetToken() {

	username := tc.GetString("username")
	password := tc.GetString("password")

	if username == "" || password == "" {
		tc.Abort("400")
	}

	// authenticate user
	if models.AuthenticateUser(username, password) == false {
		expiresIn := beego.AppConfig.DefaultInt("jwt::expiresIn", 48)

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username" : username,
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
