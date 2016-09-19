package models

// Room Pool containing all rooms currently active
var RoomPool = make([] *Room, 0)

func AddRoom(r *Room) {
	RoomPool = append(RoomPool, r)
}

// FindRoom finds a room in the RoomPool
// and returns its index and value.
// If the room does not exist in pool returns -1
// If the RoomPool is empty returns -1
func FindRoom(d string) (int, *Room) {
	for i, v := range RoomPool {
		if v.Name == d {
			return i, v
		}
	}
	return -1, &Room{}
}

// RemoveRoom removes domain from the DomainPool
func RemoveRoom(r *Room) {

	index, _ := FindRoom(r.Name)

	if index != -1 {
		RoomPool = append(RoomPool[:index], RoomPool[index + 1:]...)
	}
}

// RemoveClient removes a client from a room
// if the client is the only client in the room
// the room is also removed from the RoomPool
func RemoveClient(c *Client)  {
	for _, room := range RoomPool{
		if room.Name == c.Room {
			if room.HasClient(c) {
				room.DeleteClient(c)
				if room.ClientCount() == 0 {
					RemoveRoom(room)
				}
			}
		}
	}
}

// ListRooms return the number of registered rooms and
// a slice of strings with all room names in pool
func ListRooms() (int, []string) {
	pool := make([]string, 0)
	for _, obj := range RoomPool {
		pool = append(pool, obj.Name)
	}
	return len(RoomPool), pool
}

// PoolBroadcast send a message to all clients of the Pool
func PoolBroadcast(m Message)  {
	for _, room := range RoomPool {
		room.RoomBroadcast(m)
	}
}

// GetAllClients returns the number of all connected clients
func GetAllClients() int {

	clientSum := 0
	for _, room := range RoomPool {
		clientSum += len(room.ClientPool)
	}
	return clientSum
}

// GetNumberOfRooms get the number of all rooms
func GetNumberOfRooms()  int{
	return len(RoomPool)
}
