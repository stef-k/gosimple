package controllers
// WebSocket controller

import (
	"github.com/gorilla/websocket"
	"github.com/stef-k/gosimple/models"
	"github.com/astaxie/beego"
	"net/http"
	"fmt"
)

type WebsocketController struct {
	beego.Controller
}


// Live a WebSocket controller - upgrader
func (wc *WebsocketController) Get() {

	websocketLimits, _ := beego.AppConfig.Int("websocket::limit")

	// access if any parameters are passed
	// along with the request to connect using
	// the websockets protocol
	roomName := wc.Ctx.Input.Param("room")
	if roomName == "" {
		roomName = "lobby"
	}

	// upgrade the connection
	var upgrader = websocket.Upgrader{
		ReadBufferSize:     1024,
		WriteBufferSize:    1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Upgrade
	ws, err := upgrader.Upgrade(wc.Ctx.ResponseWriter, wc.Ctx.Request, nil)

	if _, ok := err.(websocket.HandshakeError); ok {
		panic(err)
		return
	}

	currentConnections := models.GetAllClients()

	// do not connect the client if we have reached server's limits
	if currentConnections >= websocketLimits {
		beego.Info("Server reached its limit, aborting connection")
		return
	}


	// this will probably print 127.0.0.1:somePort
	beego.Info("got new connection with IP: ", wc.Ctx.Request.RemoteAddr)
	// create a new client and do something with it
	client := models.NewClient(ws, roomName)
	// find room else create a new one
	index, room := models.FindRoom(roomName)
	if index == -1 {
		room := models.NewRoom(roomName)
		room.AddClient(client)
		models.AddRoom(room)
	} else {
		room.AddClient(client)
	}
	client.SendMessage(models.NewMessage(map[string]string{
		"1" : "Hello websocket client",
		"2" : "You are now connected",
	}))

	// client side debug, on every client connection, broadcast server status
	models.PoolBroadcast(models.NewMessage(map[string]string{
		"connectedClients": fmt.Sprintf("%v", models.GetAllClients()),
		"numberOfRooms": fmt.Sprintf("%v", models.GetNumberOfRooms()),
	}))

	// serve something so beego does not complain for missing template...
	wc.ServeJSON()
}

