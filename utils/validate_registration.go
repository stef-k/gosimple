package utils

import (
	"github.com/astaxie/beego"
	"fmt"
	"regexp"
	"unicode"
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

func usernameIsValid(registration *Registration) []string {

	feedback := make([]string, 0)

	if config["minUsername"] <= len(registration.Username) {
		feedback = append(feedback,
			fmt.Sprintf("Username must be at least %v characters long", config["minUsername"]))
	}

	return feedback
}

func passwordIsValid(registration *Registration) []string {

	feedback := make([]string, 0)

	if config["minPassword"] < len(registration.Password) {
		return false
	} else {
		feedback = append(feedback,
			fmt.Sprintf("Password length must be at least %v characters long", config["minPassword"]))
	}

	if config["passwordContainsCaps"] {
		hasCaps, _ := regexp.MatchString(".*[A-Z].*", registration.Password)
		if !hasCaps {
			feedback = append(feedback,
				fmt.Sprintf("Password must contain at least one capital letter"))
		}
	}

	if config["passwordContainsNumbers"] {
		hasCaps, _ := regexp.MatchString(".*[0-9].*", registration.Password)
		if !hasCaps {
			feedback = append(feedback,
				fmt.Sprintf("Password must contain at least one number"))
		}
	}

	if config["passwordContainsSymbols"] {
		hasSymbol := true
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

	return feedback
}

func mailIsValid(registration *Registration) []string {

	feedback := make([]string, 0)
	validMail, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, registration.Password)
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

	feedback = append(feedback, usernameIsValid(registration))
	feedback = append(feedback, passwordIsValid(registration))
	feedback = append(feedback, mailIsValid(registration))

	return feedback
}
