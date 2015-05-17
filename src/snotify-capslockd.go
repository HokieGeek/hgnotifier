package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/exec"
	"snotify"
	"strconv"
	"strings"
	"time"
)

func sendChangedState(client *rpc.Client, state string) {
	// Creating the message object
	hdr := &snotify.Header{Timestamp: time.Now()}
	msg := &snotify.Notification{Hdr: *hdr, Name: "capslock-state", Payload: []string{state}}

	// Performing the call
	var reply int
	err := client.Call("Snotify.Notify", msg, &reply)
	if err != nil {
		log.Fatal("crap:", err)
	}

	log.Println("sent caps lock state:", state)
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

	ticker := time.NewTicker(time.Millisecond * 125)
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
	// Set the log output to file
	logf, err := os.OpenFile("/tmp/snotify.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logf.Close()
	log.SetOutput(logf)

	// Load the configuration
	configFile := "/etc/snotify.config"
	config, err := snotify.LoadConfigFromFile(configFile)
	if err != nil {
		panic(err)
	}

	// Write out hte pid
	pidstr := strconv.Itoa(os.Getpid())
	ioutil.WriteFile("/tmp/snotify-capslockd.pid", []byte(pidstr), 0644)

	address := "localhost:" + strconv.Itoa(config.Port)

	// Create the connection
	log.Print("Connecting to: ", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// FIXME: if the connection fails, it should attempt to reconnect forever
	client := jsonrpc.NewClient(conn)

	// Start polling
	pollCapsLockState(func(state string) {
		sendChangedState(client, state)
	})
}
