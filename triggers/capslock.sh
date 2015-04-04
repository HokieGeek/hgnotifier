#!/bin/bash

here=$(cd `dirname $0`; pwd)

osdnotifier=${here}/../notifiers/osdnotifier/osdnotifier.py

${osdnotifier} --content '${image '${here}'/capslock.png}' \
               --alignment "top_right"
#${osdnotifier} --content 'CAPS' \
#               --color-fg "#ff0000" \
#               --alignment "top_right" \
#               --font-size 10
