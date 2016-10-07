package utils

import (
	"text/template"
	"gopkg.in/gomail.v2"
	"github.com/astaxie/beego"
	"bytes"
	"strconv"
)
// Utility functions to create and send emails
// 1. Add to recipients list Recipient objects
// 2. Use SendMailFromTemplate to send

// Email recipient
type recipient struct {
	// Recipient's name
	Name        string
	// Recipient's email address
	Address     string
}
// at init load all related configuration
// in the config map
var config = make(map[string]string)
var recipients = make([] recipient, 0)
var subject bytes.Buffer
var body bytes.Buffer
var port = -1

// load initial configuration from conf/app.conf
func init() {
	config["service"] = beego.AppConfig.String("mail::service")
	config["port"] = beego.AppConfig.String("mail::port")
	config["username"] = beego.AppConfig.String("mail::username")
	config["password"] = beego.AppConfig.String("mail::password")
	config["from"] = beego.AppConfig.String("mail::from")
	config["format"] = beego.AppConfig.DefaultString("mail::format", "text/plain")
	config["CompanyName"] = beego.AppConfig.String("mail::companyname")
	config["AppWebsite"] = beego.AppConfig.String("mail::appwebsite")
	config["CompanyMail"] = beego.AppConfig.String("mail::email")
	config["CompanyTel"] = beego.AppConfig.String("mail::tel")
}

func inList(address string) (bool, int) {
	for i, o := range recipients {
		if o.Address == address {
			return true, i
		}
	}
	return false, -1
}

func configIsOk() bool {
	if config["from"] != "" && len(recipients) != 0 && config["service"] != "" && config["port"] != "" &&
			config["username"] != "" && config["password"] != "" {
		port, _ = strconv.Atoi(config["port"])
		return true
	}
	beego.Warn("Cannot send mail. The configuration is not complete.")
	return false
}

// AddRecipient adds a recipient in the recipients list
// A recipient can have "user@example.com" format OR
func AddRecipient(name, address string) {

	if recipients == nil {
		recipients = make([] recipient, 0)
	}

	// check if recipient is already in list
	exists, _ := inList(address)
	if !exists {
		recipient := recipient{}
		recipient.Address = address
		recipient.Name = name

		recipients = append(recipients, recipient)
	}
}

// RemoveRecipient removes the given recipient from the recipients list
func RemoveRecipient(address string) {
	exists, index := inList(address)

	if exists {
		recipients = append(recipients[:index], recipients[index + 1:]...)
	}
}

// SendMailFromTemplate sends template based email
// templateName the name of the template (email body)
// subjectTemplate the filename of the template for the subject (email subject)
func SendMailFromTemplate(templateName, subjectTemplate string) {
	beego.Info("mail config: ", config)
	tplBody := template.New(templateName)
	tplSubject := template.New(subjectTemplate)
	template.Must(tplSubject.ParseFiles("./views/email/" + subjectTemplate))
	template.Must(tplBody.ParseFiles("./views/email/" + templateName))

	tplBody.Execute(&body, config)
	tplSubject.Execute(&subject, config)

	if configIsOk() {
		d := gomail.NewDialer(config["service"], port, config["username"], config["password"])
		s, err := d.Dial()
		if err != nil {
			panic(err)
		}

		m := gomail.NewMessage()

		for _, r := range recipients {
			m.SetHeader("From", config["from"])
			m.SetAddressHeader("To", r.Address, r.Name)
			m.SetHeader("Subject", subject.String())
			m.SetBody(config["format"], body.String())

			if err := gomail.Send(s, m); err != nil {
				beego.Error("Could not send email to ", r.Address, " error: ", err)
			} else {
				RemoveRecipient(r.Address)
			}
		}
		if len(recipients) != 0 {
			beego.Error("Could not send mail to all recipients list. Not send to: ", recipients)
			recipients = nil
		}
	}
}

// SendMail uses all values from config
// subject and message must be given manually
func SendMail (subject, message string) {
	if configIsOk() {
		d := gomail.NewDialer(config["service"], port, config["username"], config["password"])
		s, err := d.Dial()
		if err != nil {
			panic(err)
		}

		m := gomail.NewMessage()

		for _, r := range recipients {
			m.SetHeader("From", config["from"])
			m.SetAddressHeader("To", r.Address, r.Name)
			m.SetHeader("Subject", subject)
			m.SetBody(config["format"], message)

			if err := gomail.Send(s, m); err != nil {
				beego.Error("Could not send email to ", r.Address, " error: ", err)
			} else {
				RemoveRecipient(r.Address)
			}
		}

		if len(recipients) != 0 {
			beego.Error("Could not send mail to all recipients list. Not send to: ", recipients)
			recipients = nil
		}
	}
}

// SendSimpleMail sends mail, all values are being inserted manually
func SendSimpleMail(service string, port int, username, password, from, to, subject, message string) {
	d := gomail.NewDialer(service, port, username, password)
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetAddressHeader("To", to, to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	if err := gomail.Send(s, m); err != nil {
		beego.Error("Could not send email to ", to, " error: ", err)
	}
}
