package utils

import (
	"github.com/astaxie/beego"
	"fmt"
)

type Registration struct {
	Username   string
	Password   string
	PasswordRe string
	Email      string
}

var config = make(map[string]string)

// load registration configuration
func init() {
	config["minUsername"], _ = beego.AppConfig.Int("registration::min_username")
	config["minPassword"], _ = beego.AppConfig.Int("registration::min_password")
	config["passwordContainsCaps"], _ = beego.AppConfig.Bool("registration::password_contain_caps")
	config["passwordContainsNumbers"], _ = beego.AppConfig.Bool("registration::password_contain_numbers")
	config["passwordContainsSymbols"], _ = beego.AppConfig.Bool("registration::password_contain_symbols")
}

func usernameIsValid(registration *Registration) (bool, []string) {

	feedback := make([]string, 0)

	return config["minUsername"] >= len(registration.Username), feedback
}

func passwordIsValid(registration *Registration) (bool, []string) {

	feedback := make([]string, 0)

	if config["minPassword"] < len(registration.Password) {
		return false
	} else {
		feedback = append(feedback,
			fmt.Sprintf("Password length must be at least %v characters long", config["minPassword"]))
	}

	return true, feedback
}

func mailIsValid(registration *Registration) (bool, []string) {

	feedback := make([]string, 0)

	return false, feedback
}

// ValidateRegistration validates registration fields, returns true if the fields
// pass the validation, false otherwise. Additionally returns a slice of strings
// containing errors if validation fails
func ValidateRegistration(registration *Registration) (bool, []string) {

	feedback := make([]string, 0)

	return false, feedback
}
