package models

import (
	"time"
	"github.com/stef-k/gosimple/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id               int
	Username         string `orm:"unique"`
	Password         string
	Email            string `orm:"unique"`
	EmailConfirmed   bool
	ConfirmationCode string // used during registration confirmation TODO move to sepparate table and tight up security
	Created          time.Time `orm:"auto_now_add;type(datetime)"`
	LastLogin        time.Time `orm:"type(datetime)"`
                            // use this to set different user types, such as Admin, Staff, User, etc
                            // for more complex situtations a Roles and a Roles_Permissions models could be more of help
	Role             string
	Active           bool   // use to lock down the account
}

// New creates and returns a new User object.
// To save the new object use the Save method.
func New(username, password, email, role string) *User {
	var user User
	user.Username = username
	var err error
	if user.Password, err = utils.GeneratePassword(password); err != nil {
		beego.Error("could not generate user hashed password, ", err.Error())
	}
	user.Email = email
	user.Role = role

	return &user
}


// GetUser searches the database for a user object with the given ID
func GetUser(id int) *User {
	var user User
	user.Id = id
	o := orm.NewOrm()

	if err := o.Read(&user); err == nil {
		return &user
	} else {
		return new(User)
	}
}

// GetUserByUsername searches the database for a user object with the given username
func GetUserByUsername(username string) *User {
	var user User
	user.Username = username
	o := orm.NewOrm()

	if err := o.Read(&user); err == nil {
		return &user
	} else {
		return new(User)
	}
}

// GetUserByUsername searches the database for a user object with the given email
func GetUserByEmail(email string) *User {
	var user User
	user.Email = email
	o := orm.NewOrm()

	if err := o.Read(&user); err == nil {
		return &user
	} else {
		return new(User)
	}
}

// Authenticate authenticates a User by his username or email and his password
// Uppon successful authentication returns true, the User object and an empty string
// When unsuccessful, returns false, empty User object and a string with the error message
func AuthenticateUser(usernameOrEmail, password string) (bool, User, string){
	user := GetUserByUsername(usernameOrEmail)
	var error string
	// if user is found by his username check his password
	if (User{}) != *user {
		correctPwd := utils.CheckPassword(password, user.Password)
		if !correctPwd {
			error = "Incorrect password."
		}
		return correctPwd, *user, error
	} else {
		user := GetUserByEmail(usernameOrEmail)
		if (User{}) != *user {
			correctPwd := utils.CheckPassword(password, user.Password)
			if !correctPwd {
				error = "Incorrect password."
			}
			return correctPwd, *user, error
		} else {
			error = "Could not find a user with the secified username or password."
			return false, User{}, error
		}
	}
}

// AllUsers return all stored users from the database
func AllUsers() []*User {
	var users []*User
	o := orm.NewOrm()
	if _, err := o.QueryTable("user").All(&users); err != nil {
		beego.Warning("could not find any users in database, ", err.Error())
	}
	return users
}

// Save saves a User object to database
func (u *User) Save() error {
	o := orm.NewOrm()
	if _, err := o.Insert(u); err == nil {
		return nil
	} else {
		return err
	}
}

// Delete deletes a user object from database
func (u *User) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(u); if err == nil {
		return nil
	} else {
		return err
	}
}
