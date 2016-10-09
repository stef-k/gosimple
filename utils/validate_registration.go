package utils

import (
	"github.com/astaxie/beego"
	"fmt"
	"regexp"
	"unicode"
)

type Registration struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	PasswordRe string   `json:"passwordRe"`
	Email      string   `json:"email"`
}

var minUsername = 5
var minPassword = 6
var passwordContainsCaps = true
var passwordContainsNumbers = true
var passwordContainsSymbols = true

// load registration configuration
func init() {
	minUsername, _ = beego.AppConfig.Int("registration::min_username")
	minPassword, _ = beego.AppConfig.Int("registration::min_password")
	passwordContainsCaps, _ = beego.AppConfig.Bool("registration::password_contain_caps")
	passwordContainsNumbers, _ = beego.AppConfig.Bool("registration::password_contain_numbers")
	passwordContainsSymbols, _ = beego.AppConfig.Bool("registration::password_contain_symbols")
}

func usernameIsValid(registration *Registration) []string {

	feedback := make([]string, 0)

	if len(registration.Username) < minUsername {
		feedback = append(feedback,
			fmt.Sprintf("Username must be at least %v characters long", minUsername))
	}

	return feedback
}

func passwordIsValid(registration *Registration) []string {

	feedback := make([]string, 0)

	if minPassword > len(registration.Password) {
		feedback = append(feedback,
			fmt.Sprintf("Password length must be at least %v characters long", minPassword))
	}

	if passwordContainsCaps {
		hasCaps, _ := regexp.MatchString(".*[A-Z].*", registration.Password)
		if !hasCaps {
			feedback = append(feedback,
				fmt.Sprintf("Password must contain at least one capital letter"))
		}
	}

	if passwordContainsNumbers {
		hasNumbers, _ := regexp.MatchString(".*[0-9].*", registration.Password)
		if !hasNumbers {
			feedback = append(feedback,
				fmt.Sprintf("Password must contain at least one number"))
		}
	}

	if passwordContainsSymbols {
		hasSymbol := false
		for _, char := range registration.Password {
			if unicode.IsPunct(char) || unicode.IsSymbol(char) {
				hasSymbol = true
			}
		}
		if !hasSymbol {
			feedback = append(feedback,
				fmt.Sprintf("Password must contain at least one symbol"))
		}
	}

	if registration.Password != registration.PasswordRe {
		feedback = append(feedback, "Password confirmation does not match")
	}

	return feedback
}

func mailIsValid(registration *Registration) []string {

	feedback := make([]string, 0)
	validMail, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, registration.Email)
	if !validMail {
		feedback = append(feedback,
			fmt.Sprintf("The mail is not valid"))
	}
	return feedback
}

// ValidateRegistration validates registration fields, returns a slice of strings
// containing errors if validation fails else the slice should be empty.
func ValidateRegistration(registration *Registration) []string {

	feedback := make([]string, 0)

	feedback = append(feedback, usernameIsValid(registration)...)
	feedback = append(feedback, passwordIsValid(registration)...)
	feedback = append(feedback, mailIsValid(registration)...)

	return feedback
}
