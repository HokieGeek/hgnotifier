package main

import (
	"fmt"
	"hgnotify"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"strings"
	"time"
)

func main() {
	input := os.Args[1:]

	if len(input) <= 0 {
		fmt.Println("Can't send blank, idiot")
		return
	}

	// Connecting
	conn, err := net.Dial("tcp", "localhost:7777")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer conn.Close()

	// Creating the message object
	hdr := &hgnotify.Header{Timestamp: time.Now()}
	msg := &hgnotify.Msg{Hdr: *hdr, Payload: strings.Join(input, " ")}

	// Performing the call
	var reply int
	client := jsonrpc.NewClient(conn)
	err = client.Call("HgNotify.SendMsg", msg, &reply)
	if err != nil {
		log.Fatal("crap:", err)
	}

	fmt.Println("sent message:", input)
}
