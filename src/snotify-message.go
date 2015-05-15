package main

import (
	"fmt"
	"snotify"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Println("Can't send blank, idiot")
		return
	}

	address := os.Args[1]
	name := os.Args[2]
	payload := os.Args[3:]

	// Connecting
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer conn.Close()

	// Creating the message object
	hdr := &snotify.Header{Timestamp: time.Now()}
	msg := &snotify.Notification{Hdr: *hdr, Name: name, Payload: strings.Join(payload, " ")}

	// Performing the call
	var reply int
	client := jsonrpc.NewClient(conn)
	err = client.Call("snotify.Notify", msg, &reply)
	if err != nil {
		log.Fatal("crap:", err)
	}

	fmt.Println("sent message:", name, payload)
}
