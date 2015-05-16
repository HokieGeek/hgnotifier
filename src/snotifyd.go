package main

import (
	"log"
	"path"
	"snotify"
)

func main() {
	// FIXME: crap...
	base := "/usr"
	configFile := path.Join(base, "/etc/snotify.config")
	config, err := snotify.LoadConfigFromFile(configFile)
	if err != nil {
		log.Panic(err)
		panic("Could not load config file")
	}

	// Add the notifiers path to the config
	config.NotifiersPath = path.Join(base, "share/snotify/notifiers")

	// TODO: this is currently broken as these functions block.
	// I need for the control listener to communicate back with this dude
	// TODO: figure out how to use channels to stop
	stop := false
	snotify.StartDataListener(config)
	// snotify.StartControlListener(7778)

	for {
		if stop {
			return
		}
	}
}
