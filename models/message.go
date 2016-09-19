package models

import (
	"time"
	"strconv"
)


// Message type that will be transmitted over the wire
type Message struct {
	Timestamp string
	Data map[string]string
}

// NewMessage instantiate a new message to be transmitted
// to websoket client(s)
func NewMessage(data map[string]string) Message  {
	var message Message
	message.Timestamp = strconv.FormatInt(time.Now().UnixNano(), 10)
	message.Data = data
	return message
}
