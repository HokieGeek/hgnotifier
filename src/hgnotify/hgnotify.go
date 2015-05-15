package hgnotify

import (
    "fmt"
    "log"
    "time"
    "os/exec"
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
    log.Println("execing ", notifier, arguments)
    cmd := exec.Command("/bin/sh", "-c", notifier, arguments)
    out,err := cmd.Output()
    if err != nil {
        log.Fatal("Error executing", notifier, err)
    }
    log.Println("Notifier out: ", out)
}
