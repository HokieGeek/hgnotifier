#!/bin/sh

sudo tee /etc/udev/rules.d/51-bluetooth.rules >/dev/null << EOF
SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="0", RUN+="/usr/bin/snotify/bluetoothstate-listener off"
SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="1", RUN+="/usr/bin/snotify/bluetoothstate-listener on"
EOF
