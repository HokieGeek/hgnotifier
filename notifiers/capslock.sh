#!/bin/sh

id="hgnotifier_capslock"

case $1 in
    on)
        echo "A" | dzen2 -p -title-name ${id} \
                            -fn '-*-terminus-bold-r-*-*-15-*-*-*-*-*-*-*' \
                            -fg "#005f00" -bg "#000000" \
                            -w 20 -y -30 -x -35 \
                             -e 'raise'&
        sleep .1s && transset-df --name ${id} .65 >/dev/null 2>&1
        ;;
    off) pkill -f ${id} ;;
esac
