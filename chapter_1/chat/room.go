package chat

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/namtx/go-blueprints/chapter_1/trace"
	"github.com/stretchr/objx"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
	WriteBufferSize: messageBufferSize,
}

type Room struct {
	// forward is a channle that holds incoming messages
	// that should be forwarded to the other clients
	forward chan *Message

	join  chan *Client
	leave chan *Client

	clients map[*Client]bool

	Tracer trace.Tracer
}

func NewRoom() *Room {
	return &Room{
		forward: make(chan *Message),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
		Tracer:  trace.Off(),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			// a client joins the room
			r.clients[client] = true
			r.Tracer.Trace("New client joined")

		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			r.Tracer.Trace("Client left")

		case msg := <-r.forward:
			r.Tracer.Trace("Message received: ", msg.Message)
			for client := range r.clients {
				client.send <- msg
				r.Tracer.Trace(" -- Send to client")
			}
		}
	}
}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("Failed to get auth cookie:", err)
		return
	}

	client := &Client{
		socket:   socket,
		send:     make(chan *Message, messageBufferSize),
		room:     r,
		UserData: objx.MustFromBase64(authCookie.Value),
	}

	r.join <- client

	defer func() {
		r.leave <- client
	}()

	go client.write()

	client.read()
}
