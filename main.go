package main

import (
	"encoding/json"
	"github.com/gcaggia/golang-protobuf-server/protodef"
	"github.com/golang/protobuf/proto"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func init() {
	log.Println("App is running on port 18000")
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

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Incoming request on '/'")
		writer.Write([]byte("Received!"))
	})

	http.HandleFunc("/user/healthcheck", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Incoming request on '/user/healthcheck'")

		user := &protodef.User{
			Id: 1,
			Name: "Eric Freime",
			Email: "efreime@gmail.com",
			Active: true,
		}

		data, err := proto.Marshal(user)
		if err != nil {
			log.Fatal("marshal call error: ", err)
		}
		// log raw protobuf object
		log.Println("protobuf: " + byteToString(data))
		userData := UserData{
			User: user,
			ProtoBuf: byteToString(data),
		}
		// writer.Write([]byte("Protobuf generated for id " + strconv.FormatInt(int64(user.Id), 10)))
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(userData)
	})

	http.ListenAndServe("127.0.0.1:18000", nil)
}
