package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send         chan Message
	socket       *websocket.Conn
	checkHandler CheckHandler
}

func (c *Client) Write() {
	for msg := range c.send {
		fmt.Printf("%#v\n", msg)
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

func (c *Client) Read() {
	var msg Message
	for {
		if err := c.socket.ReadJSON(&msg); err != nil {
			break
		}
		if handler, found := c.checkHandler(msg.Name); found {
			handler(c, msg.Data)
		}
	}
	c.socket.Close()
}

func NewClient(socket *websocket.Conn, checkHandler CheckHandler) *Client {
	return &Client{
		send:         make(chan Message),
		socket:       socket,
		checkHandler: checkHandler,
	}
}
