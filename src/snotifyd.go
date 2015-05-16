package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"snotify"
	// "os/exec"
	"path"
)

func main() {
	// Load the configuration
	dir, err := os.Getwd()
	if err != nil {
		panic("WHERE AM I?!")
	}

	configFile := path.Join(path.Dir(dir), "/etc/snotify.config")
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

	// Add the notifiers path to the config
	config.NotifiersPath = path.Join(path.Dir(dir), "share/snotify/notifiers")

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
