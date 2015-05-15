package hgnotify

import (
    "fmt"
    "time"
)

type Header struct {
    Timestamp time.Time
}

type Trigger struct {
    Hdr     Header
    Name    string
    Payload string
}

type Control struct {
    Hdr     Header
    Payload string
}

type HgNotify int

// TODO: triggers := make(map[string]string, 0)

func (t *HgNotify) Notify(triggger *Trigger, reply *int) error {
    fmt.Println("Notify(", *triggger, ")")
    // TODO: match Trigger.Name to the triggers map that has executables as values
    //       cmd := exec.Command(triggers[Trigger.Name], [Trigger.payload])
    return nil
}

type HgNotifyCtrl int

func (t *HgNotifyCtrl) Stop(msg *Control, reply *int) error {
    fmt.Println("Stop(", *msg, ")")
    return nil
}

func (t *HgNotifyCtrl) Restart(msg *Control, reply *int) error {
    fmt.Println("Restart(", *msg, ")")
    return nil
}
