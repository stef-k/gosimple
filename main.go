package main

import (
	_ "github.com/stef-k/gosimple/routers"
	"github.com/stef-k/gosimple/bootstrap"
	"github.com/astaxie/beego"
)

func init() {
	// setup Logger
	bootstrap.SetupLogger()
	// setup Session
	bootstrap.SetupSession()
}

func main() {
	beego.Run()
}

