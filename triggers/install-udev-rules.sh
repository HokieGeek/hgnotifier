#!/bin/sh

sudo tee /etc/udev/rules.d/50-snotify-bluetooth.rules >/dev/null << EOF
SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="0", RUN+="/usr/share/snotify/bluetooth-osd.sh off"
SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="1", RUN+="/usr/share/snotify/bluetooth-osd.sh on"
EOF

sudo tee /etc/udev/rules.d/99-snotify-lowbattery.rules >/dev/null << EOF
SUBSYSTEM=="power_supply", ATTR{status}=="Discharging", ATTR{capacity}=="5", RUN+="/usr/share/snotify/lowbattery-osd.sh on"
SUBSYSTEM=="power_supply", ATTR{status}=="Charging", RUN+="/usr/share/snotify/lowbattery-osd.sh off"
EOF

sudo udevadm control --reload-rules
sudo udevadm trigger
