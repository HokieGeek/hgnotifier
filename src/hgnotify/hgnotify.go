package hgnotify

import (
    "fmt"
    "time"
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
    // TODO: match Trigger.Name to the triggers map that has executables as values
    //       cmd := exec.Command(triggers[Trigger.Name], [Trigger.payload])
    return nil
}

func NewHgNotify(config string) *HgNotify {
    n := new(HgNotify)
    // TODO: load config and populate notifiers
    return n
}
