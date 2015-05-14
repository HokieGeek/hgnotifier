#!/usr/bin/python

import subprocess
import os

log = open("/tmp/system-notifier.log", 'w')
# log.write(os.environ)
os.environ['HOME'] = '/home/andres'
log.write(os.environ['HOME'])
log.close()

subprocess.call("/home/andres/src/system-notifier/triggers/bluetooth.sh", shell=True, env=os.environ)

#/etc/udev/rules.d/51-bluetooth.rules
# SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="1" RUN+="/usr/local/bin/system-notifier bluetooth enabled"
# SUBSYSTEM=="rfkill", ATTR{type}=="bluetooth", ATTR{state}=="0" RUN+="/usr/local/bin/system-notifier bluetooth disabled"
