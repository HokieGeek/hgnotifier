package hgnotify

import (
	"fmt"
	"time"
)

type Header struct {
	Timestamp time.Time
}

type Msg struct {
	Hdr     Header
	Payload string
}

type HgNotify int

func (t *HgNotify) SendMsg(msg *Msg, reply *int) error {
	fmt.Println("SendMsg(", *msg, ")")
	return nil
}

type HgNotifyCtrl int

func (t *HgNotifyCtrl) Stop(msg *Msg, reply *int) error {
	fmt.Println("Stop(", *msg, ")")
	return nil
}

func (t *HgNotifyCtrl) Restart(msg *Msg, reply *int) error {
	fmt.Println("Restart(", *msg, ")")
	return nil
}
