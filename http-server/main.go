package main

func main() {
	server := NewServer(":4400")
	server.Handle("/", "GET", server.AddMiddleware(HandleRoot, Logging(), CheckAuth()))
	server.Handle("/home", "GET", HandleHome)
	server.Handle("/create", "POST", PostRequest)
	server.Listen()
}
