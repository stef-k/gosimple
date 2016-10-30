package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/stef-k/gosimple/models"
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/stef-k/gosimple/limiter"
	"fmt"
)

type TokenController struct {
	beego.Controller
}

// CustomClaims a structure to add custom fields to JWT's claim
type CustomClaims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

type jwtParameters struct {
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

	var params jwtParameters
	json.Unmarshal(tc.Ctx.Input.RequestBody, &params)

	if params.Username == "" || params.Password == "" {
		tc.Abort("400")
	}

	limiterEnabled := beego.AppConfig.DefaultBool("limiter::loginLimiterEnabled", true)
	incomingIP := tc.Ctx.Input.IP()
	lockMinutes, _ := beego.AppConfig.Int("limiter::loginLockMinutes")
	if limiter.LoginLimitReached(incomingIP) && limiterEnabled {
		tc.Data["json"] = "Login limit has been reached, please wait " + fmt.Sprintf("%v", lockMinutes) + " minutes to retry."
		tc.ServeJSON()
	}
	// authenticate user
	isAuthenticated, user, error := models.AuthenticateUser(params.Username, params.Password)
	if isAuthenticated {

		expiresIn := beego.AppConfig.DefaultInt("jwt::expiresIn", 12)

		claims := CustomClaims{
			user.Username,
			user.Role,

			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiresIn)).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// signing key
		key := beego.AppConfig.String("jwt::key")

		signedString, err := token.SignedString([]byte(key))

		if err == nil {
			tc.Data["json"] = signedString
		} else {
			tc.Data["json"] = "Could not generate JWT token"
		}

	} else {
		// with rate limiter enabled
		if limiterEnabled {
			_, attempts, _ := limiter.RecordLoginAttempt(params.Username, incomingIP)
			tc.Data["json"] = "Login failed, please try again. Reason: " + error + " Remaining attempts: " + fmt.Sprintf("%v", attempts)
		} else {
			// with rate limiter disabled
			tc.Data["json"] = "Login failed, please try again. Reason: " + error
		}
	}

	tc.ServeJSON()
}


// ValidateToken validates a JWT token
// Returns True and the Claim value uppon successful validation,
// false & empty CustomClaim otherwise
// NOTE: its uppon to you to check if the token is expired.
func ValidateToken(tokenString string) (bool, *CustomClaims) {

	signingKey := beego.AppConfig.String("jwt::key")

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	if err != nil {
		return false, &CustomClaims{}
	}

	return token.Valid, token.Claims.(*CustomClaims)

}
