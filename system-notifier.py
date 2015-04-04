#!/usr/bin/python

import subprocess

subprocess.call("/home/andres/src/system-notifier/triggers/bluetooth.sh", shell=True)

#/etc/udev/rules.d/51-bluetooth.rules
# SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="1" RUN+="/usr/local/bin/system-notifier bluetooth enabled"
# SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="0" RUN+="/usr/local/bin/system-notifier bluetooth disabled"
