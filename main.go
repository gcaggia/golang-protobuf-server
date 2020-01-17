package main

import (
	"log"
	"net/http"
)

func init() {
	log.Println("App is running on port 18000")
}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Incoming request on '/'")
		writer.Write([]byte("Received!"))
	})

	http.ListenAndServe("127.0.0.1:18000", nil)
}
