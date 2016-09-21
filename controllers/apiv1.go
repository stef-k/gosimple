package controllers

import (
	"github.com/astaxie/beego"
	"github.com/stef-k/gosimple/models"
)
// return something to test the version 1 api


type ApiV1Controller struct {
	beego.Controller
}

func (a *ApiV1Controller) URLMapping() {
	a.Mapping("Websocket", a.Websocket)
}

// Return the number of all connected clients and rooms
// @router /websocket-status/ [get]
func (a *ApiV1Controller) Websocket() {
	a.Data["json"] = map[string]int{
		"connectedClients": models.GetAllClients(),
		"numberOfRooms": models.GetNumberOfRooms()}
	a.ServeJSON()
}
