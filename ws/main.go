package main

import "net/http"

func main() {
	router := NewRouter()
	router.AddHandler("SUBSCRIBE_CHANNEL", subscribeChannel)
	http.Handle("/", router)
	http.ListenAndServe(":4432", nil)
}
