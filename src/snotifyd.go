package main

import (
	"io/ioutil"
	"log"
	"os"
	"snotify"
	"strconv"
)

func main() {
	// Set the log output to file
	logf, err := os.OpenFile("/tmp/snotify.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logf.Close()
	log.SetOutput(logf)

	// Load the config file
	configFile := "/etc/snotify.config"
	config, err := snotify.LoadConfigFromFile(configFile)
	if err != nil {
		log.Panic(err)
		panic("Could not load config file")
	}

	// Write out the pid
	pidstr := strconv.Itoa(os.Getpid())
	ioutil.WriteFile("/tmp/snotifyd.pid", []byte(pidstr), 0644)

	// Add the notifiers path to the config
	config.NotifiersPath = "/usr/share/snotify/notifiers"

	// TODO: this is currently broken as these functions block.
	// I need for the control listener to communicate back with this dude
	// TODO: figure out how to use channels to stop
	snotify.StartDataListener(config)
	// snotify.StartControlListener(7778)

	stop := false
	for {
		if stop {
			return
		}
	}
}
