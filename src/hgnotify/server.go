package hgnotify

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
)

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
	rpc.Register(new(HgNotify))
	rpc.RegisterName("com.hokiegeek.hgnotify", new(HgNotify))
	rpc.HandleHTTP()

	log.Println("Starting data listener on port:", port)

	startListener(port)
}

func StartControlListener(port int) {
	rpc.Register(new(HgNotifyCtrl))
	rpc.RegisterName("com.hokiegeek.hgnotify.control", new(HgNotifyCtrl))
	rpc.HandleHTTP()

	log.Println("Starting control listener on port:", port)

	startListener(port)
}
