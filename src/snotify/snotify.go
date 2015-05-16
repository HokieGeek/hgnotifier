package snotify

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
	"time"
)

type SnotifyScheme struct {
	Bg string `yaml:"a,omitempty"`
	Fg string `yaml:"a,omitempty"`
	Fn string `yaml:"a,omitempty"`
}

type SnotifyConfig struct {
	Port      int
	Scheme    SnotifyScheme `yaml:"a,omitempty"`
	Triggers  map[string][]string
	Notifiers map[string][]string
}

type Header struct {
	Timestamp time.Time
}

type Notification struct {
	Hdr     Header
	Name    string
	Payload string
}

type Snotify struct {
	notifiersPath string
	notifiers     map[string][]string
}

func (t *Snotify) Notify(notification *Notification, reply *int) error {
	log.Println("Notify(", notification.Name, ")")
	for _, notifier := range t.notifiers[notification.Name] {
		log.Println(" notifier:", notifier)
		exec := path.Join(t.notifiersPath, notifier)
		go execNotifier(exec, notification.Payload)
	}
	return nil
}

func execNotifier(notifier string, arguments string) {
	// log.Println("execing ", notifier, arguments)
	// cmd := exec.Command("/bin/sh", "-c", strings.Join([]string{notifier, arguments}, " ")
	cmd := exec.Command(notifier, arguments)

	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Print("Error executing", notifier, err)
	}
	// log.Println("Notifier out: ", out)
}

func execNotifier2(notifier string, arguments string) {
	log.Println("execing ", notifier, arguments)
	args := []string{notifier, arguments}

	err := syscall.Exec(notifier, args, os.Environ())
	if err != nil {
		log.Panic("Could not call notifier", err)
	}
}

func NewSnotify(config SnotifyConfig) *Snotify {
	n := new(Snotify)
	// FIXME: The path of the notifiers can't be this magical
	n.notifiersPath = "/home/andres/src/snotify/notifiers/"
	n.notifiers = config.Notifiers
	return n
}

// BLAH

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

func StartDataListener(config SnotifyConfig) {
	snotify := NewSnotify(config)
	rpc.Register(snotify)
	rpc.RegisterName("com.hokiegeek.snotify", snotify)
	rpc.HandleHTTP()

	log.Println("Starting data listener on port:", config.Port)

	startListener(config.Port)
}
