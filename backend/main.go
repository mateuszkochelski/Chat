package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type Message struct {
	Message string
}

func main() {

	var messages []Message
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(messages)
	})
	go udpListener(messages)

	http.ListenAndServe(":8080", nil)
}

func udpListener(messages []Message) {
	addr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		fmt.Println("Error during address creation")
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error during udp listening")
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error during receiving data", err)
			continue
		}

		messages = append(messages, Message{
			fmt.Sprintf("%s:%s", remoteAddr, string(buffer[:n])),
		})
	}
}
