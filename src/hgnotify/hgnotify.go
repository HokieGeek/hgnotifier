package hgnotify

import (
    "fmt"
    "log"
    "time"
    "os"
    "os/exec"
    "syscall"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "strconv"
)

type Header struct {
    Timestamp time.Time
}

type Notification struct {
    Hdr     Header
    Name    string
    Payload string
}

type HgNotify struct {
    notifiers map[string][]string
}

func (t *HgNotify) Notify(notification *Notification, reply *int) error {
    fmt.Println("Notify(", *notification, ")")
    if notification.Name == "capslock-state" {
        // TODO: notifiers[notification.Name]
        go execNotifier("/home/andres/src/hgnotifier/notifiers/capslock.sh", 
                        notification.Payload)
    }
    return nil
}

func NewHgNotify(config string) *HgNotify {
    n := new(HgNotify)
    // TODO: load config and populate notifiers
    return n
}

func execNotifier(notifier string, arguments string) {
    // log.Println("execing ", notifier, arguments)
    // cmd := exec.Command("/bin/sh", "-c", strings.Join([]string{notifier, arguments}, " ")
    cmd := exec.Command(notifier, arguments)

    out,err := cmd.Output()
    if err != nil {
        log.Fatal("Error executing", notifier, err)
    }
    log.Println("Notifier out: ", out)
}

func execNotifier2(notifier string, arguments string) {
    log.Println("execing ", notifier, arguments)
    args := []string{notifier, arguments}

    err := syscall.Exec(notifier, args, os.Environ())
    if err != nil {
        log.Panic("Could not call notifier", err)
    }
}

func startListener(port int) {
    address := ":" + strconv.Itoa(port)

    l, e := net.Listen("tcp", address)
    if e != nil {
        log.Fatal("listen error:", e)
    }

    for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
    }
}

func StartDataListener(port int) {
    hgnotify := NewHgNotify("blah")
    rpc.Register(hgnotify)
    rpc.RegisterName("com.hokiegeek.hgnotify", hgnotify)
    rpc.HandleHTTP()

    log.Println("Starting data listener on port:", port)

    startListener(port)
}
