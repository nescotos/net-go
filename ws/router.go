package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	CheckOrigin:      func(r *http.Request) bool { return true },
	HandshakeTimeout: time.Minute * 2,
}

type Router struct {
	handlers map[string]Handler
}

func (r *Router) GetHandler(name string) (Handler, bool) {
	handler, exists := r.handlers[name]
	return handler, exists
}

func (r *Router) AddHandler(name string, handler Handler) {
	r.handlers[name] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	socket, err := upgrader.Upgrade(w, request, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print(w, err.Error())
		return
	}
	client := NewClient(socket, r.GetHandler)
	go client.Write()
	go client.Read()
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[string]Handler),
	}
}
