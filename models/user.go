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
                            // use this to set different user types, such as Admin, Staff, User, etc
                            // for more complex situtations a Roles and a Roles_Permissions models could be more of help
	Usertype         string
	Active           bool   // use to lock down the account
}

