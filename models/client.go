package models

import (
	"github.com/gorilla/websocket"
	"time"
	"fmt"
	"github.com/astaxie/beego"
	"sync"
)

// Client represents an entity connected using a websocket
// Connection the websocket connection
type Client struct {
	Id         string
	Room       string
	Connection *websocket.Conn
}

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// NewClient instantiate a new websocket client
func NewClient(ws *websocket.Conn, room string) *Client {
	client := new(Client)

	// Use as client's ID the websocket's pointer address
	client.Id = fmt.Sprintf("%p", ws)
	client.Connection = ws
	// check for room name else assign it to lobby as default
	if room != "" {
		client.Room = room
	} else {
		client.Room = "lobby"
	}
	client.Connection.SetReadLimit(maxMessageSize)

	// Pong check
	client.Connection.SetPongHandler(func(string) error {
		client.Connection.SetReadDeadline(time.Now().Add(pongWait)); return nil
	})

	go func() {
		// on exit remove Client and close connection
		defer func() {
			beego.Info("Removing socket client ", client.Id)
			client.Connection.Close()
			RemoveClient(client)
		}()
		for {
			// Connection check
			_, _, err := client.Connection.ReadMessage()
			if err != nil {
				if err.Error() == "websocket: close 1001 (going away)" {
					beego.Info("client left")
					RemoveClient(client)
					// client side debug, on every client disconnection, broadcast server status
					PoolBroadcast(NewMessage(map[string]string{
						"connectedClients": fmt.Sprintf("%v", GetAllClients()),
						"numberOfRooms": fmt.Sprintf("%v", GetNumberOfRooms()),
					}))
				} else {
					beego.Warn("Sock error ", err)
				}
				break
			}
		}
	}()

	beego.Info("Spawn new websocket client")

	return client
}

// SendMessage sends a message to a connected websocket client
func (c Client) SendMessage(msg Message) {
	m := &sync.Mutex{}
	m.Lock()
	defer m.Unlock()
	c.Connection.WriteJSON(msg)
}
