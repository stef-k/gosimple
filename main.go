package main

import (
	_ "github.com/stef-k/gosimple/routers"
	"github.com/stef-k/gosimple/bootstrap"
	"github.com/astaxie/beego"
)

func init() {
	// setup Logger
	bootstrap.SetupLogger()
	// setup some opinionated defaults
	bootstrap.SetupDefaults()
	// setup Session
	bootstrap.SetupSession()
	// setup XSRF
	bootstrap.SetupXSRF()
}

func main() {
	beego.Run()
}

