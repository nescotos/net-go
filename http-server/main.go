package main

func main() {
	server := NewServer(":4400")
	server.router.Handle("/", HandleRoot)
	server.router.Handle("/home", HandleHome)
	server.Listen()
}
