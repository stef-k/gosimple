package controllers

import (
	"github.com/astaxie/beego"
	"github.com/stef-k/gosimple/models"
)
// return something to test the version 1 api


type ApiV1Controller struct {
	beego.Controller
}


// Return connected websocket clients
func (avc *ApiV1Controller) Get() {
	avc.Data["json"] = map[string]int{
		"connectedClients": models.GetAllClients(),
	"numberOfRooms": models.GetNumberOfRooms()}
	avc.ServeJSON()
}
