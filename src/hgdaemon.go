package main

import (
	"hgnotify"
)

func main() {
	// TODO: this is currently broken as these functions block.
	// I need for the control listener to communicate back with this dude
	hgnotify.StartDataListener(7777)
	hgnotify.StartControlListener(7778)
}
