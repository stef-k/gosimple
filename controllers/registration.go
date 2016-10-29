package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"github.com/stef-k/gosimple/utils"
)

type RegistrationController struct {
	beego.Controller
}

// @Title register
// @Description register a new user
// @Success 200 {string}
// @Failure 400 body is empty
// @router /register/ [post]
func (rc *RegistrationController) Register() {
	registrationIsOpen, _ := beego.AppConfig.Bool("registration::is_open")
	if registrationIsOpen {
		var re utils.Registration
		json.Unmarshal(rc.Ctx.Input.RequestBody, &re)
		utils.ValidateRegistration(&re)
		errors := utils.ValidateRegistration(&re)
		if len(errors) > 0 {
			rc.Data["json"] = errors
		} else {
			// register the user and return a success message
			rc.Data["json"] = map[string]string{"response": "success"}
		}
	} else {
		rc.Data["json"] = "Registration is closed."
	}
	rc.ServeJSON()
}
