package models

import "time"

type User struct {
	Id               int
	Username         string `orm:"unique"`
	Password         string
	Email            string `orm:"unique"`
	EmailConfirmed   bool
	ConfirmationCode string // used during registration confirmation
	Created          time.Time `orm:"auto_now_add;type(datetime)"`
	Usertype         string // use this to set different user types, such as superadmin, admin, user, etc
	Active           bool   // use to lock down the account
}

