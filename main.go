package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	log.Println("App is running on port 18000")
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
	ProtoBuf string `json:"protobuf"`
}

// global users variable, act as in memory database
var Users []User

func byteToString(b []byte) string {
	result := ""
	for _, val := range b {
		result += strconv.FormatInt(int64(val), 10) + " "
	}
	return "[" + strings.TrimSpace(result) + "]"
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the protobuf server API!")
	fmt.Println("Endpoint Hit: /")
}

func fetchUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GET /users")
	json.NewEncoder(w).Encode(Users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GET /users/:id")
	fmt.Fprintf(w, "GET /users/:id")
}

func postUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: POST /users")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	json.Unmarshal(reqBody, &user)
	// update our global users array to include new user
	Users = append(Users, user)
	// return user newly created
	json.NewEncoder(w).Encode(user)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: PUT /users/:id")
	fmt.Fprintf(w, "PUT /users/:id")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: DELETE /users/:id")
	fmt.Fprintf(w, "DELETE /users/:id")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/users",      fetchUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser   ).Methods("GET")
	router.HandleFunc("/users",      postUser  ).Methods("POST")
	router.HandleFunc("/users/{id}", editUser  ).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe("127.0.0.1:18000", router))
}

func main() {
	fmt.Println("Rest API v1.0 - protobuf server")
	handleRequests()
}
