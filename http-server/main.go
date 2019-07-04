package main

func main() {
	server := NewServer(":4400")
	server.Handle("/", server.AddMiddleware(HandleRoot, Logging(), CheckAuth()))
	server.Handle("/home", HandleHome)
	server.Listen()
}
