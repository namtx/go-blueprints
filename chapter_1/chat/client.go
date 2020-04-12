package chat

import (
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	socket *websocket.Conn
	// send is a channel on which messages are sent
	send chan *Message
	room *Room

	UserData map[string]interface{}
}

func (c *Client) read() {
	defer c.socket.Close()

	for {
		var msg Message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.UserData["name"].(string)
		if avatarUrl, ok := c.UserData["avatar_url"]; ok {
			msg.AvatarURL = avatarUrl.(string)
		}

		c.room.forward <- &msg
	}
}

func (c *Client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			break
		}
	}
}
