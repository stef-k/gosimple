package bootstrap

import "github.com/astaxie/beego"

// Logger initialization writes to ./logs directory and sets 5 different files,
// one for each of the following levels: critical, error, warning, info and debug
func SetupLogger()  {
	beego.SetLogger("multifile", `{"filename":"logs/app.log", "maxdays":30,
	"separate":["critical", "error", "warning", "info", "debug"]}`)
}
