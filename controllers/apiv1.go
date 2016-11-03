package controllers

import (
	"github.com/astaxie/beego"
	"github.com/stef-k/gosimple/models"
	"github.com/stef-k/gosimple/limiter"
	"fmt"
)
// return something to test the version 1 api


type ApiV1Controller struct {
	beego.Controller
}

// Return the number of all connected clients and rooms
// @router /websocket-status/ [get]
func (a *ApiV1Controller) Websocket() {
	a.Data["json"] = map[string]int{
		"connectedClients": models.GetAllClients(),
		"numberOfRooms": models.GetNumberOfRooms()}
	a.ServeJSON()
}

// Show an example of API rate limiting
// @router /rate-limiter/ [get]
func (a *ApiV1Controller) RateLimiter() {
	current, limit := limiter.ApiRecordRequest(a.Ctx.Input.IP())
	if !limiter.ApiLimitReached(a.Ctx.Input.IP()) {
		a.Data["json"] = map[string]int{
			"currentRequests": current,
			"requestLimit": limit}
	} else {
		apiRequestsLimit := beego.AppConfig.DefaultInt("limiter::apiRequestsLimit", 20)
		limitingTimeUnit := beego.AppConfig.DefaultString("limiter::limitingTimeUnit", "min")
		perTime := beego.AppConfig.DefaultInt("limiter::perTime", 1)
		msg := "You have reached the request limit for this API endpoint, "
		msg += fmt.Sprintf("current limits are %v requests/%v %v", apiRequestsLimit, perTime, limitingTimeUnit)
		a.Data["json"] = msg
	}

	a.ServeJSON()
}
