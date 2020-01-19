package main

import (
	"fmt"
	"github.com/gcaggia/golang-protobuf-server/protodef"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	log.Println("App is running on port 18000")
}

type UserJson struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

type UserData struct {
	User *protodef.User
	ProtoBuf string
}

func byteToString(b []byte) string {
	result := ""
	for _, val := range b {
		result += strconv.FormatInt(int64(val), 10) + " "
	}
	return "[" + strings.TrimSpace(result) + "]"
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the protobuf server API!")
	fmt.Println("Endpoint Hit: /")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":18000", router))
}

func main() {
	fmt.Println("Rest API v1.0 - protobuf server")
	handleRequests()
}
