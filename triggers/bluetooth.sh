#!/bin/bash

here=$(cd `dirname $0`; pwd)

osdnotifier=${here}/../notifiers/osdnotifier/osdnotifier.py

${osdnotifier} --content '${image '${here}'/bt.png}' \
               --color-fg "#0a3d91" \
               --alignment "bottom_right"
