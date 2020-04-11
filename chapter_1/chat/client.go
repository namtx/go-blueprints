package chat

import "github.com/gorilla/websocket"

type Client struct {
	socket *websocket.Conn
	// send is a channel on which messages are sent
	send chan []byte
	room *Room
}

func (c *Client) read() {
	defer c.socket.Close()

	for {
		_, msg, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
		c.room.forward <- msg
	}
}

func (c *Client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		err := c.socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
