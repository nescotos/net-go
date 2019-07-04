package main

import (
	"log"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Fatal("unable to run server", err)
		return err
	}
	log.Print("server running on", s.port)
	return nil
}

type Handler func(w http.ResponseWriter, r *http.Request)

type Router struct {
	rules map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]Handler),
	}
}

func (r *Router) Handle(msg string, handler Handler) {
	r.rules[msg] = handler
}

func (r *Router) FindHandler(msg string) (Handler, bool) {
	handler, exist := r.rules[msg]
	return handler, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}
