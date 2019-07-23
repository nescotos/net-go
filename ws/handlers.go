package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func subscribeChannel(client *Client, data interface{}) {
	var message Message
	mapstructure.Decode(data, &message)
	fmt.Printf("%#v\n", message)
	client.send <- message
}
