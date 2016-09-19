package main

import (
	_ "github.com/stef-k/gosimple/routers"
	"github.com/astaxie/beego"
)

func init() {
	// setup logger
	beego.SetLogger("multifile", `{"filename":"logs/app.log", "maxdays":30,
	"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
}

func main() {
	beego.Run()
}

