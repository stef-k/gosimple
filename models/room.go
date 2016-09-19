package models

import "fmt"

// A Room containing a Pool of WebSocket clients
// each time a websocket client connects, it is added
// to the ClientPool so we can reference at a later time
type Room struct {
	Name string
	Public bool
	ClientPool []*Client
}

// NewRoom create a new Room
func NewRoom(name string) *Room {
	room := new(Room)
	room.Public = true
	room.Name = name

	return room
}

// AddClient add a client in client pool
func (r *Room) AddClient(client *Client) {
	r.ClientPool = append(r.ClientPool, client)
}

// DeleteClient removes a client from the pool
func (r *Room) DeleteClient(c *Client) {
	index, _ := r.FindClient(c)
	if index != -1 {
		r.ClientPool = append(r.ClientPool[:index], r.ClientPool[index + 1:]...)
	}
}

// FindClient finds a client in the ClientPool
// and returns its index and value.
// If the room does not exists in pool returns -1
// If the ClientPool is empty returns -1
func (r *Room) FindClient(c *Client) (int, *Client) {
	if len(r.ClientPool) == 0 {
		return -1, &Client{}
	} else {
		for i, v := range r.ClientPool {
			if v.Id == c.Id {
				return i, v
			}
		}
		return -1, &Client{}
	}
}

// HasClient checks if the room has a client registered in its pool
func (r *Room) HasClient(client *Client) bool {

	for _, obj := range r.ClientPool {
		if obj.Id == client.Id {
			return true
		}
	}
	return false
}

// ListClients returns a slice with IDs of all connected clients
func (r *Room) ListClients() []string {
	pool := make([]string, 0)
	for _, obj := range r.ClientPool {
		pool = append(pool, fmt.Sprintf("%p", obj.Id))
	}
	return pool
}

// ClientCount returns the number of domain's connected clients
func (r *Room) ClientCount() int {
	return len(r.ClientPool)
}

// RoomBroadcast send a Message to all domain's connected clients
func (r *Room) RoomBroadcast(msg Message) {
	for _, client := range r.ClientPool {
		client.SendMessage(msg)
	}
}
