package controllers

import (
	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenController struct{
	beego.Controller
}

// @Title GetToken
// @Description issues a new JWT token upon client authentication
// @Success 200 {string}
// @Failure 403 body is empty
// @router /get-token/ [get]
func (tc *TokenController) GetToken() {

	expiresIn := beego.AppConfig.DefaultInt("jwt::expiresIn", 48)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo" : "bar",

		"exp" : time.Now().Add(time.Hour * time.Duration(expiresIn)).Unix(),
	})

	// signing key
	key := beego.AppConfig.String("jwt::key")

	signedString, _ := token.SignedString([]byte(key))
	tc.Data["json"] = signedString
	tc.ServeJSON()
}
