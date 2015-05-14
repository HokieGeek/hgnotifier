package main

import (
    "hgnotify"
)

func main() {
    // TODO: this is currently broken as these functions block.
    // I need for the control listener to communicate back with this dude
    stop := false

    // TODO: figure out how to use channels to stop
    hgnotify.StartDataListener(7777)
    // hgnotify.StartControlListener(7778)

    for {
        if stop {
            return
        }
    }
}
