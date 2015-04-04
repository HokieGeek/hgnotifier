#!/bin/bash

here=$(cd `dirname $0`; pwd)

osdnotifier=${here}/../notifiers/osdnotifier/osdnotifier.py

${osdnotifier} --content '${image '${here}'/bt.png}' \
               --alignment "bottom_right"
