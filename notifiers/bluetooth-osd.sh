#!/bin/bash

here=$(cd $(dirname $0); pwd)
id="hgnotifier_bluetooth"

case $1 in
    on)
        echo "^i(${here}/bt.xbm)" | dzen2 -p -title-name ${id} \
                                    -fn '-*-terminus-bold-r-*-*-12-*-*-*-*-*-*-*' \
                                    -fg "#0000ff" -bg "#ffffff" \
                                    -w 20 -h 22 -y -60 -x -38 \
                                    -e 'raise'&
        sleep .025s && transset-df --name ${id} .65 >/dev/null 2>&1
        ;;
    off) pkill -f ${id} ;;
esac
