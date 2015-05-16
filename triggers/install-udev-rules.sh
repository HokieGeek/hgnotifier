#!/bin/sh

sudo tee /etc/udev/rules.d/50-snotify-bluetooth.rules >/dev/null << EOF
SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="0", RUN+="/usr/share/snotify/triggers/bluetooth-state.sh off"
SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="1", RUN+="/usr/share/snotify/triggers/bluetooth-state.sh on"
EOF

sudo tee /etc/udev/rules.d/99-snotify-lowbattery.rules >/dev/null << EOF
SUBSYSTEM=="power_supply", ATTR{status}=="Discharging", ATTR{capacity}=="5", RUN+="/usr/share/snotify/triggers/low-battery.sh on"
SUBSYSTEM=="power_supply", ATTR{status}=="Charging", RUN+="/usr/share/snotify/triggers/low-battery.sh off"
EOF

sudo udevadm control --reload-rules
sudo udevadm trigger
