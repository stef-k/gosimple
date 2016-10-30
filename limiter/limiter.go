// Package limiter implements various rate limiting functions
// for login protection and API access limiting.
// It uses cache2go - https://github.com/muesli/cache2go
// library to limit 3d party dependencies, e.g redis.
// cache2go has an expiry functionallity so we do not
// have to store and handle how to expire stored failed logins
package limiter

import (
	"github.com/muesli/cache2go"
	"github.com/astaxie/beego"
	"time"
)

// Login failure in memory schema to store failed login attempts
// IncomingIP will be the store key
type loginAttempt struct {
	Username       string // store which user tries to login
	IncomingIP     string // store the IP of the request
	FailedAttempts int    // login attempts
	Timestamp      time.Time
}

var (
	loginAttempts int
	loginLockMinutes int
	loginLimiterEnabled bool
	loginCache cache2go.CacheTable
)

func init() {
	loginLimiterEnabled = beego.AppConfig.DefaultBool("loginLimiterEnabled", true)
	loginAttempts = beego.AppConfig.DefaultInt("loginAttempts", 5)
	loginLockMinutes = beego.AppConfig.DefaultInt("loginLockMinutes", 1)
}
