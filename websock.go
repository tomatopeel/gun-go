package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	http.HandleFunc("/ws", handler)
	http.ListenAndServe(*addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Failed to upgrade: ", err)
	}

	ticker := time.NewTicker(time.Second)

	i := 0
	for _ = range ticker.C {
		msg := []byte(fmt.Sprintf("0%d\tHello", i))
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Fatal(err)
		}
		i++
	}
}
