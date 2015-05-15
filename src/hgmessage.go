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
    if len(os.Args[1:]) <= 0 {
        fmt.Println("Can't send blank, idiot")
        return
    }

    name := os.Args[1]
    payload := os.Args[2:]

    // Connecting
    conn, err := net.Dial("tcp", "localhost:7777")
    if err != nil {
        log.Fatal("dialing:", err)
    }
    defer conn.Close()

    // Creating the message object
    hdr := &hgnotify.Header{Timestamp: time.Now()}
    msg := &hgnotify.Trigger{Hdr: *hdr, Name: name, Payload: strings.Join(payload, " ")}

    // Performing the call
    var reply int
    client := jsonrpc.NewClient(conn)
    err = client.Call("HgNotify.Notify", msg, &reply)
    if err != nil {
        log.Fatal("crap:", err)
    }

    fmt.Println("sent message:", name, payload)
}
