package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
	"os"
	"path"
	"snotify"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Println("Can't send blank, idiot")
		return
	}

	// Load the config file for he port
	base := "/usr"
	configFile := path.Join(base, "/etc/snotify.config")
	config, err := snotify.LoadConfigFromFile(configFile)
	if err != nil {
		log.Panic(err)
		panic("Could not load config file")
	}

	address := "localhost:" + strconv.Itoa(config.Port)
	// address := os.Args[1]
	name := os.Args[1]
	payload := os.Args[2:]

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
	err = client.Call("Snotify.Notify", msg, &reply)
	if err != nil {
		log.Fatal("crap:", err)
	}

	fmt.Println("sent message:", name, payload)
}
