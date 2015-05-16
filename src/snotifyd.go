package main

import (
	"gopkg.in/yaml.v2"
	"snotify"
	"io/ioutil"
    "fmt"
)

func main() {
	// Load the configuration
	// FIXME: the path can't be magical
	configFile := "/usr/etc/snotify.config"
	configBuf, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		panic("Could not read config file")
	}

	var config snotify.SnotifyConfig
	err = yaml.Unmarshal(configBuf, &config)
	if err != nil {
		panic("Could not unmarshal config")
	}

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
