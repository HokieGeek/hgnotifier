package main

import (
	"fmt"
	"snotify"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/exec"
	"strings"
	"time"
)

func sendChangedState(client *rpc.Client, state string) {
	// Creating the message object
	hdr := &snotify.Header{Timestamp: time.Now()}
	msg := &snotify.Notification{Hdr: *hdr, Name: "capslock-state", Payload: state}

	// Performing the call
	var reply int
	err := client.Call("Snotify.Notify", msg, &reply)
	if err != nil {
		log.Fatal("crap:", err)
	}

	fmt.Println("sent caps lock state:", state)
}

func getCapsLockState() string {
	cmd := exec.Command("/bin/sh", "-c", "xset -q | awk '$0 ~ /Caps Lock/ { print $4 }'")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Error getting the caps lock state:", err)
	}
	return strings.TrimSpace(string(out))
}

func pollCapsLockState(stateChange func(state string)) {
	lastState := "off"

	ticker := time.NewTicker(time.Millisecond * 350)
	for {
		select {
		case <-ticker.C:
			// log.Println("Polling caps lock state")
			currState := getCapsLockState()
			if currState != lastState {
				lastState = currState
				stateChange(lastState)
			}
		}
	}
}

func main() {
	address := os.Args[1]

	// Create the connection
	log.Print("Connecting to: ", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer conn.Close()
	// FIXME: if the connection fails, it should attempt to reconnect forever
	client := jsonrpc.NewClient(conn)

	// Start polling
	pollCapsLockState(func(state string) {
		sendChangedState(client, state)
	})
}
