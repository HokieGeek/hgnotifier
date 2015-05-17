package snotify

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"time"
)

type SnotifyScheme struct {
	Bg string
	Fg string
	Fn string
}

type SnotifyConfig struct {
	Port          int
	Scheme        SnotifyScheme
	Triggers      map[string][]string
	Notifiers     map[string][]string
	NotifiersPath string
}

func LoadConfigFromFile(file string) (SnotifyConfig, error) {
	configBuf, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		panic("Could not read config file")
	}

	var config SnotifyConfig
	err = yaml.Unmarshal(configBuf, &config)
	if err != nil {
		panic("Could not unmarshal config")
	}

	return config, err
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
	config SnotifyConfig
}

func (t *Snotify) Notify(notification *Notification, reply *int) error {
	log.Println("Notify(", notification.Name, ")")
	for _, notifier := range t.config.Notifiers[notification.Name] {
		log.Println(" notifier:", notifier)
		exec := path.Join(t.config.NotifiersPath, notifier)
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

func NewSnotify(config SnotifyConfig) *Snotify {
	n := new(Snotify)
	n.config = config
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

		runtime.Gosched()
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
