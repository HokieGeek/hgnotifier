package main

import (
	"gopkg.in/yaml.v2"
	"hgnotify"
	"io/ioutil"
)

func main() {
	// Load the configuration
	// FIXME: the path can't be magical
	configFile := "/home/aperez/src/hgnotifier/hgnotifier.config"
	configBuf, err := ioutil.ReadFile(configFile)
	if err != nil {
		// fmt.Println(err)
		panic("Could not read config file")
	}

	// fmt.Println(string(configBuf))

	var config hgnotify.HgNotifierConfig
	err = yaml.Unmarshal(configBuf, &config)
	if err != nil {
		panic("Could not unmarshal config")
	}

	// TODO: this is currently broken as these functions block.
	// I need for the control listener to communicate back with this dude
	// TODO: figure out how to use channels to stop
	stop := false
	hgnotify.StartDataListener(config)
	// hgnotify.StartControlListener(7778)

	for {
		if stop {
			return
		}
	}
}
